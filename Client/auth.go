package main

import "context"

type Atus struct {
	Appkey    string
	AppSecret string
}

func (a *Atus) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"app_key": a.Appkey, "app_secret": a.AppSecret}, nil
}

func (a *Atus) RequireTransportSecurity() bool {
	return false
}
