package filehandling

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gobuffalo/packr/v2"
)

func ReadingFilesPractice() {
	// readingFileWithAbsolutePAth()
	// readingFileWithFilePathAsCommandLine()
	// bundlingTextFileAlongWithBinary()
	// readingFilesInChunks()
	readingFilesLineByLine()
}

func readingFileWithAbsolutePAth() {
	data, err := ioutil.ReadFile("/Users/sudeepthamouniganji/GoLearn/FileHandling/test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}

func readingFileWithFilePathAsCommandLine() { //To run use: go run main.go -fpath=/Users/sudeepthamouniganji/GoLearn/FileHandling/test.txt
	fptr := flag.String("fpath", "test.txt", "file path to read from") // The first is the name of the flag, second is the default value and the third is a short description of the flag
	flag.Parse()                                                       //lag.Parse() should be called before any flag is accessed by the program.
	// fmt.Println("value of fpath is", *fptr)
	data, err := ioutil.ReadFile(*fptr)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}

func bundlingTextFileAlongWithBinary() {
	box := packr.New("fileBox", "../filehandling")
	data, err := box.FindString("test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", data)
}

func readingFilesInChunks() {
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	r := bufio.NewReader(f)
	b := make([]byte, 3)
	for {
		n, err := r.Read(b)
		if err != nil {
			fmt.Println("Error reading file:", err)
			break
		}
		fmt.Println(string(b[0:n]))
	}

}

func readingFilesLineByLine() {
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}
