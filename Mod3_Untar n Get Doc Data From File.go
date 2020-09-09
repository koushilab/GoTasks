package main

import (
	"archive/tar"
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func PrintFatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ExtractTarGz(gzipStream io.Reader) {
	uncompressedStream, err := gzip.NewReader(gzipStream)
	if err != nil {
		log.Fatal("ExtractTarGz: NewReader failed")
	}

	tarReader := tar.NewReader(uncompressedStream)

	for true {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("ExtractTarGz: Next() failed: %s", err.Error())
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.Mkdir(".\\Test2\\"+header.Name, 0755); err != nil {
				log.Fatalf("ExtractTarGz: Mkdir() failed: %s", err.Error())
			}
		case tar.TypeReg:
			fmt.Println(header.Name)
			outFile, err := os.Create(".\\Test2\\" + header.Name)
			if err != nil {
				log.Fatalf("ExtractTarGz: Create() failed: %s", err.Error())
			}
			if _, err := io.Copy(outFile, tarReader); err != nil {
				log.Fatalf("ExtractTarGz: Copy() failed: %s", err.Error())
			}
			outFile.Close()

		default:
			log.Fatalf(
				"ExtractTarGz: uknown type: %s in %s",
				header.Typeflag,
				header.Name)
		}

	}
}

func firstCount(filename string) int {
	f3, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	PrintFatalError(err)
	defer f3.Close()
	scanner := bufio.NewScanner(f3)
	line := 1
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "DOCUMENTATION = ") {
			time.Sleep(100)
			return line
		}
		line++
	}
	return line
}

func secondCount(filename string, x int) int {
	f3, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	PrintFatalError(err)
	defer f3.Close()
	scanner := bufio.NewScanner(f3)
	line := 1
	for scanner.Scan() {
		if (strings.Contains(scanner.Text(), "'''")) && line > x {
			time.Sleep(100)
			return line
		}
		if (strings.Contains(scanner.Text(), "\"\"\"")) && line > x {
			time.Sleep(100)
			return line
		}
		line++
	}
	return line
}

func main() {

	r, err := os.Open("./netapp-ontap-20.8.0.tar.gz")
	PrintFatalError(err)
	ExtractTarGz(r)

	modulePath := "E:\\Go Tasks\\Final\\Test2\\plugins\\modules\\"
	yamlPath := "E:\\Go Tasks\\Final\\Test2\\examples\\"

	files, err := ioutil.ReadDir(modulePath)
	PrintFatalError(err)

	for _, file := range files {

		filename := modulePath + file.Name()
		firstline := firstCount(filename)
		secondline := secondCount(filename, firstline)

		fmt.Println(filename, firstline)
		fmt.Println(filename, secondline)

		f2, err := os.Create(yamlPath + file.Name() + ".yaml")
		time.Sleep(100)
		PrintFatalError(err)
		defer f2.Close()

		f3, err := os.OpenFile(filename, os.O_RDONLY, 0666)
		time.Sleep(100)
		PrintFatalError(err)
		defer f3.Close()
		scanner := bufio.NewScanner(f3)

		line := 1
		for scanner.Scan() {
			line++
			if line > firstline+1 && line <= secondline {
				wr := bufio.NewWriter(f2)
				wr.WriteString(scanner.Text())
				wr.WriteString("\n")
				wr.Flush()
				time.Sleep(40)
			}

		}
	}
}
