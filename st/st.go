package st

import (
	"bufio"
	"fmt"
	rofi "github.com/DictumMortuum/gofi"
	"github.com/DictumMortuum/i3-utils/util"
	"github.com/urfave/cli"
	"io"
	"os"
	"strings"
)

func empty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func isLocalPrompt(line string) bool {
	return strings.HasPrefix(line, "❯")
}

func isLocalCommand(line string) bool {
	return isLocalPrompt(line) && len(line) > 4
}

func isRemotePrompt(line string) bool {
	return strings.HasPrefix(line, "[") && strings.Contains(line, ">")
}

func isRemoteCommand(line string) bool {
	return isRemotePrompt(line) && !strings.HasSuffix(strings.TrimSpace(line), ">")
}

func getPrompts(buffer []string) []string {
	rs := []string{}

	for _, line := range buffer {
		if isLocalCommand(line) {
			rs = append(rs, line)
		}

		if isRemoteCommand(line) {
			rs = append(rs, line)
		}
	}

	return rs
}

func stripPrompt(line string) string {
	if isLocalCommand(line) {
		return line[4:]
	}

	if isRemoteCommand(line) {
		i := strings.Index(line, ">")
		return line[i:]
	}

	return ""
}

func getCommandOutput(buffer []string, cmd string) []string {
	rs := []string{}
	capture := false

	for _, line := range buffer {
		if isLocalPrompt(line) || isRemotePrompt(line) {
			capture = false
		}

		if strings.TrimSpace(line) == strings.TrimSpace(cmd) {
			capture = true
		}

		if capture {
			rs = append(rs, line)
		}
	}

	return rs[1:]
}

func options(buffer []string) func(io.WriteCloser) {
	return func(in io.WriteCloser) {
		for _, command := range buffer {
			fmt.Fprintln(in, command)
		}
	}
}

func GetCommandOutput(c *cli.Context) error {
	out := c.Bool("command_only")

	buffer := []string{}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		buffer = append(buffer, scanner.Text())
	}

	opts := rofi.GofiOptions{
		Description:  "commands",
		ForceDesktop: true,
	}

	err, commands := rofi.FromFilter(&opts, options(getPrompts(buffer)))
	if err != nil {
		return err
	}

	for _, command := range commands {
		var rs string

		if out {
			rs = stripPrompt(command)
		} else {
			rs = strings.Join(getCommandOutput(buffer, command), "\n")
		}

		util.Clip(rs, "c")
	}

	return nil
}
