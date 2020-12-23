package main

import (
	"crypto/tls"
	//"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
	//"testing"
)

// func TestPemRequest(t *testing.T) {
func PemRequest() {
	b, _ := ioutil.ReadFile("/k8s/kubernetes/ssl/ca.pem")
	pem.Decode(b)
	var pemBlocks []*pem.Block
	var v *pem.Block
	var pkey []byte
	for {
		v, b = pem.Decode(b)
		if v == nil {
			break
		}
		if v.Type == "PRIVATE KEY" {
			pkey = pem.EncodeToMemory(v)
		} else {
			pemBlocks = append(pemBlocks, v)
		}
	}

	bytes := pem.EncodeToMemory(pemBlocks[0])
	keyString := string(pkey)
	CertString := string(bytes)
	fmt.Printf("Cert :\n %s \n Key:\n %s \n ", CertString, keyString)
	//pool := x509.NewCertPool()
	c, _ := tls.X509KeyPair(bytes, pkey)
	//pool.AppendCertsFromPEM(b)

	cfg := &tls.Config{
		Certificates: []tls.Certificate{c},
	}
	tr := &http.Transport{
		TLSClientConfig: cfg,
	}
	client := &http.Client{Transport: tr}
	//
	request, _ := http.NewRequest("GET", "https://127.0.0.1:10257/metrics", nil)
	request.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(request)
	// require.Nil(t, err)
	if err != nil {
		fmt.Println("err>>> ", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf(string(data))
	}
}
func main() {
	PemRequest()
}
