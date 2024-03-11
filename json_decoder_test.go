package GoJson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type Sample struct {
	Firsname string
	Lastname string
	Score    int
}

func TestDecoder(t *testing.T) {
	osFile, _ := os.Open("sample.json")
	decoder := json.NewDecoder(osFile)
	sample := map[string]any{}
	decoder.Decode(&sample)
	fmt.Println("sample")
	fmt.Println(sample)

}
