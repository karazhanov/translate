package translate_test

import (
	"testing"

	"github.com/karazhanov/translate/example"
	"github.com/karazhanov/translate/translate"
	"github.com/stretchr/testify/assert"
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

type Fn func()

func TestTranslate(t *testing.T) {
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

	assert.Equal(t, "tmp", example.Mnemonic)
	assert.Equal(t, "", example.Name)
	assert.Equal(t, "one", example.Mnemonic2)
	assert.Equal(t, "", example.Name2)
	assert.Equal(t, "two", example.Mnemonic3)
	assert.Equal(t, "", example.Name3)

	translate.Translate("EN", &example)

	assert.Equal(t, "tmp", example.Mnemonic)
	assert.Equal(t, "temporary", example.Name)
	assert.Equal(t, "one", example.Mnemonic2)
	assert.Equal(t, "Number one", example.Name2)
	assert.Equal(t, "two", example.Mnemonic3)
	assert.Equal(t, "2", example.Name3)

	f := func() {}
	translate.Translate("EN", &f)

	s := ""
	translate.Translate("EN", &s)

	fs := make([]Fn, 1)
	fs[0] = func() {}
	translate.Translate("EN", &fs)

	ss := make([]string, 1)
	ss[0] = ""
	translate.Translate("EN", &ss)
}
