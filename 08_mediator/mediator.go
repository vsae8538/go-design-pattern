package mediator

import (
	"fmt"
	"strings"
)

// CDDriver represents a CD driver that reads data from a CD.
type CDDriver struct {
	Data string
}

// ReadData reads data from the CD and notifies the mediator about the change.
func (c *CDDriver) ReadData() {
	c.Data = "music,image"

	fmt.Printf("CDDriver: reading data %s\n", c.Data)
	GetMediatorInstance().changed(c)
}

// CPU represents a central processing unit.
type CPU struct {
	Video string
	Sound string
}

// Process receives data from the mediator, processes it, and notifies the mediator about the change.
func (c *CPU) Process(data string) {
	sp := strings.Split(data, ",")
	c.Sound = sp[0]
	c.Video = sp[1]

	fmt.Printf("CPU: split data with Sound %s, Video %s\n", c.Sound, c.Video)
	GetMediatorInstance().changed(c)
}

// VideoCard represents a video card that displays video data.
type VideoCard struct {
	Data string
}

// Display receives video data from the mediator and displays it.
func (v *VideoCard) Display(data string) {
	v.Data = data
	fmt.Printf("VideoCard: display %s\n", v.Data)
	GetMediatorInstance().changed(v)
}

// SoundCard represents a sound card that plays audio data.
type SoundCard struct {
	Data string
}

// Play receives audio data from the mediator and plays it.
func (s *SoundCard) Play(data string) {
	s.Data = data
	fmt.Printf("SoundCard: play %s\n", s.Data)
	GetMediatorInstance().changed(s)
}

// Mediator manages communication between different components.
type Mediator struct {
	CD    *CDDriver
	CPU   *CPU
	Video *VideoCard
	Sound *SoundCard
}

var mediator *Mediator

// GetMediatorInstance returns the singleton instance of the Mediator.
func GetMediatorInstance() *Mediator {
	if mediator == nil {
		mediator = &Mediator{}
	}
	return mediator
}

// changed is called by components to notify the mediator about changes and trigger appropriate actions.
func (m *Mediator) changed(i interface{}) {
	switch inst := i.(type) {
	case *CDDriver:
		m.CPU.Process(inst.Data)
	case *CPU:
		m.Sound.Play(inst.Sound)
		m.Video.Display(inst.Video)
	}
}
