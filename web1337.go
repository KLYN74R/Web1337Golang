/*

Web1337 by KLY

For Golang devs

*/

package web1337

import (
	"bytes"
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	SIGNATURES_TYPES "github.com/KLYN74R/Web1337Golang/signatures_types"
	TXS_TYPES "github.com/KLYN74R/Web1337Golang/txs_types"
	"golang.org/x/net/proxy"
)

type Options struct {
	SymbioteID      string
	WorkflowVersion int
	NodeURL         string
	ProxyURL        string
}

type SymbioteInfo struct {
	NodeURL         string
	WorkflowVersion int
}

type Web1337 struct {
	Symbiotes       map[string]SymbioteInfo
	CurrentSymbiote string
	Proxy           http.RoundTripper
}

func NewWeb1337(options Options) (*Web1337, error) {

	web1337 := &Web1337{
		Symbiotes: make(map[string]SymbioteInfo),
	}

	if options.ProxyURL != "" {
		var transport http.RoundTripper
		if options.ProxyURL[:4] == "http" {
			proxyURL, err := url.Parse(options.ProxyURL)
			if err != nil {
				return nil, err
			}
			transport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
		} else if options.ProxyURL[:5] == "socks" {
			dialer, err := proxy.SOCKS5("tcp", options.ProxyURL[7:], nil, proxy.Direct)
			if err != nil {
				return nil, err
			}
			transport = &http.Transport{Dial: dialer.Dial}
		} else {
			return nil, errors.New("unsupported proxy URL")
		}
		web1337.Proxy = transport
	} else {
		web1337.Proxy = http.DefaultTransport
	}

	web1337.CurrentSymbiote = options.SymbioteID
	web1337.Symbiotes[options.SymbioteID] = SymbioteInfo{
		NodeURL:         options.NodeURL,
		WorkflowVersion: options.WorkflowVersion,
	}

	return web1337, nil
}

func (sdk *Web1337) getRequest(url string) ([]byte, error) {
	client := &http.Client{Transport: sdk.Proxy, Timeout: 10 * time.Second}
	resp, err := client.Get(sdk.Symbiotes[sdk.CurrentSymbiote].NodeURL + url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func (sdk *Web1337) postRequest(url string, payload interface{}) ([]byte, error) {
	client := &http.Client{Transport: sdk.Proxy, Timeout: 10 * time.Second}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	resp, err := client.Post(sdk.Symbiotes[sdk.CurrentSymbiote].NodeURL+url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func (sdk *Web1337) GetCurrentCheckpoint() ([]byte, error) {
	return sdk.getRequest("/quorum_thread_checkpoint")
}

func (sdk *Web1337) GetSymbioteInfo() ([]byte, error) {
	return sdk.getRequest("/symbiote_info")
}

func (sdk *Web1337) GetGeneralInfoAboutKLYInfrastructure() ([]byte, error) {
	return sdk.getRequest("/my_info")
}

func (sdk *Web1337) GetSyncState() ([]byte, error) {
	return sdk.getRequest("/sync_state")
}

func (sdk *Web1337) GetBlockByBlockID(blockID string) ([]byte, error) {
	return sdk.getRequest("/block/" + blockID)
}

func (sdk *Web1337) GetBlockBySID(shard string, sid string) ([]byte, error) {
	return sdk.getRequest("/block_by_sid/" + shard + "/" + sid)
}

func (sdk *Web1337) GetFromState(shard string, cellID string) ([]byte, error) {
	return sdk.getRequest("/state/" + shard + "/" + cellID)
}

func (sdk *Web1337) GetTransactionReceiptById(txID string) ([]byte, error) {
	return sdk.getRequest("/tx_receipt/" + txID)
}

func (sdk *Web1337) GetAggregatedFinalizationProofForBlock(blockID string) ([]byte, error) {
	return sdk.getRequest("/aggregated_finalization_proof/" + blockID)
}

type TransactionTemplate struct {
	V       int         `json:"v"`
	Creator string      `json:"creator"`
	Type    string      `json:"type"`
	Nonce   int         `json:"nonce"`
	Fee     int         `json:"fee"`
	Payload interface{} `json:"payload"`
	Sig     string      `json:"sig"`
}

func (sdk *Web1337) GetTransactionTemplate(workflowVersion int, creator, txType string, nonce, fee int, payload interface{}) TransactionTemplate {
	return TransactionTemplate{
		V:       workflowVersion,
		Creator: creator,
		Type:    txType,
		Nonce:   nonce,
		Fee:     fee,
		Payload: payload,
		Sig:     "",
	}
}

func (sdk *Web1337) CreateDefaultTransaction(originShard, yourAddress string, yourPrivateKey ed25519.PrivateKey, nonce, fee int, recipient string, amountInKLY int, rev_t *int) (TransactionTemplate, error) {
	workflowVersion := sdk.Symbiotes[sdk.CurrentSymbiote].WorkflowVersion

	payload := map[string]interface{}{
		"type":   SIGNATURES_TYPES.DEFAULT_SIG,
		"to":     recipient,
		"amount": amountInKLY,
	}
	if rev_t != nil {
		payload["rev_t"] = *rev_t
	}

	transaction := sdk.GetTransactionTemplate(workflowVersion, yourAddress, TXS_TYPES.TX, nonce, fee, payload)
	dataToSign := sdk.CurrentSymbiote + string(workflowVersion) + originShard + TXS_TYPES.TX + fmt.Sprintf("%v", payload) + string(nonce) + string(fee)
	transaction.Sig = base64.StdEncoding.EncodeToString(ed25519.Sign(yourPrivateKey, []byte(dataToSign)))

	return transaction, nil
}

func (sdk *Web1337) SendTransaction(transaction TransactionTemplate) ([]byte, error) {
	return sdk.postRequest("/transaction", transaction)
}
