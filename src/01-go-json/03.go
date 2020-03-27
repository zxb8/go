package main

import (
"encoding/json"
	"fmt"
	"os"
)

type Member struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	enc := json.NewEncoder(os.Stdout)

	letters := []string{"a", "b", "c", "d", "e", "f"}

	for i, letter := range letters {
		key := Member{letter, i}

		var err = enc.Encode(&key)
		if err != nil {
			fmt.Println(err)
		}
	}
}
