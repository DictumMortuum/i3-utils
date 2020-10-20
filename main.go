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
	app.Version = "8.0.0"

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
						i3.Next(i3.SetCurrentWorkspace, i3.EmptyWorkspace)
					},
				},
				{
					Name: "prev",
					Action: func(c *cli.Context) {
						i3.Prev(i3.SetCurrentWorkspace, i3.EmptyWorkspace)
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
			Name: "layout",
			Subcommands: []cli.Command{
				{
					Name: "dock",
					Action: func(c *cli.Context) {
						xrandr.Heads().Dock()
					},
				},
				{
					Name: "restore",
					Action: func(c *cli.Context) {
						if c.Bool("full") {
							xrandr.Heads().Restore()
						} else {
							xrandr.Heads().Active([]string{"--auto"})
						}
					},
					Flags: []cli.Flag{
						cli.BoolFlag{
							Name: "full",
						},
					},
				},
				{
					Name: "diablo",
					Action: func(c *cli.Context) {
						xrandr.Heads().Active([]string{"--mode", "800x600"})
					},
				},
				{
					Name: "list",
					Action: func(c *cli.Context) {
						fmt.Print(xrandr.Heads())
					},
				},
			},
		},
	}

	app.Run(os.Args)
}
