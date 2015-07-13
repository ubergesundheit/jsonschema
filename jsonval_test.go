// jsonval project jsonval.go
package jsonval

import (
	"github.com/xeipuuv/gojsonschema"
	"testing"
)

const (
	JSONSchema  = "file:///home/peza/Documents/DevProjects/nodeplay/jsonschema/schemas/json-schema.json"
	MovenSchema = "file:/home/peza/Documents/DevProjects/jsonschema/src/jsonval/cmd/validatejs/moven-dataservice-schema.json"
	Sample1     = "file:///home/peza/Documents/DevProjects/jsonschema/src/jsonval/cmd/validatejs/sample-message-1.json"
)

func TestValidateMovenSchema(t *testing.T) {

	schemaLoader := gojsonschema.NewReferenceLoader(MovenSchema)

	_, err := gojsonschema.NewSchema(schemaLoader)
	if err != nil {
		t.Errorf("Error loading schema: %s", err.Error())
	}
	t.Log("MovenSchema is valid.")

}

func TestValidateSample(t *testing.T) {

	schemaLoader := gojsonschema.NewReferenceLoader(MovenSchema)
	documentLoader := gojsonschema.NewReferenceLoader(Sample1)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	if result.Valid() {
		t.Logf("The document is valid\n")
	} else {
		t.Logf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			t.Logf("- %s\n", desc)
		}
	}
}
