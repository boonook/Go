package user

import (
	"encoding/json"
)

type Transport struct {
	Time  string
	MAC   string
	Id    string
	Rssid string
}

func returnUser() string {
	var st []Transport
	t1 := Transport{Time: "22", MAC: "33", Id: "44", Rssid: "55"}
	st = append(st, t1)
	t2 := Transport{Time: "66", MAC: "77", Id: "88", Rssid: "99"}
	st = append(st, t2)
	buf, _ := json.Marshal(st)
	return string(buf)
}
