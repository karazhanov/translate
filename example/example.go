package example

type Phrases = map[string]string

type LocalizatorInMemory struct {
	memory map[string]Phrases
}

func Stub() *LocalizatorInMemory {
	memoryStub := make(map[string]Phrases)
	langEN := make(Phrases)
	langEN["tmp"] = "temporary"
	langEN["one"] = "Number one"
	langEN["two"] = "2"
	memoryStub["EN"] = langEN
	return &LocalizatorInMemory{
		memory: memoryStub,
	}
}

func (l *LocalizatorInMemory) Translate(lang string, mnemo string) string {
	if translateByLang, foundLang := l.memory[lang]; foundLang {
		if translateByMnemo, foundMnemo := translateByLang[mnemo]; foundMnemo {
			return translateByMnemo
		}
	}
	return mnemo
}
