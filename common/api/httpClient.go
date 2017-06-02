package api

import (
	"net/http"
	"log"
	"io/ioutil"
	"strings"
)

type HttpClient struct {
	BaseUrl string
}

func (cli *HttpClient)GetCall (url string) ([]byte) {
	resp, err := http.Get(cli.BaseUrl+url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}

func (cli *HttpClient) PostCall (url string, json string)   {
	_ , err := http.Post(cli.BaseUrl+url,"application/json",strings.NewReader(json))
	if err != nil {
		log.Fatal(err)
	}

}

func (cli *HttpClient) PutCall (url string, json string)  {
	client := &http.Client{}

	request, err := http.NewRequest("PUT", cli.BaseUrl+url, strings.NewReader(json))
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	_ , err = client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
}

func (cli *HttpClient) DeleteCall (url string)  {
	client := &http.Client{}

	request, err := http.NewRequest("DELETE", cli.BaseUrl+url, nil)
	if err != nil {
		log.Fatal(err)
	}
	_ , err = client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
}

