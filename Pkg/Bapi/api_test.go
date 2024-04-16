package Bapi

import (
	"context"
	"fmt"
	"log"
	"testing"
)

func TestAPI_Gettaglist(t *testing.T) {
	url := "https://restapi.amap.com/v3/weather/"
	api := NewAPi(url)
	datas, err := api.Gettaglist(context.Background(), "weatherInfo")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(string(datas))

}
