package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getTemperature() {
	dat, err := ioutil.ReadFile("/sys/class/thermal/thermal_zone0/temp")
	check(err)
	tempStr := string(dat)
	temp, err := strconv.Atoi(tempStr[:len(tempStr)-1])
	check(err)
	temp = temp / 1000
	fmt.Println("Temperatura da CPU:\t", temp, "°C")
}

func main() {
	var watchTime int
	flag.IntVar(&watchTime, "t", 0, "tempo(em segundos) entre cada interação do comando.")

	flag.Parse()
	watch := watchTime > 0
	if !watch {
		getTemperature()
	}
	for watch {
		print("\033[H\033[2J")
		fmt.Print(time.Now().Format("2006-01-02 15:04:05"), "\n\n")
		getTemperature()
		time.Sleep(time.Duration(watchTime) * time.Second)
	}

}
