package xrandr

import (
	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/randr"
	"github.com/BurntSushi/xgb/xinerama"
	"github.com/BurntSushi/xgb/xproto"
	"log"
	"sort"
)

var (
	xgbConn *xgb.Conn
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

func GetXineramaConfiguration() []int {
	err := xinerama.Init(xgbConn)
	if err != nil {
		log.Fatal(err)
	}

	reply, err := xinerama.QueryScreens(xgbConn).Reply()
	if err != nil {
		log.Fatal(err)
	}

	type Monitor struct {
		Order int
		XOrg  int16
	}

	monitors := []Monitor{}

	for i, monitor := range reply.ScreenInfo {
		monitors = append(monitors, Monitor{i, monitor.XOrg})
	}

	sort.Slice(monitors, func(i, j int) bool { return monitors[i].XOrg < monitors[j].XOrg })

	ret := []int{}

	for _, monitor := range monitors {
		ret = append(ret, monitor.Order)
	}

	return ret
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

	for output, _ := range getOutputConfiguration() {
		ret = append(ret, output)
	}

	return ret
}

func ActiveOutputs() []string {
	ret := []string{}

	for output, status := range getOutputConfiguration() {
		if status {
			ret = append(ret, output)
		}
	}

	return ret
}
