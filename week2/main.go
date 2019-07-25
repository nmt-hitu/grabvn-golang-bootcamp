package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

func RemoveSpecialCharacter(text string) string {

	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	validCharacter := reg.ReplaceAllString(text, " ")

	return validCharacter
}

//Read the content of files and remove special character
func ReadContentByFile(file_name string, channel chan []string) {
	file, err := os.Open(file_name)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	content, _ := ioutil.ReadAll(reader)
	validate_text := RemoveSpecialCharacter(string(content))
	valid_string := strings.Split(validate_text, " ")

	fmt.Printf("Valid text: %s\n", valid_string)
	channel <- valid_string
}

//Count ocurrency of data
func CountOfOccurences(data []string, counts map[string]int) {
	for _, i := range data {
		counts[i]++
	}
}

func main() {
	//Create list to store for all file name
	var InputFileList []string

	//Path of files
	InputFilePath := "./DataTest"

	//creating chanel to count
	counts := make(map[string]int)
	//channel between routines
	channel := make(chan []string)

	//Create files from the directory
	files, err := ioutil.ReadDir(InputFilePath)
	//If no file exist will return error
	if err != nil {
		log.Fatal(err)
	}

	//Save file to the empty list that we created before
	for _, file := range files {
		InputFileList = append(InputFileList, file.Name())
	}

	//Read content of files
	//Number of routines depend on how many files exist
	for _, input := range InputFileList {
		go ReadContentByFile(InputFilePath+"/"+input, channel)
	}
	//Count the occurency of each character
	go func() {
		for {
			CountOfOccurences(<-channel, counts)
		}
	}()
	//Time wait for all routines
	time.Sleep(5 * time.Second)
	//Print the result
	for key, value := range counts {
		fmt.Printf("Count of character %s is : %d \n", key, value)
	}

}
