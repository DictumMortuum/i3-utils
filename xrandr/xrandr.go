package xrandr

// Mostly stolen from https://github.com/BurntSushi/gohead/blob/master/main.go

import (
	"fmt"
	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/randr"
	"github.com/BurntSushi/xgb/xproto"
	rofi "github.com/DictumMortuum/gofi"
	prmt "github.com/gitchander/permutation"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strings"
)

type head struct {
	id                  randr.Output
	output              string
	x, y, width, height int
}

type heads struct {
	primary      *head
	heads        []head
	off          []string
	disconnected []string
}

func Heads() heads {
	X, err := xgb.NewConn()
	if err != nil {
		log.Fatal(err)
	}

	err = randr.Init(X)
	if err != nil {
		log.Fatal(err)
	}

	return newHeads(X)
}

func newHeads(X *xgb.Conn) heads {
	var primaryHead *head
	var primaryOutput randr.Output

	root := xproto.Setup(X).DefaultScreen(X).Root
	resources, err := randr.GetScreenResourcesCurrent(X, root).Reply()
	if err != nil {
		log.Fatalf("Could not get screen resources: %s.", err)
	}

	primaryOutputReply, _ := randr.GetOutputPrimary(X, root).Reply()
	if primaryOutputReply != nil {
		primaryOutput = primaryOutputReply.Output
	}

	hds := make([]head, 0, len(resources.Outputs))
	off := make([]string, 0)
	disconnected := make([]string, 0)
	for i, output := range resources.Outputs {
		oinfo, err := randr.GetOutputInfo(X, output, 0).Reply()
		if err != nil {
			log.Fatalf("Could not get output info for screen %d: %s.", i, err)
		}
		outputName := string(oinfo.Name)

		if oinfo.Connection != randr.ConnectionConnected {
			disconnected = append(disconnected, outputName)
			continue
		}
		if oinfo.Crtc == 0 {
			off = append(off, outputName)
			continue
		}

		crtcinfo, err := randr.GetCrtcInfo(X, oinfo.Crtc, 0).Reply()
		if err != nil {
			log.Fatalf("Could not get crtc info for screen (%d, %s): %s.",
				i, outputName, err)
		}

		head := newHead(output, outputName, crtcinfo)
		if output == primaryOutput {
			primaryHead = &head
		}
		hds = append(hds, head)
	}
	if primaryHead == nil && len(hds) > 0 {
		tmp := hds[0]
		primaryHead = &tmp
	}

	hdsPrim := heads{
		primary:      primaryHead,
		heads:        hds,
		off:          off,
		disconnected: disconnected,
	}
	sort.Sort(hdsPrim)
	return hdsPrim
}

func (hs heads) Dock() {
	for _, head := range hs.heads {
		if isLaptopScreen(head.output) {
			exec.Command("xrandr", "--output", head.output, "--off").Run()
		}
	}
}

func (hs heads) Restore(interactive bool) {
	outputs := []string{}
	args := []string{}

	for _, head := range hs.heads {
		outputs = append(outputs, head.output)
	}

	for _, head := range hs.off {
		outputs = append(outputs, head)
	}

	if interactive {
		var err error

		opts := rofi.GofiOptions{
			Description: "monitors",
		}

		err, outputs = rofi.FromFilter(&opts, func(in io.WriteCloser) {
			permutations := prmt.New(prmt.StringSlice(outputs))

			for permutations.Next() {
				fmt.Fprintln(in, strings.Join(outputs, " "))
			}
		})
		if err != nil {
			log.Fatal(err)
		}

		if len(outputs) == 0 {
			os.Exit(1)
		}
	}

	for i, head := range outputs {
		if i == 0 {
			args = append(args, "--output", head, "--auto", "--primary")
		} else {
			args = append(args, "--output", head, "--auto", "--right-of", outputs[i-1])
		}
	}

	exec.Command("xrandr", args...).Run()
}

func (hs heads) Active(mode []string) {
	args := []string{}

	for i, head := range hs.heads {
		if i == 0 {
			args = append(args, "--output", head.output, "--primary")
			args = append(args, mode...)
		} else {
			args = append(args, "--output", head.output, "--auto", "--right-of", hs.heads[i-1].output)
		}
	}

	exec.Command("xrandr", args...).Run()
}

func (hs heads) String() string {
	log.Print(hs.disconnected)
	log.Print(hs.off)

	lines := make([]string, len(hs.heads))
	for i, head := range hs.heads {
		lines[i] = fmt.Sprintf("%d: %s (%d, %d) %dx%d",
			i, head.output, head.x, head.y, head.width, head.height)
	}
	return strings.Join(lines, "\n")
}

func (hs heads) Len() int {
	return len(hs.heads)
}

func (hs heads) Less(i, j int) bool {
	h1, h2 := hs.heads[i], hs.heads[j]
	return h1.x < h2.x || (h1.x == h2.x && h1.y < h2.y)
}

func (hs heads) Swap(i, j int) {
	hs.heads[i], hs.heads[j] = hs.heads[j], hs.heads[i]
}

func newHead(id randr.Output, name string, info *randr.GetCrtcInfoReply) head {
	return head{
		id:     id,
		output: name,
		x:      int(info.X),
		y:      int(info.Y),
		width:  int(info.Width),
		height: int(info.Height),
	}
}

func isLaptopScreen(output string) bool {
	re := regexp.MustCompile(`eDP`)
	return re.MatchString(output)
}
