package main

import (
	"flag"
	"time"
	"net/http"
	"log"
	"encoding/csv"
)

func main() {
	//marathonUrl := flag.String("marathon-url", "localhost:8080", "url where marathon API is exposed")
	haproxyUrl := flag.String("haproxy-url", "gosec-int06.stratio.com:9090", "url where haproxy API is exposed")

	flag.Parse()

	ticker := time.NewTicker(time.Second * 2)

	for range ticker.C {
		resp, err := http.Get("http://" + *haproxyUrl + "/haproxy?stats;csv")
		if err != nil {
			log.Fatal(err)
		}

		r := csv.NewReader(resp.Body)

		records, err := r.ReadAll()
		if err != nil {
			log.Fatal(err)
		}

		tmp := make(map[string]string)

		for i := range records {
			tmp[i] = records[i]
		}
	}


}

