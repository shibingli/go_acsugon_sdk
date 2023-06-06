package http

import (
	ct "acsugon_sdk/constant"
	aErr "acsugon_sdk/errors"
	"acsugon_sdk/utils/os"
	"bytes"
	"context"
	"errors"
	"fmt"
	gub "gitee.com/shibingli/goutils/bytes"
	"gitee.com/shibingli/goutils/json"
	"io"
	"net/http"
	"strings"
)

func NewRequest(method, reqUrl string, headers map[string]string, body ...interface{}) (request *http.Request, err error) {

	method = strings.TrimSpace(method)
	method = strings.ToUpper(method)

	reqUrl = strings.TrimSpace(reqUrl)

	if method == "" || reqUrl == "" {
		return nil, aErr.ErrInvalidParameter
	}

	request, err = http.NewRequestWithContext(context.Background(), method, reqUrl, nil)
	if err != nil {
		return
	}

	if len(body) > 0 && body[0] != nil {
		var b []byte

		b, err = json.Marshal(body[0])
		if err != nil {
			return
		}

		request, err = http.NewRequestWithContext(context.Background(), method, reqUrl, bytes.NewReader(b))
		if err != nil {
			return
		}
	}

	request.Header.Set("User-Agent", ct.DefaultUserAgent)
	request.Header.Set("Client-Type", "Web")
	request.Header.Set("Client-Version", "1.0")
	request.Header.Set("Accept-Language", "zh-CN,zh;q=0.8,zh;q=0.6")

	if http.MethodPost == strings.ToUpper(method) {
		request.Header.Set("Content-Type", "application/json")
	}

	mid, err := os.MachineID(ct.AppName)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Device-Id", mid)

	if headers != nil {
		for key, value := range headers {
			key = strings.TrimSpace(key)
			value = strings.TrimSpace(value)

			if key != "" && value != "" {
				request.Header.Set(key, value)
			}
		}
	}

	return
}

func Request(client *http.Client, method, url string, headers map[string]string, result interface{}, data ...interface{}) (response string, err error) {
	method = strings.TrimSpace(method)
	url = strings.TrimSpace(url)

	req, err := NewRequest(method, url, headers, data...)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer func() { _ = resp.Body.Close() }()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(b, result)
	if err != nil {
		return "", err
	}

	body := gub.UnsafeString(b)

	switch resp.StatusCode {
	case http.StatusNoContent:
		return "", nil
	case http.StatusOK, http.StatusCreated:
		return body, nil
	default:
		return "", errors.New(fmt.Sprintf("%d", resp.StatusCode))
	}
}
