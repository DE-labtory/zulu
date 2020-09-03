package bitcoin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/DE-labtory/zulu/account/bitcoin/chaincfg"
	"github.com/DE-labtory/zulu/types"
)

type Adapter interface {
	ListUTXO(addr string) ([]Unspent, error)
	EstimateFeeRate() (float64, error)
	SendRawTx(raw string) (TxResult, error)
}

type adapter struct {
	httpClient *httpClient
}

func NewAdapter(network types.Network) *adapter {
	return &adapter{
		httpClient: newHttpClient(chaincfg.Supplier[network].NodeUrl),
	}
}

func (a *adapter) ListUTXO(addr string) ([]Unspent, error) {
	return a.httpClient.ListUTXO(addr)
}

func (a *adapter) EstimateFeeRate() (float64, error) {
	feeRates, err := a.httpClient.GetFeeRateEstimates()
	if err != nil {
		return 0, err
	}
	if _, ok := feeRates["1"]; !ok {
		return defaultFeeRate, nil
	}
	return feeRates["1"], nil
}

func (a *adapter) SendRawTx(raw string) (TxResult, error) {
	txId, err := a.httpClient.SendRawTxData(raw)
	if err != nil {
		return TxResult{}, err
	}
	return TxResult{
		TxId: txId,
	}, nil
}

type httpClient struct {
	*http.Client
	baseUrl string
}

func newHttpClient(baseUrl string) *httpClient {
	return &httpClient{
		Client:  http.DefaultClient,
		baseUrl: baseUrl,
	}
}

func (c *httpClient) ListUTXO(addr string) ([]Unspent, error) {
	result, err := c.requestTemplate("GET", fmt.Sprintf("/address/%s/utxo", addr), nil,
		func(resp []byte) (interface{}, error) {
			var utxos []Unspent
			if err := json.Unmarshal(resp, &utxos); err != nil {
				return nil, err
			}
			return utxos, nil
		})
	if err != nil {
		return nil, err
	}
	return result.([]Unspent), nil
}

func (c *httpClient) GetFeeRateEstimates() (map[string]float64, error) {
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

func (c *httpClient) SendRawTxData(raw string) (string, error) {
	result, err := c.requestTemplate("POST", "/tx", bytes.NewBufferString(raw),
		func(resp []byte) (interface{}, error) {
			// TODO: change response type
			return string(resp), nil
		})
	if err != nil {
		return "", err
	}
	return result.(string), nil
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
