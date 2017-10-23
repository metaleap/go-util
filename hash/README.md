# uhash

Go programming helpers for common hashing needs.

Also implements a selection of the hashing algorithms described at:
http://eternallyconfuzzled.com/tuts/algorithms/jsw_tut_hashing.aspx

## Usage

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

#### func  WriteAndSum

```go
func WriteAndSum(h hash.Hash, data, b []byte) (sum []byte, err error)
```
Convenience short-hand for `h.Write(data)`, then `h.Sum(b)`.
