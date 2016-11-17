package ksynth

import "testing"

func TestCalculateFrequency(t *testing.T) {
	octave := int(4)
	for note, tone := range Notation {
		freq := CalculateFrequency(tone, octave)
		t.Logf("%s%d = %f Hz", note, octave, freq)
	}
}

func TestCalculateDuration(t *testing.T) {
	beatsPerMinute := 120.
	beatsPerNote := 1.
	expectedDuration := 0.5
	duration, err := CalculateDuration(beatsPerNote, beatsPerMinute)
	if duration != expectedDuration {
		t.Errorf("Expected %f got %f", expectedDuration, duration)
	}
	if err != nil {
		t.Error(err)
	}
	negBeatsPerNote := -1.
	_, err = CalculateDuration(negBeatsPerNote, beatsPerMinute)
	if err == nil {
		t.Errorf("Expected failure with beatsPerNote: %f beatsPerMinute: %f", negBeatsPerNote, beatsPerMinute)
	}
	_, err = CalculateDuration(beatsPerMinute, negBeatsPerNote)
	if err == nil {
		t.Errorf("Expected failure with beatsPerNote: %f beatsPerMinute: %f", beatsPerMinute, negBeatsPerNote)
	}
}

func TestParseString(t *testing.T) {

	inputString := "A3q A3q E3q E3q F#3q F#3q E3h"
	bpm := 120.
	notes := ParseString(inputString, bpm)
	for _, note := range notes {
		if note.Error != nil {
			t.Error(note.Error)
		} else {
			t.Logf("f: %f d: %f", note.Frequency, note.Duration)
		}
	}
	failString := "Not E#3w Bb2t"
	failNotes := ParseString(failString, bpm)
	for i, note := range failNotes {
		if note.Error == nil {
			t.Errorf("Expected failure on %d note in %s got %v", i, failString, note)
		}
	}
}
