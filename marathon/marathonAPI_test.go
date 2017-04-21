package marathon

import (
	"testing"
	"fmt"
	"io/ioutil"
	"os"
	"encoding/json"
)

var maraApp MarathonApp

func TestGetApps(*testing.T){
	file, e := ioutil.ReadFile("resources/basicApp.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	error := json.Unmarshal(file,&maraApp)
	if error != nil {
		println(error)
	}

	fmt.Println(maraApp)

	CreateApp(maraApp)
	apps := GetApps()

	fmt.Println(apps.Apps[0].ID)

 	SetAppInstances(maraApp, 3)
}
