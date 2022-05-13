package main

import (
	"fmt"

	"github.com/karazhanov/translate/example"
	"github.com/karazhanov/translate/translate"
)

const tagName = "translate_to"

type Example struct {
	Mnemonic string `translate_to:"Name"`
	Name     string

	Mnemonic2 string `translate_to:"Name2"`
	Name2     string

	Mnemonic3 string `translate_to:"Name3"`
	Name3     string
}

func main() {
	stub := example.Stub()
	translate := translate.New(stub, tagName)

	example := Example{
		Mnemonic:  "tmp",
		Name:      "",
		Mnemonic2: "one",
		Name2:     "",
		Mnemonic3: "two",
		Name3:     "",
	}
	fmt.Printf("Origin    : %+v\n", example)

	translate.Translate("EN", &example)

	fmt.Printf("Translated: %+v\n", example)
}
