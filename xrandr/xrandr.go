package xrandr

import (
	"fmt"
	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/randr"
	"github.com/BurntSushi/xgb/xproto"
	"log"
)

var (
	xgbConn                 *xgb.Conn
	lastOutputConfiguration map[string]bool
)

func Init() {
	var err error
	xgbConn, err = xgb.NewConn()
	if err != nil {
		log.Fatal(err)
	}

	err = randr.Init(xgbConn)
	if err != nil {
		log.Fatal(err)
	}
}

func getOutputConfiguration() map[string]bool {
	config := make(map[string]bool)

	root := xproto.Setup(xgbConn).DefaultScreen(xgbConn).Root
	resources, err := randr.GetScreenResources(xgbConn, root).Reply()

	if err != nil {
		log.Fatal(err)
	}

	for _, output := range resources.Outputs {
		info, err := randr.GetOutputInfo(xgbConn, output, 0).Reply()
		if err != nil {
			log.Fatal(err)
		}

		config[string(info.Name)] = info.Connection == randr.ConnectionConnected
	}

	return config
}

func AllOutputs() []string {
	ret := []string{}

	for output, status := range getOutputConfiguration() {
		fmt.Printf("%-10s %v\n", output, status)
		ret = append(ret, output)
	}

	return ret
}

func ActiveOutputs() []string {
	ret := []string{}

	for output, status := range getOutputConfiguration() {
		if status {
			fmt.Println(output)
			ret = append(ret, output)
		}
	}

	return ret
}
