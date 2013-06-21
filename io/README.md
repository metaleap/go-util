# uio
--
    import "github.com/metaleap/go-util/io"

Various line-savers for common I/O needs.

## Usage

```go
var (
	//	The permission bits used in EnsureDirExists(), WriteBinaryFile() and WriteTextFile()
	ModePerm = os.ModePerm
)
```

#### func  ClearDirectory

```go
func ClearDirectory(path string, keepNamePatterns ...string) (err error)
```
Removes anything in path (but not path itself), except those whose name matches
any of the specified keepNamePatterns.

#### func  CopyAll

```go
func CopyAll(srcDirPath, destDirPath string, skipDirs *ustr.Matcher) (err error)
```
Copies all files and directories inside srcDirPath to destDirPath. All
sub-directories whose name is matched by skipDirs (optional) are skipped.

#### func  CopyFile

```go
func CopyFile(srcFilePath, destFilePath string) (err error)
```
Performs an io.Copy from the specified local source file to the specified local
destination file.

#### func  DirExists

```go
func DirExists(path string) bool
```
Returns true if a directory exists at the specified path.

#### func  DirsFilesExist

```go
func DirsFilesExist(dirPath string, dirOrFileNames ...string) (allExist bool)
```
Returns true if all dirOrFileNames exist in dirPath.

#### func  EnsureDirExists

```go
func EnsureDirExists(path string) (err error)
```
If a directory does not exist at the specified path, attempts to create it.

#### func  FileExists

```go
func FileExists(path string) bool
```
Returns true if a file (not a directory) exists at the specified path.

#### func  FindFileInfo

```go
func FindFileInfo(dirPath string, fileBaseName string, fileExts []string, tryLower bool, tryUpper bool) (fullFilePath string, fileInfo *os.FileInfo)
```
If a file with a given base-name and one of a set of extensions exists in the
specified directory, returns details on it. The tryLower and tryUpper flags also
test for upper-case and lower-case variants of the specified fileBaseName.

#### func  IsNewerThan

```go
func IsNewerThan(srcFilePath, dstFilePath string) (newer bool, err error)
```
Returns whether `srcFilePath` has been modified later than `dstFilePath`. NOTE:
be aware that `newer` will be returned as `true` if `err` is returned as *not*
`nil`, since that is often very convenient for many use-cases.

#### func  ReadBinaryFile

```go
func ReadBinaryFile(filePath string, panicOnError bool) []byte
```
Reads and returns the binary contents of a file with non-idiomatic error
handling, mostly for one-off `package main`s.

#### func  ReadTextFile

```go
func ReadTextFile(filePath string, panicOnError bool, defVal string) string
```
Reads and returns the contents of a text file with non-idiomatic error handling,
mostly for one-off `package main`s.

#### func  SaveToFile

```go
func SaveToFile(r io.Reader, filename string) (err error)
```
Performs an io.Copy() from the specified io.Reader to the specified local file.

#### func  WalkAllDirs

```go
func WalkAllDirs(dirPath string, visitor WalkerVisitor) []error
```
Calls `visitor` for `dirPath` and all descendent directories (but not files).

#### func  WalkAllFiles

```go
func WalkAllFiles(dirPath string, visitor WalkerVisitor) []error
```
Calls `visitor` for all files (but not directories) directly or indirectly
descendent to `dirPath`.

#### func  WalkDirsIn

```go
func WalkDirsIn(dirPath string, visitor WalkerVisitor) []error
```
Calls `visitor` for all directories (but not files) in `dirPath`, but not their
sub-directories and not `dirPath` itself.

#### func  WalkFilesIn

```go
func WalkFilesIn(dirPath string, visitor WalkerVisitor) []error
```
Calls `visitor` for all files (but not directories) in `dirPath`, but not for
any in sub-directories.

#### func  WriteBinaryFile

```go
func WriteBinaryFile(filePath string, contents []byte) error
```
A short-hand for ioutil.WriteFile, without needing to specify os.ModePerm. Also
ensures the target file's directory exists.

#### func  WriteTextFile

```go
func WriteTextFile(filePath, contents string) error
```
A short-hand for ioutil.WriteFile, without needing to specify os.ModePerm or
string-conversion. Also ensures the target file's directory exists.

#### type DirWalker

```go
type DirWalker struct {
	//	Walk() returns a slice of all errors encountered but
	//	to cancel walking upon the first error, set this to true.
	BreakOnError bool

	//	After invoking DirVisitor on the specified directory, by default
	//	its files get visited first before visiting its sub-directories.
	//	If VisitDirsFirst is true, then files get visited last, after
	//	having visited all sub-directories.
	VisitDirsFirst bool

	//	If false, only the items in the specified directory get visited
	//	(and the directory itself if `VisitSelf`), but no items inside its sub-directories.
	VisitSubDirs bool

	//	Defaults to `true` if initialized via `NewDirWalker()`.
	VisitSelf bool

	//	Called for every directory being visited during Walk().
	DirVisitor WalkerVisitor

	//	Called for every file being visited during Walk().
	FileVisitor WalkerVisitor
}
```

Provides recursive directory walking with a variety of options.

#### func  NewDirWalker

```go
func NewDirWalker(deep bool, dirVisitor, fileVisitor WalkerVisitor) (me *DirWalker)
```
Initializes and returns a new DirWalker with the specified (optional) visitors.
The deep argument sets the VisitSubDirs field.

#### func (*DirWalker) Walk

```go
func (me *DirWalker) Walk(dirPath string) (errs []error)
```
Initiates me walking through the specified directory.

#### type NoopWriter

```go
type NoopWriter struct {
}
```

Implements io.Writer and discards/ignores all Write() calls.

#### func (*NoopWriter) Write

```go
func (_ *NoopWriter) Write(_ []byte) (n int, err error)
```
No-op

#### type WalkerVisitor

```go
type WalkerVisitor func(fullPath string) (keepWalking bool)
```

Used for DirWalker.DirVisitor and DirWalker.FileVisitor. Always return
keepWalking as true unless you want to immediately terminate a Walk() early.

#### type Watcher

```go
type Watcher struct {
}
```

File-watching is not allowed and not necessary on Google App Engine. So this is
a "polyfil" empty struct with no-op methods.

#### func  NewWatcher

```go
func NewWatcher() (me *Watcher, err error)
```
Always returns a new Watcher, even if err is not nil.

#### func (*Watcher) Close

```go
func (me *Watcher) Close() (err error)
```
No-op

#### func (*Watcher) Go

```go
func (me *Watcher) Go()
```
Starts watching. A never-ending loop designed to be called in a new go-routine.

#### func (*Watcher) WatchIn

```go
func (me *Watcher) WatchIn(dirPath string, namePattern ustr.Pattern, runHandlerNow bool, handler WatcherHandler) (errs []error)
```
Watches dirs/files (whose name matches the specified pattern) inside the
specified dirPath for change events.

handler is invoked whenever a change event is observed, providing the full file
path.

runHandlerNow allows immediate one-off invokation of handler. This will Walk()
dirPath. This is for the use-case pattern "load those files now, then reload in
exactly the same way whenever they are modified"

#### type WatcherHandler

```go
type WatcherHandler func(path string)
```

--
**godocdown** http://github.com/robertkrimen/godocdown