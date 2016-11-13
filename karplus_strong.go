package ksynth

import (
	"encoding/binary"
	"os"
)

var scalingFactor = 0.75 // the scaling factor for volume between 0.0 and 1.0

// SetVolume sets the scaling factor for volume between 0.0 and 1.0
func SetVolume(newVolume float64) {
	if newVolume > 1.0 { // clipped to range
		newVolume = 1.0
	}
	if newVolume < 0.0 {
		newVolume = 0.0
	}
	scalingFactor = newVolume
}

// KarplusStrong creates samples and writes them to WAV files
type KarplusStrong struct {
	samplesPerSecond uint32
	bitsPerSample    uint8
	numChannels      uint8
	data             []byte
}

// NewKarplusStrong initializes a KarplusStrong object
func NewKarplusStrong(samplesPerSecond uint32, bitsPerSample uint8, numChannels uint8) KarplusStrong {
	return KarplusStrong{
		samplesPerSecond: samplesPerSecond,
		bitsPerSample:    bitsPerSample,
		numChannels:      numChannels,
		data:             []byte{},
	}
}

// AddNote adds samples for the given note with frequency in Hz and duration in seconds to the internal data buffer
func (ks *KarplusStrong) AddNote(frequency float64, duration float64) {
	bufferSize := uint(float64(ks.samplesPerSecond)/frequency + 0.5)
	numberSamples := int(float64(ks.samplesPerSecond)*duration + 0.5)
	buffer := NewKSBuffer(bufferSize)
	for i := 0; i < numberSamples; i++ {
		val := buffer.Value()
		volume := scalingFactor * float64(uint64(1)<<(ks.bitsPerSample-1))
		for i := uint8(0); i < ks.numChannels; i++ { // write one sample for each channel
			u32Sample := uint32(val * volume) // bits per sample only supported multiples of 8 up to 32
			if ks.bitsPerSample == 8 {
				u32Sample += (1 << 7) // 8-bit is offset encoded
				u32Sample %= (1 << 8)
			}
			for i := uint8(0); i < ks.bitsPerSample/8; i++ {
				ks.data = append(ks.data, byte(u32Sample&0xFF))
				u32Sample = u32Sample >> 8
			}
		}
		buffer.Update()
		buffer = buffer.Next()
	}
}

// WriteWAV writes out the current data buffer to a WAV file
func (ks KarplusStrong) WriteWAV(filename string) error {
	data := ks.data
	if len(ks.data)%2 != 0 {
		data = append(data, byte(0)) // pad a zero byte if the length is not even
	}
	M := ks.bitsPerSample / 8     // Bytes per sample
	Nc := ks.numChannels          // number of channels
	dataSize := uint32(len(data)) // the total number of bytes, with padding
	header := waveHeader{
		riffChunkSize:         uint32(36 + dataSize),
		formatChunkSize:       uint32(16),
		waveFormatTag:         uint16(0x0001),
		numberOfChannels:      uint16(Nc),
		samplesPerSecond:      uint32(ks.samplesPerSecond),
		averageBytesPerSecond: ks.samplesPerSecond * uint32(M) * uint32(Nc),
		blockAlign:            uint16(M) * uint16(Nc),
		bitsPerSample:         uint16(ks.bitsPerSample),
		dataChunkSize:         dataSize,
	}
	copy(header.riffChunkID[:], riffTag)
	copy(header.waveChunkID[:], waveTag)
	copy(header.formatChunkID[:], fmtTag)
	copy(header.dataChunkID[:], dataTag)

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	err = binary.Write(file, binary.LittleEndian, header)
	if err != nil {
		return err
	}
	err = binary.Write(file, binary.LittleEndian, ks.data)
	if err != nil {
		return err
	}

	return nil
}

type waveHeader struct {
	riffChunkID           [4]byte // "RIFF"
	riffChunkSize         uint32  // 4 + (8 + formatChunkSize) + (8 + dataChunkSize) = 36 + dataChunkSize
	waveChunkID           [4]byte // "WAVE"
	formatChunkID         [4]byte // "fmt "
	formatChunkSize       uint32  // 16 for PCM
	waveFormatTag         uint16  // 0x0001 for PCM
	numberOfChannels      uint16  // Nc
	samplesPerSecond      uint32  // sampling frequency, e.g. 48000
	averageBytesPerSecond uint32  // F*M*Nc
	blockAlign            uint16  // M*Nc
	bitsPerSample         uint16  // 8*M
	dataChunkID           [4]byte // "data"
	dataChunkSize         uint32  // M*Nc*Ns
}

const (
	riffTag = "RIFF" // RIFF tag header for entire file
	waveTag = "WAVE" // WAVE tag header identifying type of RIFF
	fmtTag  = "fmt " // fmt tag header for format chunk
	dataTag = "data" // data tag header for data chunk
)
