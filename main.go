package main

import (
  "fmt"
  "os"
  "github.com/urfave/cli"
  "github.com/DictumMortuum/i3-utils/xrandr"
  "github.com/DictumMortuum/i3-utils/i3"
)

func all(output string, status bool) {
  fmt.Printf("%-10s %v\n", output, status)
}

func active(output string, status bool) {
  if status == true {
    fmt.Println(output)
  }
}

func inactive(output string, status bool) {
  if status == false {
    fmt.Println(output)
  }
}

func main() {
  app := cli.NewApp()
  app.Name = "i3-util"
  app.Usage = "Utilities for the i3wm"
  app.Version = "2.0.0"

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
      Name: "next",
      Action: func(c *cli.Context) {
				i3.Next()
      },
		},
    {
      Name: "prev",
      Action: func(c *cli.Context) {
				i3.Prev()
      },
    },
    {
      Name: "display",
      Usage: "Show display outputs",
      Subcommands: []cli.Command{
        {
          Name: "all",
          Action: func(c *cli.Context) {
            xrandr.Outputs(all)
          },
        },
        {
          Name: "active",
          Action: func(c *cli.Context) {
            xrandr.Outputs(active)
          },
        },
        {
          Name: "inactive",
          Action: func(c *cli.Context) {
            xrandr.Outputs(inactive)
          },
        },
      },
    },
  }

  app.Run(os.Args)
}
