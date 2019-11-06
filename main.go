package main

import (
  "fmt"
  "os"
  "github.com/urfave/cli"
  "i3-utils/xrandr"
  "i3-utils/i3"
)

func main() {
  app := cli.NewApp()
  app.Name = "i3-util"
  app.Usage = "Utilities for the i3wm"
  app.Version = "1.0.0"

  xrandr.Init()

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
        return i3.MoveWorkspace(workspace, display)
      },
    },
    {
      Name: "display",
      Usage: "Show display outputs",
      Subcommands: []cli.Command{
        {
          Name: "all",
          Action: func(c *cli.Context) {
            for output, status := range xrandr.Outputs() {
              fmt.Printf("%-10s %v\n", output, status)
            }
          },
        },
        {
          Name: "active",
          Action: func(c *cli.Context) {
            for output, status := range xrandr.Outputs() {
              if status == true {
                fmt.Println(output)
              }
            }
          },
        },
        {
          Name: "inactive",
          Action: func(c *cli.Context) {
            for output, status := range xrandr.Outputs() {
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
