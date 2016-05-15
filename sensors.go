// Copyright ©2016 The ev3go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package brickpi

import (
	"fmt"

	"github.com/ev3go/ev3dev"
)

const portDriver = "brickpi-in-port"

type mode struct {
	mode   string
	setDev bool
}

// modeLookup maps driver names to lego-port modes.
//
// See documentation at http://www.ev3dev.org/docs/ports/brickpi-in-port/
// and http://www.ev3dev.org/docs/sensors/#supported-sensors.
var modeLookup = map[string]mode{
	"nxt-analog":       mode{"nxt-analog", true},
	"di-dflex":         mode{"nxt-analog", true},
	"ht-nxt-eopd":      mode{"nxt-analog", true},
	"ht-nxt-force":     mode{"nxt-analog", true},
	"ht-nxt-gyro":      mode{"nxt-analog", true},
	"ht-nxt-mag":       mode{"nxt-analog", true},
	"lego-nxt-touch":   mode{"nxt-analog", true},
	"lego-nxt-light":   mode{"nxt-analog", true},
	"lego-nxt-sound":   mode{"nxt-analog", true},
	"ms-nxt-touch-mux": mode{"nxt-analog", true},

	"lego-nxt-color": mode{"nxt-color", false},

	// BUG(kortschak): nxt-i2c devices do not appear to work.
	// The documentation at http://www.ev3dev.org/docs/ports/brickpi-in-port/
	// says that setting the mode to nxt-i2c should be associated with
	// a write to set_device. This second write fails with EINVAL.
	//
	// See https://github.com/ev3dev/ev3dev/issues/640
	"nxt-i2c":            mode{"nxt-i2c", true},
	"pixy-lego":          mode{"nxt-i2c", true},
	"ht-nxt-color":       mode{"nxt-i2c", true},
	"ht-nxt-angle":       mode{"nxt-i2c", true},
	"ht-nxt-accel":       mode{"nxt-i2c", true},
	"ht-nxt-barometric":  mode{"nxt-i2c", true},
	"ht-nxt-color-v2":    mode{"nxt-i2c", true},
	"ht-nxt-ir-link":     mode{"nxt-i2c", true},
	"ht-nxt-ir-receiver": mode{"nxt-i2c", true},
	"ht-nxt-pir":         mode{"nxt-i2c", true},
	"ht-nxt-compass":     mode{"nxt-i2c", true},
	"ht-nxt-ir-seek-v2":  mode{"nxt-i2c", true},
	"ht-nxt-smux":        mode{"nxt-i2c", true},
	"ht-super-pro":       mode{"nxt-i2c", true},
	"lego-power-storage": mode{"nxt-i2c", true},
	"lego-nxt-us":        mode{"nxt-i2c", true},
	"mi-xg1300l":         mode{"nxt-i2c", true},
	"ms-absolute-imu":    mode{"nxt-i2c", true},
	"ms-angle":           mode{"nxt-i2c", true},
	"ms-ev3-smux":        mode{"nxt-i2c", true},
	"ms-light-array":     mode{"nxt-i2c", true},
	"ms-line-leader":     mode{"nxt-i2c", true},
	"ms-nxtcam":          mode{"nxt-i2c", true},
	"ms-nxtmmx":          mode{"nxt-i2c", true},
	"ms-8ch-servo":       mode{"nxt-i2c", true},
	"ms-pps58-nx":        mode{"nxt-i2c", true},
	"ms-pixy-adapter":    mode{"nxt-i2c", true},

	"lego-ev3-touch": mode{"ev3-analog", false},

	// BUG(DexterInd): EV3 UART sensors do not work correctly, apparently due to
	// timing problems in the Dexter Industries firmware for the BrickPi devices.
	//
	// See https://github.com/DexterInd/BrickPi/issues/24
	"lego-ev3-us":    mode{"ev3-uart", true},
	"lego-ev3-color": mode{"ev3-uart", true},
	"lego-ev3-gyro":  mode{"ev3-uart", true},
	"lego-ev3-ir":    mode{"ev3-uart", true},
}

// SensorFor returns an ev3dev.Sensor for the given BrickPi port name
// and driver.
// When the sensor is no longer being used, Unregister should be called
// for the port.
//
// BrickPi sensor port names are:
//  - ttyAMA0:S1
//  - ttyAMA0:S2
//  - ttyAMA0:S3
//  - ttyAMA0:S4
func SensorFor(port, driver string) (*ev3dev.Sensor, error) {
	m, ok := modeLookup[driver]
	if !ok {
		return nil, fmt.Errorf("brickpi: driver not supported: %q", driver)
	}

	p, err := ev3dev.LegoPortFor(port, portDriver)
	if err != nil {
		return nil, err
	}
	err = p.SetMode(m.mode).Err()
	if err != nil {
		return nil, err
	}
	if m.setDev {
		err = p.SetDevice(driver).Err()
		if err != nil {
			p.SetMode("none")
			return nil, err
		}
	}

	s, err := ev3dev.SensorFor(port, driver)
	if err != nil {
		p.SetMode("none")
		return nil, err
	}
	return s, err
}

// Unregister sets the lego-port mode to "none".
func Unregister(port string) error {
	p, err := ev3dev.LegoPortFor(port, portDriver)
	if err != nil {
		return err
	}
	return p.SetMode("none").Err()
}
