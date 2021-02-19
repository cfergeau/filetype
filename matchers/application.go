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
	var dexMagic = []byte{
		0x64, 0x65, 0x78, 0x0A,
	}
	// https://source.android.com/devices/tech/dalvik/dex-format#dex-file-magic
	return len(buf) > 36 &&
		// magic
		compareBytes(buf, dexMagic, 0) &&
		// file size
		buf[36] == 0x70
}

// Dey Optimized Dalvik Executable(ODEX)
func Dey(buf []byte) bool {
	var deyMagic = []byte{
		0x64, 0x65, 0x78, 0x0A,
	}
	return len(buf) > 100 &&
		// dey magic
		compareBytes(buf, deyMagic, 0) &&
		// dex
		Dex(buf[40:100])
}
