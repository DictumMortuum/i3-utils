package browser

import (
	rofi "github.com/DictumMortuum/gofi"
	"github.com/DictumMortuum/i3-utils/util"
	"github.com/urfave/cli"
	"os/exec"
	"regexp"
)

var (
	re = regexp.MustCompile(`([A-Z]+-[0-9]+)`)
)

func GetTickets(c *cli.Context) error {
	out1, err := util.GetClip("--clipboard")
	if err != nil {
		return err
	}

	out2, err := util.GetClip("--primary")
	if err != nil {
		return err
	}

	opts := rofi.GofiOptions{
		Description:  "tickets",
		ForceDesktop: true,
	}

	tickets, err := rofi.FromArray(&opts, re.FindAllString(out1+out2, -1))
	if err != nil {
		return err
	}

	path, err := exec.LookPath("xdg-open")
	if err != nil {
		return err
	}

	for _, ticket := range tickets {
		cmd := exec.Command(path, "https://jira.sgdigital.com/browse/"+ticket)
		err = cmd.Run()
		if err != nil {
			return err
		}
	}

	return nil
}
