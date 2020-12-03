package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const japan = "Japan"

const japanHelloPrefix = "こんにちは,"    //日语你好
const frenchHelloPrefix = "Bonjour," //法语你好
const spanishHelloPrefix = "Hola,"   //西班牙语你好
const defaultHelloPrefix = "hello,"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case japan:
		prefix = japanHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = defaultHelloPrefix
	}
	return
}
func main() {
	fmt.Print(Hello("", ""))
}
