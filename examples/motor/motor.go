package main

import (
	"fmt"
	"log"

	"github.com/ev3go/ev3dev"
)

func main() {

	// degrees to rotate motor by
	rotateDegree := 180

	// Motor Port can be A, B, C or D
	motorPort := "D"

	motorStopAction := "hold" // can be brake, hold etc

	// Motor driver .This value works for both nxt and ev3 large motor
	motorDriver := "lego-nxt-motor"

	// motor Port String format
	motorPortStringPrefix := "serial0-0:M"

	motorPortString := fmt.Sprintf("%s%s", motorPortStringPrefix, motorPort)
	fmt.Printf("motorPortString: %s\n", motorPortString)

	outPort, err := ev3dev.TachoMotorFor(motorPortString, motorDriver)
	if err != nil {
		log.Fatalf("failed to find  motor on port %s: %v", motorPort, err)
	}
	err = outPort.SetStopAction("hold").Err()
	if err != nil {
		log.Fatalf("failed to set hold stop for motor on port %s: %v", motorPort, err)
	}

	// rotate motor by rotateDegree angle
	outPort.SetSpeedSetpoint(500).SetPositionSetpoint(rotateDegree).SetStopAction(motorStopAction).Command("run-to-rel-pos")
}
