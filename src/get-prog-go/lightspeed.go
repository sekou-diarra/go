package main

import "fmt"

func main(){
	const lightSpeed = 299792 // km/S
	var distance = 56000000

	fmt.Println(distance/lightSpeed, "seconds")

	distance = 401000000
	fmt.Println(distance/lightSpeed, "seconds")
}