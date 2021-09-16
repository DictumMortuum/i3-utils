package util

import (
	"bytes"
	"errors"
	"log"
	"os"
	"os/exec"
	"strings"
)

func Bash(args []string) (string, error) {
	var out bytes.Buffer

	path, err := exec.LookPath("bash")
	if err != nil {
		return "", err
	}

	cmd := exec.Command(path, append([]string{"-c"}, args...)...)
	log.Print(args)
	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil {
		return "", err
	}

	return out.String(), nil
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

func GetClip(sel string) (string, error) {
	path, err := exec.LookPath("xsel")
	if err != nil {
		return "", nil
	}

	cmd := exec.Command(path, "-o", sel)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	result := string(out)
	return result, nil
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
