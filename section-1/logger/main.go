package main

import "fmt"

type LogLevel int

const(
	LevelTrace LogLevel = iota
	LevelDebug
	LevelInfo
	LevelWarning
	LevelError
)

var levelNames = []string{"Trace", "Debug", "Info", "Warning", "Error"}

// This is the same as kotlin's extension function, though
// String() is special because it implements the Stringer interface from the fmt package. 
// When fmt.Println, fmt.Printf, etc. need to print a value, they automatically check 
// if it has a String() method and use it. With any other name, you lose this automatic behavior.

// For method that are going to be exported (visible on outside package), it uses PascalCase
func (l LogLevel) String() string {
	if l < LevelTrace || l > LevelError {
		return "Unknown"
	}

	return levelNames[l]
}

// private unexported method uses camelCase
func printLogLevel(level LogLevel) {
	fmt.Printf("Log level: %d %s\n", level, level)
}

func main() {
	printLogLevel(LevelTrace)
	printLogLevel(LevelDebug)
	printLogLevel(LevelError)
	printLogLevel(LevelInfo)
	printLogLevel(LevelWarning)
}