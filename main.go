package main

import goroutine "main-mode/pkg/14_goroutine_15_WaitGroups"

// Выполняется самая первая
func init() {
	var initMsg string = "Init - выполняется быстрее main\n=======================\n"
	println(initMsg)
}

// Функция main - сама по себе как горутина
func main() {
	goroutine.Main()
}
