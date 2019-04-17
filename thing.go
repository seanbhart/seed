package seed

import (
	"encoding/json"
	"fmt"
	"time"
)

// Thing is a standard data format
type Thing struct {
	Address     string `json:"address" binding:"required"`
	Version     string `json:"version" binding:"required"`
	ThingType   int    `json:"type" binding:"required"`
	Registry    int    `json:"registry"`
	Tradable    bool   `json:"tradable" binding:"required"`
	Spawner     bool   `json:"spawner" binding:"required"`
	Cert        string `json:"cert" binding:"required"`
	Proof       string `json:"proof"`
	Hash        []byte `json:"hash"`
	Data        []byte `json:"data"`
	FeaturesRaw []byte `json:"features"`
	UpdatedRaw  int    `json:"updated"` //Timestamp
	Features    []Feature
	Updated     time.Time
}

// NewThing initializes a new Thing object with the mandatory data
func NewThing(address string, version string, thingType int, tradable bool, spawner bool, cert string) *Thing {

	var thing = Thing{Address: address, Version: version, ThingType: thingType, Tradable: tradable, Spawner: spawner, Cert: cert}

	// If the json contained a map for Features, convert each to a Feature object and add to the Thing

	return &thing
}

// NewThingFromJSON initializes a new Thing from json
func NewThingFromJSON(jsonThing []byte) *Thing {

	var thing Thing
	err := json.Unmarshal(jsonThing, &thing)
	if err != nil {
		fmt.Printf("THING json unmarshal ERROR: " + err.Error())
	}

	// If the json contained a map for Features, convert each to a Feature object and add to the Thing
	if len(thing.FeaturesRaw) > 0 {
		var features []Feature
		for f := 1; f <= len(thing.FeaturesRaw); f++ {
			var feature = NewFeature(int(thing.FeaturesRaw[f]["order"]), thing.FeaturesRaw[f]["address"])
			if int(thing.FeaturesRaw[f]["type"])
			feature.ThingType
			features[f] = feature
		}
	}

	return &thing
}

// Feature is a reference to another Thing at a different Address
type Feature struct {
	Order     int    `json:"order" binding:"required"`
	Address   string `json:"address" binding:"required"`
	ThingType int    `json:"type"`
	Title     string `json:"title"`
	FormatRaw
	Format FeatureFormat
}

// NewFeature constructor for Feature defaults to type 0 (a Thing with nested Features)
// The Format object must be added manually
func NewFeature(order int, address string) *Feature {
	return &Feature{Order: order, Address: address, ThingType: 0}
}

// FeatureFormat contains the general format settings for a Thing
type FeatureFormat struct {
	Group     bool
	Title     bool
	Display   bool
	TextAlign string
	FontSize  float32
	Regex     map[int]FeatureFormatRegex
}

// NewFeatureFormat constructor for FeatureFormat defaults the text alignment and size
// Regex must be added manually
func NewFeatureFormat(group bool, title bool, display bool) *FeatureFormat {
	return &FeatureFormat{Group: group, Title: title, Display: display, TextAlign: "l", FontSize: 12.0}
}

// FeatureFormatRegex is regex settings for a Feature field to ensure display consistency accross apps
type FeatureFormatRegex struct {
	Order     int
	RegexType string
	Pattern   string
	Template  string
}

// NewFeatureFormatRegex constructor for FeatureFormatRegex
func NewFeatureFormatRegex(order int, regexType string, pattern string, template string) *FeatureFormatRegex {
	return &FeatureFormatRegex{Order: order, RegexType: regexType, Pattern: pattern, Template: template}
}
