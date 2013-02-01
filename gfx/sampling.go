package ugfx

//	Categorizes the kinds of wrapping used in SamplerWrapping.
type WrapKind int

const (
	//	Ignores the integer part of texture coordinates, using only the fractional part and tiling the
	//	texture at every integer junction. For example, for u values between 0 and 3, the texture is
	//	repeated three times; no mirroring is performed.
	WrapKindRepeat WrapKind = 0x2901

	//	First mirrors the texture coordinate. The mirrored coordinate is then clamped as described for
	//	WrapKindClamp. Flips the texture at every integer junction. For u values between 0 and 1,
	//	for example, the texture is addressed normally; between 1 and 2, the texture is flipped (mirrored);
	//	between 2 and 3, the texture is normal again; and so on.
	WrapKindMirror WrapKind = 0x8370

	//	Clamps texture coordinates at all MIPmap levels such that
	//	the texture filter never samples a border texel.
	WrapKindClamp WrapKind = 0x812F

	//	Clamps texture coordinates at all MIPmaps such that the texture filter always samples border
	//	texels for fragments whose corresponding texture coordinate is sufficiently far outside
	//	the range [0, 1]. Much like WrapKindClamp, except texture coordinates outside
	//	the range [0.0, 1.0] are set to the border color.
	WrapKindBorder WrapKind = 0x812D

	//	Takes the absolute value of the texture coordinate (thus, mirroring around 0),
	//	and then clamps to the maximum value.
	WrapKindMirrorOnce WrapKind = 41
)

//	Controls texture repeating and clamping.
type SamplerWrapping struct {
	//	When reading past the edge of the texture address space
	//	based on the wrap modes involving clamps, this color takes over.
	BorderColor Rgba32

	//	Controls texture repeating and clamping of the S coordinate.
	//	Must be one of the WrapKind* enumerated constants.
	WrapS WrapKind

	//	Controls texture repeating and clamping of the T coordinate.
	//	Must be one of the WrapKind* enumerated constants.
	WrapT WrapKind

	//	Controls texture repeating and clamping of the P coordinate.
	//	Must be one of the WrapKind* enumerated constants.
	WrapP WrapKind
}

//	Initializes a new SamplerWrapping with the specified coordinate wrappings
//	and borderColor. If no borderColor is specified, black (0, 0, 0, 1) is used.
func NewSamplerWrapping(stp WrapKind, borderColor *Rgba32) (me *SamplerWrapping) {
	if borderColor == nil {
		borderColor = NewRgba32(0, 0, 0, 1)
	}
	me = &SamplerWrapping{WrapS: stp, WrapT: stp, WrapP: stp, BorderColor: *borderColor}
	return
}
