package main

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if len(name) == 0 {
		return englishHelloPrefix + "World"
	}
	return string(englishHelloPrefix + name)
}
