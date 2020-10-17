package firefox

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/frioux/leatherman/pkg/mozlz4"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"os/user"
	"sort"
	"time"
)

type FirefoxRecovery struct {
	Windows []FirefoxWindow `json:"windows"`
}

type FirefoxWindow struct {
	Tabs     []FirefoxTab `json:"tabs"`
	Selected int          `json:"selected"`
}

type FirefoxTab struct {
	Entries      []FirefoxEntry `json:"entries"`
	LastAccessed int            `json:"lastAccessed"`
}

type FirefoxEntry struct {
	Url   string `json:"url"`
	Title string `json:"title"`
}

func CurrentUrl() {
	path, _ := readInstall()
	recovery := path + "/sessionstore-backups/recovery.jsonlz4"

	fileinfo, _ := os.Stat(recovery)
	modtime := fileinfo.ModTime().UnixNano() / int64(time.Millisecond)
	now := time.Now().UnixNano() / int64(time.Millisecond)
	threshold := 15000 - now + modtime

	if threshold > 0 {
		time.Sleep(time.Duration(threshold) * time.Millisecond)
	} else {
		time.Sleep(2000 * time.Millisecond)
	}

	file, err := os.Open(recovery)
	if err != nil {
		log.Fatal(err)
	}

	r, err := mozlz4.NewReader(file)
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	var temp FirefoxRecovery
	json.Unmarshal(buf.Bytes(), &temp)
	tabs := temp.Windows[0].Tabs

	sort.Slice(tabs, func(i, j int) bool {
		return tabs[i].LastAccessed > tabs[j].LastAccessed
	})

	l := len(tabs[0].Entries)
	log.Print(now - modtime)
	fmt.Println(tabs[0].Entries[l-1].Url)
}

func readInstall() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}

	path := u.HomeDir + "/.mozilla/firefox/installs.ini"

	cfg, err := ini.Load(path)
	if err != nil {
		return "", err
	}

	for _, section := range cfg.Sections() {
		if section.HasKey("Default") {
			key, _ := section.GetKey("Default")
			return u.HomeDir + "/.mozilla/firefox/" + key.String(), nil
		}
	}

	return "", errors.New("No default profile found in installs.ini")
}
