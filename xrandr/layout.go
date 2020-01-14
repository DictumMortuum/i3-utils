package xrandr

import (
	"bytes"
	"fmt"
	rofi "github.com/DictumMortuum/gofi"
	prmt "github.com/gitchander/permutation"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func getXrandrFilePath(i int) string {
	home := os.Getenv("HOME")
	return fmt.Sprintf("%s/.cache/screenlayout/xrandr.%d", home, i)
}

func generateXrandrConfig(outputs []string) string {
	buf := bytes.NewBufferString("#!/bin/bash")
	fmt.Fprintf(buf, "xrandr --setprovideroutputsource modesetting NVIDIA-0\n")

	for i, m := range outputs {
		if i == 0 {
			fmt.Fprintf(buf, "xrandr --output %s --auto\n", m)
		} else {
			fmt.Fprintf(buf, "xrandr --output %s --auto --right-of %s\n", m, outputs[i-1])
		}
	}

	fmt.Fprintf(buf, "xrandr --output %s --primary\n", outputs[len(outputs)-1])

	return buf.String()
}

func generateXrandrFile(outputs []string) {
	data := generateXrandrConfig(outputs)
	buffer := []byte(data)
	self := getXrandrFilePath(len(outputs))
	ioutil.WriteFile(self, buffer, 0700)
}

func Detect() string {
	outputs := ActiveOutputs()
	return getXrandrFilePath(len(outputs))
}

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
