package ksynth

import (
	"fmt"
	"math"
)

// This file contains basic music functions for calculating frequencies and durations

// CalculateFrequency calculates the frequency of any note and octave
// integer notes 1 to 12 are the well-tempered scale with A4=440Hz
func CalculateFrequency(note float64, octave uint) float64 {
	return 440.0 * math.Pow(2.0, (float64(octave-4)*12+note)/12.0)
}

// Notation contains the mapping from ascii notes to numerical values
var Notation = map[string]float64{
	"A":  0,
	"A#": 1,
	"B":  2,
	"C":  3,
	"C#": 4,
	"D":  5,
	"D#": 6,
	"E":  7,
	"F":  8,
	"F#": 9,
	"G":  10,
	"G#": 11,

	// additional flat notations
	"Bb": 1,
	"Db": 4,
	"Eb": 6,
	"Gb": 9,
	"Ab": 11,
}

// NoteToTone converts a string e.g. A# to a numerical value 1.0
func NoteToTone(note string) (float64, error) {
	if tone, ok := Notation[note]; ok {
		return tone, nil
	}
	return 0.0, fmt.Errorf("%s is an unknown note, valid values are: A A# Bf B C C# etc", note)
}

// BeatsPerNote contains the mapping from notes, whole, half, etc. to fraction of beat in 4/4
var BeatsPerNote = map[string]float64{
	"w": 4,
	"h": 2,
	"q": 1,
	"e": 0.5,
	"s": 0.25,
}

// NoteToBeats converts a note to how many beats it takes in 4/4 time
func NoteToBeats(note string) (float64, error) {
	if beats, ok := BeatsPerNote[note]; ok {
		return beats, nil
	}
	return 0.0, fmt.Errorf("%s is an unknown note type, valid values are: w h q e s", note)
}

// CalculateDuration returns how long a note lasts at a given tempo in bpm
func CalculateDuration(beatsPerNote float64, beatsPerMinute float64) (float64, error) {
	if beatsPerMinute <= 0.0 || beatsPerNote <= 0 {
		return 0.0, fmt.Errorf("Both beatsPerNote: %f and beatsPerNote: %f must be positive", beatsPerNote, beatsPerMinute)
	}
	return beatsPerNote / (beatsPerMinute / 60.0), nil
}

// FrequencyDuration returns the frequency in Hz and length in seconds for a given note at given bpm
func FrequencyDuration(pitch string, octave uint, note string, bpm float64) (float64, float64, error) {
	beats, err := NoteToBeats(note)
	if err != nil {
		return 0.0, 0.0, err
	}
	duration, err := CalculateDuration(beats, bpm)
	if err != nil {
		return 0.0, 0.0, err
	}
	tone, err := NoteToTone(pitch)
	if err != nil {
		return 0.0, 0.0, err
	}
	frequency := CalculateFrequency(tone, octave)
	return frequency, duration, nil
}
