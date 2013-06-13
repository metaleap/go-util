package uio

import (
	"archive/zip"
	"encoding/binary"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	usl "github.com/metaleap/go-util/slice"
	ustr "github.com/metaleap/go-util/str"
)

type WatcherHandler func(path string)

var (
	//	The permission bits used in EnsureDirExists(), WriteBinaryFile() and WriteTextFile()
	ModePerm = os.ModePerm
)

//	Implements io.Writer and discards/ignores all Write() calls.
type NoopWriter struct {
}

//	No-op
func (_ *NoopWriter) Write(_ []byte) (n int, err error) {
	return
}

//	Removes anything in path (but not path itself), except those whose name matches any of the specified keepNamePatterns.
func ClearDirectory(path string, keepNamePatterns ...string) (err error) {
	var fileInfos []os.FileInfo
	var matcher ustr.Matcher
	matcher.AddPatterns(keepNamePatterns...)
	if fileInfos, err = ioutil.ReadDir(path); err == nil {
		for _, fi := range fileInfos {
			if fn := fi.Name(); !matcher.IsMatch(fn) {
				if err = os.RemoveAll(filepath.Join(path, fn)); err != nil {
					return
				}
			}
		}
	}
	return
}

//	Copies all files and directories inside srcDirPath to destDirPath.
//	All sub-directories whose name is matched by skipDirs (optional) are skipped.
func CopyAll(srcDirPath, destDirPath string, skipDirs *ustr.Matcher) (err error) {
	var (
		srcPath, destPath string
		fileInfos         []os.FileInfo
	)
	if fileInfos, err = ioutil.ReadDir(srcDirPath); err == nil {
		EnsureDirExists(destDirPath)
		for _, fi := range fileInfos {
			if srcPath, destPath = filepath.Join(srcDirPath, fi.Name()), filepath.Join(destDirPath, fi.Name()); fi.IsDir() {
				if skipDirs == nil || !skipDirs.IsMatch(fi.Name()) {
					CopyAll(srcPath, destPath, skipDirs)
				}
			} else {
				CopyFile(srcPath, destPath)
			}
		}
	}
	return
}

//	Performs an io.Copy from the specified local source file to the specified local destination file.
func CopyFile(srcFilePath, destFilePath string) (err error) {
	var src *os.File
	if src, err = os.Open(srcFilePath); err != nil {
		return
	}
	defer src.Close()
	err = SaveToFile(src, destFilePath)
	return
}

//	Returns true if a directory exists at the specified path.
func DirExists(path string) bool {
	if stat, err := os.Stat(path); err == nil {
		return stat.IsDir()
	}
	return false
}

//	Returns true if all dirOrFileNames exist in dirPath.
func DirsFilesExist(dirPath string, dirOrFileNames ...string) (allExist bool) {
	allExist = true
	var (
		err  error
		stat os.FileInfo
	)
	for _, name := range dirOrFileNames {
		if stat, err = os.Stat(filepath.Join(dirPath, name)); err != nil || stat == nil {
			allExist = false
			break
		}

	}
	return
}

//	If a directory does not exist at the specified path, attempts to create it.
func EnsureDirExists(path string) (err error) {
	if !DirExists(path) {
		if err = EnsureDirExists(filepath.Dir(path)); err == nil {
			err = os.Mkdir(path, ModePerm)
		}
	}
	return
}

//	Extracts a ZIP archive to the local file system.
//	zipFilePath: full file path to the ZIP archive file.
//	targetDirPath: directory path where un-zipped archive contents are extracted to.
//	deleteZipFile: deletes the ZIP archive file upon successful extraction.
func ExtractZipFile(zipFilePath, targetDirPath string, deleteZipFile bool, fileNamesPrefix string, fileNamesToExtract ...string) error {
	var (
		fnames      []string
		fnprefix    string
		efile       *os.File
		zfile       *zip.File
		zfileReader io.ReadCloser
		unzip, err  = zip.OpenReader(zipFilePath)
	)
	if unzip != nil {
		if (err == nil) && (unzip.File != nil) {
			if (fileNamesToExtract != nil) && (len(fileNamesToExtract) > 0) {
				fnames = fileNamesToExtract
				for i, fn := range fnames {
					if strings.HasPrefix(fn, fileNamesPrefix) {
						fnames[i] = fn[len(fileNamesPrefix):]
						fnprefix = fileNamesPrefix
					}
				}
			}
			for _, zfile = range unzip.File {
				if (fnames == nil) || (usl.StrAt(fnames, zfile.FileHeader.Name) >= 0) {
					zfileReader, err = zfile.Open()
					if zfileReader != nil {
						if err == nil {
							efile, err = os.Create(filepath.Join(targetDirPath, fnprefix+zfile.FileHeader.Name))
							if efile != nil {
								if err == nil {
									_, err = io.Copy(efile, zfileReader)
								}
								efile.Close()
							}
						}
						zfileReader.Close()
					}
				}
				if err != nil {
					break
				}
			}
		}
		unzip.Close()
		if deleteZipFile && (err == nil) {
			err = os.Remove(zipFilePath)
		}
	}
	return err
}

//	Returns true if a file (not a directory) exists at the specified path.
func FileExists(path string) bool {
	if stat, err := os.Stat(path); err == nil {
		return !stat.IsDir()
	}
	return false
}

//	If a file with a given base-name and one of a set of extensions exists in the specified directory, returns details on it.
//	The tryLower and tryUpper flags also test for upper-case and lower-case variants of the specified fileBaseName.
func FindFileInfo(dirPath string, fileBaseName string, fileExts []string, tryLower bool, tryUpper bool) (fullFilePath string, fileInfo *os.FileInfo) {
	var (
		stat        os.FileInfo
		err         error
		fext, fpath string
	)
	for _, fext = range fileExts {
		fpath = filepath.Join(dirPath, fileBaseName+fext)
		if stat, err = os.Stat(fpath); err != nil {
			if tryUpper {
				fpath = filepath.Join(dirPath, strings.ToUpper(fileBaseName)+fext)
				stat, err = os.Stat(fpath)
			}
			if (err != nil) && tryLower {
				fpath = filepath.Join(dirPath, strings.ToLower(fileBaseName)+fext)
				stat, err = os.Stat(fpath)
			}
		}
		if (err == nil) && !stat.IsDir() {
			return fpath, &stat
		}
	}
	return "", nil
}

//	Returns true if srcFilePath has been modified after dstFilePath.
//	NOTE: be aware that newer will be returned as true if err is returned as NOT nil,
//	as this is often very convenient for many use-cases.
func IsNewerThan(srcFilePath, dstFilePath string) (newer bool, err error) {
	var out, src os.FileInfo
	newer = true
	if out, err = os.Stat(dstFilePath); err == nil && out != nil {
		if src, err = os.Stat(srcFilePath); err == nil && src != nil {
			newer = src.ModTime().UnixNano() > out.ModTime().UnixNano()
		}
	}
	return
}

//	Reads and returns the binary contents of a file with non-idiomatic error handling:
//	filePath: full local file path
//	panicOnError: true to panic() if an error occurred reading the file
func ReadBinaryFile(filePath string, panicOnError bool) []byte {
	bytes, err := ioutil.ReadFile(filePath)
	if panicOnError && (err != nil) {
		panic(err)
	}
	return bytes
}

//	Reads binary data into the specified interface{} from the specified io.ReadSeeker at the specified offset using the specified binary.ByteOrder.
//	Returns false if data could not be successfully read as specified, otherwise true.
func ReadFromBinary(readSeeker io.ReadSeeker, offset int64, byteOrder binary.ByteOrder, ptr interface{}) bool {
	o, err := readSeeker.Seek(offset, 0)
	if (o != offset) || (err != nil) {
		return false
	}
	if err = binary.Read(readSeeker, byteOrder, ptr); err != nil {
		return false
	}
	return true
}

//	Reads and returns the contents of a text file with non-idiomatic error handling:
func ReadTextFile(filePath string, panicOnError bool, defVal string) string {
	bytes, err := ioutil.ReadFile(filePath)
	if err == nil {
		return string(bytes)
	}
	if panicOnError && (err != nil) {
		panic(err)
	}
	return defVal
}

//	Performs an io.Copy() from the specified io.Reader to the specified local file.
func SaveToFile(r io.Reader, filename string) (err error) {
	var file *os.File
	file, err = os.Create(filename)
	if file != nil {
		defer file.Close()
		if err == nil {
			_, err = io.Copy(file, r)
		}
	}
	return
}

func WalkAllDirs(dirPath string, visitor WalkerVisitor) []error {
	return NewDirWalker(true, visitor, nil).Walk(dirPath)
}

func WalkAllFiles(dirPath string, visitor WalkerVisitor) []error {
	return NewDirWalker(true, nil, visitor).Walk(dirPath)
}

func WalkDirsIn(dirPath string, visitor WalkerVisitor) []error {
	w := NewDirWalker(false, visitor, nil)
	w.VisitSelf = false
	return w.Walk(dirPath)
}

func WalkFilesIn(dirPath string, visitor WalkerVisitor) []error {
	w := NewDirWalker(false, nil, visitor)
	w.VisitSelf = false
	return w.Walk(dirPath)
}

//	A short-hand for ioutil.WriteFile, without needing to specify os.ModePerm.
//	Also ensures the target file's directory exists.
func WriteBinaryFile(filePath string, contents []byte) error {
	EnsureDirExists(filepath.Dir(filePath))
	return ioutil.WriteFile(filePath, contents, ModePerm)
}

//	A short-hand for ioutil.WriteFile, without needing to specify os.ModePerm or string-conversion.
//	Also ensures the target file's directory exists.
func WriteTextFile(filePath, contents string) error {
	return WriteBinaryFile(filePath, []byte(contents))
}

func watchRunHandler(dirPath string, namePattern ustr.Pattern, handler WatcherHandler) []error {
	vis := func(fullPath string) (keepWalking bool) {
		keepWalking = true
		if namePattern.IsMatch(filepath.Base(fullPath)) {
			handler(fullPath)
		}
		return
	}
	w := NewDirWalker(false, vis, vis)
	w.VisitSelf = false
	w.VisitDirsFirst = true
	return w.Walk(dirPath)
}
