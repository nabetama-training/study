package main

import (
	"encoding/json"
	"fmt"
)

// Servers original type Server list.
type Servers []map[hostName]Server // 変なJSONをよしなにラップする型を提議してあげる
type hostName string               // 可読性のため

// HogeHoge is API response from /api/pair_status.XXXXX
type HogeHoge struct {
	HOGE struct {
		Groups map[string]Group `json:"groups"`
	} `json:"hoge"`
}

// Group is information of ap server.
type Group struct {
	Port    int `json:port`
	Servers Servers
}

type Server struct {
	Connections int    `json:"connections"`
	RemovedBy   string `json:"removed_by"`
	Version     string `json:"version"`
	CPU         int    `json:"cpu"`
}

// JSONデータとしてmapしにくい、すなわち扱いにくいJSON
func main() {
	b := []byte(`
{
  "hoge": {
    "groups": {
      "81": {
        "port": 81,
        "servers": [
          {
            "some.host.invalid.key": {
              "connections": 999999,
              "removed_by": "",
              "version": "1702161129",
              "cpu": 99
            }
          }
        ]
      }
    }
  }
}
`)

	var h HogeHoge
	err := json.Unmarshal(b, &h)
	if err != nil {
		fmt.Errorf("Couldn't convert json: %v", err)
	}
	fmt.Println(connections("some.host.invalid.key", &h))
}

func connections(k hostName, h *HogeHoge) int {
	return h.HOGE.Groups["81"].Servers[0][k].Connections
}
