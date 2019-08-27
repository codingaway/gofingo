package main

import "fmt"

type Car interface {
	Drive()
}

type CarFactory interface {
	CreateCar(string) (Car, error)
}

type MazdaCarFactory struct{}

func (m *MazdaCarFactory) CreateCar(model string) (Car, error) {
	switch model {
	case "mx5":
		return &MazdaMX5{}, nil
	case "6":
		return &Mazda6{}, nil
	default:
		return nil, fmt.Errorf("invalid car model for Mazda")
	}
}

type MazdaMX5 struct {
}

func (m *MazdaMX5) Drive() {
	fmt.Println("driving a Mazda MX5")
}

type Mazda6 struct {
}

func (m *Mazda6) Drive() {
	fmt.Println("driving a Mazda 6")
}

type HyundaiFactory struct{}

func (m *HyundaiFactory) CreateCar(model string) (Car, error) {
	switch model {
	case "coupe":
		return &HyundaiCoupe{}, nil
	case "i30":
		return &HyundaiI30{}, nil
	default:
		return nil, fmt.Errorf("invalid car model Hyundai")
	}
}

type HyundaiCoupe struct {
}

func (m *HyundaiCoupe) Drive() {
	fmt.Println("driving a Hyundai Coupe")
}

type HyundaiI30 struct {
}

func (m *HyundaiI30) Drive() {
	fmt.Println("driving a Hyundai i30")
}
