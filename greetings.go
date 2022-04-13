package greetings

import "fmt"
func GreetingHello(name string) string {
    greeting := fmt.Sprintf("Hello %s", name)
    return greeting
}
