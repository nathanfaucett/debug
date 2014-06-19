package debugger

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	sec    = int64(1000)
	min    = int64(60 * 1000)
	hour   = int64(60 * min)
	prev   = make(map[string]int64)
	colors = [...]int{
		32, // green
		33, // yellow
		34, // blue
		35, // magenta
		36, // cyan
		37, // white
		38,
	}
	prevColor = 0
)

func stylize(str string, color int) string {
	return "\x1B[" + strconv.Itoa(color) + "m" + str + "\x1B[39m"
}

func color() int {
	c := colors[prevColor%len(colors)]
	prevColor++
	return c
}

func humanize(ms int64) string {
	if ms >= hour {
		return fmt.Sprintf("%.2f", float64(ms)/float64(hour)) + "h"
	}
	if ms >= min {
		return fmt.Sprintf("%.2f", float64(ms)/float64(min)) + "m"
	}
	if ms >= sec {
		return fmt.Sprintf("%.3f", float64(ms)/float64(sec)) + "s"
	}

	return fmt.Sprintf("%d", ms) + "ms"
}

type debugInterface interface {
	Log(string)
	Warning(string)
	Error(string)
}

type debug struct {
	name  string
	color int
}

func newDebug(name string) *debug {
	this := new(debug)

	this.name = name
	this.color = color()

	return this
}

func (this *debug) timeDifference() string {
	current := time.Now().UnixNano() / 1e6
	last := prev[this.name]
	ms := current - last
	prev[this.name] = current

	return humanize(ms)
}

func (this *debug) Log(msg string) {
	fmt.Println(stylize(this.name+": ", this.color) + stylize(msg, 90) + stylize(" +"+this.timeDifference(), this.color))
}

func (this *debug) Warning(msg string) {
	fmt.Println(stylize(this.name, this.color) + stylize(" Warning: "+msg, 33) + stylize(" +"+this.timeDifference(), this.color))
}

func (this *debug) Error(msg string) {
	fmt.Println(stylize(this.name, this.color) + stylize(" Error: "+msg, 31) + stylize(" +"+this.timeDifference(), this.color))
}

type emptyDebug struct{}

func newEmptyDebug() *emptyDebug {
	this := new(emptyDebug)
	return this
}

func (this *emptyDebug) Log(msg string)     {}
func (this *emptyDebug) Warning(msg string) {}
func (this *emptyDebug) Error(msg string)   {}

func Debug(name string) debugInterface {
	env := os.Getenv("GO_ENVIRONMENT_NAME")

	if env != "Dev" && env != "Development" {
		return newEmptyDebug()
	}
	prev[name] = time.Now().UnixNano() / 1e6

	return newDebug(name)
}
