package main

import (
	"encoding/json"
	"fmt"
)

type Player struct {
	Username string `json: "username"`
	Level    uint   `json: "level"`
	Status   string `json: "status"`
	Class    string `json: "class"`
}

// Username, Level is field
// `json: "class"` is bult-in json

func (p Player) GetUsername() string {
	return p.Username
}

// ต้องใช้ pointer เพื่อทำให้ค่าเปลี่ยน
func (p *Player) LevelUp() {
	p.Level++
}

func main() {

	player1 := Player{
		Username: "player1",
		Level:    1,
		Status:   "active",
		Class:    "Warrior",
	}

	fmt.Println(player1)

	jsonData, _ := json.Marshal(&player1)
	fmt.Println(string(jsonData))

	// Marshal ควร pass by reference
	jsonData2, _ := json.MarshalIndent(&player1, "", "\t")
	fmt.Println(string(jsonData2))

	fmt.Println(player1.GetUsername())

	player1.LevelUp()
	fmt.Println(player1.Level)
}
