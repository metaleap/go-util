package io

import (
	"archive/zip"
	"encoding/binary"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	ustr "github.com/metaleap/go-util/str"
)

//	Performs an io.Copy from the specified local source file to the specified local destination file.
func CopyFile (srcFilePath, destFilePath string) (err error) {
	var src *os.File
	if src, err = os.Open(srcFilePath); err != nil { return }
	defer src.Close()
	err = SaveToFile(src, destFilePath)
	return
}

//	Returns true if a directory exists at the specified path.
func DirExists (path string) bool {
	var stat, err = os.Stat(path)
	if (err == nil) && (stat != nil) { return stat.IsDir() }
	return false
}

//	If a directory does not exist at the specified path, attempts to create it.
func EnsureDirExists (path string) error {
	var err error
	if !DirExists(path) {
		if err = EnsureDirExists(filepath.Dir(path)); err == nil {
			err = os.Mkdir(path, os.ModePerm)
		}
	}
	return err
}

//	Extracts a ZIP archive to the local file system.
//	zipFilePath: full file path to the ZIP archive file
//	targetDirPath: directory path where un-zipped archive contents are extracted to
//	deleteZipFile: deletes the ZIP archive file upon successful extraction
func ExtractZipFile (zipFilePath, targetDirPath string, deleteZipFile bool, fileNamesPrefix string, fileNamesToExtract ... string) error {
	var fnames []string = nil
	var fnprefix string = ""
	var efile *os.File
	var zfile *zip.File
	var zfileReader io.ReadCloser
	var unzip, err = zip.OpenReader(zipFilePath)
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
				if (fnames == nil) || (ustr.InSliceAt(fnames, zfile.FileHeader.Name) >= 0) {
					zfileReader, err = zfile.Open()
					if zfileReader != nil {
						if (err == nil) {
							efile, err = os.Create(filepath.Join(targetDirPath, fnprefix + zfile.FileHeader.Name))
							if efile != nil {
								if (err == nil) {
									_, err = io.Copy(efile, zfileReader)
								}
								efile.Close()
							}
						}
						zfileReader.Close()
					}
				}
				if (err != nil) {
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

//	Returns true if a file exists at the specified path.
func FileExists (path string) bool {
	var stat, err = os.Stat(path)
	if (err == nil) && (stat != nil) { return !stat.IsDir() }
	return false
}

//	If a file with a given base-name and one of a number of extensions exists in the specified directory, returns details on it.
//	The tryLower and tryUpper flags also test for upper-case and lower-case variants of the specified fileBaseName.
func FileExistsPath (dirPath string, fileBaseName string, fileExts []string, tryLower bool, tryUpper bool) (fullFilePath string, modTime time.Time, size int64) {
	var stat os.FileInfo
	var err error
	var fext, fpath string
	for _, fext = range fileExts {
		fpath = filepath.Join(dirPath, fileBaseName + fext)
		if stat, err = os.Stat(fpath); err != nil {
			if tryUpper {
				fpath = filepath.Join(dirPath, strings.ToUpper(fileBaseName) + fext)
				stat, err = os.Stat(fpath)
			}
			if (err != nil) && tryLower {
				fpath = filepath.Join(dirPath, strings.ToLower(fileBaseName) + fext)
				stat, err = os.Stat(fpath)
			}
		}
		if (err == nil) && (stat != nil) && !stat.IsDir() {
			return fpath, stat.ModTime(), stat.Size()
		}
	}
	return "", time.Time {}, 0
}

//	Reads and returns the binary contents of a file with non-idiomatic error handling:
//	filePath: full local file path
//	panicOnError: true to panic() if an error occurred reading the file
func ReadBinaryFile (filePath string, panicOnError bool) []byte {
	var bytes, err = ioutil.ReadFile(filePath)
	if panicOnError && (err != nil) { panic(err) }
	return bytes
}

//	Reads binary data into the specified interface{} from the specified io.ReadSeeker at the specified offset using the specified binary.ByteOrder.
//	Returns false if data could not be successfully read as specified, otherwise true.
func ReadFromBinary (readSeeker io.ReadSeeker, offset int64, byteOrder binary.ByteOrder, ptr interface{}) bool {
	var o, err = readSeeker.Seek(offset, 0)
	if (o != offset) || (err != nil) { return false }
	if err = binary.Read(readSeeker, byteOrder, ptr); err != nil { return false }
	return true
}

//	Reads and returns the contents of a text file with non-idiomatic error handling:
//	filePath: full local file path
//	panicOnError: true to panic() if an error occurred reading the file, or false to return defVal in the case of error
//	defVal: the string value to return if the file couldn't be read successfully
func ReadTextFile (filePath string, panicOnError bool, defVal string) string {
	var bytes, err = ioutil.ReadFile(filePath)
	if err == nil { return string(bytes) }
	if panicOnError && (err != nil) { panic(err) }
	return defVal
}

//	Performs an io.Copy() from the specified io.Reader to the specified local file.
func SaveToFile (r io.Reader, filename string) (err error) {
	var file *os.File
	file, err = os.Create(filename)
	if file != nil {
		defer file.Close()
		if err == nil { _, err = io.Copy(file, r) }
	}
	return
}

//	Recursively walks along a directory hierarchy, calling the specified callback function for each file encountered.
//	dirPath: the path of the directory in which to start walking
//	fileSuffix: optional; if specified, fileFunc is only called for files whose name has this suffix
//	fileFunc: callback function called per file. Returns true to keep recursing into sub-dirs. Arguments: full file path and current recurseSubDirs value
//	recurseSubDirs: true to recurse into sub-directories.
func WalkDirectory (dirPath, fileSuffix string, fileFunc func (string, bool) bool, recurseSubDirs bool) error {
	var fileInfos, err = ioutil.ReadDir(dirPath)
	if err == nil {
		for _, fi := range fileInfos {
			if !fi.IsDir() {
				if (len(fileSuffix) == 0) || strings.HasSuffix(fi.Name(), fileSuffix) {
					recurseSubDirs = fileFunc(filepath.Join(dirPath, fi.Name()), recurseSubDirs)
				}
			}
		}
		if recurseSubDirs {
			for _, fi := range fileInfos {
				if fi.IsDir() {
					if err = WalkDirectory(filepath.Join(dirPath, fi.Name()), fileSuffix, fileFunc, recurseSubDirs); err != nil {
						break
					}
				}
			}
		}
	}
	return err
}

//	A short-hand for ioutil.WriteFile, without needing to specify os.ModePerm.
func WriteBinaryFile (filePath string, contents []byte) error {
	return ioutil.WriteFile(filePath, contents, os.ModePerm)
}

//	A short-hand for ioutil.WriteFile, without needing to specify os.ModePerm or string-conversion.
func WriteTextFile (filePath, contents string) error {
	return WriteBinaryFile(filePath, []byte(contents))
}
