package marathon

import (
	"net/http"
	"log"
	"io/ioutil"
	"strings"
)

type HtClient struct {
	baseurl string
}

func (cli *HtClient)GetCall (url string) ([]byte) {
	resp, err := http.Get(cli.baseurl+url)
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

func (cli *HtClient) PostCall (url string, json string)   {
	http.Post(cli.baseurl+url,"application/json",strings.NewReader(json))
	//if error != nil {
	//	log.Fatal(error)
	//}
}

func (cli *HtClient) PutCall (url string, json string)  {
	client := &http.Client{}

	request, err := http.NewRequest("PUT", cli.baseurl+url, strings.NewReader(json))
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
}

func (cli *HtClient) DeleteCall (url string)  {
	client := &http.Client{}

	request, err := http.NewRequest("DELETE", cli.baseurl+url, nil)
	if err != nil {
		log.Fatal(err)
	}
	client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
}

