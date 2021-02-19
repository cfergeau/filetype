package matchers

import "bytes"

var (
	TypeMp4  = newType("mp4", "video/mp4")
	TypeM4v  = newType("m4v", "video/x-m4v")
	TypeMkv  = newType("mkv", "video/x-matroska")
	TypeWebm = newType("webm", "video/webm")
	TypeMov  = newType("mov", "video/quicktime")
	TypeAvi  = newType("avi", "video/x-msvideo")
	TypeWmv  = newType("wmv", "video/x-ms-wmv")
	TypeMpeg = newType("mpg", "video/mpeg")
	TypeFlv  = newType("flv", "video/x-flv")
	Type3gp  = newType("3gp", "video/3gpp")
)

var Video = Map{
	TypeMp4:  mp4,
	TypeM4v:  m4v,
	TypeMkv:  mkv,
	TypeWebm: webm,
	TypeMov:  mov,
	TypeAvi:  avi,
	TypeWmv:  bytePrefixMatcher(wmvMagic),
	TypeMpeg: mpeg,
	TypeFlv:  bytePrefixMatcher(flvMagic),
	Type3gp:  match3gp,
}

var (
	wmvMagic = []byte{
		0x30, 0x26, 0xB2, 0x75, 0x8E, 0x66, 0xCF, 0x11,
		0xA6, 0xD9,
	}
	flvMagic = []byte{0x46, 0x4C, 0x56, 0x01}
)

func m4v(buf []byte) bool {
	var m4vMagic = []byte{0x66, 0x74, 0x79, 0x70, 0x4D, 0x34, 0x56}
	return compareBytes(buf, m4vMagic, 4)
}

var mkvMagic = []byte{0x1A, 0x45, 0xDF, 0xA3}

func mkv(buf []byte) bool {
	return compareBytes(buf, mkvMagic, 0) &&
		containsMatroskaSignature(buf, []byte{'m', 'a', 't', 'r', 'o', 's', 'k', 'a'})
}

func webm(buf []byte) bool {
	return compareBytes(buf, mkvMagic, 0) &&
		containsMatroskaSignature(buf, []byte{'w', 'e', 'b', 'm'})
}

func mov(buf []byte) bool {

	var (
		movMagic1 = []byte{0x00, 0x00, 0x00, 0x14, 0x66, 0x74, 0x79, 0x70}
		movMagic2 = []byte{0x6d, 0x6f, 0x6f, 0x76}
		movMagic3 = []byte{0x6d, 0x64, 0x61, 0x74}
		movMagic4 = []byte{0x6d, 0x64, 0x61, 0x74}
	)
	return compareBytes(buf, movMagic1, 0) ||
		compareBytes(buf, movMagic2, 4) ||
		compareBytes(buf, movMagic3, 4) ||
		compareBytes(buf, movMagic4, 12)
}

func avi(buf []byte) bool {
	var (
		aviMagic1 = []byte{0x52, 0x49, 0x46, 0x46}
		aviMagic2 = []byte{0x41, 0x56, 0x49}
	)
	return compareBytes(buf, aviMagic1, 0) &&
		compareBytes(buf, aviMagic2, 8)
}

func mpeg(buf []byte) bool {
	return len(buf) > 3 &&
		buf[0] == 0x0 && buf[1] == 0x0 &&
		buf[2] == 0x1 && buf[3] >= 0xb0 &&
		buf[3] <= 0xbf
}

func mp4(buf []byte) bool {
	return len(buf) > 11 &&
		(buf[4] == 'f' && buf[5] == 't' && buf[6] == 'y' && buf[7] == 'p') &&
		((buf[8] == 'a' && buf[9] == 'v' && buf[10] == 'c' && buf[11] == '1') ||
			(buf[8] == 'd' && buf[9] == 'a' && buf[10] == 's' && buf[11] == 'h') ||
			(buf[8] == 'i' && buf[9] == 's' && buf[10] == 'o' && buf[11] == '2') ||
			(buf[8] == 'i' && buf[9] == 's' && buf[10] == 'o' && buf[11] == '3') ||
			(buf[8] == 'i' && buf[9] == 's' && buf[10] == 'o' && buf[11] == '4') ||
			(buf[8] == 'i' && buf[9] == 's' && buf[10] == 'o' && buf[11] == '5') ||
			(buf[8] == 'i' && buf[9] == 's' && buf[10] == 'o' && buf[11] == '6') ||
			(buf[8] == 'i' && buf[9] == 's' && buf[10] == 'o' && buf[11] == 'm') ||
			(buf[8] == 'm' && buf[9] == 'm' && buf[10] == 'p' && buf[11] == '4') ||
			(buf[8] == 'm' && buf[9] == 'p' && buf[10] == '4' && buf[11] == '1') ||
			(buf[8] == 'm' && buf[9] == 'p' && buf[10] == '4' && buf[11] == '2') ||
			(buf[8] == 'm' && buf[9] == 'p' && buf[10] == '4' && buf[11] == 'v') ||
			(buf[8] == 'm' && buf[9] == 'p' && buf[10] == '7' && buf[11] == '1') ||
			(buf[8] == 'M' && buf[9] == 'S' && buf[10] == 'N' && buf[11] == 'V') ||
			(buf[8] == 'N' && buf[9] == 'D' && buf[10] == 'A' && buf[11] == 'S') ||
			(buf[8] == 'N' && buf[9] == 'D' && buf[10] == 'S' && buf[11] == 'C') ||
			(buf[8] == 'N' && buf[9] == 'S' && buf[10] == 'D' && buf[11] == 'C') ||
			(buf[8] == 'N' && buf[9] == 'D' && buf[10] == 'S' && buf[11] == 'H') ||
			(buf[8] == 'N' && buf[9] == 'D' && buf[10] == 'S' && buf[11] == 'M') ||
			(buf[8] == 'N' && buf[9] == 'D' && buf[10] == 'S' && buf[11] == 'P') ||
			(buf[8] == 'N' && buf[9] == 'D' && buf[10] == 'S' && buf[11] == 'S') ||
			(buf[8] == 'N' && buf[9] == 'D' && buf[10] == 'X' && buf[11] == 'C') ||
			(buf[8] == 'N' && buf[9] == 'D' && buf[10] == 'X' && buf[11] == 'H') ||
			(buf[8] == 'N' && buf[9] == 'D' && buf[10] == 'X' && buf[11] == 'M') ||
			(buf[8] == 'N' && buf[9] == 'D' && buf[10] == 'X' && buf[11] == 'P') ||
			(buf[8] == 'N' && buf[9] == 'D' && buf[10] == 'X' && buf[11] == 'S') ||
			(buf[8] == 'F' && buf[9] == '4' && buf[10] == 'V' && buf[11] == ' ') ||
			(buf[8] == 'F' && buf[9] == '4' && buf[10] == 'P' && buf[11] == ' '))
}

func match3gp(buf []byte) bool {
	var threegpMagic = []byte{0x66, 0x74, 0x79, 0x70, 0x33, 0x67, 0x70}
	return compareBytes(buf, threegpMagic, 4)
}

func containsMatroskaSignature(buf, subType []byte) bool {
	limit := 4096
	if len(buf) < limit {
		limit = len(buf)
	}

	index := bytes.Index(buf[:limit], subType)
	if index < 3 {
		return false
	}

	return buf[index-3] == 0x42 && buf[index-2] == 0x82
}
