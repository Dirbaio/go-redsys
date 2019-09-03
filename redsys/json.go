package redsys

import (
	"encoding/base64"
	"encoding/json"
)

type SignedJSONRequest struct {
	Request          string
	Signature        string
	SignatureVersion string
}

func (c *client) SignRequestJSON(req *Request) (*SignedJSONRequest, error) {
	j, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}

	j64 := base64.StdEncoding.EncodeToString(j)
	sig := c.signData([]byte(j64), req.OperationID)

	return &SignedJSONRequest{
		Request:          j64,
		Signature:        sig,
		SignatureVersion: "HMAC_SHA256_V1",
	}, nil
}
