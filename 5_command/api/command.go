package api

type ICommand interface {
	Execute()
}

type lightOnCommand struct {
	light *Light
}

/**
 * The constructor is passed the specific
 * Light that this command is going to
 * control - say the living room Light
 * - and stashes it in the Light instance
 * variable. When Execute gets called, this
 * is the Light object that is going to be
 * the Receiver of the request.
 */
func NewLightOnCommand(l *Light) *lightOnCommand {
	return &lightOnCommand{
		light: l,
	}
}

/**
 * The execute method calls the
 * on() method on the receiving
 * object, which is the Light we
 * are controlling
 */
func (l *lightOnCommand) Execute() {
	l.light.on()
}

/**
 * The LightOffCommand works exactly
 * the same way as the LightOnCommand,
 * except that we are binding the receiver
 * to a different action: the off() method.
 */
type LightOffCommand struct {
	light *Light
}

func NewLightOffCommand(l *Light) *LightOffCommand {
	return &LightOffCommand{
		light: l,
	}
}

func (l *LightOffCommand) Execute() {
	l.light.off()
}

type GarageDoorOpenCommand struct {
	garage *Garage
}

func NewGarageDoorOpenCommand(g *Garage) *GarageDoorOpenCommand {
	return &GarageDoorOpenCommand{
		garage: g,
	}
}

func (g *GarageDoorOpenCommand) Execute() {
	g.garage.up()
	g.garage.lightOn()
}

type GarageDoorCloseCommand struct {
	garage *Garage
}

func NewGarageDoorCloseCommand(g *Garage) *GarageDoorCloseCommand {
	return &GarageDoorCloseCommand{
		garage: g,
	}
}

func (g *GarageDoorCloseCommand) Execute() {
	g.garage.lightOff()
	g.garage.down()
}

/**
 * Just like the LightOnCommand, we get
 * passed the instance of the Stereo we
 * are going to be controlling and we store
 * it in a local instance variable.
 */
type StereoOnCommand struct {
	stereo *Stereo
}

func NewStereoOnCommand(s *Stereo) *StereoOnCommand {
	return &StereoOnCommand{
		stereo: s,
	}
}

/**
 * To carry out this request, we need to call three
 * methods on the Stereo: first, turn it on, then set
 * it to play the CD, and finally set the volume to 11.
 */
func (s *StereoOnCommand) Execute() {
	s.stereo.on()
	s.stereo.setCD()
	s.stereo.setVolume(11)
}

type StereoOffCommand struct {
	stereo *Stereo
}

func NewStereoOffCommand(s *Stereo) *StereoOffCommand {
	return &StereoOffCommand{
		stereo: s,
	}
}

func (s *StereoOffCommand) Execute() {
	s.stereo.off()
}

type CeilingFanOnCommand struct {
	ceilingFan *CeilingFan
}

func NewCeilingFanOnCommand(c *CeilingFan) *CeilingFanOnCommand {
	return &CeilingFanOnCommand{
		ceilingFan: c,
	}
}

func (c *CeilingFanOnCommand) Execute() {
	c.ceilingFan.on()
}

type CeilingFanOffCommand struct {
	ceilingFan *CeilingFan
}

func NewCeilingFanOffCommand(c *CeilingFan) *CeilingFanOffCommand {
	return &CeilingFanOffCommand{
		ceilingFan: c,
	}
}

func (c *CeilingFanOffCommand) Execute() {
	c.ceilingFan.off()
}

type MacroCommand struct {
	commands []ICommand
}

func NewMacroCommand(c []ICommand) *MacroCommand {
	return &MacroCommand{
		commands: c,
	}
}

func (m *MacroCommand) addCommand(c ICommand) {
	m.commands = append(m.commands, c)
}

func (m *MacroCommand) Execute() {
	for _, command := range m.commands {
		command.Execute()
	}
}
