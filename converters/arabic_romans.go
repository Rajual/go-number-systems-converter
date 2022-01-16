package converters

import (
	"fmt"
)

/*
I->1
V->5
X->10
L->50
C->100
D->500
M->1000
*/

var ConstArabicToRomans = map[string]uint8{
	"I": 1,
}

func ArabicToRomans(number uint16) {

	fmt.Println(Int2Array(11110))
}
