package main

import "testing"
import "encoding/json"

func TestConnections(t *testing.T) {
	b := []byte(`
{
  "hoge": {
    "groups": {
      "81": {
        "port": 81,
        "servers": [
          {
            "not.json.style.key": {
              "connections": 42,
              "removed_by": "",
              "version": "17",
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
	_ = json.Unmarshal(b, &h) // 無視
	var expected = 42
	res := connections("not.json.style.key", &h)
	if expected != res {
		t.Errorf("conncetions expected %d, but returns %d", expected, res)
	}

}
