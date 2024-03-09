package GoJson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data any) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestEncode(t *testing.T) {
	logJson("arifin")
	logJson(1)
	logJson(false)
	logJson([]string{"arifin", "azriel", "MJ"})

}
