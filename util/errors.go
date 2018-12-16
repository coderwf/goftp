package util

import "fmt"

func PanicToError(fn func()) (err error){
	defer func() {
		if r := recover() ; r != nil{
			err = fmt.Errorf("Panic:%v",err)
		}//if
	}()

	fn()
	return
}
