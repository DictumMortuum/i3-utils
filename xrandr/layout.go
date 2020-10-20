package xrandr

import (
	//	rofi "github.com/DictumMortuum/gofi"
	//	prmt "github.com/gitchander/permutation"
	//	"io"
	//	"os/exec"
	"regexp"
	//	"strings"
)

func isLaptopScreen(output string) bool {
	re := regexp.MustCompile(`eDP`)
	return re.MatchString(output)
}

/*
func Layout() {
	outputs := ActiveOutputs()

	selection := rofi.Plain("monitors", func(in io.WriteCloser) {
		permutations := prmt.New(prmt.StringSlice(outputs))

		for permutations.Next() {
			fmt.Fprintln(in, strings.Join(outputs, " "))
		}
	})

	generateXrandrFile(strings.Split(selection, " "))
}

func DynamicLayout() {
	outputs := ActiveOutputs()

	output := rofi.Plain("monitors", func(in io.WriteCloser) {
		for _, tmp := range outputs {
			fmt.Fprintln(in, tmp)
		}
	})

	if output == "" {
		os.Exit(1)
	}

	modes := [2]string{"off", "on"}

	mode := rofi.Plain("mode", func(in io.WriteCloser) {
		for _, tmp := range modes {
			fmt.Fprintln(in, tmp)
		}
	})

	if mode == "" {
		os.Exit(1)
	}

	if mode == "off" {
		exec.Command("xrandr", "--output", output, "--off").Run()
	} else {
		rightof := rofi.Plain("monitors", func(in io.WriteCloser) {
			for _, tmp := range outputs {
				fmt.Fprintln(in, tmp)
			}
		})
		exec.Command("xrandr", "--output", output, "--auto", "--right-of", rightof).Run()
	}
}
*/
