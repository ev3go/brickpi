![Gopherbrick](gopherbrick.png)
# brickpi provides BrickPi-specific functions for the Go ev3dev interface

[![Build Status](https://travis-ci.org/ev3go/brickpi.svg?branch=master)](https://travis-ci.org/ev3go/brickpi) [![Coverage Status](https://coveralls.io/repos/ev3go/brickpi/badge.svg?branch=master&service=github)](https://coveralls.io/github/ev3go/brickpi?branch=master) [![GoDoc](https://godoc.org/github.com/ev3go/brickpi?status.svg)](https://godoc.org/github.com/ev3go/brickpi)

github.com/ev3go/brickpi depends on an ev3dev kernel v4.1.21-10-ev3dev-rpi or v4.1.21-10-ev3dev-rpi2 or better (See http://www.ev3dev.org/news/2016/04/11/Kernel-Release-Cycle-10/).

## Example code

```
package main

import (
	"log"
	"time"

	"github.com/ev3go/brickpi"
)

func main() {
	var bright byte
	var err error
	for i := 0; i < 10; i++ {
		err = brickpi.Blue1.SetBrightness(int(bright)).Err()
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second)

		bright = ^bright

		err = brickpi.Blue2.SetBrightness(int(bright)).Err()
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second)
	}
}
```