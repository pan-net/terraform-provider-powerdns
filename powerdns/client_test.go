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
	url, err := sanitizeURL(URLMissingSchemaAndNotEndingWithSlash)
	assert.NoError(t, err)

	expectedURL := DefaultSchema + "://" + URLMissingSchemaAndNotEndingWithSlash
	assert.Equal(t, url, expectedURL,
		"Expected '"+expectedURL+"' but got '"+url+"'")
}

func TestURLMissingSchemaAndEndingWithSlash(t *testing.T) {
	url, err := sanitizeURL(URLMissingSchemaAndEndingWithSlash)
	assert.NoError(t, err)

	expectedURL := DefaultSchema + "://" +
		strings.TrimSuffix(URLMissingSchemaAndEndingWithSlash, "/")

	assert.Equal(t, url, expectedURL,
		"Expected '"+expectedURL+"' but got '"+url+"'")
}

func TestURLWithSchemaAndEndingWithSlash(t *testing.T) {
	url, err := sanitizeURL(URLWithSchemaAndEndingWithSlash)
	assert.NoError(t, err)

	expectedURL := strings.TrimSuffix(URLWithSchemaAndEndingWithSlash, "/")

	assert.Equal(t, url, expectedURL,
		"Expected '"+expectedURL+"' but got '"+url+"'")
}

func TestURLWithSchemaAndNotEndingWithSlash(t *testing.T) {
	url, err := sanitizeURL(URLWithSchemaAndNotEndingWithSlash)
	assert.NoError(t, err)

	expectedURL := URLWithSchemaAndNotEndingWithSlash

	assert.Equal(t, url, expectedURL,
		"Expected '"+expectedURL+"' but got '"+url+"'")
}

func TestURLMissingSchemaHasPort(t *testing.T) {
	url, err := sanitizeURL(URLMissingSchemaHasPort)
	assert.NoError(t, err)

	expectedURL := DefaultSchema + "://" + URLMissingSchemaHasPort

	assert.Equal(t, url, expectedURL,
		"Expected '"+expectedURL+"' but got '"+url+"'")
}

func TestURLMissingSchemaHasPortAndEndsWithSlash(t *testing.T) {
	url, err := sanitizeURL(URLMissingSchemaHasPortAndEndsWithSlash)
	assert.NoError(t, err)

	expectedURL := DefaultSchema + "://" +
		strings.TrimSuffix(URLMissingSchemaHasPortAndEndsWithSlash, "/")

	assert.Equal(t, url, expectedURL,
		"Expected '"+expectedURL+"' but got '"+url+"'")
}

func TestURLWithSchemaHasPort(t *testing.T) {
	url, err := sanitizeURL(URLWithSchemaHasPort)
	assert.NoError(t, err)

	expectedURL := URLWithSchemaHasPort

	assert.Equal(t, url, expectedURL,
		"Expected '"+expectedURL+"' but got '"+url+"'")
}

func TestURLWithSchemaHasPortAndEndsWithSlash(t *testing.T) {
	url, err := sanitizeURL(URLWithSchemaHasPortAndEndsWithSlash)
	assert.NoError(t, err)

	expectedURL := strings.TrimSuffix(URLWithSchemaHasPortAndEndsWithSlash, "/")

	assert.Equal(t, url, expectedURL,
		"Expected '"+expectedURL+"' but got '"+url+"'")
}
