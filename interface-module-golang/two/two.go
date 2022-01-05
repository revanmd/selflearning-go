package main

import "fmt"

type Persegi struct {
	panjang, lebar float64
}

func (p *Persegi) Init(panjang float64, lebar float64) {
	if p.panjang < 0 || p.lebar < 0 {
		fmt.Println("Value dibawah 0")

		p.panjang = 0
		p.lebar = 0
	} else {
		p.panjang = panjang
		p.lebar = lebar
	}
}

func (p *Persegi) IsType() string {
	if p.panjang == p.lebar {
		return "Bujur Sangkar"
	} else {
		return "Persegi Panjang"
	}
}

func main() {
	var panjang float64
	var lebar float64

	fmt.Print("Masukan panjang : ")
	fmt.Scanln(&panjang)

	fmt.Print("Masukan lebar : ")
	fmt.Scanln(&lebar)

	persegi := new(Persegi)
	persegi.Init(panjang, lebar)

	fmt.Println("Tipe persegi adalah : ", persegi.IsType())

}
