package xrandr

import (
  "log"
  "reflect"
  "github.com/BurntSushi/xgb"
  "github.com/BurntSushi/xgb/randr"
  "github.com/BurntSushi/xgb/xproto"
  "github.com/DictumMortuum/i3-utils/i3"
)

var (
  xgbConn *xgb.Conn
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

func Refresh() {
  currentOutputConfiguration := getOutputConfiguration()

  if reflect.DeepEqual(currentOutputConfiguration, lastOutputConfiguration) {
    return
  }

  currentWorkspace, err := i3.GetCurrentWorkspaceNumber()
  if err != nil {
    log.Fatal(err)
  }

  err = i3.SetCurrentWorkspace(currentWorkspace)
  if err != nil {
    log.Fatal(err)
  }

  lastOutputConfiguration = currentOutputConfiguration
}

func ListenEvents() {
  defer xgbConn.Close()

  root := xproto.Setup(xgbConn).DefaultScreen(xgbConn).Root
  err := randr.SelectInputChecked(xgbConn, root,
    randr.NotifyMaskScreenChange|randr.NotifyMaskCrtcChange|randr.NotifyMaskOutputChange).Check()

  if err != nil {
    log.Fatal(err)
  }

  for {
    ev, err := xgbConn.WaitForEvent()
    if err != nil {
      log.Fatal(err)
    }

    switch ev.(type) {
    case randr.ScreenChangeNotifyEvent:
      Refresh()
    }
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

func Outputs(fn func(string, bool)) {
  for output, status := range getOutputConfiguration() {
    fn(output, status)
  }
}
