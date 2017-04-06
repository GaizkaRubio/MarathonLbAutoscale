package marathon

import (
	"testing"
	"fmt"
)

func TestGetApps(*testing.T){
	apps := GetApps()
	fmt.Println(apps.Apps[0])
	SetAppInstances(apps.Apps[0], 3)
}
