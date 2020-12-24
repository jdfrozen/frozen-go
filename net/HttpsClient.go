package main

import (
	"crypto/tls"
	"crypto/x509"
	"log"

	"fmt"
	"io/ioutil"
	"net/http"
)

func client() {
	pool := x509.NewCertPool()
	caCrt, err := ioutil.ReadFile("ca.pem")
	if err != nil {
		log.Fatal("read ca.crt file error:", err.Error())
	}
	pool.AppendCertsFromPEM(caCrt)
	cliCrt, err := tls.LoadX509KeyPair("client.pem", "client-key.pem")
	if err != nil {
		log.Fatalln("LoadX509KeyPair error:", err.Error())
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:      pool,
			Certificates: []tls.Certificate{cliCrt},
		},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://127.0.0.1:8081/")
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
func client1() {
	pool := x509.NewCertPool()
	caCrt, err := ioutil.ReadFile("ca.pem")
	if err != nil {
		log.Fatal("read ca.crt file error:", err.Error())
	}
	pool.AppendCertsFromPEM(caCrt)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs: pool,
		},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://127.0.0.1:8081/")
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func main() {
	client1()
}
