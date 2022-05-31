package translate

import "reflect"

type Translator struct {
	translator Localizator
	tagname    string
}

func New(
	translator Localizator,
	tagname string) *Translator {
	return &Translator{
		translator: translator,
		tagname:    tagname,
	}
}

func (t *Translator) TranslateMnemo(lang string, mnemo string) string {
	return t.translator.Translate(lang, mnemo)
}

func (t *Translator) Translate(lang string, translate any) {
	value := reflect.ValueOf(translate)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	if value.Kind() == reflect.Slice {
		for i := 0; i < value.Len(); i++ {
			t.translate(lang, value.Index(i))
		}
	} else {
		t.translate(lang, value)
	}
}

func (t *Translator) translate(lang string, val reflect.Value) {
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		switch field.Kind() {
		case reflect.Struct:
			t.translate(lang, field)
		case reflect.Slice:
			for j := 0; j < field.Len(); j++ {
				t.translate(lang, field.Index(j))
			}
		case reflect.String:
			if tag, find := val.Type().Field(i).Tag.Lookup(t.tagname); find {
				translatedField := val.FieldByName(tag)
				if translatedField != (reflect.Value{}) {
					translatedField.SetString(t.TranslateMnemo(lang, field.String()))
				}
			}
		}
	}
}
