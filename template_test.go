package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
	"text/template"

	"github.com/PuerkitoBio/goquery"
	"github.com/flosch/pongo2/v6"
	"github.com/hoisie/mustache"
	"github.com/valyala/fasttemplate"
)

var styles = map[string]string{
	"h1":         "text-7xl",
	"h2":         "text-4xl",
	"h3":         "text-3xl",
	"h4":         "text-2xl",
	"p":          "text-xl leading-loose",
	"em":         "text-xl",
	"ul":         "text-xl",
	"blockquote": "text-xl",
}

func readTestFile(filename string) (string, error) {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(fileContent), nil
}

func BenchmarkPango2(b *testing.B) {
	fileContent, err := readTestFile("fasttemplate.html")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	context := pongo2.Context{}
	for key, value := range styles {
		context[key] = value
	}

	b.ResetTimer()

	tpl, err := pongo2.FromString(fileContent)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	_, err = tpl.Execute(context)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func BenchmarkGoquery(b *testing.B) {
	fileContent, err := readTestFile("blame.html")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v", err)
	}

	b.ResetTimer()

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(fileContent))
	if err != nil {
		fmt.Fprintf(os.Stderr, "goquery error: %v\n", err)
	}

	for tag, style := range styles {
		doc.Find(tag).SetAttr("class", style)
	}

	doc.Html()
}

func BenchmarkGoTemplate(b *testing.B) {
	fileContent, err := readTestFile("blame_template.html")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	b.ResetTimer()

	tpl, err := template.New("").Parse(fileContent)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	var result bytes.Buffer
	tpl.Execute(&result, styles)

	result.String()
}

func BenchmarkMustache(b *testing.B) {
	fileContent, err := readTestFile("fasttemplate.html")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	b.ResetTimer()

	mustache.Render(fileContent, styles)
}

func BenchmarkFastTemplate(b *testing.B) {
	fileContent, err := readTestFile("fasttemplate.html")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	context := make(map[string]interface{})
	for key, value := range styles {
		context[key] = value
	}

	b.ResetTimer()

	fasttemplate.New(fileContent, "{{", "}}").ExecuteString(context)
}
