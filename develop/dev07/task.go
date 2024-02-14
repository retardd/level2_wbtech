package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone)

		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-or(append(channels[2:], orDone)...):
			}
		}
	}()

	return orDone
}

func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите количество каналов: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	numChannels, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	channels := make([]<-chan interface{}, numChannels)
	times := make([]int, numChannels)

	for i := 0; i < numChannels; i++ {
		fmt.Printf("Введите время работы для канала %d (в секундах): ", i+1)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		duration, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}
		times[i] = duration
	}
	for i := 0; i < numChannels; i++ {
		channels[i] = sig(time.Duration(times[i]) * time.Second)
	}
	start := time.Now()
	<-or(channels...)
	fmt.Printf("done after %v\n", time.Since(start))
}
