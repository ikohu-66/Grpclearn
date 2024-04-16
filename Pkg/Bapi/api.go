package Bapi

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type API struct {
	URL string
}

const (
	//API_KEY    = "grace"
	Acesstoken = "4ad08daeeea6b6d0e53d7e9c6497ac02"
)

type acessToken struct {
	Token string `json:"Token"`
}

func NewAPi(url string) *API {
	return &API{URL: url}
}

func (a *API) httpGet(ctx context.Context, path string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("%s%s", a.URL, path))
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil

}
func (a *API) Gettaglist(ctx context.Context, name string) ([]byte, error) {
	body, err := a.httpGet(ctx, fmt.Sprintf("%s?city=110101&key=%s", name, Acesstoken))
	if err != nil {
		return nil, err
	}
	return body, nil
}
