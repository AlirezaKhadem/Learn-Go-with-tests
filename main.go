package main

const englishLanguage = "en"
const spanishLanguage = "es"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Holla, "

var languagePrefixMapping = map[string]string{
	englishLanguage: englishHelloPrefix,
	spanishLanguage: spanishHelloPrefix,
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
