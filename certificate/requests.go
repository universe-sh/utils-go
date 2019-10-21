package certificate

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

var oidEmailAddress = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 1}

type CSRequest struct {
	Length int
	Key    *rsa.PrivateKey
	CSR    []byte
}

// https://stackoverflow.com/questions/26043321/create-a-certificate-signing-request-csr-with-an-email-address-in-go
func NewCSR(length int, name string) (CSRequest, error) {
	var (
		request      CSRequest
		emailAddress string = "bot@universe.sh"
		subj                = pkix.Name{
			CommonName:   fmt.Sprintf("system:node:%s", name),
			Organization: []string{"system:nodes"},
		}
		err error = nil
	)

	if length < 2048 {
		return CSRequest{}, errors.New("RSA key is too weak")
	}

	if length > 8192 {
		return CSRequest{}, errors.New("RSA key size too large")
	}

	request.Length = length
	request.Key, err = rsa.GenerateKey(rand.Reader, length)
	if err != nil {
		return CSRequest{}, err
	}

	rawSubj := subj.ToRDNSequence()
	rawSubj = append(rawSubj, []pkix.AttributeTypeAndValue{{Type: oidEmailAddress, Value: emailAddress}})

	asn1Subj, _ := asn1.Marshal(rawSubj)
	template := x509.CertificateRequest{
		RawSubject:         asn1Subj,
		EmailAddresses:     []string{emailAddress},
		SignatureAlgorithm: x509.SHA256WithRSA,
	}

	csrBytes, _ := x509.CreateCertificateRequest(rand.Reader, &template, request.Key)
	request.CSR = pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE REQUEST", Bytes: csrBytes})

	return request, nil
}

// String encoding certificate
func (c *CSRequest) String() string {
	return string(c.CSR)
}

// Base64 encoding certificate
func (c *CSRequest) Base64() string {
	return base64.StdEncoding.EncodeToString(c.CSR)
}
