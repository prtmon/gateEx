package gateEx

import (
	"context"
	"fmt"
	"github.com/gateio/gateapi-go/v6"
	gatews "github.com/gateio/gatews/go"
	"github.com/prtmon/tools"
	"net/http"
	"net/url"
	"strings"
)

type ExchangeConfig struct {
	Name          string `json:"name"`           //交易所名称 name
	ApiKey        string `json:"api_key"`        //交易所api key
	ApiSecretKey  string `json:"api_secret_key"` //交易所secret key
	ApiPassphrase string `json:"api_passphrase"` //交易所API密码
	ProxyUrl      string `json:"proxy_url"`      //代理服务器地址
	SettleCoin    string `json:"settle_coin"`    //现货为计价币种,合约为结算币种
	Leverage      int64  `json:"leverage"`       //杠杆倍数,部分交易所支持小数位的倍数系数
}

type Exchange struct {
	Name          string             `json:"name"`           //交易所名称 name
	ApiKey        string             `json:"api_key"`        //交易所api key
	ApiSecretKey  string             `json:"api_secret_key"` //交易所secret key
	ApiPassphrase string             `json:"api_passphrase"` //交易所API密码
	ProxyUrl      string             `json:"proxy_url"`      //代理服务器地址
	SettleCoin    string             `json:"settle_coin"`    //现货为计价币种,合约为结算币种
	Leverage      int64              `json:"leverage"`       //杠杆倍数,部分交易所支持小数位的倍数系数
	apiClient     *gateapi.APIClient //api client
	wsClient      *gatews.WsService  //websocket client
	apiCtx        context.Context    //全局context
	Uid           string             `json:"uid"` //用户ID
}

func NewExchange(conf *ExchangeConfig) *Exchange {
	ex := &Exchange{
		Name:         conf.Name,
		ApiKey:       conf.ApiKey,
		ApiSecretKey: conf.ApiSecretKey,
		SettleCoin:   conf.SettleCoin,
		ProxyUrl:     conf.ProxyUrl,
		Leverage:     conf.Leverage,
	}

	accountDetail, err := ex.GetAccountDetail()
	if err != nil {
		panic(err.Error())
	}
	ex.Uid = tools.Int64ToString(accountDetail.UserId)
	return ex
}

func (e *Exchange) ApiClient() *gateapi.APIClient {
	if e.apiClient == nil {
		cfg := gateapi.NewConfiguration()
		cfg.HTTPClient = &http.Client{
			Transport: &http.Transport{
				Proxy: func(request *http.Request) (*url.URL, error) {
					if e.ProxyUrl == "" {
						return nil, nil
					}
					proxy, err := url.Parse(e.ProxyUrl)
					if err != nil {
						return nil, nil
					}
					return proxy, nil
				},
				MaxConnsPerHost:     2,
				MaxIdleConnsPerHost: 2,
			},
		}
		e.apiClient = gateapi.NewAPIClient(cfg)
	}
	return e.apiClient
}

func (e *Exchange) ApiCtx() context.Context {
	if e.apiCtx == nil {
		e.apiCtx = context.WithValue(context.Background(), gateapi.ContextGateAPIV4, gateapi.GateAPIV4{
			Key:    e.ApiKey,
			Secret: e.ApiSecretKey,
		})
	}
	return e.apiCtx
}

func (e *Exchange) WsClient() *gatews.WsService {
	if e.wsClient == nil {
		wsUrl := gatews.FuturesUsdtUrl
		if strings.ToUpper(e.SettleCoin) == "USD" || strings.ToUpper(e.SettleCoin) == "BTC" {
			wsUrl = gatews.FuturesBtcUrl
		}
		var err error
		e.wsClient, err = gatews.NewWsService(nil, nil, gatews.NewConnConfFromOption(&gatews.ConfOptions{
			URL:           wsUrl,
			Key:           e.ApiKey,
			Secret:        e.ApiSecretKey,
			MaxRetryConn:  10, // default value is math.MaxInt64, set it when needs
			SkipTlsVerify: false,
		}))
		if err != nil {
			fmt.Printf("Init websocket client fail,error:%+v", err)
		}
	}
	return e.wsClient
}
