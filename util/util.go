package util

import (
	"bytes"
	"errors"
	"log"
	"os"
	"os/exec"
	"strings"
)

func Bash(cmd string) string {
	c := exec.Command("bash", "-c", cmd)

	log.Print(cmd)

	var out bytes.Buffer
	c.Stdout = &out
	err := c.Run()
	if err != nil {
		log.Fatal(err)
	}

	return out.String()
}

func Spawn(args string) {
	ssh := strings.Split(args, " ")
	tmp := exec.Command(ssh[0], ssh[1:]...)

	tmp.Stdout = os.Stdout
	tmp.Stderr = os.Stderr
	tmp.Stdin = os.Stdin

	err := tmp.Start()
	if err != nil {
		log.Fatal(err)
	}

	err = tmp.Wait()
	if err != nil {
		log.Fatal(err)
	}
}

func Clip(output, sel string) error {
	path, err := exec.LookPath("xclip")
	if err != nil {
		return nil
	}

	cmd := exec.Command(path, "-selection", sel)
	in, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	_, err = in.Write([]byte(output))
	if err != nil {
		return err
	}

	err = in.Close()
	if err != nil {
		return err
	}

	err = cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}

func Type(in string) error {
	path, err := exec.LookPath("xdotool")
	if err != nil {
		return errors.New("Please install xdotool.")
	}

	args := []string{"type", "--delay", "3", "--clearmodifiers", in}

	err = exec.Command(path, args...).Run()
	if err != nil {
		return err
	}

	return nil
}
