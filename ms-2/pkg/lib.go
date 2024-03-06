package lib

import (
	"context"
	"crypto/tls"
	"io"
	"net/http"
)

func CallHttpGet(ctx context.Context, address string) ([]byte, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, err := http.NewRequestWithContext(ctx, "GET", address, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func CallHttpGetWeather(ctx context.Context, city string) ([]byte, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, err := http.NewRequestWithContext(ctx, "GET", "https://api.weatherapi.com/v1/current.json", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("key", "2345b10c20c34affa4c170312241602")
	q.Add("q", city)
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
