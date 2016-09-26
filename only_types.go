// START MAIN OMIT
package main

import (
	"fmt"

	"github.com/tux-eithel/linuxday2016-golang/human"
	"github.com/tux-eithel/linuxday2016-golang/pets"
)

type MyCustomSpeak string

func (m MyCustomSpeak) MyName() string {
	return string(m)
}

func (m MyCustomSpeak) Speak() string {
	return "..."
}

func main() {

	billy := &human.Human{}
	billy.SetName("Billy")
	billy.SetSentence("I've so many things to say...")

	matilde := &pets.Duck{
		pets.Pets{"Matilde"},
		"Duck Duck",
	}

	SaySomething(billy)
	SaySomething(matilde)
	SaySomething(MyCustomSpeak("hello"))

}

// END MAIN OMIT

func SaySomething(petsOrHuman interface{}) { // HL

	switch talk := petsOrHuman.(type) {

	case *human.Human:
		fmt.Println(talk.MyName() + " - " + talk.Speak())

	case *pets.Duck:
		fmt.Println(talk.MyName() + " - " + talk.Speak())

	default:
		panic("not supported")
	}
}
