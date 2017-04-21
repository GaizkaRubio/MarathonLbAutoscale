package marathon

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"strings"
)

func GetApps() AutoGenerated {
	resp, err := http.Get("http://localhost:8080/v2/apps")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var apps AutoGenerated

	err = json.Unmarshal(body, &apps)

	if err != nil {
		panic(err)
	}

	return apps
}

func CreateApp(app MarathonApp) {
	b, err := json.Marshal(app)
	if err != nil {
		log.Fatal(err)
	}
	http.Post("http://localhost:8080/v2/apps","application/json",strings.NewReader(string(b)))
}

func SetAppInstances(app MarathonApp, instances int) {
	app.Instances = instances
	client := &http.Client{}
	//b, err := json.Marshal(app)
	//if err != nil {
	//	log.Fatal(err)
	//}

	request, err := http.NewRequest("PUT", "http://localhost:8080/v2/apps/"+app.ID, strings.NewReader("{\"instances\": 3}"))
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response)
}