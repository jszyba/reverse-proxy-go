package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

func DisplayHeaders(header http.Header, headerType string)  {
	fmt.Printf("\n\n%s\n",time.Now().Format("15:04:05.000000"))
	fmt.Printf("%s Headers ==> ==> ==>\n", strings.ToUpper(headerType))
	for key, vals := range header {
		fmt.Printf("[%s]: %s\n", key, vals)
	}
}

func GetTransport() *TransportStruct {
	//fmt.Println("GetTransport() executing ==> ==> ==>")
	target, err := url.Parse(SocksProxyURI)
	if err != nil {
		err = errors.Errorf("Error parsing proxy URI, %s", SocksProxyURI, err)
		fmt.Println(err)
	}
	transport := &http.Transport{Proxy: http.ProxyURL(target)}
	transportStruct := &TransportStruct{transport, nil}
	return transportStruct
}

type TransportStruct struct {
	*http.Transport
	header *http.Header
}

func StartReverseProxy(app string) {

	var target *url.URL
	var err error

	target, err = url.Parse(app)
	if err != nil {
		fmt.Printf("Error parsing app uri: [%s], %+v\n", app, err)
	}

	proxy := &httputil.ReverseProxy{

		Director: func(req *http.Request) {

			//fmt.Println("Director() executing ==> ==> ==>")

			//fmt.Printf("Creating new request with parameters ==>\n")
			//fmt.Printf("New Request Scheme: [%s]\n", target.Scheme)
			//fmt.Printf("New Request host: [%s]\n", target.Host)
			//fmt.Printf("Request URL Path: [%s]\n", req.URL.Path)
			//fmt.Printf("Request URL RawQuery: [%s]\n", req.URL.RawQuery)

			err = req.ParseForm()
			if err != nil {
				fmt.Printf("Error parsing Form: %+v", err)
			}

			//fmt.Printf("Form: %+v\n", req.Form)

			//req.Header = http.Header{}
			//host := fmt.Sprintf("%s://%s", HTTP, req.Host)
			//req.Header.Add("X-Forwarded-Host", host)
			//req.Header.Add("X-Origin-Host", host)

			req.Header.Add("X-Forwarded-Host", req.Host)
			req.Header.Add("X-Origin-Host", target.Host)

			//req.Header.Add("Accept", "*/*")
			//req.Header.Add("User-Agent", "curl/7.61.1")

			req.URL.Scheme = HTTP
			req.URL.Host = target.Host

			DisplayHeaders(req.Header, "Request from Director()")

		},
		Transport: GetTransport(),
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})

	http.HandleFunc("/rm/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/main.html")
	})

	log.Fatal(http.ListenAndServe(":9001", nil))
}

func (transport TransportStruct) RoundTrip(req *http.Request) (*http.Response, error) {
	//fmt.Println("RoundTrip() executing ==> ==> ==>")
	var headerPtr *http.Header
	resp, err := transport.Transport.RoundTrip(req)
	if err != nil {
		fmt.Printf("Error with Transport RoundTrip(), %v", err)
	}
	headerPtr = &resp.Header
	DisplayHeaders(*headerPtr, "response from roundtrip()")

	delete(*headerPtr, "X-Frame-Options")
	//delete(*headerPtr, "X-Xss-Protection")

	return resp, err
}