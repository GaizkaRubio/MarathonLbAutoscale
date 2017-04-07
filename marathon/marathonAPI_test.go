package marathon

import (
	"testing"
	"fmt"
)

func TestGetApps(*testing.T){
	apps := GetApps()
	//fmt.Println(apps.Apps[0])
	fmt.Println(apps.Apps[0].ID)
 	SetAppInstances(apps.Apps[0], 3)
}
