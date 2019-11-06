package xrandr

import (
  "log"
  "github.com/BurntSushi/xgb"
  "github.com/BurntSushi/xgb/randr"
  "github.com/BurntSushi/xgb/xproto"
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

func Outputs() map[string]bool {
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
