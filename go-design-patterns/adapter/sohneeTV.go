package main

import "fmt"

type SohneeTV struct {
	vol     int
	channel int
	isOn    bool
}

func (s *SohneeTV) turnOn() {
	fmt.Println("SohneeTV is now on")
	s.isOn = true
}

func (s *SohneeTV) turnOff() {
	fmt.Println("SohneeTV is now off")
	s.isOn = false
}

func (s *SohneeTV) volumeUp() int {
	s.vol++
	fmt.Println("Increasing SohneeTV volume to", s.vol)
	return s.vol
}

func (s *SohneeTV) volumeDown() int {
	s.vol--
	fmt.Println("Decreasing SohneeTV volume to", s.vol)
	return s.vol
}

func (s *SohneeTV) channelUp() int {
	s.channel++
	fmt.Println("Decreasing SohneeTV channel to", s.channel)
	return s.channel
}

func (s *SohneeTV) channelDown() int {
	s.channel--
	fmt.Println("Increasing SohneeTV channel to", s.channel)
	return s.channel
}

func (s *SohneeTV) goToChannel(cha int) {
	fmt.Println("Setting channel to", cha)
	s.channel = cha
}
