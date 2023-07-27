package main

import (
	"fmt"
	"log"
)

type CComplexArray struct {
	values []complex128
	len    int
}

func (cArr *CComplexArray) Append_C(val complex128) {
	cArr.values = append(cArr.values, val)
	cArr.len++
}

func (cArr *CComplexArray) Print_CArr() {
	fmt.Printf("[")

	for i := 0; i < cArr.len; i++ {
		fmt.Printf("%v+%vi, ", real(cArr.values[i]), imag(cArr.values[i]))
	}

	fmt.Printf("\b\b]\n")
}

func (cArr *CComplexArray) Abs_C() []float64 {
	abs_c := make([]float64, cArr.len)

	for i := 0; i < cArr.len; i++ {
		tmp := real(cArr.values[i])*real(cArr.values[i]) + imag(cArr.values[i])*imag(cArr.values[i])
		abs_c = append(abs_c, tmp)
	}

	return abs_c
}

func Add_C(cArr1, cArr2 CComplexArray) CComplexArray {
	arrLen := 0

	if cArr1.len == cArr2.len {
		arrLen = cArr1.len
	} else if cArr1.len < cArr2.len {
		arrLen = cArr2.len
	} else if cArr1.len > cArr2.len {
		arrLen = cArr1.len
		cArr1, cArr2 = cArr2, cArr1
	}

	add_c := make([]complex128, arrLen)

	for i := 0; i < arrLen; i++ {
		tmp := 0 + 0i

		if i < cArr1.len {
			tmp = cArr1.values[i] + cArr2.values[i]
		} else {
			tmp = cArr2.values[i]
		}

		add_c = append(add_c, tmp)
	}

	return CComplexArray{add_c, arrLen}
}

func Scaling_C(cArr CComplexArray, val complex128) CComplexArray {
	scaled_c := make([]complex128, cArr.len)

	// Add go routines and check if speed up is possible :)

	for i := 0; i < cArr.len; i++ {
		scaled_c = append(scaled_c, cArr.values[i]*val)
	}

	return CComplexArray{scaled_c, cArr.len}
}

func Sub_C(cArr1, cArr2 CComplexArray) CComplexArray {
	return Add_C(cArr1, Scaling_C(cArr2, -1+0i))
}

func Multiply_El_C(cArr1, cArr2 CComplexArray) CComplexArray {
	if cArr1.len != cArr2.len {
		log.Fatalln("Error: Input complex arrays have different lengths!")
	}

	mul_el_c := make([]complex128, cArr1.len)

	for i := 0; i < cArr1.len; i++ {
		mul_el_c = append(mul_el_c, cArr1.values[i]*cArr2.values[i])
	}

	return CComplexArray{mul_el_c, cArr1.len}
}

func DotProduct_El_C(cArr1, cArr2 CComplexArray) complex128 {
	if cArr1.len != cArr2.len {
		log.Fatalln("Error: Input complex arrays have different lengths!")
	}

	dotProduct_c := 0 + 0i

	for i := 0; i < cArr1.len; i++ {
		dotProduct_c = dotProduct_c + cArr1.values[i]*cArr2.values[i]
	}

	return dotProduct_c
}

func main() {
	var cArr *CComplexArray = &CComplexArray{[]complex128{}, 0}

	cArr.Append_C(1 + 2i)
	cArr.Append_C(2 + 3i)
	cArr.Append_C(3 + 4i)
	cArr.Append_C(4 + 5i)
	cArr.Append_C(5 + 6i)

	cArr.Print_CArr()

	fmt.Println(cArr.Abs_C())
}
