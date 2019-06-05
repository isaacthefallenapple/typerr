package main

import (
	"fmt"
	"time"
)

type result struct {
	input    []string
	length   int
	mistakes int
	err      error
	time     time.Duration
}

func (r *result) Add(s string, n int, err error) {
	r.input = append(r.input, s)
	r.length += len(s)
	r.mistakes += n
	r.err = err
}
func (r *result) Ratio() float64 {
	if r.length == 0 {
		return 0
	}
	return float64(r.mistakes) / float64(r.length)
}
func (r *result) MPS() float64 {
	return float64(r.mistakes) / r.time.Seconds()
}
func (r result) String() string {
	plural := "s"
	n := r.mistakes
	if n == 1 { plural = "" }
	return fmt.Sprintf("You made %d mistake%s in %s. That is a %.2f%% error rate at %.2f errors per second.",
		n, plural, r.time.Truncate(time.Millisecond).String(), r.Ratio()*100, r.MPS())
}
