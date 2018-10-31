package tmeapi

import (
	"strings"
	"net/url"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/http"
)

type TMEApi struct {
    Token string
    Secret string
}

func Client(token string, secret string) *TMEApi {
		return &TMEApi{token, secret}
}

func (tmeApi TMEApi) Request(requestUrl string, formValues url.Values) (*http.Response, error) {
    formValues.Add("Token", tmeApi.Token)
    formValues.Add("ApiSignature", CalculateSignature("POST", requestUrl, formValues, tmeApi.Secret))

    return http.PostForm(requestUrl, formValues)
}

func CalculateSignature(method string, requestUrl string, urlValues url.Values, secret string) string {
    signatureBase := strings.Join([]string{method, url.QueryEscape(requestUrl), url.QueryEscape(urlValues.Encode())}, "&")
    mac := hmac.New(sha1.New, []byte(secret))
    mac.Write([]byte(signatureBase))

    return base64.StdEncoding.EncodeToString([]byte(mac.Sum(nil)))
}
