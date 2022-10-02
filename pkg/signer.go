package pkg

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Sign interface {
	getSignature(body string) (string, string)
	Get(method, point string) ([]byte, error)
	Post(method, point string, body map[string]string) ([]byte, error)
}

type Signer struct {
	Site      string `json:"site,omitempty"`
	apiKey    string
	apiSecret string
	client    *http.Client
}

func NewSigner() *Signer {
	return &Signer{
		Site:      os.Getenv("API_SITE"),
		apiKey:    os.Getenv("API_KEY"),
		apiSecret: os.Getenv("API_SECRET"),
		client:    http.DefaultClient}
}

func (a *Signer) getSignature(body string) (string, string) {
	t := time.Now().UnixMilli()
	ts := strconv.FormatInt(t, 10)
	s := ts + a.apiKey + "5000" + body

	h := hmac.New(sha256.New, []byte(a.apiSecret))
	h.Write([]byte(s))

	return fmt.Sprintf("%x", h.Sum(nil)), ts
}

func (a *Signer) Get(method, point string) ([]byte, error) {
	url := a.Site + point
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
	}

	sign, timestamp := a.getSignature("")

	request.Header.Add("X-BAPI-SIGN-TYPE", "2")
	request.Header.Add("X-BAPI-SIGN", sign)
	request.Header.Add("X-BAPI-API-KEY", a.apiKey)
	request.Header.Add("X-BAPI-TIMESTAMP", timestamp)
	request.Header.Add("X-BAPI-RECV-WINDOW", "5000")

	response, err := a.client.Do(request)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	readAll, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return readAll, err
}

func (a *Signer) Post(method, point string, body map[string]string) ([]byte, error) {
	marshal, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	url := a.Site + point
	request, err := http.NewRequest(method, url, bytes.NewBuffer(marshal))
	if err != nil {
		fmt.Println(err)
	}

	sign, timestamp := a.getSignature(string(marshal))

	request.Header.Add("X-BAPI-SIGN-TYPE", "2")
	request.Header.Add("X-BAPI-SIGN", sign)
	request.Header.Add("X-BAPI-API-KEY", a.apiKey)
	request.Header.Add("X-BAPI-TIMESTAMP", timestamp)
	request.Header.Add("X-BAPI-RECV-WINDOW", "5000")

	response, err := a.client.Do(request)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	readAll, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return readAll, err

}
