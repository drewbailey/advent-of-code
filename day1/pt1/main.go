package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	count := 0
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		i, err := strconv.Atoi(scan.Text())
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		count += i
	}
	log.Printf("calibration: %v", count)
}
