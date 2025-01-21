package main

import "fmt"

type Singleton struct{}

var instance *Singleton

func GetInstance() *Singleton {
	if instance == nil {
		instance = &Singleton{}
		fmt.Println("Singleton created")
	} else {
		fmt.Println("Singleton already created")
	}
	return instance
}
