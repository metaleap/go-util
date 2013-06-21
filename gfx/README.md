# ugfx
--
    import "github.com/metaleap/go-util/gfx"

A few helpers for working with graphics and images.

## Usage

#### func  Index2D

```go
func Index2D(x, y, ysize int) int
```
If 2 dimensions are represented in a 1-dimensional linear array, this function
provides a way to return a 1D index addressing the specified 2D coordinate.

#### func  Index3D

```go
func Index3D(x, y, z, xsize, ysize int) int
```
If 3 dimensions are represented in a 1-dimensional linear array, this function
provides a way to return a 1D index addressing the specified 3D coordinate.

#### func  PreprocessImage

```go
func PreprocessImage(src image.Image, dst Picture, flipY, toBgra, toLinear bool)
```
Processes the specified `Image` and writes the result to the specified
`Picture`. Unless `flipY` is `true`, `dst` and `src` may well be the same
object. If `flipY` is `true`, all pixel rows are inverted (`dst` becomes `src`
vertically mirrored). If `toBgra` is `true`, all pixels' red and blue components
are swapped. If `toLinear` is `true`, all pixels are converted from gamma/sRGB
to linear space -- only use this if you're certain that `src` is not already in
linear space.

#### func  SavePngImageFile

```go
func SavePngImageFile(img image.Image, filePath string) error
```
Saves any given `Image` as a local PNG file.

#### type Picture

```go
type Picture interface {
	image.Image

	//	Set pixel at x,y to the specified color.
	Set(int, int, color.Color)
}
```

The interface that the image package was missing...

#### func  CloneImage

```go
func CloneImage(src image.Image, copyPixels bool) (dst Picture, pix []byte)
```
Creates and returns a copy of `src`. If `copyPixels` is `true`, pixels in `src`
are copied to `dst`, otherwise `dst` will be an empty/black image of the same
dimensions, color format, stride/offset/etc as `src`.

#### type Rgba32

```go
type Rgba32 struct {
	//	Red component
	R float32
	//	Green component
	G float32
	//	Blue component
	B float32
	//	Alpha component
	A float32
}
```

Describes a literal color using four 32-bit floating-point numbers in RGBA
order.

#### func  NewRgba32

```go
func NewRgba32(vals ...float64) (me *Rgba32)
```
Converts the specified `vals` to a newly initialized `Rgba32` instance. The
first 4 `vals` are used for `R`, `G`, `B`, and `A` in that order, if present.
`A` is set to 1 if `vals[3]` is not present.

#### type Rgba64

```go
type Rgba64 struct {
	//	Red component
	R float64
	//	Green component
	G float64
	//	Blue component
	B float64
	//	Alpha component
	A float64
}
```

Describes a literal color using four 64-bit floating-point numbers in RGBA
order.

#### func  NewRgba64

```go
func NewRgba64(vals ...float64) (me *Rgba64)
```
Converts the specified `vals` to a newly initialized `Rgba64` instance. The
first 4 `vals` are used for `R`, `G`, `B`, and `A` in that order, if present.
`A` is set to 1 if `vals[3]` is not present.

--
**godocdown** http://github.com/robertkrimen/godocdown