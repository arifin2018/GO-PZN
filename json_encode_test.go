package GoJson

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	osfile, _ := os.Create("sample_output.json")
	encoder := json.NewEncoder(osfile)

	sample := &Sample{
		Firsname: "nur",
		Lastname: "arifin",
		Score:    32,
	}

	encoder.Encode(sample)
}
