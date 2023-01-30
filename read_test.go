package libx_test

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/miajio/libx/read"
)

type XmlRoot struct {
	XMLName xml.Name `xml:"note"`
	XmlNode
}

type XmlNode struct {
	To      string `xml:"to" json:"to"`
	From    string `xml:"from" json:"from"`
	Heading string `xml:"heading" json:"heading"`
	Body    string `xml:"body" json:"body"`
}

type JsonRoot struct {
	Note XmlNode `json:"note"`
}

func TestReadXml(t *testing.T) {
	base, _ := os.Getwd()
	fmt.Println(base)
	xn := XmlRoot{}
	if err := read.ReadXml(base+"/hello.xml", &xn); err != nil {
		fmt.Printf("read xml fail: %v\n", err)
	}
	r, _ := json.Marshal(xn)
	fmt.Println(string(r))
}

func TestReadJson(t *testing.T) {
	base, _ := os.Getwd()
	fmt.Println(base)
	xn := JsonRoot{}
	if err := read.ReadJson(base+"/hello.json", &xn); err != nil {
		fmt.Printf("read xml fail: %v\n", err)
	}
	r, _ := json.Marshal(xn)
	fmt.Println(string(r))
}

func TestReadText(t *testing.T) {
	base, _ := os.Getwd()
	fmt.Println(base)
	if xn, err := read.ReadText(base + "/hello.json"); err != nil {
		fmt.Printf("read xml fail: %v\n", err)
	} else {
		fmt.Println(xn)
	}
}

func TestReadIO(t *testing.T) {
	resp, err := http.Get("https://miajio.oss-cn-hangzhou.aliyuncs.com/group1/M00/01/3B/hello.json")
	if err != nil {
		fmt.Printf("get failed, err: %v", err)
		return
	}
	defer resp.Body.Close()

	if xn, err := read.ReadIO(read.TEXT, resp.Body, nil); err != nil {
		fmt.Printf("read io fail: %v\n", err)
	} else {
		fmt.Println(string(xn))
	}
}

func TestReadIOJson(t *testing.T) {
	resp, err := http.Get("https://miajio.oss-cn-hangzhou.aliyuncs.com/group1/M00/01/3B/hello.json")
	if err != nil {
		fmt.Printf("get failed, err: %v", err)
		return
	}
	defer resp.Body.Close()
	xn := JsonRoot{}
	if _, err := read.ReadIO(read.JSON, resp.Body, &xn); err != nil {
		fmt.Printf("read io fail: %v\n", err)
	} else {
		r, _ := json.Marshal(xn)
		fmt.Println(string(r))
	}
}

func TestReadIOXml(t *testing.T) {
	resp, err := http.Get("https://miajio.oss-cn-hangzhou.aliyuncs.com/group1/M00/01/3B/hello.xml")
	if err != nil {
		fmt.Printf("get failed, err: %v", err)
		return
	}
	defer resp.Body.Close()
	xn := XmlRoot{}
	if _, err := read.ReadIO(read.XML, resp.Body, &xn); err != nil {
		fmt.Printf("read io fail: %v\n", err)
	} else {
		r, _ := json.Marshal(xn)
		fmt.Println(string(r))
	}
}
