// Package human defines a person
package human

// Human general struct
type Human struct {
	name     string // HL
	sentence string // HL
}

// MyName returns the person's name
func (h *Human) MyName() string {
	return h.name
}

// Speak now or forever hold your peace
func (h *Human) Speak() string {
	return h.sentence
}

// SetName sets the name
func (h *Human) SetName(name string) {
	h.name = name
}

// SetSentence sets the sentence
func (h *Human) SetSentence(sentence string) {
	h.sentence = sentence
}
