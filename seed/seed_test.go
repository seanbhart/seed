// -----------------------------
// Copyright TangoJ Labs, LLC
// Apache 2.0 License
// -----------------------------

package seed

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/seanbhart/seed/thing"
)

// TestSeed creates a Seed Thing manually and tests requesting all nested Things
func TestSeed(t *testing.T) {
	// Get the json for the Seed Thing
	jsonBytes, _ := ioutil.ReadFile("../example/seed.json")

	// Create the Seed Thing
	seedThing := thing.NewThingFromJSON(0, jsonBytes)

	// Create the Seed
	seed := NewSeed(seedThing)

	PrintBranchContent(&seed.Tree)
}

func PrintBranchContent(branch *Branch) {

	log.Printf(branch.Thing.GetAddress())

	// Print the child branch content
	for _, b := range branch.Branches {
		PrintBranchContent(&b)
	}
}

func TestRetrieve(t *testing.T) {

	// Get the json for the Seed Thing
	jsonBytes, _ := ioutil.ReadFile("../example/seed.json")

	// Create the Seed Thing
	seedThing := thing.NewThingFromJSON(0, jsonBytes)

	// Create the Seed
	seed := NewSeed(seedThing)

	log.Printf(seed.Thing.GetAddress())
}
