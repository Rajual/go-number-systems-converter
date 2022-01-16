package converters

import (
	"fmt"
	"math"
	"time"
)

func Int2Array(number int) []int {
	_time := time.Now()
	var arrayAux []int
	for i := 10; i <= number*10; i *= 10 {
		_, d := math.Modf(float64(number) / float64(i))
		arrayAux = append(arrayAux, int(math.Round(d*10)))
	}
	var array []int
	for i := len(arrayAux) - 1; i >= 0; i-- {
		array = append(array, arrayAux[i]*int(math.Pow10(i)))
	}

	fmt.Println(time.Since(_time))
	return array
}
