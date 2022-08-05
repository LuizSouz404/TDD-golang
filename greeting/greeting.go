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

	prefixGreeting := prefixGreetingInPortuguese

	switch language {
	case english:
		prefixGreeting = prefixGreetingInEnglish
	case french:
		prefixGreeting = prefixGreetingInFrench
	}

	return prefixGreeting + name
}
