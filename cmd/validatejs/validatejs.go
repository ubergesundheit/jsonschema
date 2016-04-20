package main

import (
	"flag"
	"log"
	"os"

	"github.com/xeipuuv/gojsonschema"
)

var schema = flag.String("schema", "file://schema.json", "Path of the schema file. (file:// or http(s)://)")
var json = flag.String("json", "file://json.json", "Path of the json file to check against the schema. (file:// or http(s)://)")

// Moven Message schema.
var movenSchemaValidator *gojsonschema.Schema

// LoadSchema loads the schema from a string.
func LoadSchema(schema string) error {
	var err error
	schemaLoader := gojsonschema.NewReferenceLoader(schema)

	movenSchemaValidator, err = gojsonschema.NewSchema(schemaLoader)
	return err
}

// ValidateJson validates the Json string agains the movenSchemaValidator.
func ValidateJson(json string) (*gojsonschema.Result, error) {
	documentLoader := gojsonschema.NewReferenceLoader(json)
	result, err := movenSchemaValidator.Validate(documentLoader)
	return result, err
}

// Main function, way to large and ugly.
func main() {
	flag.Parse()

	err := LoadSchema(string(*schema))
	if err != nil {
		log.Fatalf("Failed to load the schema validator: %s", err.Error())
	}

	res, err := ValidateJson(string(*json))
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
		os.Exit(1)
	}
	os.Exit(0)
}
