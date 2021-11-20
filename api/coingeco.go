package api

import (
	"context"
	"cryptolist/common"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	marketEndpoint      string = "coins/markets"
	marketChartEndpoint string = "coins/%s/market_chart"
)

type Header struct {
	Key   string
	Value string
}

type Options struct {
	Headers []*Header
}

type CoinGeco struct {
	endpoint string
}

func NewCoinGeco(endpoint string) *CoinGeco {
	return &CoinGeco{
		endpoint: endpoint,
	}
}

func (cg *CoinGeco) get(ctx context.Context, path string, params url.Values, opts *Options) (*http.Response, error) {
	u, err := url.Parse(cg.endpoint)
	if err != nil {
		return nil, err
	}
	u.Path += path
	u.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	if opts != nil {
		if len(opts.Headers) > 0 {
			for _, h := range opts.Headers {
				req.Header.Set(h.Key, h.Value)
			}
		}
	}

	client := &http.Client{}
	return client.Do(req.WithContext(ctx))
}

func (cg *CoinGeco) Markets(ctx context.Context, currency string, perPage int) ([]*common.Market, error) {
	resp, err := cg.get(ctx, marketEndpoint, url.Values{
		"vs_currency": {currency},
		"per_page":    {fmt.Sprint(perPage)},
	}, nil)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("error: %d - %s", resp.StatusCode, body))
	}

	var markets []*common.Market
	if err = json.Unmarshal(body, &markets); err != nil {
		return nil, err
	}

	return markets, nil
}

func (cg *CoinGeco) MarketChart(ctx context.Context, id, currency string, days uint) (*common.MarketChart, error) {
	resp, err := cg.get(ctx, fmt.Sprintf(marketChartEndpoint, id), url.Values{
		"vs_currency": {currency},
		"days":        {fmt.Sprint(days)},
	}, nil)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("error: %d - %s", resp.StatusCode, body))
	}

	var chart *common.MarketChart
	if err = json.Unmarshal(body, &chart); err != nil {
		return nil, err
	}

	chart.Name = id
	chart.Currency = currency

	return chart, nil
}
