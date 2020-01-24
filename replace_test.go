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

// BenchmarkOneBytesReplace to check Replace with bytes
// https://golang.org/pkg/bytes/#Replace
func BenchmarkF1_OneBytesReplace(b *testing.B) {
	data := dat
	toFind := []byte("\"\n\"")
	toReplace := []byte("")

	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		bytes.Replace(data, toFind, toReplace, -1) // -1 replace all
	}
}

// BenchmarkStringsOneReplace to check Replace with strings
// https://golang.org/pkg/strings/#Replace
func BenchmarkF1_StringsOneReplace(b *testing.B) {
	data := str
	toFind := "\"\n\""
	toReplace := ""

	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		strings.Replace(data, toFind, toReplace, -1) // -1 replace all
	}
}

// BenchmarkStringsOneReplaceConvert to check Replace with strings
// https://golang.org/pkg/strings/#Replace
func BenchmarkF1_StringsOneReplaceConvert(b *testing.B) {
	toFind := "\"\n\""
	toReplace := ""

	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		str := string(dat)
		strings.Replace(str, toFind, toReplace, -1) // -1 replace all
	}
}

// BenchmarkBytesRegexp to check ReplaceAll with regexp
// https://golang.org/pkg/regexp/#Regexp.ReplaceAll
func BenchmarkF1_BytesOneRegexp(b *testing.B) {
	data := dat
	toReplace := []byte("")

	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		r, _ := regexp.Compile("\"\n\"")
		r.ReplaceAll(data, toReplace)
	}
}

func BenchmarkF1_BytesManyRegexp(b *testing.B) {
	data := dat
	toReplace := []byte("")

	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		r, _ := regexp.Compile("(\"\n\"|MERCADOLIBRE|R01FR078|K0101FC|anuncios)")
		r.ReplaceAll(data, toReplace)
	}
}

func BenchmarkF1_BytesManyReplace(b *testing.B) {
	data := dat
	toFind1 := []byte("\"\n\"")
	toFind2 := []byte("MERCADOLIBRE")
	toFind3 := []byte("R01FR078")
	toFind4 := []byte("K0101FC")
	toFind5 := []byte("anuncios")
	toReplace := []byte("")

	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		bytes.Replace(data, toFind1, toReplace, -1) // -1 replace all
		bytes.Replace(data, toFind2, toReplace, -1) // -1 replace all
		bytes.Replace(data, toFind3, toReplace, -1) // -1 replace all
		bytes.Replace(data, toFind4, toReplace, -1) // -1 replace all
		bytes.Replace(data, toFind5, toReplace, -1) // -1 replace all
	}
}

func BenchmarkF1_StringsManyReplacer(b *testing.B) {
	data := str
	toFind := "\"\n\""
	toReplace := ""

	b.ResetTimer()
	b.ReportAllocs()
	replacer := strings.NewReplacer(toFind, toReplace, "MERCADOLIBRE", toReplace, "R01FR078", toReplace, "K0101FC", toReplace, "anuncios", toReplace)
	for n := 0; n < b.N; n++ {
		replacer.Replace(data)
	}
}

func BenchmarkF1_StringsManyReplace(b *testing.B) {
	data := str
	toFind := "\"\n\""
	toReplace := ""

	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		strings.Replace(data, toFind, toReplace, -1)         // -1 replace all
		strings.Replace(data, "MERCADOLIBRE", toReplace, -1) // -1 replace all
		strings.Replace(data, "R01FR078", toReplace, -1)     // -1 replace all
		strings.Replace(data, "K0101FC", toReplace, -1)      // -1 replace all
		strings.Replace(data, "anuncios", toReplace, -1)     // -1 replace all
	}
}

func BenchmarkF1_StringsManyReplaceConvert(b *testing.B) {
	toFind := "\"\n\""
	toReplace := ""
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		data := string(dat)
		strings.Replace(data, toFind, toReplace, -1)         // -1 replace all
		strings.Replace(data, "MERCADOLIBRE", toReplace, -1) // -1 replace all
		strings.Replace(data, "R01FR078", toReplace, -1)     // -1 replace all
		strings.Replace(data, "K0101FC", toReplace, -1)      // -1 replace all
		strings.Replace(data, "anuncios", toReplace, -1)     // -1 replace all
	}
}
