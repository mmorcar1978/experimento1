package main

import "fmt"
import "github.com/mmorcar/experimento1/mod1"

func init () {

	fmt.Println("Ready to launch...")
}

func main () {

saludo:= "Hola mundo!!"
fmt.Println(saludo)
fmt.Println(mod1.Translate(saludo))
}
