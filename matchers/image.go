package matchers

import "github.com/h2non/filetype/matchers/isobmff"

var (
	TypeJpeg     = newType("jpg", "image/jpeg")
	TypeJpeg2000 = newType("jp2", "image/jp2")
	TypePng      = newType("png", "image/png")
	TypeGif      = newType("gif", "image/gif")
	TypeWebp     = newType("webp", "image/webp")
	TypeCR2      = newType("cr2", "image/x-canon-cr2")
	TypeTiff     = newType("tif", "image/tiff")
	TypeBmp      = newType("bmp", "image/bmp")
	TypeJxr      = newType("jxr", "image/vnd.ms-photo")
	TypePsd      = newType("psd", "image/vnd.adobe.photoshop")
	TypeIco      = newType("ico", "image/vnd.microsoft.icon")
	TypeHeif     = newType("heif", "image/heif")
	TypeDwg      = newType("dwg", "image/vnd.dwg")
)

var Image = Map{
	TypeJpeg:     bytePrefixMatcher(jpegMagic),
	TypeJpeg2000: bytePrefixMatcher(jpeg2000Magic),
	TypePng:      bytePrefixMatcher(pngMagic),
	TypeGif:      bytePrefixMatcher(gifMagic),
	TypeWebp:     Webp,
	TypeCR2:      CR2,
	TypeTiff:     Tiff,
	TypeBmp:      bytePrefixMatcher(bmpMagic),
	TypeJxr:      bytePrefixMatcher(jxrMagic),
	TypePsd:      bytePrefixMatcher(psdMagic),
	TypeIco:      bytePrefixMatcher(icoMagic),
	TypeHeif:     Heif,
	TypeDwg:      bytePrefixMatcher(dwgMagic),
}

var (
	jpegMagic = []byte{
		0xFF, 0xD8, 0xFF,
	}
	jpeg2000Magic = []byte{
		0x00, 0x00, 0x00, 0x0C, 0x6A, 0x50, 0x20, 0x20,
		0x0D, 0x0A, 0x87, 0x0A, 0x00,
	}
	pngMagic = []byte{
		0x89, 0x50, 0x4E, 0x47,
	}
	gifMagic = []byte{
		0x47, 0x49, 0x46,
	}
	bmpMagic = []byte{
		0x42, 0x4D,
	}
	jxrMagic = []byte{
		0x49, 0x49, 0xBC,
	}
	psdMagic = []byte{
		0x38, 0x42, 0x50, 0x53,
	}
	icoMagic = []byte{
		0x00, 0x00, 0x01, 0x00,
	}
	dwgMagic = []byte{
		0x41, 0x43, 0x31, 0x30,
	}
)

func Webp(buf []byte) bool {
	return len(buf) > 11 &&
		buf[8] == 0x57 && buf[9] == 0x45 &&
		buf[10] == 0x42 && buf[11] == 0x50
}

func CR2(buf []byte) bool {
	return len(buf) > 10 &&
		((buf[0] == 0x49 && buf[1] == 0x49 && buf[2] == 0x2A && buf[3] == 0x0) || // Little Endian
			(buf[0] == 0x4D && buf[1] == 0x4D && buf[2] == 0x0 && buf[3] == 0x2A)) && // Big Endian
		buf[8] == 0x43 && buf[9] == 0x52 && // CR2 magic word
		buf[10] == 0x02 // CR2 major version
}

func Tiff(buf []byte) bool {
	return len(buf) > 10 &&
		((buf[0] == 0x49 && buf[1] == 0x49 && buf[2] == 0x2A && buf[3] == 0x0) || // Little Endian
			(buf[0] == 0x4D && buf[1] == 0x4D && buf[2] == 0x0 && buf[3] == 0x2A)) && // Big Endian
		!CR2(buf) // To avoid conflicts differentiate Tiff from CR2
}

func Heif(buf []byte) bool {
	if !isobmff.IsISOBMFF(buf) {
		return false
	}

	majorBrand, _, compatibleBrands := isobmff.GetFtyp(buf)
	if majorBrand == "heic" {
		return true
	}

	if majorBrand == "mif1" || majorBrand == "msf1" {
		for _, compatibleBrand := range compatibleBrands {
			if compatibleBrand == "heic" {
				return true
			}
		}
	}

	return false
}
