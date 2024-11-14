package web1337

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/net/proxy"
)

type Options struct {
	ChainID         string
	WorkflowVersion uint
	NodeURL         string
	ProxyURL        string
}

type ChainInfo struct {
	NodeURL         string
	WorkflowVersion uint
}

type Web1337 struct {
	Chains       map[string]ChainInfo
	CurrentChain string
	Proxy        http.RoundTripper
}

func mapToJSON(data map[string]interface{}) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return ""
	}

	return string(jsonData)
}

func NewWeb1337(options Options) (*Web1337, error) {
	web1337 := &Web1337{
		Chains: make(map[string]ChainInfo),
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

	web1337.CurrentChain = options.ChainID

	web1337.Chains[options.ChainID] = ChainInfo{
		NodeURL:         options.NodeURL,
		WorkflowVersion: options.WorkflowVersion,
	}

	return web1337, nil
}

func (sdk *Web1337) getRequest(url string) ([]byte, error) {
	client := &http.Client{Transport: sdk.Proxy, Timeout: 10 * time.Second}

	resp, err := client.Get(sdk.Chains[sdk.CurrentChain].NodeURL + url)
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

	resp, err := client.Post(sdk.Chains[sdk.CurrentChain].NodeURL+url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
