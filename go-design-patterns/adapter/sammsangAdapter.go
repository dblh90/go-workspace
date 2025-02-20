package main

type sammSangAdapter struct {
	sstv *SammsangTV
}

func (ss *sammSangAdapter) volumeUp() int {
	ss.sstv.currentVolume++
	return ss.sstv.getVolume()
}

func (ss *sammSangAdapter) volumeDown() int {
	ss.sstv.currentVolume--
	return ss.sstv.getVolume()
}

func (ss *sammSangAdapter) channelUp() int {
	ss.sstv.currentChan--
	return ss.sstv.currentChan
}

func (ss *sammSangAdapter) channelDown() int {
	ss.sstv.currentChan++
	return ss.sstv.currentChan
}

func (ss *sammSangAdapter) turnOn() {
	ss.sstv.setOnState(true)
}

func (ss *sammSangAdapter) turnOff() {
	ss.sstv.setOnState(false)
}

func (ss *sammSangAdapter) goToChannel(ch int) {
	ss.sstv.currentChan = ch
}
