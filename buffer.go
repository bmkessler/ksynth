package ksynth

import "math/rand"

// KSBuffer is a linked-list implementation of a circular buffer
// for use in Karplus-Strong synthesis
type KSBuffer struct {
	value float64
	next  *KSBuffer
}

// Decay is how much each sample decays during play
var Decay = 0.996

// NewKSBuffer initializes the buffer of size with random noise
func NewKSBuffer(size uint) *KSBuffer {
	// get a random number between -1.0 and 1.0
	initialNode := KSBuffer{value: rand.Float64()*2.0 - 1.0}
	currentNode := initialNode
	for i := uint(0); i < size; i++ {
		newNode := KSBuffer{value: rand.Float64()*2.0 - 1.0}
		currentNode.next = &newNode
		currentNode = newNode
	}
	currentNode.next = &initialNode // point back to the beginning
	return &currentNode
}

// Value returns the value of the current member
func (buffer KSBuffer) Value() float64 {
	return buffer.value
}

// Next returns the next member of the buffer
func (buffer KSBuffer) Next() KSBuffer {
	return *buffer.next
}

// Update sets the value of the current member to the average of it
// and the next value multiplied by the decay factor
func (buffer *KSBuffer) Update() {
	buffer.value = Decay * (buffer.value + buffer.next.value) / 2.0
}

// Insert places a new node into the buffer between the current and next node
// the value is an average of the two nodes to interpolate them
func (buffer *KSBuffer) Insert() {
	newNode := KSBuffer{
		value: (buffer.value + buffer.next.value) / 2.0,
		next:  buffer.next,
	}
	buffer.next = &newNode
}

// Delete removes the next node from the buffer
func (buffer *KSBuffer) Delete() {
	buffer.next = buffer.next.next
}
