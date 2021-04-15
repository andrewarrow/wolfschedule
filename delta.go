package main

import "time"

type Delta struct {
	Val   int
	Text  string
	Month int
	Time  time.Time
}

func NewDelta(val int, text string, month int) Delta {
	d := Delta{val, text, month, time.Now()}
	return d
}
