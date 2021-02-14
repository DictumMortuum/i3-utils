package servus

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gen2brain/beeep"
	"github.com/urfave/cli"
	"io/ioutil"
	"net/http"
)

type RouterResponse struct {
	Error    string `json:"error"`
	Response struct {
		Uptime      int64 `json:"Uptime"`
		CurrentDown int64 `json:"CurrentDown"`
		InitialDown int64 `json:"InitialDown"`
	} `json:"Response"`
}

func GetRouter(c *cli.Context) error {
	var rs RouterResponse

	res, err := http.Get("https://servus.dictummortuum.com/router/latest")
	if err != nil {
		return err
	}

	raw, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw, &rs)
	if err != nil {
		return err
	}

	if rs.Error != "" {
		return errors.New(rs.Error)
	}

	msg := fmt.Sprintf("%d/%d", rs.Response.CurrentDown, rs.Response.InitialDown)
	err = beeep.Notify("Current Internet Speed", msg, "/usr/share/servus/globe.svg")
	if err != nil {
		return err
	}

	return nil
}
