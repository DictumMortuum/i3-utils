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

func Clip(output, sel string) {
	var copyCmd *exec.Cmd

	path, err := exec.LookPath("xclip")
	if err != nil {
		log.Println("Please install xclip if you want to automatically copy " + sel + "to your clipboard")
		return
	}

	copyCmd = exec.Command(path, "-selection", sel)
	in, err := copyCmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	err = copyCmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	_, err = in.Write([]byte(output))
	if err != nil {
		log.Fatal(err)
	}

	err = in.Close()
	if err != nil {
		log.Fatal(err)
	}

	copyCmd.Wait()
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
