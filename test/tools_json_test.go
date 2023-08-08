package test

import (
	"gogogadget/pkg/tools"
	"testing"
)

func TestSimpleJSON(t *testing.T) {
	jsonStr := `{"name": "John Doe","age": 30,"country": "USA","address": {"street": "123 Main St","city": "New York"},"hobbies": ["reading", "swimming"]}`
	tabSize := 4

	formattedJSON, err := tools.IndentJSON(jsonStr, tabSize)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	expected := `{
	"name": "John Doe",
	"age": 30,
	"country": "USA",
	"address": {
		"street": "123 Main St",
		"city": "New York"
	},
	"hobbies": [
		"reading",
		"swimming"
	]
}`
	if formattedJSON != expected {
		t.Errorf("Formatted JSON doesn't match expected output:\n\nExpected:\n%s\n\nActual:\n%s", expected, formattedJSON)
	}
}

func TestNestedJSONWithArrays(t *testing.T) {
	jsonStr := `{
	"name": "Alice",
	"age": 25,
	"addresses": [
		{"street": "456 Elm St","city": "San Francisco"},
		{"street": "789 Oak St","city": "Los Angeles"}
	]
}`
	tabSize := 2

	formattedJSON, err := tools.IndentJSON(jsonStr, tabSize)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	expected := `{
  "name": "Alice",
  "age": 25,
  "addresses": [
    {
      "street": "456 Elm St",
      "city": "San Francisco"
    },
    {
      "street": "789 Oak St",
      "city": "Los Angeles"
    }
  ]
}`
	if formattedJSON != expected {
		t.Errorf("Formatted JSON doesn't match expected output:\n\nExpected:\n%s\n\nActual:\n%s", expected, formattedJSON)
	}
}
func TestEmptyJSON(t *testing.T) {
	jsonStr := "{}"
	tabSize := 3

	formattedJSON, err := tools.IndentJSON(jsonStr, tabSize)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	expected := "{}"
	if formattedJSON != expected {
		t.Errorf("Formatted JSON doesn't match expected output:\n\nExpected:\n%s\n\nActual:\n%s", expected, formattedJSON)
	}
}

func TestMalformedJSON(t *testing.T) {
	jsonStr := `{"name": "John Doe", "age": 30`
	tabSize := 2

	_, err := tools.IndentJSON(jsonStr, tabSize)
	if err == nil {
		t.Errorf("Expected an error for malformed JSON, but got none")
	}
}

func TestLargeJSONWithLongStrings(t *testing.T) {
	jsonStr := `{
    "description": "This is a very long description that goes on and on. It's so long that it spans multiple lines. We want to see how well the JSON formatting works with such long strings.",
    "data": [1, 2, 3, 4, 5]
}`
	tabSize := 2

	formattedJSON, err := tools.IndentJSON(jsonStr, tabSize)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	expected := `{
  "description": "This is a very long description that goes on and on. It's so long that it spans multiple lines. We want to see how well the JSON formatting works with such long strings.",
  "data": [
    1,
    2,
    3,
    4,
    5
  ]
}`
	if formattedJSON != expected {
		t.Errorf("Formatted JSON doesn't match expected output:\n\nExpected:\n%s\n\nActual:\n%s", expected, formattedJSON)
	}
}
