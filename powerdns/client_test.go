package powerdns

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	URLMissingSchemaAndNotEndingWithSlash   = "powerdnsapi.com"
	URLMissingSchemaAndEndingWithSlash      = "powerdnsapi.com/"
	URLWithSchemaAndEndingWithSlash         = "http://powerdnsapi.com/"
	URLWithSchemaAndNotEndingWithSlash      = "http://powerdnsapi.com"
	URLWithSchemaAndPath                    = "https://powerdnsapi.com/api/v2"
	URLMissingSchemaHasPort                 = "powerdnsapi.com:443"
	URLMissingSchemaHasPortAndEndsWithSlash = "powerdnsapi.com:443/"
	URLWithSchemaHasPort                    = "http://powerdnsapi.com:443"
	URLWithSchemaHasPortAndEndsWithSlash    = "http://powerdnsapi.com:443/"
)

func TestURLMissingSchema(t *testing.T) {
	client, err := NewClient(URLMissingSchemaAndNotEndingWithSlash,
		"secretapikey", true)
	assert.NoError(t, err)

	expectedURL := DefaultSchema + "://" + URLMissingSchemaAndNotEndingWithSlash

	assert.Equal(t, client.ServerUrl, expectedURL,
		"Expected '"+expectedURL+"' but got '"+client.ServerUrl+"'")
}

func TestURLMissingSchemaAndEndingWithSlash(t *testing.T) {
	client, err := NewClient(URLMissingSchemaAndEndingWithSlash,
		"secretapikey", true)
	assert.NoError(t, err)

	expectedURL := DefaultSchema + "://" +
		strings.TrimSuffix(URLMissingSchemaAndEndingWithSlash, "/")

	assert.Equal(t, client.ServerUrl, expectedURL,
		"Expected '"+expectedURL+"' but got '"+client.ServerUrl+"'")
}

func TestURLWithSchemaAndEndingWithSlash(t *testing.T) {
	client, err := NewClient(URLWithSchemaAndEndingWithSlash,
		"secretapikey", true)
	assert.NoError(t, err)

	expectedURL := strings.TrimSuffix(URLWithSchemaAndEndingWithSlash, "/")

	assert.Equal(t, client.ServerUrl, expectedURL,
		"Expected '"+expectedURL+"' but got '"+client.ServerUrl+"'")
}

func TestURLWithSchemaAndNotEndingWithSlash(t *testing.T) {
	client, err := NewClient(URLWithSchemaAndNotEndingWithSlash,
		"secretapikey", true)
	assert.NoError(t, err)

	expectedURL := URLWithSchemaAndNotEndingWithSlash

	assert.Equal(t, client.ServerUrl, expectedURL,
		"Expected '"+expectedURL+"' but got '"+client.ServerUrl+"'")
}

func TestURLMissingSchemaHasPort(t *testing.T) {
	client, err := NewClient(URLMissingSchemaHasPort, "secretapikey", true)
	assert.NoError(t, err)

	expectedURL := DefaultSchema + "://" + URLMissingSchemaHasPort

	assert.Equal(t, client.ServerUrl, expectedURL,
		"Expected '"+expectedURL+"' but got '"+client.ServerUrl+"'")
}

func TestURLMissingSchemaHasPortAndEndsWithSlash(t *testing.T) {
	client, err := NewClient(URLMissingSchemaHasPortAndEndsWithSlash,
		"secretapikey", true)
	assert.NoError(t, err)

	expectedURL := DefaultSchema + "://" +
		strings.TrimSuffix(URLMissingSchemaHasPortAndEndsWithSlash, "/")

	assert.Equal(t, client.ServerUrl, expectedURL,
		"Expected '"+expectedURL+"' but got '"+client.ServerUrl+"'")
}

func TestURLWithSchemaHasPort(t *testing.T) {
	client, err := NewClient(URLWithSchemaHasPort,
		"secretapikey", true)
	assert.NoError(t, err)

	expectedURL := URLWithSchemaHasPort

	assert.Equal(t, client.ServerUrl, expectedURL,
		"Expected '"+expectedURL+"' but got '"+client.ServerUrl+"'")
}

func TestURLWithSchemaHasPortAndEndsWithSlash(t *testing.T) {
	client, err := NewClient(URLWithSchemaHasPortAndEndsWithSlash,
		"secretapikey", true)
	assert.NoError(t, err)

	expectedURL := strings.TrimSuffix(URLWithSchemaHasPortAndEndsWithSlash, "/")

	assert.Equal(t, client.ServerUrl, expectedURL,
		"Expected '"+expectedURL+"' but got '"+client.ServerUrl+"'")
}
