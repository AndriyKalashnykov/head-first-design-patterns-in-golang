package api

import "fmt"

/**
 * Here is the composition;
 * these are all the components of the
 * subsystem we are going to use
 */
type HomeTheaterFacade struct {
	Amp    *Amplifier
	Dvd    *DvdPlayer
	Proj   *Projector
	Lights *TheaterLights
	Screen *Screen
	Popper *PopcornPopper
}

/**
 * Instantiating all the components of the subsystem
 */
func NewHomeTheaterFacade() *HomeTheaterFacade {
	return &HomeTheaterFacade{
		Amp:    &Amplifier{},
		Dvd:    &DvdPlayer{},
		Proj:   &Projector{},
		Lights: &TheaterLights{},
		Screen: &Screen{},
		Popper: &PopcornPopper{},
	}
}

/**
 * WatchMovie() follows the same sequence
 * we had to do by hand before, but wraps
 * it Up in a handy method that does all
 * the work. Notice that for each task we
 * are delegating the responsibility to the
 * corresponding component in the subsystem.
 */
func (h *HomeTheaterFacade) WatchMovie(movie string) {
	fmt.Println("\nGet ready to watch a movie...")
	h.Popper.On()
	h.Popper.Pop()
	h.Lights.Dim(10)
	h.Screen.Down()
	h.Proj.On()
	h.Proj.WideScreenMode()
	h.Amp.On()
	h.Amp.SetSurroundSound()
	h.Amp.SetVolume(5)
	h.Dvd.On()
	h.Dvd.Play(movie)
}

/**
 * EndMovie() takes care
 * of shutting everything Down
 * for us. Again, each task is
 * delegated to the appropriate
 * component in the subsystem.
 */
func (h *HomeTheaterFacade) EndMovie() {
	fmt.Println("\nShutting movie theater Down...")
	h.Popper.Off()
	h.Lights.On()
	h.Screen.Up()
	h.Proj.Off()
	h.Amp.Off()
	h.Dvd.Stop()
	h.Dvd.Eject()
	h.Dvd.Off()
}
