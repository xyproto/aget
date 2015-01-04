package main

/*
 *
 * Colored text output
 *
 * Only supports a few selected colors
 *
 */

import (
	"fmt"
	"os"
)

type Output struct {
	color   bool
	enabled bool
}

func NewOutput(color bool, enabled bool) *Output {
	return &Output{color, enabled}
}

func (o *Output) Err(msg string) {
	if o.enabled {
		fmt.Fprintf(os.Stderr, o.DarkRed(msg)+"\n")
	}
}

// Print the message as red text and exit with errorlevel 1
func (o *Output) Exit(msg string) {
	o.Err(msg)
	os.Exit(1)
}

func (o *Output) Println(msg string) {
	if o.enabled {
		fmt.Println(msg)
	}
}

func (o *Output) IsEnabled() bool {
	return o.enabled
}

func (o *Output) colorOn(num1 int, num2 int) string {
	if o.color {
		return fmt.Sprintf("\033[%d;%dm", num1, num2)
	}
	return ""
}

func (o *Output) colorOff() string {
	if o.color {
		return "\033[0m"
	}
	return ""
}

// TODO: Consider generating the following functions as closures instead

func (o *Output) DarkRed(s string) string {
	return o.colorOn(0, 31) + s + o.colorOff()
}

func (o *Output) LightGreen(s string) string {
	return o.colorOn(1, 32) + s + o.colorOff()
}

func (o *Output) DarkGreen(s string) string {
	return o.colorOn(0, 32) + s + o.colorOff()
}

func (o *Output) LightYellow(s string) string {
	return o.colorOn(1, 33) + s + o.colorOff()
}

func (o *Output) DarkYellow(s string) string {
	return o.colorOn(0, 33) + s + o.colorOff()
}

func (o *Output) LightBlue(s string) string {
	return o.colorOn(1, 34) + s + o.colorOff()
}

func (o *Output) DarkBlue(s string) string {
	return o.colorOn(0, 34) + s + o.colorOff()
}

func (o *Output) LightCyan(s string) string {
	return o.colorOn(1, 36) + s + o.colorOff()
}

func (o *Output) LightPurple(s string) string {
	return o.colorOn(1, 35) + s + o.colorOff()
}

func (o *Output) DarkPurple(s string) string {
	return o.colorOn(0, 35) + s + o.colorOff()
}

func (o *Output) DarkGray(s string) string {
	return o.colorOn(1, 30) + s + o.colorOff()
}
