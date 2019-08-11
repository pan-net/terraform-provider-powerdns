package powerdns

import (
	"crypto/tls"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	tlsConfig                               = &tls.Config{}
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
		"secretapikey", tlsConfig)
	assert.NoError(t, err)

	expectedURL := DefaultSchema + "://" + URLMissingSchemaAndNotEndingWithSlash

	assert.Equal(t, client.ServerURL, expectedURL,
		"Expected '"+expectedURL+"' but got '"+client.ServerURL+"'")
}

func TestURLMissingSchemaAndEndingWithSlash(t *testing.T) {
	client, err := NewClient(URLMissingSchemaAndEndingWithSlash,
		"secretapikey", tlsConfig)
	assert.NoError(t, err)

	expectedURL := DefaultSchema + "://" +
		strings.TrimSuffix(URLMissingSchemaAndEndingWithSlash, "/")

	assert.Equal(t, client.ServerURL, expectedURL,
		"Expected '"+expectedURL+"' but got '"+client.ServerURL+"'")
}

func TestURLWithSchemaAndEndingWithSlash(t *testing.T) {
	client, err := NewClient(URLWithSchemaAndEndingWithSlash,
		"secretapikey", tlsConfig)
	assert.NoError(t, err)

	expectedURL := strings.TrimSuffix(URLWithSchemaAndEndingWithSlash, "/")

	assert.Equal(t, client.ServerURL, expectedURL,
		"Expected '"+expectedURL+"' but got '"+client.ServerURL+"'")
}

func TestURLWithSchemaAndNotEndingWithSlash(t *testing.T) {
	client, err := NewClient(URLWithSchemaAndNotEndingWithSlash,
		"secretapikey", tlsConfig)
	assert.NoError(t, err)

	expectedURL := URLWithSchemaAndNotEndingWithSlash

	assert.Equal(t, client.ServerURL, expectedURL,
		"Expected '"+expectedURL+"' but got '"+client.ServerURL+"'")
}

func TestURLMissingSchemaHasPort(t *testing.T) {
	client, err := NewClient(URLMissingSchemaHasPort, "secretapikey", tlsConfig)
	assert.NoError(t, err)

	expectedURL := DefaultSchema + "://" + URLMissingSchemaHasPort

	assert.Equal(t, client.ServerURL, expectedURL,
		"Expected '"+expectedURL+"' but got '"+client.ServerURL+"'")
}

func TestURLMissingSchemaHasPortAndEndsWithSlash(t *testing.T) {
	client, err := NewClient(URLMissingSchemaHasPortAndEndsWithSlash,
		"secretapikey", tlsConfig)
	assert.NoError(t, err)

	expectedURL := DefaultSchema + "://" +
		strings.TrimSuffix(URLMissingSchemaHasPortAndEndsWithSlash, "/")

	assert.Equal(t, client.ServerURL, expectedURL,
		"Expected '"+expectedURL+"' but got '"+client.ServerURL+"'")
}

func TestURLWithSchemaHasPort(t *testing.T) {
	client, err := NewClient(URLWithSchemaHasPort,
		"secretapikey", tlsConfig)
	assert.NoError(t, err)

	expectedURL := URLWithSchemaHasPort

	assert.Equal(t, client.ServerURL, expectedURL,
		"Expected '"+expectedURL+"' but got '"+client.ServerURL+"'")
}

func TestURLWithSchemaHasPortAndEndsWithSlash(t *testing.T) {
	client, err := NewClient(URLWithSchemaHasPortAndEndsWithSlash,
		"secretapikey", tlsConfig)
	assert.NoError(t, err)

	expectedURL := strings.TrimSuffix(URLWithSchemaHasPortAndEndsWithSlash, "/")

	assert.Equal(t, client.ServerURL, expectedURL,
		"Expected '"+expectedURL+"' but got '"+client.ServerURL+"'")
}
