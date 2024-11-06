package main

import "fmt"

type Currency interface {
	RupeeConvert() float64
	DollarConvert() float64
}

type Rupee struct {
	amount           float64
	devaluation_rate float64
}

type Dollar struct {
	amount         float64
	increment_rate float64
}

func (r Rupee) RupeeConvert() float64 {
	return (r.amount * 80.2) / r.devaluation_rate
}

func (d Dollar) DollarConvert() float64 {
	return (d.amount / 80.2) * d.increment_rate
}

func getResults(i Currency) {
	fmt.Printf("", i.DollarConvert())
}

func main() {

}