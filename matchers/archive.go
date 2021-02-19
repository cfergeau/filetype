package matchers

var (
	TypeEpub   = newType("epub", "application/epub+zip")
	TypeZip    = newType("zip", "application/zip")
	TypeTar    = newType("tar", "application/x-tar")
	TypeRar    = newType("rar", "application/vnd.rar")
	TypeGz     = newType("gz", "application/gzip")
	TypeBz2    = newType("bz2", "application/x-bzip2")
	Type7z     = newType("7z", "application/x-7z-compressed")
	TypeXz     = newType("xz", "application/x-xz")
	TypeZstd   = newType("zst", "application/zstd")
	TypePdf    = newType("pdf", "application/pdf")
	TypeExe    = newType("exe", "application/vnd.microsoft.portable-executable")
	TypeSwf    = newType("swf", "application/x-shockwave-flash")
	TypeRtf    = newType("rtf", "application/rtf")
	TypeEot    = newType("eot", "application/octet-stream")
	TypePs     = newType("ps", "application/postscript")
	TypeSqlite = newType("sqlite", "application/vnd.sqlite3")
	TypeNes    = newType("nes", "application/x-nintendo-nes-rom")
	TypeCrx    = newType("crx", "application/x-google-chrome-extension")
	TypeCab    = newType("cab", "application/vnd.ms-cab-compressed")
	TypeDeb    = newType("deb", "application/vnd.debian.binary-package")
	TypeAr     = newType("ar", "application/x-unix-archive")
	TypeZ      = newType("Z", "application/x-compress")
	TypeLz     = newType("lz", "application/x-lzip")
	TypeRpm    = newType("rpm", "application/x-rpm")
	TypeElf    = newType("elf", "application/x-executable")
	TypeDcm    = newType("dcm", "application/dicom")
	TypeIso    = newType("iso", "application/x-iso9660-image")
	TypeMachO  = newType("macho", "application/x-mach-binary") // Mach-O binaries have no common extension.
)

var Archive = Map{
	TypeEpub:   bytePrefixMatcher(epubMagic),
	TypeZip:    Zip,
	TypeTar:    Tar,
	TypeRar:    Rar,
	TypeGz:     bytePrefixMatcher(gzMagic),
	TypeBz2:    bytePrefixMatcher(bz2Magic),
	Type7z:     bytePrefixMatcher(sevenzMagic),
	TypeXz:     bytePrefixMatcher(xzMagic),
	TypeZstd:   bytePrefixMatcher(zstdMagic),
	TypePdf:    bytePrefixMatcher(pdfMagic),
	TypeExe:    bytePrefixMatcher(exeMagic),
	TypeSwf:    Swf,
	TypeRtf:    bytePrefixMatcher(rtfMagic),
	TypeEot:    Eot,
	TypePs:     bytePrefixMatcher(psMagic),
	TypeSqlite: bytePrefixMatcher(sqliteMagic),
	TypeNes:    bytePrefixMatcher(nesMagic),
	TypeCrx:    bytePrefixMatcher(crxMagic),
	TypeCab:    Cab,
	TypeDeb:    bytePrefixMatcher(debMagic),
	TypeAr:     bytePrefixMatcher(arMagic),
	TypeZ:      Z,
	TypeLz:     bytePrefixMatcher(lzMagic),
	TypeRpm:    Rpm,
	TypeElf:    Elf,
	TypeDcm:    Dcm,
	TypeIso:    Iso,
	TypeMachO:  MachO,
}

var (
	epubMagic = []byte{
		0x50, 0x4B, 0x03, 0x04, 0x6D, 0x69, 0x6D, 0x65,
		0x74, 0x79, 0x70, 0x65, 0x61, 0x70, 0x70, 0x6C,
		0x69, 0x63, 0x61, 0x74, 0x69, 0x6F, 0x6E, 0x2F,
		0x65, 0x70, 0x75, 0x62, 0x2B, 0x7A, 0x69, 0x70,
	}
	gzMagic     = []byte{0x1F, 0x8B, 0x08}
	bz2Magic    = []byte{0x42, 0x5A, 0x68}
	sevenzMagic = []byte{0x37, 0x7A, 0xBC, 0xAF, 0x27, 0x1C}
	pdfMagic    = []byte{0x25, 0x50, 0x44, 0x46}
	exeMagic    = []byte{0x4D, 0x5A}
	rtfMagic    = []byte{0x7B, 0x5C, 0x72, 0x74, 0x66}
	nesMagic    = []byte{0x4E, 0x45, 0x53, 0x1A}
	crxMagic    = []byte{0x43, 0x72, 0x32, 0x34}
	psMagic     = []byte{0x25, 0x21}
	xzMagic     = []byte{0xFD, 0x37, 0x7A, 0x58, 0x5A, 0x00}
	sqliteMagic = []byte{0x53, 0x51, 0x4C, 0x69}
	debMagic    = []byte{
		0x21, 0x3C, 0x61, 0x72, 0x63, 0x68, 0x3E, 0x0A,
		0x64, 0x65, 0x62, 0x69, 0x61, 0x6E, 0x2D, 0x62,
		0x69, 0x6E, 0x61, 0x72, 0x79,
	}
	arMagic   = []byte{0x21, 0x3C, 0x61, 0x72, 0x63, 0x68, 0x3E}
	zstdMagic = []byte{0x28, 0xB5, 0x2F, 0xFD}
	lzMagic   = []byte{0x4C, 0x5A, 0x49, 0x50}
)

func bytePrefixMatcher(magicPattern []byte) Matcher {
	return func(data []byte) bool {
		return compareBytes(data, magicPattern, 0)
	}
}

func Zip(buf []byte) bool {
	return len(buf) > 3 &&
		buf[0] == 0x50 && buf[1] == 0x4B &&
		(buf[2] == 0x3 || buf[2] == 0x5 || buf[2] == 0x7) &&
		(buf[3] == 0x4 || buf[3] == 0x6 || buf[3] == 0x8)
}

func Tar(buf []byte) bool {
	var tarMagic = []byte{0x75, 0x73, 0x74, 0x61, 0x72}
	return compareBytes(buf, tarMagic, 257)
}

func Rar(buf []byte) bool {
	var (
		rarMagic1 = []byte{0x52, 0x61, 0x72, 0x21, 0x1A, 0x07, 0x00}
		rarMagic2 = []byte{0x52, 0x61, 0x72, 0x21, 0x1A, 0x07, 0x01}
	)
	return compareBytes(buf, rarMagic1, 0) ||
		compareBytes(buf, rarMagic2, 0)
}

func Swf(buf []byte) bool {
	var (
		swfMagic1 = []byte{0x43, 0x57, 0x53}
		swfMagic2 = []byte{0x46, 0x57, 0x53}
	)
	return compareBytes(buf, swfMagic1, 0) ||
		compareBytes(buf, swfMagic2, 0)
}

func Cab(buf []byte) bool {
	var (
		cabMagic1 = []byte{0x4D, 0x53, 0x43, 0x46}
		cabMagic2 = []byte{0x49, 0x53, 0x63, 0x28}
	)
	return compareBytes(buf, cabMagic1, 0) ||
		compareBytes(buf, cabMagic2, 0)
}

func Eot(buf []byte) bool {
	return len(buf) > 35 &&
		buf[34] == 0x4C && buf[35] == 0x50 &&
		((buf[8] == 0x02 && buf[9] == 0x00 &&
			buf[10] == 0x01) || (buf[8] == 0x01 &&
			buf[9] == 0x00 && buf[10] == 0x00) ||
			(buf[8] == 0x02 && buf[9] == 0x00 &&
				buf[10] == 0x02))
}

func Z(buf []byte) bool {
	var (
		zMagic1 = []byte{0x1F, 0xA0}
		zMagic2 = []byte{0x1F, 0x9D}
	)
	return compareBytes(buf, zMagic1, 0) ||
		compareBytes(buf, zMagic2, 0)
}

func Rpm(buf []byte) bool {
	var rpmMagic = []byte{0xED, 0xAB, 0xEE, 0xDB}
	return len(buf) > 96 &&
		compareBytes(buf, rpmMagic, 0)
}

func Elf(buf []byte) bool {
	var elfMagic = []byte{
		0x7F, 0x45, 0x4C, 0x46,
	}
	return len(buf) > 52 &&
		compareBytes(buf, elfMagic, 0)
}

func Dcm(buf []byte) bool {
	var dcmMagic = []byte{0x44, 0x49, 0x43, 0x4D}
	return compareBytes(buf, dcmMagic, 128)
}

func Iso(buf []byte) bool {
	var isoMagic = []byte{0x43, 0x44, 0x30, 0x30, 0x31}
	return compareBytes(buf, isoMagic, 32769)
}

func MachO(buf []byte) bool {
	return len(buf) > 3 && ((buf[0] == 0xFE && buf[1] == 0xED && buf[2] == 0xFA && buf[3] == 0xCF) ||
		(buf[0] == 0xFE && buf[1] == 0xED && buf[2] == 0xFA && buf[3] == 0xCE) ||
		(buf[0] == 0xBE && buf[1] == 0xBA && buf[2] == 0xFE && buf[3] == 0xCA) ||
		// Big endian versions below here...
		(buf[0] == 0xCF && buf[1] == 0xFA && buf[2] == 0xED && buf[3] == 0xFE) ||
		(buf[0] == 0xCE && buf[1] == 0xFA && buf[2] == 0xED && buf[3] == 0xFE) ||
		(buf[0] == 0xCA && buf[1] == 0xFE && buf[2] == 0xBA && buf[3] == 0xBE))
}
