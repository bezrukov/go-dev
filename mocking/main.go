package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

const FinalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer, sleep Sleeper)  {
	for i := countdownStart; i > 0; i-- {
		sleep.Sleep()
		fmt.Fprintln(out, i)
	}

	sleep.Sleep()
	fmt.Fprint(out, FinalWord)
}
