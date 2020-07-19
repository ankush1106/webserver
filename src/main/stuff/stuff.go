package main

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Stuff struct {
// 	Name string `json:"name"`
// }

// func main() {
// 	s := Stuff{Name: "Ankush"}
// 	r := s.Print().PrintAgain()
// 	fmt.Println(r)
// 	t, _ := json.Marshal(s)
// 	fmt.Println(string(t))
// 	var p *interface{}
// 	json.Unmarshal(t, &p)
// 	fmt.Println(*p)

// }

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	Data    []Data `json:"data"`
	HasNext bool   `json:"has_next"`
}
type Data struct {
	Completed bool   `json:"completed"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	UserID    int    `json:"userId"`
}

func main() {
	baseURL := "http://localhost:3001/todos"
	c := make(chan Result)
	var wg sync.WaitGroup
	totalPages := 20
	for x := 1; x <= totalPages; x++ {
		wg.Add(1)
		url := fmt.Sprintf(baseURL+"?page=%d", x)
		go callURL(url, c, &wg)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	// this shorthand loop is syntactic sugar for an endless loop that just waits for results to come in through the 'c' channel
	for msg := range c {
		fmt.Println(msg)
	}
}

func callURL(url string, c chan Result, wg *sync.WaitGroup) {
	defer (*wg).Done()

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
		c <- Result{}
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
		c <- Result{}
	}

	result := Result{}
	jsonErr := json.Unmarshal(body, &result)
	if jsonErr != nil {
		log.Fatal(jsonErr)
		c <- Result{}
	}
	c <- result
}
