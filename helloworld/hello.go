package helloworld

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Holla, "
	frenchHelloPrefix  = "Bonjour, "

	english = "en"
	spanish = "sp"
	french  = "fn"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	return getGreetingPrefix(language) + name + "!"
}

var languagePrefixMap = map[string]string{
	"en": englishHelloPrefix,
	"sp": spanishHelloPrefix,
	"fn": frenchHelloPrefix,
}

func getGreetingPrefix(language string) string {
	if _, found := languagePrefixMap[language]; !found {
		language = english
	}
	return languagePrefixMap[language]
}
