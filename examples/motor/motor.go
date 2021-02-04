// Copyright ©2016 The ev3go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// motor demonstrates using a motor on a BrickPi.
//
// The program should be run from the command line after attaching a nxt or ev3 motor
// to the BrickPi on a output port (MA, MB, MC or MD).
package main

import (
	"fmt"
	"log"

	"github.com/ev3go/ev3dev"
)

func main() {
	// Motor driver .This value works for both nxt and ev3 large motor
	motorDriver := "lego-nxt-motor"

	// Prefix string for all motor ports when using brickpi
	motorPortStringPrefix := "serial0-0:M"

	// Motor Port can be A, B, C or D, corresponding to the brickpi output ports MA,MB,MC and MD
	motorPort := "A"

	// motorPortString is of the format serial0-0:MA, serial0-0:MB etc
	motorPortString := fmt.Sprintf("%s%s", motorPortStringPrefix, motorPort)

	m, err := ev3dev.TachoMotorFor(motorPortString, motorDriver)
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
