package main

import (
	"testing"
	"time"

	"github.com/beevik/ntp"
)

func TestFetchNTPTime(t *testing.T) {
	// Подключаемся к NTP серверу
	ntpTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		t.Fatalf("Failed to fetch time from NTP server: %v", err)
	}

	// Получаем текущее локальное время
	localTime := time.Now()

	// Проверяем, что разница времени между NTP и локальным временем не слишком велика
	// Погрешность времени устанавливаем в 5 секунд
	delta := localTime.Sub(ntpTime)
	if delta > 5*time.Second || delta < -5*time.Second {
		t.Errorf("Time difference between NTP and local time is too big: %v", delta)
	}
}
