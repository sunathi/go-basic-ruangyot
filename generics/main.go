package main

import "fmt"

func main(){
	numsInt := []int{1,2,3}
	numsFloat64 := []float64{1.1,2.2,3.3}

	sumInt := sum(numsInt)
	sumFloat := sum(numsFloat64)
	fmt.Println(sumInt)
	fmt.Println(sumFloat)
}


// type Number interface {
// 	int | float64
// }

type (
	Number interface {
	int | float64
	}
	Player[T Number] struct{
		Name string
		Age int
		Damage T
	}
	Database[T Number] interface {

	}
)

func sum[T Number](nums []T) T {
	var sum T
	for _, n := range nums {
		sum += n
	}
	return sum
}