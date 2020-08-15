package node

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/DE-labtory/zulu/account/bitcoin/chaincfg"
	"github.com/DE-labtory/zulu/types"
)

type Adapter interface {
	ListUTXO(addr string) ([]UTXO, error)
	EstimateFees() (float64, error)
}

type adapter struct {
	httpClient *httpClient
}

func NewAdapter(network types.Network) *adapter {
	return &adapter{
		httpClient: NewHttpClient(chaincfg.Supplier[network].NodeUrl),
	}
}

func (a *adapter) ListUTXO(addr string) ([]UTXO, error) {
	return a.httpClient.ListUTXO(addr)
}

func (a *adapter) EstimateFees() (float64, error) {
	estimates, err := a.httpClient.GetFeeEstimates()
	if err != nil {
		return 0, err
	}
	// TODO: Maybe we can get estimates different by confirmation number
	return estimates["1"], nil
}

type UTXO struct {
	Txid  string `json:"txid"`
	Vout  int    `json:"vout"`
	Value int    `json:"value"`
}

type httpClient struct {
	*http.Client
	baseUrl string
}

func NewHttpClient(baseUrl string) *httpClient {
	return &httpClient{
		Client:  http.DefaultClient,
		baseUrl: baseUrl,
	}
}

func (c *httpClient) ListUTXO(addr string) ([]UTXO, error) {
	result, err := c.requestTemplate("GET", fmt.Sprintf("/address/%s/utxo", addr), nil,
		func(resp []byte) (interface{}, error) {
			var utxos []UTXO
			if err := json.Unmarshal(resp, &utxos); err != nil {
				return nil, err
			}
			return utxos, nil
		})
	if err != nil {
		return nil, err
	}
	return result.([]UTXO), nil
}

func (c *httpClient) GetFeeEstimates() (map[string]float64, error) {
	result, err := c.requestTemplate("GET", "/fee-estimates", nil,
		func(resp []byte) (interface{}, error) {
			var estimates map[string]float64
			if err := json.Unmarshal(resp, &estimates); err != nil {
				return nil, err
			}
			return estimates, nil
		})
	if err != nil {
		return nil, err
	}
	return result.(map[string]float64), nil
}

func (c *httpClient) requestTemplate(method, path string, body io.Reader,
	callback respResolveFunc) (interface{}, error) {
	req, err := c.makeRequest(method, path, body)
	if err != nil {
		return nil, err
	}
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	raw, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return callback(raw)
}

func (c *httpClient) makeRequest(method, path string, body io.Reader) (*http.Request, error) {
	return http.NewRequest(method, fmt.Sprintf("%s%s", c.baseUrl, path), body)
}

type respResolveFunc func(resp []byte) (interface{}, error)
