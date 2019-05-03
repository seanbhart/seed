// -----------------------------
// Copyright TangoJ Labs, LLC
// Apache 2.0 License
// -----------------------------

package thing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

type JSONThings struct {
	JSONThings []json.RawMessage `json:"things"`
}

func TestJsonThings(t *testing.T) {

	jsonBytes, _ := ioutil.ReadFile("../example/thing.json")

	var theseThings JSONThings
	err := json.Unmarshal(jsonBytes, &theseThings)
	if err != nil {
		log.Printf("THING json unmarshal ERROR: " + err.Error())
	}

	// Loop through the Things and print the address to check
	for t := 0; t < len(theseThings.JSONThings); t++ {
		fmt.Println(string(theseThings.JSONThings[t]))

		thing := NewThingFromJSON(0, theseThings.JSONThings[t])
		fmt.Println(thing.DataString())
	}
}
