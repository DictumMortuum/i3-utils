// function st-superuser() {
//   local user=$(echo -e "openbet\ninformix\nobdba\ncentos\nroot" | rofi-select)
//   [[ -z $user ]] && exit 0
//   type-text "sudo -iu ${user}"
// }
package st

import (
	rofi "github.com/DictumMortuum/gofi"
	"github.com/DictumMortuum/i3-utils/util"
	"github.com/urfave/cli"
	"strings"
)

func Type(c *cli.Context) error {
	buffer := []string{
		"sudo -iu openbet",
		"sudo -iu informix",
		"sudo -iu obdba",
		"sudo -iu centos",
		"sudo -iu root",
		"curl ifconfig.me",
	}

	opts := rofi.GofiOptions{
		Description:  "type",
		ForceDesktop: true,
	}

	err, commands := rofi.FromFilter(&opts, options(buffer))
	if err != nil {
		return err
	}

	return util.Type(strings.Join(commands, " ; "))
}
