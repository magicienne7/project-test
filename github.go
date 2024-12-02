package main

import "fmt"

func main() {
number := 2
power := 3
result := 1
for i := 0; i < power; i++ {
result *= number
}
fmt.Println(result)
}
