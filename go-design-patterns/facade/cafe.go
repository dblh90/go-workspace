package main

import "fmt"

func makeAmericano(size float32) {
	fmt.Println("\nMaking an Americano\n......")
	americano := &CoffeeMachine{}

	beanAmount := (size / 8.0) * 5.0
	americano.startCoffee(beanAmount, 2)
	americano.grindBeans()
	americano.useHotWater(size)
	americano.endCoffee()

	fmt.Println("Americano is ready!")
}

func makeLatee(size float32, foam bool) {
	fmt.Println("\nMaking a Latee\n......")
	latte := &CoffeeMachine{}
	beansAmount := (size / 8.0) * 5.0
	latte.startCoffee(beansAmount, 4)
	latte.grindBeans()
	latte.useHotWater(size)

	milkAmt := (size / 8.0) * 2.0
	latte.useMilk(milkAmt)
	latte.doFoam(foam)
	latte.endCoffee()
	fmt.Println("Latee is ready!")
}
