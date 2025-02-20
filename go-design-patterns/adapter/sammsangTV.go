package main

import "fmt"

type SammsangTV struct {
	currentChan   int
	currentVolume int
	tvOn          bool
}

func (tv *SammsangTV) getVolume() int {
	fmt.Println("SammysangTV volume is", tv.currentVolume)
	return tv.currentVolume
}

func (tv *SammsangTV) setVolume(vol int) {
	fmt.Println("Setting SammsangTV volume to", vol)

	tv.currentVolume = vol
}

func (tv *SammsangTV) setOnState(tvOn bool) {

	if tvOn {
		fmt.Println("SammsangTV is on")
	} else {
		fmt.Println("SammsangTV is off")
	}

	tv.tvOn = tvOn
}
