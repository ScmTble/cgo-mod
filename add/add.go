package add

//#include "number.h"
import "C"

func Add(a, b, mod int) int {
	return int(C.number_add_mod(10, 5, 12))
}
