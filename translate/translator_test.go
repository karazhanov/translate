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
}
