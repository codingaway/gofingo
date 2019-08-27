package main

func main() {

	carFactory := &HyundaiFactory{}
	carA, _ := carFactory.CreateCar("coupe")
	carB, _ := carFactory.CreateCar("i30")

	carA.Drive()
	carB.Drive()
}
