// Написать свой интерфейс и создать несколько структур, удовлетворяющих ему.

package main

import "fmt"

type Vehicle interface {
	Speed() float32
	Mileage(t uint32) float32
	Stop()
}

type bike struct {
	speed float32
}

func (b *bike) Speed() float32 {
	return b.speed
}

func (b *bike) Mileage(t uint32) float32 {
	return b.speed * float32(t)
}

func (b *bike) Stop() {
	b.speed = 0
}

type car struct {
	speed float32
}

func (c *car) Speed() float32 {
	return c.speed
}

func (c *car) Mileage(t uint32) float32 {
	return c.speed * float32(t)
}

func (c *car) Stop() {
	c.speed = 0
}

func Mileage(v Vehicle, t uint32) float32 {
	return v.Mileage(t)
}

func main() {
	b := &bike{
		speed: 10.0,
	}
	c := &car{
		speed: 20.0,
	}
	fmt.Println("Bike mileage", Mileage(b, 10))
	fmt.Println("Car mileage", Mileage(c, 20))
}
