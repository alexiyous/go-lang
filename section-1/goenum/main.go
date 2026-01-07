package main

import "fmt"

// default enum
const (
	Sunday = iota
	Monday
	Tuesday
)

// Custom typed enum
type LogLevel int
const (
	LogError = iota
	LogWarn
	LogInfo
	LogDebug
	LogFatal
)


func main()  {
	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
}