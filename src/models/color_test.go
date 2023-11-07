package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColorJSONMarshaling(t *testing.T) {
	color := Color{
		Name: "red",
		Hex:  "#FF0000",
	}

	jsonBytes, err := json.Marshal(color)
	assert.NoError(t, err, "marshaling to JSON should not produce an error")

	var decodedColor Color
	err = json.Unmarshal(jsonBytes, &decodedColor)
	assert.NoError(t, err, "unmarshaling JSON should not produce an error")

	assert.Equal(t, color, decodedColor, "the decoded color should match the original")
}

func TestColorJSONUnmarshaling(t *testing.T) {
	jsonString := `{"name":"blue","hex":"#0000FF"}`

	var color Color
	err := json.Unmarshal([]byte(jsonString), &color)
	assert.NoError(t, err, "unmarshaling JSON should not produce an error")

	assert.Equal(t, "blue", color.Name, "the name should be set correctly")
	assert.Equal(t, "#0000FF", color.Hex, "the hex value should be set correctly")
}
