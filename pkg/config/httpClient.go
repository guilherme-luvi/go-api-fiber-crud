package config

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type HttpClientService struct {
	Client *http.Client
}

func NewHttpClientService() *HttpClientService {
	return &HttpClientService{
		Client: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (s *HttpClientService) Get(url string) ([]byte, error) {
	resp, err := s.Client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer a solicitação GET: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler o corpo da resposta: %v", err)
	}

	return body, nil
}

func (s *HttpClientService) Post(url string, body io.Reader) ([]byte, error) {
	resp, err := s.Client.Post(url, "application/json", body)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer a solicitação POST: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler o corpo da resposta: %v", err)
	}

	return respBody, nil
}
