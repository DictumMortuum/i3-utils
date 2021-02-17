package st

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli"
	"os"
	"strings"
)

func startedFromTmux() bool {
	return os.Getenv("TMUX") != ""
}

func empty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func getPrompt(buffer []string) string {
	pos := len(buffer) - 1
	if startedFromTmux() {
		pos = len(buffer) - 2
	}

	for i := pos; i != 0; i-- {
		if !empty(buffer[i]) {
			return buffer[i]
		}
	}

	return ""
}

func getCommands(buffer []string) []string {
	rs := []string{}
	prompt := getPrompt(buffer)

	for _, line := range buffer {
		if strings.HasPrefix(line, prompt) && line != prompt {
			rs = append(rs, line)
		}
	}

	return rs
}

func getCommandOutput(buffer []string, prompt string, cmd string) []string {
	rs := []string{}
	capture := false

	for _, line := range buffer {
		if strings.HasPrefix(line, prompt) {
			capture = false
		}
		if line == cmd {
			capture = true
		}
		if capture {
			rs = append(rs, line)
		}
	}

	return rs[1:]
}

func CheckGlob(c *cli.Context) error {
	buffer := []string{}
	fmt.Printf("started from tmux: %v\n", startedFromTmux())
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		buffer = append(buffer, scanner.Text())
	}

	for i, line := range getCommands(buffer) {
		fmt.Println(i, line)
		fmt.Println(getCommandOutput(buffer, getPrompt(buffer), line))
	}

	return nil
}
