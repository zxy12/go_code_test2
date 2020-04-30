package number

//#cgo CFLAGS: -I../lib2/number1
//#cgo LDFLAGS: -L../lib2/number1 -lnumber
//
//#include "number.h"
import "C"

func Num() int {
	ret := C.number_add_mod(10,
		5, 12)
	return int(ret)
}
