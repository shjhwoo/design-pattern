package singleton

import (
	"fmt"
	"sync"
)

var once sync.Once

type single2 struct {
}

var singleInstance2 *single2

func getInstance2() *single2 {
	fmt.Println("*")
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance now._2")
				singleInstance2 = &single2{}
			})
	} else {
		fmt.Println("Single instance already created._2")
	}

	return singleInstance2
}
