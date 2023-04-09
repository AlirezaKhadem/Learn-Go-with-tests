package helloworld

const (
	english = "en"
	spanish = "es"
	french  = "fr"
)

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Holla, "
	frenchHelloPrefix  = "Bonjour, "
)

var languagePrefixMapping = map[string]string{
	english: englishHelloPrefix,
	spanish: spanishHelloPrefix,
	french:  frenchHelloPrefix,
}

func Hello(
	name string,
	language string,
) string {
	helloPrefix := getHelloPrefix(language)
	if len(name) == 0 {
		name = "World"
	}

	return helloPrefix + name
}

func getHelloPrefix(language string) (helloPrefix string) {
	if _, ok := languagePrefixMapping[language]; !ok {
		helloPrefix = englishHelloPrefix
	} else {
		helloPrefix = languagePrefixMapping[language]
	}
	return helloPrefix
}
