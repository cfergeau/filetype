package matchers

var (
	TypeWasm = newType("wasm", "application/wasm")
	TypeDex  = newType("dex", "application/vnd.android.dex")
	TypeDey  = newType("dey", "application/vnd.android.dey")
)

var Application = Map{
	TypeWasm: bytePrefixMatcher(wasmMagic),
	TypeDex:  Dex,
	TypeDey:  Dey,
}

var (
	wasmMagic = []byte{
		0x00, 0x61, 0x73, 0x6D, 0x01, 0x00, 0x00, 0x00,
	}
)

// Dex detects dalvik executable(DEX)
func Dex(buf []byte) bool {
	// https://source.android.com/devices/tech/dalvik/dex-format#dex-file-magic
	return len(buf) > 36 &&
		// magic
		buf[0] == 0x64 && buf[1] == 0x65 && buf[2] == 0x78 && buf[3] == 0x0A &&
		// file sise
		buf[36] == 0x70
}

// Dey Optimized Dalvik Executable(ODEX)
func Dey(buf []byte) bool {
	return len(buf) > 100 &&
		// dey magic
		buf[0] == 0x64 && buf[1] == 0x65 && buf[2] == 0x79 && buf[3] == 0x0A &&
		// dex
		Dex(buf[40:100])
}
