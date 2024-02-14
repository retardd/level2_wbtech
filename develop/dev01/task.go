package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	// Получаем точное время с помощью библиотеки NTP
	ntpTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting NTP time: %v\n", err)
		os.Exit(1)
	}

	// Печатаем точное время
	fmt.Println("Точное время (NTP):", ntpTime.Format(time.RFC3339))
}
