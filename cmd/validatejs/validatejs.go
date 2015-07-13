package main

import (
	"flag"
	"github.com/xeipuuv/gojsonschema"
	"io/ioutil"
	_ "jsonval"
	"log"
	"os"
)

var schema = flag.String("schema", "schema.json", "Name of the schema file.")
var json = flag.String("json", "json.json", "Name of the json file to check against the schema.")

// Moven Message schema.
var movenSchemaValidator *gojsonschema.Schema

// LoadSchema loads the schema from a string.
func LoadSchema(schema string) error {
	var err error
	schemaLoader := gojsonschema.NewStringLoader(schema)

	movenSchemaValidator, err = gojsonschema.NewSchema(schemaLoader)
	return err
}

// ValidateJson validates the Json string agains the movenSchemaValidator.
func ValidateJson(json string) (*gojsonschema.Result, error) {
	documentLoader := gojsonschema.NewStringLoader(json)
	result, err := movenSchemaValidator.Validate(documentLoader)
	return result, err
}

// Main function, way to large and ugly.
func main() {
	flag.Parse()

	// Perform the checks: schema must exist.
	fi, err := os.Stat(*schema)
	if err != nil {
		log.Fatalf("Cannot open schema file: %s", *schema)
	}
	if !fi.Mode().IsRegular() {
		log.Fatalf("Schema file is not a regular file: %s", *schema)
	}

	// Perform the checks: json file must exist.
	fi, err = os.Stat(*json)
	if err != nil {
		log.Fatalf("%s. Cannot find json file: %s.", err.Error(), *json)
	}
	if !fi.Mode().IsRegular() {
		log.Fatalf("%s. Json file is not a regular file: %s", err.Error(), *json)
	}

	// Get the schema as a string.
	schemaStr, err := ioutil.ReadFile(*schema)
	if err != nil {
		log.Fatalf("Failed to load the schema content: %s", err.Error())
	}

	// Get the json as a string.
	jsonStr, err := ioutil.ReadFile(*json)
	if err != nil {
		log.Fatalf("Failed to load the json content: %s", err.Error())
	}

	err = LoadSchema(string(schemaStr))
	if err != nil {
		log.Fatalf("Failed to load the schema validator: %s", err.Error())
	}

	res, err := ValidateJson(string(jsonStr))
	if err != nil {
		log.Fatalf("Error validating json: %s", err.Error())
	}

	if res.Valid() {
		log.Printf("The document is valid\n")
	} else {
		log.Printf("The document is not valid. see errors :\n")
		for _, desc := range res.Errors() {
			log.Printf("- %s\n", desc)
		}
	}
}
