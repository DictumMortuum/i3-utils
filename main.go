package main

import (
	"fmt"
	"github.com/DictumMortuum/i3-utils/i3"
	"github.com/DictumMortuum/i3-utils/xrandr"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "i3-util"
	app.Usage = "Utilities for the i3wm"
	app.Version = "7.0.1"

	xrandr.Init()

	app.Commands = []cli.Command{
		{
			Name: "focus",
			Subcommands: []cli.Command{
				{
					Name: "next",
					Action: func(c *cli.Context) {
						i3.Next(i3.SetCurrentWorkspace, i3.SameScreen)
					},
				},
				{
					Name: "prev",
					Action: func(c *cli.Context) {
						i3.Prev(i3.SetCurrentWorkspace, i3.SameScreen)
					},
				},
			},
		},
		{
			Name: "create",
			Subcommands: []cli.Command{
				{
					Name: "next",
					Action: func(c *cli.Context) {
						i3.Next(i3.SetCurrentWorkspace, i3.AnyScreen)
					},
				},
				{
					Name: "prev",
					Action: func(c *cli.Context) {
						i3.Prev(i3.SetCurrentWorkspace, i3.AnyScreen)
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
						i3.Next(i3.MoveContainer, i3.AnyScreen)
					},
				},
				{
					Name: "prev",
					Action: func(c *cli.Context) {
						i3.Prev(i3.MoveContainer, i3.AnyScreen)
					},
				},
			},
		},
		{
			Name: "display",
			Subcommands: []cli.Command{
				{
					Name: "all",
					Action: func(c *cli.Context) {
						tmp := xrandr.AllOutputs()

						for _, output := range tmp {
							fmt.Printf("%-10s\n", output)
						}
					},
				},
				{
					Name: "active",
					Action: func(c *cli.Context) {
						tmp := xrandr.ActiveOutputs()

						for _, output := range tmp {
							fmt.Println(output)
						}
					},
				},
			},
		},
		{
			Name: "layout",
			Subcommands: []cli.Command{
				{
					Name: "detect",
					Action: func(c *cli.Context) {
						tmp := xrandr.Detect()
						fmt.Println(tmp)
					},
				},
				{
					Name: "change",
					Action: func(c *cli.Context) {
						xrandr.Layout()
					},
				},
				{
					Name: "conky",
					Action: func(c *cli.Context) {
						tmp := xrandr.GetXineramaConfiguration()
						fmt.Println(tmp[len(tmp)-1])
					},
				},
			},
		},
	}

	app.Run(os.Args)
}
