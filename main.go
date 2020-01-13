package main

import (
	"fmt"
	"github.com/DictumMortuum/i3-utils/i3"
	"github.com/DictumMortuum/i3-utils/xrandr"
	"github.com/urfave/cli"
	"os"
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
	app.Version = "3.0.0"

	xrandr.Init()

	app.Commands = []cli.Command{
		{
			Name: "focus",
			Subcommands: []cli.Command{
				{
					Name: "next",
					Action: func(c *cli.Context) {
						i3.Next(i3.SetCurrentWorkspace)
					},
				},
				{
					Name: "prev",
					Action: func(c *cli.Context) {
						i3.Prev(i3.SetCurrentWorkspace)
					},
				},
			},
		},
		{
			Name: "move",
			Subcommands: []cli.Command{
				{
					Name: "next",
					Action: func(c *cli.Context) {
						i3.Next(i3.MoveContainer)
					},
				},
				{
					Name: "prev",
					Action: func(c *cli.Context) {
						i3.Prev(i3.MoveContainer)
					},
				},
			},
		},
		{
			Name:  "display",
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
