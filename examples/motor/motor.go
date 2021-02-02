package main

import (
	"log"

	"github.com/ev3go/ev3dev"
)

func main() {

	m, err := ev3dev.TachoMotorFor("serial0-0:MA", "lego-nxt-motor")
	if err != nil {
		log.Fatalf("failed to find lego-nxt-motor motor port: %v", err)
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
