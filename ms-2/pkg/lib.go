package lib

import (
	"context"
	"crypto/tls"
	"io"
	"net/http"
)

type AppError struct {
	ErrorMsg error
	Message  string
	Code     int
}

func (a *AppError) Error() string {
	return a.ErrorMsg.Error()
}

func NewAppError(errorMsg error, message string, code int) *AppError {
	return &AppError{
		ErrorMsg: errorMsg,
		Message:  message,
		Code:     code,
	}
}

type Response struct {
	Message interface{}
	Code    int
}

func CallHttpGet(ctx context.Context, address string, params map[string]string) ([]byte, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, address, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
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
