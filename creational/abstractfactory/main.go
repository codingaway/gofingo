package main

func main() {
	client1 := NewClient(&ConcreteFactory1{})
	client1.DoA()
	client1.DoB()
	client2 := NewClient(&ConcreteFactory2{})
	client2.DoA()
	client2.DoB()
}
