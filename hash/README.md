# uhash
--
    import "github.com/metaleap/go-util/hash"

Implements the most useful of the hashing algorithms described at:
http://eternallyconfuzzled.com/tuts/algorithms/jsw_tut_hashing.aspx

## Usage

#### func  EncodeToString

```go
func EncodeToString(h hash.Hash, data []byte, enc *base64.Encoding) string
```
Returns the `enc`'s (or if `nil`, the `base64.URLEncoding`'s) string-encoding of
the specified `Hash` for `data`.

#### func  Fnv1

```go
func Fnv1(vals []int) (h int)
```
Fowler/Noll/Vo '1'

#### func  Fnv1a

```go
func Fnv1a(vals []int) (h int)
```
Fowler/Noll/Vo '1a'

#### func  ModifiedBernstein

```go
func ModifiedBernstein(vals []int) (h int)
```
A minor update to Bernstein's hash replaces addition with XOR for the combining
step.

#### func  OneAtATime

```go
func OneAtATime(vals []int) (h int)
```
Bob Jenkins

#### func  RotatingAdd

```go
func RotatingAdd(vals []int) (h int)
```
The rotating hash with SUM

#### func  RotatingXor

```go
func RotatingXor(vals []int) (h int)
```
The rotating hash with XOR

--
**godocdown** http://github.com/robertkrimen/godocdown