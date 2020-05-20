package mod1

import (
"fmt"
"math/rand"
"time"
)

func init () {
fmt.Println ("Modulo 1: calentado motores")
rand.Seed(time.Now().Unix())
}

func Translate(cadena string) string {

runes := []rune(cadena)
rand.Shuffle(len(runes), func(i, j int){
runes[i], runes[j] = runes [j], runes [i]
  })
  return string(runes)
}
