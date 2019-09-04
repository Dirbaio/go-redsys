package redsys

import (
	"context"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func (c *client) DoXMLRequest(ctx context.Context, req *Request) (*Response, error) {
	req.MerchantCode = c.MerchantCode
	req.Terminal = c.Terminal

	data := url.Values{}
	data.Add("entrada", c.signRequestXML(req))
	httpReq, err := http.NewRequest("POST", c.URL+"/operaciones", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed creating Redsys request: %v", err)
	}

	httpReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	httpReq = httpReq.WithContext(ctx)

	httpResp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error doing Redsys request: %v", err)
	}
	defer httpResp.Body.Close()

	bytes, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading Redsys response: %v", err)
	}

	var resp Response
	if err := xml.Unmarshal(bytes, &resp); err != nil {
		return nil, fmt.Errorf("error unmarshaling Redsys response: %v", err)
	}

	return &resp, nil
}

func (c *client) signRequestXML(req *Request) string {
	x, err := xml.Marshal(req)
	if err != nil {
		panic(err)
	}

	res, err := xml.Marshal(&struct {
		XMLName          xml.Name `xml:"REQUEST"`
		Request          []byte   `xml:",innerxml"`
		SignatureVersion string   `xml:"DS_SIGNATUREVERSION"`
		Signature        string   `xml:"DS_SIGNATURE"`
	}{
		Request:          x,
		Signature:        base64.StdEncoding.EncodeToString(c.signData(x, req.OperationID)),
		SignatureVersion: "HMAC_SHA256_V1",
	})
	if err != nil {
		panic(err)
	}

	return string(res)
}
