package main
//
//import (
//	"fmt"
//	"github.com/pkg/errors"
//	"log"
//	"net/http"
//	"net/http/httputil"
//	"net/url"
//)
//

//
//func GetTransport() *TransportStruct {
//	target, err := url.Parse(SocksProxyURI)
//	if err != nil {
//		err = errors.Errorf("Error parsing proxy URI, %s", SocksProxyURI, err)
//		fmt.Println(err)
//	}
//	transport := &http.Transport{Proxy: http.ProxyURL(target)}
//	transportStruct := &TransportStruct{transport, nil}
//	//return transport
//	return transportStruct
//}
//
//type TransportStruct struct {
//	*http.Transport
//	header *http.Header
//	//CapturedTransport http.RoundTripper
//}
//
//func (transport TransportStruct) RoundTrip(req *http.Request) (*http.Response, error) {
//
//	resp, err := transport.Transport.RoundTrip(req)
//
//	//fmt.Println("RoundTrip() executing ==> ==> ==>")
//	if err != nil {
//		fmt.Printf("Error with Transport RoundTrip(), %v", err)
//	}
//
//	//bytes, err := httputil.DumpResponse(resp, true)
//	//if err != nil {
//	//	fmt.Println(err)
//	//}
//	//fmt.Println(string(bytes))
//
//	//originalResponseHeader := resp.Header
//	//delete(originalResponseHeader, "X-Frame-Options")
//	//fmt.Printf("\nRequest [%s] ==> ==> ==>\n", resp.Request.RequestURI)
//	//
//	//fmt.Println("\nRequest Headers ==> ==> ==>")
//	//for key, vals := range req.Header {
//	//	fmt.Printf("[%s]: %s\n", key, vals)
//	//}
//	//
//	//fmt.Println("\nResponse Headers ==> ==> ==>")
//	//for key, vals := range resp.Header {
//	//	fmt.Printf("[%s]: %s\n", key, vals)
//	//}
//
//	delete(resp.Header, "X-Frame-Options")
//
//	return resp, err
//}
//
//func httpbin(app string) {
//
//	target, _ := url.Parse(app)
//
//	proxy := &httputil.ReverseProxy{
//
//		Director: func(req *http.Request) {
//
//			fmt.Printf("Creating new request with parameters ==>\n")
//			fmt.Printf("New Request Scheme: [%s]\n", target.Scheme)
//			fmt.Printf("New Request host: [%s]\n", target.Host)
//			fmt.Printf("Request URL Path: [%s]\n", req.URL.Path)
//			fmt.Printf("Request URL RawQuery: [%s]\n", req.URL.RawQuery)
//
//			err := req.ParseForm()
//			if err != nil {
//				fmt.Printf("Error parsing Form: %+v", err)
//			}
//
//			fmt.Printf("Form: %+v\n", req.Form)
//
//			req.Header = http.Header{}
//			//req.Header.Add("X-Forwarded-Host", req.Host)
//			//req.Header.Add("X-Origin-Host", target.Host)
//			//req.Header.Add("Accept", "*/*")
//			//req.Header.Add("User-Agent", "curl/7.61.1")
//
//			//GET /api/v1/applications?limit=2147483647&status=completed HTTP/1.1
//			//Host: ip-10-42-24-247.ec2.internal:18080
//			//User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:81.0) Gecko/20100101 Firefox/81.0
//			//Accept: application/json, text/javascript, */*; q=0.01
//			//Accept-Language: en-US,en;q=0.5
//			//Accept-Encoding: gzip, deflate
//			//X-Requested-With: XMLHttpRequest
//			//DNT: 1
//			//Connection: keep-alive
//			//Referer: http://ip-10-42-24-247.ec2.internal:18080/
//			//Cookie: _xsrf=2|29580759|e8848305d8f9893bb39a05a8eb4f1781|1600134271
//			//Cache-Control: max-age=0
//
//			//req.Host = target.Host
//			req.URL.Scheme = "http"
//			req.URL.Host = target.Host
//			//req.URL.Path = req.RequestURI
//			//req.URL.Path = target.Path
//
//		},
//		Transport: GetTransport(),
//	}
//
//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//
//		//target, err := url.Parse(r.RequestURI)
//		//if err != nil {
//		//	fmt.Printf("Error parsing initial request, %+v\n", err)
//		//}
//		//fmt.Printf("Target EMR App Interface URI: %+v\n", target)
//
//		proxy.ServeHTTP(w, r)
//	})
//
//	log.Fatal(http.ListenAndServe(":9001", nil))
//}