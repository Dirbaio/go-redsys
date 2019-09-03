package redsys

import (
	"context"
	"crypto/cipher"
	"crypto/des"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
)

type Options struct {
	MerchantCode string
	Terminal     int
	Key          string

	URL        string
	HTTPClient *http.Client
}

type Client interface {
	DoRESTRequest(ctx context.Context, req *Request) (*Response, error)
	SignRequestJSON(req *Request) (*SignedJSONRequest, error)
}

type client struct {
	URL        string
	HTTPClient *http.Client

	MerchantCode string
	Terminal     int

	tdes cipher.Block
}

func New(opts Options) (Client, error) {
	key, err := base64.StdEncoding.DecodeString(opts.Key)
	if err != nil {
		return nil, fmt.Errorf("Key is invalid base64: %v", err)
	}

	tdes, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, fmt.Errorf("error creating 3DES cipher: %v", err)
	}

	if opts.HTTPClient == nil {
		opts.HTTPClient = http.DefaultClient
	}
	if opts.URL == "" {
		opts.URL = "https://sis.redsys.es/sis"
	}
	return &client{
		URL:          opts.URL,
		HTTPClient:   opts.HTTPClient,
		MerchantCode: opts.MerchantCode,
		Terminal:     opts.Terminal,
		tdes:         tdes,
	}, nil
}

func (c *client) signData(data []byte, id string) string {
	var plaintext [16]byte
	copy(plaintext[:], []byte(id))

	var zeros [8]byte
	var key2 [16]byte
	cipher.NewCBCEncrypter(c.tdes, zeros[:]).CryptBlocks(key2[:], plaintext[:])

	h := hmac.New(sha256.New, key2[:])
	_, _ = h.Write(data)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
