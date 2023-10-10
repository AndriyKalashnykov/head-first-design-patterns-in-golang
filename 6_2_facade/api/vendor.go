package api

import "fmt"

type Amplifier struct{}

func (a *Amplifier) On() {
	fmt.Println("Top-O-Line Amplifier On")
}

func (a *Amplifier) SetSurroundSound() {
	fmt.Println("Top-O-Line Amplifier surround sound On (5 speakers, 1 subwoofer)")
}

func (a *Amplifier) SetVolume(n int) {
	fmt.Printf("Top-O-Line Amplifier setting volume to %d\n", n)
}

func (a *Amplifier) Off() {
	fmt.Println("Top-O-Line Amplifier Off")
}

type DvdPlayer struct{}

func (d *DvdPlayer) On() {
	fmt.Println("Top-O-Line DVD Player On")
}

func (d *DvdPlayer) Play(movie string) {
	fmt.Printf("Top-O-Line DVD Player playing \"%s\"\n", movie)
}

func (d *DvdPlayer) Stop() {
	fmt.Println("Top-O-Line DVD Player stopped")
}

func (d *DvdPlayer) Eject() {
	fmt.Println("Top-O-Line DVD Player Eject")
}

func (d *DvdPlayer) Off() {
	fmt.Println("Top-O-Line DVD Player Off")
}

type Projector struct{}

func (p *Projector) On() {
	fmt.Println("Top-O-Line Projector On")
}

func (p *Projector) WideScreenMode() {
	fmt.Println("Top-O-Line Projector in widescreen mode (16x9 aspect ratio)")
}

func (p *Projector) Off() {
	fmt.Println("Top-O-Line Projector Off")
}

type TheaterLights struct{}

func (t *TheaterLights) Dim(n int) {
	fmt.Printf("Theater Ceiling Lights dimming to %d%%\n", n)
}

func (t *TheaterLights) On() {
	fmt.Println("Theater Ceiling Lights On")
}

type Screen struct{}

func (s *Screen) Down() {
	fmt.Println("Theater Screen going Down")
}

func (s *Screen) Up() {
	fmt.Println("Theater Screen going Up")
}

type PopcornPopper struct{}

func (p *PopcornPopper) On() {
	fmt.Println("Popcorn Popper On")
}

func (p *PopcornPopper) Pop() {
	fmt.Println("Popcorn Popper popping popcorn!")
}

func (p *PopcornPopper) Off() {
	fmt.Println("Popcorn Popper Off")
}
