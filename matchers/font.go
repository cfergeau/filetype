package matchers

var (
	TypeWoff  = newType("woff", "application/font-woff")
	TypeWoff2 = newType("woff2", "application/font-woff")
	TypeTtf   = newType("ttf", "application/font-sfnt")
	TypeOtf   = newType("otf", "application/font-sfnt")
)

var Font = Map{
	TypeWoff:  bytePrefixMatcher(woffMagic),
	TypeWoff2: bytePrefixMatcher(woff2Magic),
	TypeTtf:   bytePrefixMatcher(ttfMagic),
	TypeOtf:   bytePrefixMatcher(otfMagic),
}

var (
	woffMagic = []byte{
		0x77, 0x4F, 0x46, 0x46, 0x00, 0x01, 0x00, 0x00,
	}
	woff2Magic = []byte{
		0x77, 0x4F, 0x46, 0x32, 0x00, 0x01, 0x00, 0x00,
	}
	ttfMagic = []byte{
		0x00, 0x01, 0x00, 0x00, 0x00,
	}
	otfMagic = []byte{
		0x4F, 0x54, 0x54, 0x4F, 0x00,
	}
)
