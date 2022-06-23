package main

// TODO: Add support for loading file

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"os"
// )

// func readFile(filepath string) string {
// 	body, err := ioutil.ReadFile(filepath)
// 	if err != nil {
// 		log.Fatalf("unable to read file: %v", err)
// 	}
// 	return string(body)
// }

// func writeFile(filepath string, content string) {
// 	file, err := os.Create(filepath)
// 	if err != nil {
// 		log.Fatalf("unable to wrote file: %v", err)
// 	}
// 	defer file.Close()
// 	_, err2 := file.WriteString(content)

// 	if err2 != nil {
// 		log.Fatalf("unable to wrote file: %v", err)
// 	}
// 	fmt.Println("done")
// }