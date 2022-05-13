package translate

type Localizator interface {
	Translate(lang string, mnemo string) string
}
