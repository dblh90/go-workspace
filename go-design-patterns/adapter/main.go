package main

import "fmt"

func main() {
	tv1 := &SammsangTV{
		currentChan:   13,
		currentVolume: 35,
		tvOn:          true,
	}

	tv2 := &SohneeTV{
		vol:     20,
		channel: 9,
		isOn:    true,
	}

	tv2.turnOn()
	tv2.volumeUp()
	tv2.volumeDown()
	tv2.channelUp()
	tv2.channelDown()
	tv2.goToChannel(23)
	tv2.turnOff()

	fmt.Println("-----------------------------------")

	ssAdapt := &sammSangAdapter{
		sstv: tv1,
	}

	ssAdapt.turnOn()
	ssAdapt.volumeUp()
	ssAdapt.volumeDown()
	ssAdapt.channelUp()
	ssAdapt.channelDown()
	ssAdapt.goToChannel(68)
	ssAdapt.turnOff()
}
