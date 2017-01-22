package hello

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func openfileWithIoutil() {
	filename := "./extras/rabbits.txt"

	// read in file, one command loads file into content variable
	// if an error occurs returns it as the second return value
	// if no error, err will be nil
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		// log is a nifty little utility which can also output
		// in this case, a fatal log will halt the program
		log.Fatalln("Error reading file", filename)
	}

	// content returned as []byte not string
	// so must cast to string first and then can display
	// fmt.Println(string(content))

	// write back to new file
	// see documentation for which methods take what type
	outfile := "output.txt"
	err = ioutil.WriteFile(outfile, content, 0644)
	if err != nil {
		log.Fatalln("Error writing file: ", err)
	}

}

func openfilewithsys() {
	filepath := "./extras/rabbits.txt"

	// read the file by using buffer io
	// The Scan() method scans token by token, default token is ScanLines
	// ScanWords and other methods available, See http://golang.org/pkg/bufio/
	f, _ := os.Open(filepath)
	// defer is a nifty bit of magic which automatically runs
	// the command before exiting in this case close file
	// good practice to defer right after opening
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// returns text prior to token
		line := scanner.Text()
		line += "1"
		// fmt.Println(line)
	}

	// check if any errors occurred
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Write file by creating and then writing string by string
	f, err := os.Create("output1.txt")
	if err != nil {
		log.Fatalln("Error creating file: ", err)
	}
	defer f.Close()

	for _, str := range []string{"apple", "banana", "carrot"} {
		bytes, err := f.WriteString(str)
		if err != nil {
			log.Fatalln("Error writing string: ", err)
		}
		fmt.Printf("Wrote %d bytes to file\n", bytes)
	}
}
