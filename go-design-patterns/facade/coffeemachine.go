package main

import "fmt"

type CoffeeMachine struct {
	beanAmount   float32 // amount in gram
	grinderLevel int     // granularity of grinder
	waterTemp    int
	waterAmt     float32
	milkAmount   float32
	addFoam      bool
}

func (c *CoffeeMachine) startCoffee(beanAmount float32, grind int) {
	c.beanAmount = beanAmount
	c.grinderLevel = grind
	fmt.Println("Starting coffee order with beans", beanAmount, "and grind level of", grind)
}

func (c *CoffeeMachine) endCoffee() {
	fmt.Println("Ending coffee order")
}

func (c *CoffeeMachine) grindBeans() bool {
	fmt.Println("Grinding beans, amount:", c.beanAmount, "beans at", c.grinderLevel)
	return true
}

func (c *CoffeeMachine) useMilk(amount float32) float32 {
	fmt.Println("Adding milk:", amount, "g")
	c.milkAmount = amount
	return amount
}

func (c *CoffeeMachine) setWaterTemp(temp int) {
	fmt.Println("Setting water temp:", temp)
}

func (c *CoffeeMachine) useHotWater(amount float32) float32 {
	fmt.Println("Adding hot water:", amount)
	c.waterAmt = amount
	return amount
}

func (c *CoffeeMachine) doFoam(useFoam bool) {
	if useFoam {
		fmt.Println("Using foam!")
	} else {
		fmt.Println("Not using foam!")
	}
	c.addFoam = useFoam
}
