package ksynth

import (
	"math"
	"testing"
)

func TestNewKSBuffer(t *testing.T) {
	bufferSize := uint(10)
	buffer := NewKSBuffer(bufferSize)
	// check that it is a circular buffer
	initialNode := buffer
	buffer = buffer.Next() // take an initial step
	i := uint(1)
	for ; buffer != initialNode; buffer = buffer.Next() {
		if math.Abs(buffer.Value()) > 1.0 {
			t.Errorf("Buffer contains value %f larger than 1.0 in absolute value", buffer.Value())
		}
		i++ // one step around the buffer
	}
	if i != bufferSize {
		t.Errorf("Expected buffer to repeat after %d steps not %d", bufferSize, i)
	}
}

func TestUpdate(t *testing.T) {
	bufferSize := uint(10)
	buffer := NewKSBuffer(bufferSize)
	// go around the whole buffer once
	initialNode := buffer
	for ; buffer != initialNode; buffer = buffer.Next() {
		expectedValue := Decay * (buffer.value + buffer.next.value) / 2.0
		buffer.Update()
		if buffer.Value() != expectedValue {
			t.Errorf("Expected buffer to update to %f but got %f", buffer.Value(), expectedValue)
		}
	}
}

func TestInsert(t *testing.T) {
	bufferSize := uint(10)
	buffer := NewKSBuffer(bufferSize)
	buffer.Insert()
	// check that it is a circular buffer
	initialNode := buffer
	buffer = buffer.Next() // take an initial step
	i := uint(1)
	for ; buffer != initialNode; buffer = buffer.Next() {
		i++ // one step around the buffer
	}
	if i != bufferSize+1 {
		t.Errorf("Expected buffer to repeat after %d steps not %d", bufferSize, i)
	}
}

func TestDelete(t *testing.T) {
	bufferSize := uint(10)
	buffer := NewKSBuffer(bufferSize)
	buffer.Delete()
	// check that it is a circular buffer
	initialNode := buffer
	buffer = buffer.Next() // take an initial step
	i := uint(1)
	for ; buffer != initialNode; buffer = buffer.Next() {
		i++ // one step around the buffer
	}
	if i != bufferSize-1 {
		t.Errorf("Expected buffer to repeat after %d steps not %d", bufferSize, i)
	}
}
