// -----------------------------
// Copyright TangoJ Labs, LLC
// Apache 2.0 License
// -----------------------------

package thing

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

// Thing is a standard data format
type Thing interface {
	GetAddress() string
	GetType() int
	DataString() string
}

// Common contains all the shared components of a Thing
type Common struct {
	Address    string `json:"address" binding:"required"`
	Version    string `json:"version" binding:"required"`
	ThingType  int    `json:"type" binding:"required"`
	Registry   string `json:"registry"`
	Tradable   bool   `json:"tradable" binding:"required"`
	Spawner    bool   `json:"spawner" binding:"required"`
	Cert       []byte `json:"cert" binding:"required"`
	Proof      string `json:"proof"`
	Hash       []byte `json:"hash"`
	UpdatedRaw int    `json:"updated"` //Timestamp
	Updated    time.Time
}

// DefaultThing is the default version of a Thing (contains Features - references to nested Things)
type DefaultThing struct {
	*Common
	Features []Feature `json:"features"`
}

// GetAddress returns the address
func (dt *DefaultThing) GetAddress() string {
	return dt.Common.Address
}

// GetType returns the ThingType as int
func (dt *DefaultThing) GetType() int {
	return dt.Common.ThingType
}

// DataString returns the data as a string
func (dt *DefaultThing) DataString() string {
	return fmt.Sprintln("DATA IS ALWAYS NIL FOR DEFAULT THINGS")
}

// GetFeatures returns all Features
func (dt *DefaultThing) GetFeatures() []Feature {
	return dt.Features
}

// NumberThing is the number version of the Thing standard
type NumberThing struct {
	*Common
	Data float64 `json:"data"`
}

// GetAddress returns the address
func (nt *NumberThing) GetAddress() string {
	return nt.Common.Address
}

// GetType returns the ThingType as int
func (nt *NumberThing) GetType() int {
	return nt.Common.ThingType
}

// DataString returns the floating point number as a string
func (nt *NumberThing) DataString() string {
	return fmt.Sprintf("%f\n", nt.Data)
}

// StringThing is the string version of the Thing standard
type StringThing struct {
	*Common
	Data string `json:"data"`
}

// GetAddress returns the address
func (st *StringThing) GetAddress() string {
	return st.Common.Address
}

// GetType returns the ThingType as int
func (st *StringThing) GetType() int {
	return st.Common.ThingType
}

// DataString returns string
func (st *StringThing) DataString() string {
	return fmt.Sprintf("%s\n", st.Data)
}

// ImageThing is the image version of the Thing standard (image byte array)
type ImageThing struct {
	*Common
	compression string
	Data        []byte `json:"data"`
}

// DataString returns the image byte array as a string for printing
func (it *ImageThing) DataString() string {
	return fmt.Sprintf("%s\n", string(it.Data))
}

// NewThing initializes a new Thing object with the mandatory data
func NewThing(address string, version string, thingType int, tradable bool, spawner bool, cert []byte) *Common {

	var thing = Common{Address: address, Version: version, ThingType: thingType, Tradable: tradable, Spawner: spawner, Cert: cert}

	// If the json contained a map for Features, convert each to a Feature object and add to the Thing

	return &thing
}

// NewThingFromJSON initializes a new Thing from json
func NewThingFromJSON(thingType int, jsonThing []byte) Thing {

	// Check the type of Thing and create as necessary
	switch thingType {
	case 1:
		// string type
		fmt.Println("TYPE 1")
		thing := &StringThing{}
		err := json.Unmarshal(jsonThing, &thing)
		if err != nil {
			log.Printf("THING TYPE 1 json unmarshal ERROR: " + err.Error())
		}
		return thing
	case 2:
		// number type
		fmt.Println("TYPE 2")
		thing := &NumberThing{}
		err := json.Unmarshal(jsonThing, &thing)
		if err != nil {
			log.Printf("THING TYPE 2 json unmarshal ERROR: " + err.Error())
		}
		return thing
	default:
		// default should be 0 (nested Things)
		fmt.Println("TYPE 0")
		thing := &DefaultThing{}
		err := json.Unmarshal(jsonThing, &thing)
		if err != nil {
			log.Printf("THING TYPE 0 json unmarshal ERROR: " + err.Error())
		}
		return thing
	}
}

// Feature is a reference to another Thing at a different Address
type Feature struct {
	Order     int    `json:"order" binding:"required"`
	Address   string `json:"address" binding:"required"`
	ThingType int    `json:"type"`
	Title     string `json:"title"`
}

// NewFeature constructor for Feature defaults to type 0 (a Thing with nested Features)
// The Format object must be added manually
func NewFeature(order int, address string) *Feature {
	return &Feature{Order: order, Address: address, ThingType: 0}
}

// Retrieve will use a Feature to recall the referenced Thing
func (f *Feature) Retrieve() (*Thing, error) {
	// Connect to the address and get a json response with the the Thing
	jsonBytes, err := ioutil.ReadFile("../example/" + f.Address + ".json")
	if err != nil {
		return nil, errors.New("ERROR: package thing: Feature Retrieve - unable to access file")
	}

	// Convert JSON to Thing
	var thing = NewThingFromJSON(f.ThingType, jsonBytes)
	return &thing, nil
}

func (f *Feature) RetrieveHTTP() {
	client := http.Client{
		Timeout: time.Second * 10, // Maximum of 10 secs
	}

	// urlString := "http://localhost:3000/api/test"
	urlString := "http://localhost:3000/api/thing"

	log.Printf("THING ADDRESS: %s", f.Address)

	data := url.Values{}
	data.Set("address", f.Address)

	req, err := http.NewRequest(http.MethodPost, urlString, bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Printf("ERROR: package thing: RetrieveHTTP - NewRequest failure")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") // This makes it work

	res, err := client.Do(req)
	if err != nil {
		log.Printf("ERROR: package thing: RetrieveHTTP - Do failure")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("ERROR: package thing: RetrieveHTTP - ReadAll failure")
	}
	res.Body.Close()
	if err != nil {
		log.Printf("ERROR: package thing: RetrieveHTTP - Body Close failure")
	}
	log.Printf("RetrieveHTTP: body: %s", string(body))

}
