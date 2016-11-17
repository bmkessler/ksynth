package ksynth

import "testing"

func TestA440(t *testing.T) {
	ks := NewKarplusStrong(44100, 16, 1)
	ks.AddNote(440, 1)
	ks.WriteWAV("A440_1sec.wav")
}

func TestA4SlideB4(t *testing.T) {
	ks := NewKarplusStrong(44100, 16, 1)
	ks.AddSlide(440., 493.883301, 0.25, 0.25, 1.0)
	ks.WriteWAV("A4_B4_slide.wav")
}

func TestBb4Bend(t *testing.T) {
	ks := NewKarplusStrong(44100, 16, 1)
	ks.AddSlide(493.883301, 466.163762, 0.05, 0.1, 1.0)
	ks.WriteWAV("Bb4_bend.wav")
}

func TestLongSlide(t *testing.T) {
	ks := NewKarplusStrong(44100, 16, 1)
	ks.AddSlide(493.883301, 466.163762, 0.25, 1.0, 1.0)
	ks.WriteWAV("LongSlide.wav")
}

func TestShortSlide(t *testing.T) {
	ks := NewKarplusStrong(44100, 16, 1)
	ks.AddSlide(440., 880., 0.25, 0.1, 1.0)
	ks.WriteWAV("ShortSlide.wav")
}
