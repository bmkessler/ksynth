package ksynth

import "testing"

func TestA440(t *testing.T) {
	ks := NewKarplusStrong(44100, 16, 1)
	ks.AddNote(440, 1)
	ks.WriteWAV("A440_1sec.wav")
}
