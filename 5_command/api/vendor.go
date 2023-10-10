package api

import "fmt"

type Light struct {
	RoomName string
}

func (l *Light) on() {
	fmt.Printf("%s Light is turned on\n", l.RoomName)
}

func (l *Light) off() {
	fmt.Printf("%s Light is turned off\n", l.RoomName)
}

type CeilingFan struct {
	RoomName string
}

func (c *CeilingFan) on() {
	fmt.Printf("%s ceiling fan is turned on\n", c.RoomName)
}

func (c *CeilingFan) off() {
	fmt.Printf("%s ceiling fan is turned off\n", c.RoomName)
}

type Garage struct{}

func (g *Garage) up() {
	fmt.Println("The Garage is open")
}

func (g *Garage) lightOn() {
	fmt.Println("The Garage Lights are on")
}

func (g *Garage) down() {
	fmt.Println("The Garage is down")
}

func (g *Garage) lightOff() {
	fmt.Println("The Garage Lights are off")
}

type Stereo struct {
	RoomName string
}

func (s *Stereo) on() {
	fmt.Printf("%s Stereo is on\n", s.RoomName)
}

func (s *Stereo) setCD() {
	fmt.Printf("%s Stereo is set for CD input\n", s.RoomName)
}

func (s *Stereo) setVolume(volume int) {
	fmt.Printf("%s Stereo volume set to %d\n", s.RoomName, volume)
}

func (s *Stereo) off() {
	fmt.Printf("%s Stereo is off\n", s.RoomName)
}
