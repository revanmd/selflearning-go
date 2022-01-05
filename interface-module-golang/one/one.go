package main

import "fmt"

type Persegi struct {
	panjang, lebar float64
}

func (p *Persegi) Validate() {
	if p.panjang < 0 || p.lebar < 0 {
		fmt.Println("Value dibawah 0")

		p.panjang = 0
		p.lebar = 0
	}
}

func (p *Persegi) Luas() float64 {
	return (p.panjang * p.lebar)
}

func main() {
	persegi := Persegi{100, 2}
	persegi.Validate()

	fmt.Println("Luas persegi adalah : ", persegi.Luas())

}
