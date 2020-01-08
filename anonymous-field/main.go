package main

import "fmt"

type Discount struct {
	percent float32
}

func (d Discount) Calculate(originalPrice float32) float32 {
	return originalPrice * d.percent
}

type PremiumDiscount struct {
	Discount
	additional float32
}

func (d *PremiumDiscount) Calculate(originalPrice float32) float32 {
	return d.Discount.Calculate(originalPrice) - d.additional
}

func main() {
	var originalPrice float32 = 400.0

	discount := Discount{
		percent: 0.8,
	}

	discountedPrice := discount.Calculate(originalPrice)

	fmt.Printf("Price after discount: %f\n", discountedPrice)

	premiumDiscount := &PremiumDiscount{
		Discount:   discount,
		additional: 40,
	}

	premiumPrice := premiumDiscount.Calculate(originalPrice)
	fmt.Printf("Price after premium discount: %f\n", premiumPrice)
}
