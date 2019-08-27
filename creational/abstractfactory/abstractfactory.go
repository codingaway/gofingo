package main

import "fmt"

type ProductA interface {
	DoA()
}
type ProductB interface {
	DoB()
}

type AbstractFactory interface {
	CreateProductA() ProductA
	CreateProductB() ProductB
}

type ConcreteFactory1 struct {
}

func (f *ConcreteFactory1) CreateProductA() ProductA {
	return &productA1{}
}

func (f *ConcreteFactory1) CreateProductB() ProductB {
	return &productB1{}
}

type ConcreteFactory2 struct {
}

func (f *ConcreteFactory2) CreateProductA() ProductA {
	return &productA2{}
}

func (f *ConcreteFactory2) CreateProductB() ProductB {
	return &productB2{}
}

type productA1 struct{}

func (p *productA1) DoA() {
	fmt.Println("ProductA1 doing A")
}

type productA2 struct{}

func (p *productA2) DoA() {
	fmt.Println("ProductA2 doing A")
}

type productB1 struct{}

func (p *productB1) DoB() {
	fmt.Println("ProductB1 doing B")
}

type productB2 struct{}

func (p *productB2) DoB() {
	fmt.Println("ProductB2 doing B")
}

type Client interface {
	ProductA
	ProductB
}

type client struct {
	prodA ProductA
	prodB ProductB
}

func (c *client) DoA() {
	c.prodA.DoA()
}
func (c *client) DoB() {
	c.prodB.DoB()
}

func NewClient(f AbstractFactory) Client {
	return &client{
		prodA: f.CreateProductA(),
		prodB: f.CreateProductB(),
	}
}
