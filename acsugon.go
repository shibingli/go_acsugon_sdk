package acsugon_sdk

import (
	"acsugon_sdk/api"
	ct "acsugon_sdk/constant"
	"acsugon_sdk/entity"
	"acsugon_sdk/errors"
	hu "acsugon_sdk/utils/http"
	"crypto/tls"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Proxy struct {
	URL string `json:"url"`
}

type Option struct {
	HttpClientTimeout *time.Duration `json:"http_client_timeout,omitempty"`

	Proxy *Proxy `json:"-"`
}

type ACSugon struct {
	Endpoint string                  `json:"endpoint"`
	User     string                  `json:"user"`
	Password string                  `json:"password"`
	OrgID    string                  `json:"orgId"`
	Tokens   map[string]entity.Token `json:"tokens,omitempty"`

	Client *http.Client `json:"-"`
}

func NewACSugon(endpoint, user, password, orgID string, opts ...*Option) (sg *ACSugon, err error) {
	endpoint = strings.TrimSpace(endpoint)
	endpoint = strings.TrimRight(endpoint, ct.HttpPathSeparator)

	user = strings.TrimSpace(user)
	password = strings.TrimSpace(password)

	orgID = strings.TrimSpace(orgID)

	if endpoint == "" || user == "" || password == "" || orgID == "" {
		return nil, errors.ErrInvalidParameter
	}

	sg = &ACSugon{
		Endpoint: endpoint,
		User:     user,
		Password: password,
		OrgID:    orgID,
	}

	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	var proxy *Proxy

	if len(opts) > 0 {
		opt := opts[0]

		if opt.HttpClientTimeout != nil {
			client.Timeout = *opt.HttpClientTimeout
		}

		if opt.Proxy != nil {
			proxy = opt.Proxy
		}
	}

	if proxy != nil {
		client, err = upgradeHttpClient(client, proxy)
		if err != nil {
			return nil, err
		}
	} else {
		client, err = upgradeHttpClient(client)
		if err != nil {
			return nil, err
		}
	}

	sg.Client = client

	return sg, nil
}

func upgradeHttpClient(client *http.Client, proxy ...*Proxy) (*http.Client, error) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	if len(proxy) > 0 {
		p := proxy[0]
		pUrl := strings.TrimSpace(p.URL)

		if pUrl != "" {
			u, err := url.Parse(pUrl)
			if err != nil {
				return nil, err
			}

			tr.Proxy = http.ProxyURL(u)
		}
	}

	client.Transport = tr

	return client, nil
}

func (a *ACSugon) Login() (sg *ACSugon, err error) {
	ep := a.Endpoint + "/" + api.Tokens.Path

	resp := entity.Response[[]entity.Token]{}

	_, err = hu.Request(
		a.Client, api.Tokens.Method, ep,
		map[string]string{"user": a.User, "password": a.Password, "orgId": a.OrgID},
		&resp,
	)
	if err != nil {
		return nil, err
	}

	if resp.Code != api.CodeOK.Key {
		return nil, api.ResultCode(resp.Code).Error()
	}

	tokens := make(map[string]entity.Token)

	for _, token := range resp.Data {
		tokens[token.ClusterId] = token
	}

	a.Tokens = tokens

	return a, nil
}
