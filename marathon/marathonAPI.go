package marathon

import (
	"net/http"
	"log"
	"fmt"
	"io/ioutil"
	"encoding/json"
)



func GetApps() {
	resp, err := http.Get("http://localhost:8080/v2/apps")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var apps MarathonApp

	err = json.Unmarshal(body, &apps)

	if err != nil {
		panic(err)
	}


	fmt.Println(apps)

}
