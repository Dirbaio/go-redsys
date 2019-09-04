package redsys

import (
	"crypto/subtle"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

type SignedJSON struct {
	MerchantParameters string
	Signature          string
	SignatureVersion   string
}

func (c *client) SignJSONRequest(req *Request) (*SignedJSON, error) {
	req.MerchantCode = c.MerchantCode
	req.Terminal = c.Terminal

	j, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}

	j64 := base64.StdEncoding.EncodeToString(j)
	sig := c.signData([]byte(j64), req.OperationID)

	return &SignedJSON{
		MerchantParameters: j64,
		Signature:          base64.StdEncoding.EncodeToString(sig),
		SignatureVersion:   "HMAC_SHA256_V1",
	}, nil
}

func (c *client) ValidateJSONResponse(s *SignedJSON) (*Response, error) {
	if s.SignatureVersion != "HMAC_SHA256_V1" {
		return nil, fmt.Errorf("Invalid signature version '%s'", s.SignatureVersion)
	}

	// Redsys only accepts signatures in StdEncoding, but it sends signatures in URLEncoding...???
	// Replace this so we accept both.
	b64 := s.Signature
	b64 = strings.Replace(b64, "-", "+", -1)
	b64 = strings.Replace(b64, "_", "/", -1)

	gotSig, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return nil, fmt.Errorf("Invalid base64 in signature: %v", err)
	}

	// Do the same for data, just in case
	b64 = s.MerchantParameters
	b64 = strings.Replace(b64, "-", "+", -1)
	b64 = strings.Replace(b64, "_", "/", -1)
	data, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return nil, fmt.Errorf("Invalid base64 in MerchantParameters: %v", err)
	}

	var resp Response
	if err = json.Unmarshal(data, &resp); err != nil {
		return nil, fmt.Errorf("Invalid json in MerchantParameters: %v", err)
	}

	correctSig := c.signData([]byte(s.MerchantParameters), resp.OperationID)

	if subtle.ConstantTimeCompare(gotSig, correctSig) != 1 {
		return nil, fmt.Errorf("Wrong signature")
	}

	return &resp, nil
}
