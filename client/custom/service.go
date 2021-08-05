// Package custom contains the methods to interact with the Custom Bitlinks in Bitly
package custom

import (
	"fmt"
	"github.com/retgits/bitly/client"
	"net/http"
)

const (
	createCustomEndpoint   = "custom_bitlinks"
	updateCustomEndpoint   = "custom_bitlinks/%s"
	retrieveCustomEndpoint = "custom_bitlinks/%s"
)

type CustomBitlinks struct {
	*client.Client
}

func New(c *client.Client) *CustomBitlinks {
	return &CustomBitlinks{
		c,
	}
}

func (b *CustomBitlinks) CreateCustomBitlink(bitlink *CustomBitlinkRequest) (CustomBitlink, error) {
	payload, err := bitlink.marshal()
	if err != nil {
		return CustomBitlink{}, err
	}

	data, err := b.Call(createCustomEndpoint, http.MethodPost, payload)
	if err != nil {
		return CustomBitlink{}, err
	}

	return unmarshalCustomBitlink(data)
}

func (b *CustomBitlinks) UpdateCustomBitlink(bitlink string, update *UpdateCustomBitlinkRequest) (CustomBitlink, error) {
	payload, err := update.marshal()
	if err != nil {
		return CustomBitlink{}, err
	}

	data, err := b.Call(fmt.Sprintf(updateCustomEndpoint, bitlink), http.MethodPatch, payload)
	if err != nil {
		return CustomBitlink{}, err
	}

	return unmarshalCustomBitlink(data)
}

func (b *CustomBitlinks) RetrieveCustomBitlink(bitlink string) (CustomBitlink, error) {
	data, err := b.Call(fmt.Sprintf(retrieveCustomEndpoint, bitlink), http.MethodGet, nil)
	if err != nil {
		return CustomBitlink{}, err
	}

	return unmarshalCustomBitlink(data)
}
