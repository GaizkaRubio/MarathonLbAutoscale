package marathon

import (
	"testing"
	"fmt"
	"io/ioutil"
	"os"
	"encoding/json"
)

var apps MarathonApp

func TestGetApps(*testing.T){
	file, e := ioutil.ReadFile("resources/basicApp.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	error := json.Unmarshal(file,&apps)
	if error != nil {
		println(error)
	}
	CreateApp(apps)
	apps := GetApps()

	fmt.Println(apps.Apps[0].ID)
 	SetAppInstances(apps.Apps[0], 3)
}
