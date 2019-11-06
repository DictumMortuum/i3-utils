package main

import (
  "fmt"
  "os"
  "log"
  "go.i3wm.org/i3"
  "github.com/urfave/cli"
  "github.com/BurntSushi/xgb"
  "github.com/BurntSushi/xgb/randr"
  "github.com/BurntSushi/xgb/xproto"
)

var (
  xgbConn *xgb.Conn
)

func xinit() {
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

func UpdateWorkspaces(workspace, display string) error {
  command := fmt.Sprintf("workspace %s; move workspace to %s", workspace, display)
  _, err := i3.RunCommand(command)

  if err != nil {
    return err
  }

  return nil
}

func main() {
  app := cli.NewApp()
  app.Name = "i3-util"
  app.Usage = "Utilities for the i3wm"
  app.Version = "1.0.0"

  xinit()

  app.Commands = []cli.Command{
    {
      Name: "move",
      Usage: "Move a workspace to an output",
      Action: func(c *cli.Context) error {
        if c.NArg() < 2 {
          fmt.Println("move subcommand requires at least two arguments")
          return nil
        }
        workspace := c.Args().Get(0)
        display := c.Args().Get(1)
        return UpdateWorkspaces(workspace, display)
      },
    },
    {
      Name: "display",
      Usage: "Show display outputs",
      Subcommands: []cli.Command{
        {
          Name: "all",
          Action: func(c *cli.Context) {
            for output, status := range getOutputConfiguration() {
              fmt.Printf("%-10s %v\n", output, status)
            }
          },
        },
        {
          Name: "active",
          Action: func(c *cli.Context) {
            for output, status := range getOutputConfiguration() {
              if status == true {
                fmt.Println(output)
              }
            }
          },
        },
        {
          Name: "inactive",
          Action: func(c *cli.Context) {
            for output, status := range getOutputConfiguration() {
              if status == false {
                fmt.Println(output)
              }
            }
          },
        },
      },
    },
  }

  app.Run(os.Args)
}
