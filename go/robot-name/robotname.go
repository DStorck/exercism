package main

// package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

type Robot string

var robotNames = make(map[*Robot]string)

func (r *Robot) Name() string {
	if robotNames[r] != "" {
		return robotNames[r]
	}
	randNums := randNumber()
	randLetters := letter() + letter()
	temp_name := randLetters + randNums
	robotNames[r] = temp_name
	fmt.Println(robotNames)
	return temp_name
}

func randNumber() string {
	return fmt.Sprintf("%03d", rand.Intn(1000))
}

func (r *Robot) Reset() {
	robotNames[r] = ""
}

func letter() string {
	return string(rand.Intn('Z'-'A') + 'A')
}

func main() {
	rand.Seed(time.Now().UnixNano())
	tim := Robot{}
	fmt.Println(tim)
	a := tim.Name()
	fmt.Println(a)
	b := tim.Name()
	fmt.Println(b)
	tom := Robot{}
	c := tom.Name()
	fmt.Println(c)
	fmt.Println(robotNames)
	fmt.Println(letter())
}
