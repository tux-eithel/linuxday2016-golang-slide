package main

import (
	"fmt"

	"github.com/tux-eithel/linuxday2016-golang-slide/human"
	"github.com/tux-eithel/linuxday2016-golang-slide/pets"
)

type MyCustomSpeak string

func (m MyCustomSpeak) MyName() string {
	return string(m)
}

func (m MyCustomSpeak) Speak() string {
	return "..."
}

// START MAIN OMIT
func main() {

	billy := &human.Human{}
	billy.SetName("Billy")
	billy.SetSentence("I've so many things to say...")

	matilde := &pets.Duck{pets.Pets{"Matilde"}, "Duck Duck"}

	SaySomethingSpeaker(billy)
	SaySomethingSpeaker(matilde)
	SaySomethingSpeaker(MyCustomSpeak("hello"))

}

// END MAIN OMIT

// START I OMIT
type speaker interface {
	MyName() string
	Speak() string
}

func SaySomethingSpeaker(aSpeaker speaker) { // HL
	fmt.Println(aSpeaker.MyName() + " - " + aSpeaker.Speak())
}

// END I OMIT
