package main

import (
	"database/sql"
	"fmt"
)

type I1 interface {
	Ha()
}
type I2 interface {
	He()
}
type Dog struct {
	Name string
}

func (d Dog) Ha() {
	fmt.Println("Ha")
}

func (d Dog) He() {
	fmt.Println("He")
}

func main() {
	var i I1
	i = Dog{}
	i.Ha()
	ii := i.(I2)
	ii.He()
	db, err := sql.Open("mysql", "")
	if nil != err {

	}
	db.SetConnMaxLifetime(1)
	//s := make([]int64, 0, 10)
	//fmt.Println(fmt.Sprintf("%v", s))
	//s[0] = 100
	//s = append(s, 10)
	//s = append(s, 11)
	//s = append(s, 12)
	//fmt.Println(fmt.Sprintf("%v", s))

}
