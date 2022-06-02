package componen

import (
	"fmt"
	"time"
)

var months = [...]string{
	"Januari",
	"Februari",
	"Maret",
	"April",
	"Mei",
	"Juni",
	"Juli",
	"Agustus",
	"September",
	"Oktober",
	"November",
	"Desember",
}

var days = [...]string{
	"Minggu",
	"Senin",
	"Selasa",
	"Rabu",
	"Kamis",
	"Jumat",
	"Sabtu",
}

func DateToday() *time.Time {
	now := time.Now()
	return &now
}

func DateTodayLocal() *time.Time {
	now := time.Now().Local()
	return &now
}

func DateTodayRange() (*time.Time, *time.Time) {
	now := DateTodayLocal()
	dateStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	return &dateStart, now
}

func DateBackwardMonthRange(month int) (*time.Time, *time.Time) {
	now := DateTodayLocal()
	dateBackward := now.AddDate(0, -month, 0)
	return now, &dateBackward
}

func FormatDateIdn(t *time.Time) string {

	st := fmt.Sprintf("%s, %d %s %d",
		days[t.Weekday()], t.Day(), months[t.Month()-1], t.Year())

	if st == "Senin, 1 Januari 1" {
		return "-"
	}

	return st
}
func FormatDateIdnNotDay(t *time.Time) string {
	return fmt.Sprintf("%d %s %d",
		t.Day(), months[t.Month()-1], t.Year())
}

func FormatDateIDInMinut(t *time.Time) string {
	return fmt.Sprintf("%s, %d %s %d %d:%d:%d",
		days[t.Weekday()], t.Day(), months[t.Month()-1], t.Year(), t.Hour(), t.Minute(), t.Second())
}
func FormatIDInMinut(t *time.Time) string {
	return fmt.Sprintf("%d:%d:%d",
		t.Hour(), t.Minute(), t.Second())
}
