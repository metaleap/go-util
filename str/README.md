# ustr
--
    import "github.com/metaleap/go-util/str"

Various line-savers for common string-processing needs.

## Usage

#### func  ClearMap

```go
func ClearMap(m map[string]string)
```
Sets all values in m to the empty string.

#### func  Concat

```go
func Concat(vals ...string) string
```
Does a strings.Join() on the specified string values.

#### func  ContainsOnce

```go
func ContainsOnce(str1, str2 string) bool
```
Returns true if str2 is contained in str1 exactly once.

#### func  Distance

```go
func Distance(s1, s2 string) int
```
A simple string-similarity algorithm.

#### func  ExtractAllIdentifiers

```go
func ExtractAllIdentifiers(src, prefix string) (identifiers []string)
```
Extracts all identifiers (no duplicates, ordered by occurrence) starting with
prefix occurring in src.

#### func  ExtractFirstIdentifier

```go
func ExtractFirstIdentifier(src, prefix string, minPos int) (identifier string)
```
Extracts the first occurrence (at or after minPos) of an identifier starting
with prefix in src.

#### func  First

```go
func First(predicate func(s string) bool, step int, vals ...string) string
```
Returns the first string in vals to match the specified predicate. step: 1 to
test all values. A higher value to skip n values after each test. Negative for
reverse slice traversal. Or use 0 to get stuck in an infinite loop.

#### func  FirstNonEmpty

```go
func FirstNonEmpty(vals ...string) (val string)
```
Returns the first non-empty string in vals.

#### func  HasAnyPrefix

```go
func HasAnyPrefix(s string, prefixes ...string) bool
```
Returns true if s starts with any one of the specified prefixes.

#### func  HasAnySuffix

```go
func HasAnySuffix(s string, suffixes ...string) bool
```
Returns true if s ends with any one of the specified suffixes.

#### func  Ifm

```go
func Ifm(cond bool, ifTrue, ifFalse map[string]string) map[string]string
```
Returns ifTrue if cond is true, otherwise returns ifFalse.

#### func  Ifs

```go
func Ifs(cond bool, ifTrue string, ifFalse string) string
```
Returns ifTrue if cond is true, otherwise returns ifFalse.

#### func  IndexAny

```go
func IndexAny(s string, seps ...string) (pos int)
```
For all seps, computes the index of first occurrence in s, then returns the
smallest index.

#### func  IsAscii

```go
func IsAscii(str string) bool
```
Returns true if str is ASCII-compatible.

#### func  IsLower

```go
func IsLower(s string) bool
```
Returns true if all Letter-runes in s are lower-case.

#### func  IsOneOf

```go
func IsOneOf(s string, all ...string) bool
```
Returns true if s is in all.

#### func  IsUpper

```go
func IsUpper(s string) bool
```
Returns true if all Letter-runes in s are upper-case.

#### func  LettersOnly

```go
func LettersOnly(s string) (ret string)
```
Returns a string representation of s with all non-Letter-runes removed.

#### func  MatchesAny

```go
func MatchesAny(value string, patterns ...string) bool
```
Uses a Matcher to determine whether value matches any one of the specified
simple-patterns.

#### func  NonEmpties

```go
func NonEmpties(breakAtFirstEmpty bool, vals ...string) (slice []string)
```
Returns a slice that contains the non-empty strings in vals.

#### func  Pluralize

```go
func Pluralize(s string) string
```
A most simplistic (not linguistically-correct) English-language pluralizer that
may be useful for code or doc generation. If s ends with "s", only appends "es":
bus -> buses, mess -> messes If s ends with "y" (but not "ay", "ey", "oy", "uy"
or "iy"), removes "y" and appends "ies": autonomy -> autonomies, dictionary ->
dictionaries etc. Otherwise, appends "s".

#### func  PrefixWithSep

```go
func PrefixWithSep(prefix, sep, v string) string
```
Prepends prefix + sep to v only if prefix isn't empty.

#### func  PrependIf

```go
func PrependIf(s, p string) string
```
Prepends p to s only if s doesn't already have that prefix.

#### func  Replace

```go
func Replace(str string, repls map[string]string) string
```
Replaces in str all occurrences of all repls map keys with their associated
(mapped) value.

#### func  SafeIdentifier

```go
func SafeIdentifier(s string) (ret string)
```
Creates a Pascal-cased "identifier" version of the specified string.

#### func  Split

```go
func Split(v, s string) (sl []string)
```
Returns an empty slice is v is emtpy, otherwise like strings.Split()

#### func  StripPrefix

```go
func StripPrefix(val, prefix string) string
```
Strips prefix off val if possible.

#### func  StripSuffix

```go
func StripSuffix(val, suffix string) string
```
Strips suffix off val if possible.

#### func  ToLowerIfUpper

```go
func ToLowerIfUpper(s string) string
```
Returns the lower-case representation of s only if it is currently fully
upper-case as per IsUpper().

#### func  ToUpperIfLower

```go
func ToUpperIfLower(s string) string
```
Returns the upper-case representation of s only if it is currently fully
lower-case as per IsLower().

#### type Buffer

```go
type Buffer struct {
	bytes.Buffer
}
```

A convenience wrapper for bytes.Buffer.

#### func (*Buffer) Write

```go
func (me *Buffer) Write(format string, fmtArgs ...interface{})
```
Convenience short-hand for bytes.Buffer.WriteString(fmt.Sprintf(format,
fmtArgs...))

#### func (*Buffer) Writeln

```go
func (me *Buffer) Writeln(format string, fmtArgs ...interface{})
```
Convenience short-hand for bytes.Buffer.WriteString(fmt.Sprintf(format+"\n",
fmtArgs...))

#### type Matcher

```go
type Matcher struct {
}
```

Matches a string against 'simple-patterns': patterns that can have asterisk (*)
wildcards only at the beginning ("ends-with"), at the end ("begins-with"), or
both ("contains"), or not at all ("equals").

For more complex pattern-matching needs, go and unleash the full force of
regular expressions. But I found that in a big portion of pattern-matching
use-cases, I'm just doing "begins-or-ends-or-contains-or-equals" testing. Hence
the conception of the 'simple-pattern'.

There is also an alternative Pattern type in this package. Use Matcher to match
strings against multiple patterns at once, especially if the patterns don't
change often and the matchings occur frequently / repeatedly. In simpler, rarer
one-off matchings, Pattern may be used for simpler "setup-less" matching.

#### func (*Matcher) AddPatterns

```go
func (me *Matcher) AddPatterns(patterns ...string)
```
Adds the specified simple-pattern to me.

#### func (*Matcher) HasWildcardPatterns

```go
func (me *Matcher) HasWildcardPatterns() bool
```
Returns whether any of the simple-patterns specified for this Matcher contains a
*-wildcard.

#### func (*Matcher) IsMatch

```go
func (me *Matcher) IsMatch(s string) bool
```
Matches s against all patterns in me.

#### type Pattern

```go
type Pattern string
```

An 'leaner' alternative to Matcher (see type docs for Matcher). This represents
a single 'simple-pattern' and provides matching methods for one or multiple
values.

#### func (Pattern) AllMatch

```go
func (me Pattern) AllMatch(values ...string) (allMatch bool)
```
Returns true if all specified values match this simple-pattern.

#### func (Pattern) AnyMatches

```go
func (me Pattern) AnyMatches(values ...string) (firstMatch string)
```
Returns the first of the specified values to match this simple-pattern, or empty
if none of them match.

#### func (Pattern) IsMatch

```go
func (me Pattern) IsMatch(value string) bool
```
Returns true if the specified value matches this simple-pattern.

--
**godocdown** http://github.com/robertkrimen/godocdown