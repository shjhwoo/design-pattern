package singleton

import (
	"fmt"
	"sync"
)

//유의사항
/*
인스턴스 생성 전에 해당 인스턴스가 있는지 없는지를 반드시 확인해야 한다
여러 개의 고루틴이 각각 인스턴스를 생성할 위험이 있기 때문이다

생성할 떄는 lock을 걸어줘야 한다
이 이후에도 한번 더 nil 체크를 해야 하는데,
하나 이상의 고루틴이 첫번쨰 체크를 통과한 경우에도
단 하나의 고루틴만이 인스턴스를 만들어야 하기 떄문이다
*/

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}
