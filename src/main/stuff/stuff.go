package main

import (
	"encoding/json"
	"fmt"
)

type Stuff struct {
	Name string `json:"name"`
}

func (s *Stuff) Print() *Stuff {
	fmt.Println(s.Name)
	return s
}

func (s *Stuff) PrintAgain() string {
	fmt.Println("pringing again")
	return "sdfsdfsg"

}

func main() {
	s := Stuff{Name: "Ankush"}
	r := s.Print().PrintAgain()
	fmt.Println(r)
	t, _ := json.Marshal(s)
	fmt.Println(string(t))
	var p *interface{}
	json.Unmarshal(t, &p)
	fmt.Println(*p)

}
