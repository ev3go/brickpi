// Copyright ©2016 The ev3go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package brickpi

import (
	"fmt"

	"github.com/ev3go/ev3dev"
)

// LED handles for BrickPi devices.
var (
	// Blue1 and Blue2 are the blue LEDs on
	// the BrickPi board.
	Blue1 = &ev3dev.LED{Name: brickPiLED(1)}
	Blue2 = &ev3dev.LED{Name: brickPiLED(2)}

	// Green and Red are the green and red LEDs
	// on the Raspberry Pi board.
	//
	// These LED devices are normally writable
	// only by root, so permissions will need
	// to be correctly set prior to use.
	Green = &ev3dev.LED{Name: raspberryPiLED(0)}
	Red   = &ev3dev.LED{Name: raspberryPiLED(1)}
)

// brickPiLED is a fmt.Stringer LED name for the BrickPi LEDs.
type brickPiLED int

func (l brickPiLED) String() string { return fmt.Sprintf("led%d:blue:brick-status", l) }

// raspberryPiLED is a fmt.Stringer LED name for the Raspberry Pi LEDs.
type raspberryPiLED int

func (l raspberryPiLED) String() string { return fmt.Sprintf("led%d", l) }
