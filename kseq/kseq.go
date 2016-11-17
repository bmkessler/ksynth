package main

import "github.com/bmkessler/ksynth"
import "flag"
import "io/ioutil"
import "log"
import "path"

func main() {
	inputFile := flag.String("f", "input_seq.txt", "The input file to parse and create input_seq.wav")
	bpm := flag.Float64("bpm", 120, "The beats per minute to record the file at")
	// WAV file parameters
	sampleRate := flag.Uint("sr", 48000, "Sample rate in samples per second")
	bitRate := flag.Uint("br", 16, "Bit rate in bits per sample, 8, 16, 24, and 32 supported")
	numChannels := flag.Uint("nc", 1, "Number of audio channels to record")

	flag.Parse()
	buf, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	notes := ksynth.ParseString(string(buf), *bpm)
	ks := ksynth.NewKarplusStrong(uint32(*sampleRate), uint8(*bitRate), uint8(*numChannels))
	for _, note := range notes {
		if note.Error == nil {
			ks.AddNote(note.Frequency, note.Duration)
		} else {
			log.Print(err)
		}
	}
	// trim any extension and append .wav to the input file
	outputFile := (*inputFile)[:len(*inputFile)-len(path.Ext(*inputFile))] + ".wav"
	err = ks.WriteWAV(outputFile)

	if err != nil {
		log.Fatal(err)
	}

}
