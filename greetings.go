package greetings

import "fmt"
func GreetingHello(name string) string {
    greeting := fmt.Sprintf("Hi there,  %s", name)
    return greeting
}
