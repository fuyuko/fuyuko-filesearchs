package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

//FileWithDir struct with its functions
type FileWithDir struct {
	f fs.DirEntry
	d string
}

func (fwd *FileWithDir) File() fs.DirEntry {
	return fwd.f
}

func (fwd *FileWithDir) Dir() string {
	return fwd.d
}

func (fwd *FileWithDir) Name() string {
	return fwd.f.Name()
}

func (fwd *FileWithDir) NameWithPath() string {
	return fwd.Dir() + "/" + fwd.Name()
}

func (fwd *FileWithDir) FileExt() string {
	return filepath.Ext(fwd.Name())
}

func (fwd *FileWithDir) Size() int64 {
	fileStat, err := os.Stat(fwd.NameWithPath())

	if err != nil {
		fmt.Println("error occurred.")
		log.Fatal(err)
	}

	return fileStat.Size()
}

//conditions
/*func FileTypeCond(ext string, fwd *FileWithDir) bool {
	return fwd.FileExt() == ext
}

func FileMinSizeCond(size int64, fwd *FileWithDir) bool{
	return fwd.Size() >= size
}
*/

func findFiles(dir string) []FileWithDir {
	var files []FileWithDir

	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range dirEntries {
		fileStat, err := os.Stat(dir + "/" + file.Name())

		//test
		fmt.Println("direntry name: ", file.Name())

		if err != nil {
			fmt.Println("error occurred.")
			log.Fatal(err)
		}

		if fileStat.IsDir() {
			//handle subdir names
			subdir := dir + "/" + file.Name()
			files = append(files, findFiles(subdir)...)
		} else {
			//apply conditions here

			files = append(files, FileWithDir{f: file, d: dir})
		}

	}

	return files
}

func main() {
	filesFound := findFiles(".")

	fmt.Println("File search completed")

	for _, file := range filesFound {

		fmt.Println("File Name:", file.Name()) // Base name of the file
		fmt.Println("File Size:", file.Size())
	}
}
