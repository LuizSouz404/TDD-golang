package main

func main() {}

const english = "english"
const french = "french"
const prefixGreetingInPortuguese = "Olá, "
const prefixGreetingInEnglish = "Hello, "
const prefixGreetingInFrench = "Bonjour, "

func Greeting(name string, language string) string {
	if name == "" {
		name = "mundo"
	}

	return prefixLanguage(language) + name
}

func prefixLanguage(language string) (prefix string) {
	switch language {
	case english:
		prefix = prefixGreetingInEnglish
	case french:
		prefix = prefixGreetingInFrench
	default:
		prefix = prefixGreetingInPortuguese
	}

	return
}
