package logoai

import (
	"bytes"
	"context"
	"encoding/json"
)

type LogoServiceOp struct {
	client *Client
}

type LogoRequest struct {
	Name    string  `json:"name"`
	Slogan  *string `json:"slogan"`
	Page    *string `json:"page"`
	ColorID *string `json:"colorId"`
	FontID  *string `json:"fontId"`
	StyleID *string `json:"styleId"`
	Size    *string `json:"size"`
	Format  *string `json:"format"`
}

type LogoResponse struct {
	Data    interface{} `json:"data"`
	ErrCode int         `json:"errCode"`
	ErrMsg  string      `json:"errMsg"`
	Success string      `json:"success"`
}

type LogosService interface {
	LogosGeneration(context.Context, LogoRequest) (*LogoResponse, error)
}

func (s *LogoServiceOp) LogosGeneration(ctx context.Context, req LogoRequest) (*LogoResponse, error) {

	var reqResponse []byte

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	errRequest := s.client.Request("GET", "https://api.logoai.com/logoapi/async/logo", bytes.NewBuffer(reqBody), &reqResponse)
	if errRequest != nil {
		return nil, errRequest
	}

	var response LogoResponse
	err = json.Unmarshal(reqResponse, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
