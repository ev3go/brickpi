![Gopherbrick](gopherbrick.png)
# brickpi provides BrickPi-specific functions for the Go ev3dev interface

[![Build Status](https://travis-ci.org/ev3go/brickpi.svg?branch=master)](https://travis-ci.org/ev3go/brickpi) [![Coverage Status](https://coveralls.io/repos/ev3go/brickpi/badge.svg?branch=master&service=github)](https://coveralls.io/github/ev3go/brickpi?branch=master) [![GoDoc](https://godoc.org/github.com/ev3go/brickpi?status.svg)](https://godoc.org/github.com/ev3go/brickpi)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fev3go%2Fbrickpi.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fev3go%2Fbrickpi?ref=badge_shield)

github.com/ev3go/ev3dev depends on ev3dev stretch.

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


## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fev3go%2Fbrickpi.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fev3go%2Fbrickpi?ref=badge_large)