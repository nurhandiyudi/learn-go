package main

import "fmt"

func main() {
	var angka int = 10
	angkalagi := 10
	fmt.Println(angka)
	fmt.Printf("angka pertama :  %d\n", angka)
	fmt.Println(angkalagi)
	fmt.Printf("angka kedua :  %d\n", angkalagi)
	if angka > 5 {
		fmt.Printf("lebih besar dari 5  %d\n")
	}
	for a := 0; a <= 10; a++ {
		fmt.Printf("angka urut :  %d\n", a)
	}
	var a int
	for a < 10 {
		a++
		fmt.Printf("angka urut :  %d\n", a)
	}
	numbers := [7]int{1, 2, 3, 5, 6}
	for i, x := range numbers {
		fmt.Printf("value of x = %d at %d\n", x, i)
	}

	//var x int = 10
	//var y int = 15
	var fun int = max(10, 15)
	fmt.Printf("angka terbesar %d\n", fun)

	var sum int = sum(10, 15)
	fmt.Printf("jumlah %d\n", sum)

	var kata = "Belajar Go"
	fmt.Printf("Normal String: ")
	fmt.Printf("%s \n", kata)

	var len = len(kata)
	fmt.Printf("jumlah huruf", len)

	const sampleText = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", sampleText)
	fmt.Printf("\n")

	var n [10]int
	var i, j int

	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}

	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
}

func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func sum(a, b int) int {
	return a + b
}
