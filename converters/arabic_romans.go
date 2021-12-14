package converters

import "fmt"

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
	var digitos []uint16
	for i := number; i <= 0; {
		i /= 10
		digitos = append(digitos, i)
	}
	fmt.Println(digitos)
}
