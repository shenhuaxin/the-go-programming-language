package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g℃", c)
}

func main() {
	fmt.Printf("%g\n", BoilingC-FreezingC)
	boilingF1 := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF1-CToF(FreezingC))

	//fmt.Printf("%g\n", boilingF1-FreezingC)

	c := FToC(212)
	fmt.Println(c.String())

	fmt.Printf("%v\n", c)   // "100°C"; no need to call String explicitly
	fmt.Printf("%s\n", c)   // "100°C"
	fmt.Println(c)          // "100°C"
	fmt.Printf("%g\n", c)   // "100"; does not call String
	fmt.Println(float64(c)) // "100"; does not call String

}
