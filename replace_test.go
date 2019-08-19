package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"testing"
)

var result interface{}
var dat []byte
var str string

func TestMain(m *testing.M) {
	var err error
	dat, err = ioutil.ReadFile("correios_pac_error.zpl")
	if err != nil {
		fmt.Printf("Error read file %v\n", err)
	}
	str = string(dat)

	os.Exit(m.Run())
}

// BenchmarkBytesReplace to check Replace with bytes
// https://golang.org/pkg/bytes/#Replace
func BenchmarkBytesReplace(b *testing.B) {
	data := dat
	toFind := []byte("220A22")
	toReplace := []byte("")

	b.ResetTimer()
	b.ReportAllocs()
	var r []byte
	for n := 0; n < b.N; n++ {
		r = bytes.Replace(data, toFind, toReplace, -1) // -1 replace all
	}

	result = r
}

// BenchmarkStringsReplace to check Replace with strings
// https://golang.org/pkg/strings/#Replace
func BenchmarkStringsReplace(b *testing.B) {
	data := str
	toFind := "\"\n\""
	toReplace := ""

	b.ResetTimer()
	b.ReportAllocs()
	var r string
	for n := 0; n < b.N; n++ {
		r = strings.Replace(data, toFind, toReplace, -1) // -1 replace all
	}

	result = r
}

// BenchmarkBytesRegexp to check ReplaceAll with regexp
// https://golang.org/pkg/regexp/#Regexp.ReplaceAll
func BenchmarkBytesRegexp(b *testing.B) {
	data := dat
	toReplace := []byte("")

	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		r, _ := regexp.Compile("\"\n\"")
		r.ReplaceAll(data, toReplace)
	}

	result = data
}
