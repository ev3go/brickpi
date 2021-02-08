// Copyright ©2021 The ev3go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// motor demonstrates using a motor on a BrickPi.
//
// The program should be run from the command line after attaching a nxt or ev3 motor
// to the BrickPi on a output port (MA, MB, MC or MD).
package main

import (
	"log"

	"github.com/ev3go/ev3dev"
)

const (
	// motorDriver is the default BrickPi motorDriver. For details, see
	// http://docs.ev3dev.org/projects/lego-linux-drivers/en/ev3dev-stretch/brickpi.html#output-ports
	motorDriver = "lego-nxt-motor"

	// motorPortPrefix is the prefix for BrickPi output port addresses. The
	// prefix must have the specific port appended. Posts can be  A, B, C or D,
	// corresponding to the BrickPi output ports MA, MB, MC and MD.
	motorPortPrefix = "serial0-0:M"
)

func main() {
	m, err := ev3dev.TachoMotorFor(motorPortPrefix+"A", motorDriver)
	if err != nil {
		log.Fatalf("failed to find motor port: %v", err)
	}

	err = m.SetSpeedSetpoint(500).
		SetPositionSetpoint(180).
		SetStopAction("hold").
		Command("run-to-rel-pos").
		Err()
	if err != nil {
		log.Fatalf("error during motor operation: %v", err)
	}
}
