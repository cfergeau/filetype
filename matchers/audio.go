package matchers

var (
	TypeMidi = newType("mid", "audio/midi")
	TypeMp3  = newType("mp3", "audio/mpeg")
	TypeM4a  = newType("m4a", "audio/m4a")
	TypeOgg  = newType("ogg", "audio/ogg")
	TypeFlac = newType("flac", "audio/x-flac")
	TypeWav  = newType("wav", "audio/x-wav")
	TypeAmr  = newType("amr", "audio/amr")
	TypeAac  = newType("aac", "audio/aac")
)

var Audio = Map{
	TypeMidi: bytePrefixMatcher(midiMagic),
	TypeMp3:  Mp3,
	TypeM4a:  M4a,
	TypeOgg:  bytePrefixMatcher(oggMagic),
	TypeFlac: bytePrefixMatcher(flacMagic),
	TypeWav:  Wav,
	TypeAmr:  Amr,
	TypeAac:  Aac,
}

var (
	midiMagic = []byte{0x4D, 0x54, 0x68, 0x64}
	oggMagic  = []byte{0x4F, 0x67, 0x67, 0x53}
	flacMagic = []byte{0x66, 0x4C, 0x61, 0x43}
)

func Mp3(buf []byte) bool {
	var (
		mp3Magic1 = []byte{0x49, 0x44, 0x33}
		mp3Magic2 = []byte{0xff, 0xfb}
	)
	return compareBytes(buf, mp3Magic1, 0) ||
		compareBytes(buf, mp3Magic2, 0)
}

func M4a(buf []byte) bool {
	var (
		m4aMagic1 = []byte{0x66, 0x74, 0x79, 0x70, 0x4D, 0x34, 0x41}
		m4aMagic2 = []byte{0x4D, 0x34, 0x41, 0x20}
	)
	return compareBytes(buf, m4aMagic1, 4) ||
		compareBytes(buf, m4aMagic2, 0)
}

func Wav(buf []byte) bool {
	var (
		wavMagic1 = []byte{0x52, 0x49, 0x46, 0x46}
		wavMagic2 = []byte{0x57, 0x41, 0x56, 0x45}
	)
	return compareBytes(buf, wavMagic1, 0) &&
		compareBytes(buf, wavMagic2, 8)
}

func Amr(buf []byte) bool {
	var amrMagic = []byte{0x23, 0x21, 0x41, 0x4D, 0x52, 0x0A}
	return len(buf) > 11 &&
		compareBytes(buf, amrMagic, 0)
}

func Aac(buf []byte) bool {
	var (
		aacMagic1 = []byte{0xFF, 0xF1}
		aacMagic2 = []byte{0xFF, 0xF9}
	)
	return compareBytes(buf, aacMagic1, 0) ||
		compareBytes(buf, aacMagic2, 8)
}
