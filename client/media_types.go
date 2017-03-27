// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/nubunto/autobus-web/design
// --out=$(GOPATH)/src/github.com/nubunto/autobus-web
// --version=v1.1.0-dirty
//
// API "autobus-web": Application Media Types
//
// The content of this file is auto-generated, DO NOT MODIFY

package client

import (
	"net/http"
)

// The GPS data media type (default view)
//
// Identifier: autobus.web.platform/gps.media+json; view=default
type GpsMedia struct {
	Date      *string `form:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	ID        *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Latitude  *string `form:"latitude,omitempty" json:"latitude,omitempty" xml:"latitude,omitempty"`
	Longitude *string `form:"longitude,omitempty" json:"longitude,omitempty" xml:"longitude,omitempty"`
	Status    *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	Time      *string `form:"time,omitempty" json:"time,omitempty" xml:"time,omitempty"`
}

// DecodeGpsMedia decodes the GpsMedia instance encoded in resp body.
func (c *Client) DecodeGpsMedia(resp *http.Response) (*GpsMedia, error) {
	var decoded GpsMedia
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
