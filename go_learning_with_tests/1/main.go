// hello world with tdd
package main

import "fmt"

const (
	spanish            = "Spanish"
	french             = "French"
	english            = "English"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := greetingPrefix(language)
	return fmt.Sprintf("%s%s!", prefix, name)
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}

func main() {
	fmt.Println(Hello("World", "English"))
}
