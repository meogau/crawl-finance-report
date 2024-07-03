package service

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"strings"
)

type AuthService struct {
}

func (auth *AuthService) GetCookieAndToken() (string, string) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Println("Error creating cookie jar:", err)
	}
	client := &http.Client{
		Jar: jar,
	}
	url := "https://www.barchart.com/crypto/quotes/%5EBTCUSDT/performance"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	cookies := jar.Cookies(req.URL)
	var cookieHeader string
	var token string
	for _, cookie := range cookies {
		cookieHeader += cookie.Name + "=" + cookie.Value + "; "
		if cookie.Name == "XSRF-TOKEN" {
			token = strings.ReplaceAll(cookie.Value, "%3D", "=")
		}
	}
	return cookieHeader, token
}
