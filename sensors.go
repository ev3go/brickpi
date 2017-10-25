// Copyright ©2016 The ev3go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package brickpi

import (
	"fmt"

	"github.com/ev3go/ev3dev"
)

const portDriver = "brickpi-in-port"

type device struct {
	mode   string
	setDev bool
	addr   int
}

// deviceLookup maps driver names to lego-port modes, devices and addresses.
//
// See documentation at http://www.ev3dev.org/docs/ports/brickpi-in-port/
// and http://www.ev3dev.org/docs/sensors/#supported-sensors.
var deviceLookup = map[string]device{
	"nxt-analog":       {mode: "nxt-analog", setDev: true},
	"di-dflex":         {mode: "nxt-analog", setDev: true},
	"ht-nxt-eopd":      {mode: "nxt-analog", setDev: true},
	"ht-nxt-force":     {mode: "nxt-analog", setDev: true},
	"ht-nxt-gyro":      {mode: "nxt-analog", setDev: true},
	"ht-nxt-mag":       {mode: "nxt-analog", setDev: true},
	"lego-nxt-touch":   {mode: "nxt-analog", setDev: true},
	"lego-nxt-light":   {mode: "nxt-analog", setDev: true},
	"lego-nxt-sound":   {mode: "nxt-analog", setDev: true},
	"ms-nxt-touch-mux": {mode: "nxt-analog", setDev: true},

	"lego-nxt-color": {mode: "nxt-color"},

	"nxt-i2c":            {mode: "nxt-i2c", setDev: true}, // User must call SensorForAddr.
	"pixy-lego":          {mode: "nxt-i2c", setDev: true, addr: 0x01},
	"ht-nxt-color":       {mode: "nxt-i2c", setDev: true, addr: 0x01},
	"ht-nxt-angle":       {mode: "nxt-i2c", setDev: true, addr: 0x01},
	"ht-nxt-accel":       {mode: "nxt-i2c", setDev: true, addr: 0x01},
	"ht-nxt-barometric":  {mode: "nxt-i2c", setDev: true, addr: 0x01},
	"ht-nxt-color-v2":    {mode: "nxt-i2c", setDev: true, addr: 0x01},
	"ht-nxt-ir-link":     {mode: "nxt-i2c", setDev: true, addr: 0x01},
	"ht-nxt-ir-receiver": {mode: "nxt-i2c", setDev: true, addr: 0x01},
	"ht-nxt-pir":         {mode: "nxt-i2c", setDev: true, addr: 0x01},
	"ht-nxt-compass":     {mode: "nxt-i2c", setDev: true, addr: 0x01},
	"ht-nxt-ir-seek-v2":  {mode: "nxt-i2c", setDev: true, addr: 0x08},
	"ht-super-pro":       {mode: "nxt-i2c", setDev: true, addr: 0x08},
	"lego-power-storage": {mode: "nxt-i2c", setDev: true, addr: 0x02},
	"lego-nxt-us":        {mode: "nxt-i2c", setDev: true, addr: 0x01},
	"ms-angle":           {mode: "nxt-i2c", setDev: true, addr: 0x18},
	"ms-light-array":     {mode: "nxt-i2c", setDev: true, addr: 0x0a},
	"ms-line-leader":     {mode: "nxt-i2c", setDev: true, addr: 0x01},
	"ms-nxtcam":          {mode: "nxt-i2c", setDev: true, addr: 0x01},
	"ms-pixy-adapter":    {mode: "nxt-i2c", setDev: true, addr: 0x01},

	"lego-ev3-touch": {mode: "ev3-analog"},

	// BUG(kortschak): Device drivers for NXT/I2C devices with special operations
	// will not work with the BrickPi.

	// The following have special operations and will not work on BrickPi.
	"ht-nxt-smux":     {mode: "nxt-i2c", setDev: true, addr: 0x08},
	"mi-xg1300l":      {mode: "nxt-i2c", setDev: true, addr: 0x01},
	"ms-absolute-imu": {mode: "nxt-i2c", setDev: true, addr: 0x11},
	"ms-ev3-smux":     {mode: "nxt-i2c", setDev: true}, // addr: 0x50, 0x51 and 0x52. User must call SensorForAddr.
	"ms-nxtmmx":       {mode: "nxt-i2c", setDev: true, addr: 0x03},
	"ms-8ch-servo":    {mode: "nxt-i2c", setDev: true, addr: 0x58},
	"ms-pps58-nx":     {mode: "nxt-i2c", setDev: true, addr: 0x0c},

	// BUG(DexterInd): EV3 UART sensors do not work correctly, apparently due to
	// timing problems in the Dexter Industries firmware for the BrickPi devices.
	//
	// See https://github.com/DexterInd/BrickPi/issues/24
	"lego-ev3-us":    {mode: "ev3-uart", setDev: true},
	"lego-ev3-color": {mode: "ev3-uart", setDev: true},
	"lego-ev3-gyro":  {mode: "ev3-uart", setDev: true},
	"lego-ev3-ir":    {mode: "ev3-uart", setDev: true},
}

// SensorFor returns an ev3dev.Sensor for the given BrickPi port name
// and driver. SensorFor uses the default address for the device.
// When the sensor is no longer being used, Unregister should be called
// for the port.
//
// BrickPi and BrickPi+ sensor port names are:
//  - ttyAMA0:S1
//  - ttyAMA0:S2
//  - ttyAMA0:S3
//  - ttyAMA0:S4
//
// BrickPi3 sensor port names are:
//  - spi0.1:S1
//  - spi0.1:S2
//  - spi0.1:S3
//  - spi0.1:S4
func SensorFor(port, driver string) (*ev3dev.Sensor, error) {
	dev, ok := deviceLookup[driver]
	if !ok {
		return nil, fmt.Errorf("brickpi: driver not supported: %q", driver)
	}
	return sensorFor(port, dev.mode, driver, dev.setDev, dev.addr)
}

// SensorForAddr returns an ev3dev.Sensor for the given BrickPi port name,
// driver and address.
// When the sensor is no longer being used, Unregister should be called
// for the port.
func SensorForAddr(port, driver string, addr int) (*ev3dev.Sensor, error) {
	dev, ok := deviceLookup[driver]
	if !ok {
		return nil, fmt.Errorf("brickpi: driver not supported: %q", driver)
	}
	return sensorFor(port, dev.mode, driver, dev.setDev, addr)
}

func sensorFor(port, mode, driver string, setDev bool, addr int) (*ev3dev.Sensor, error) {
	p, err := ev3dev.LegoPortFor(port, portDriver)
	if err != nil {
		return nil, err
	}
	err = p.SetMode(mode).Err()
	if err != nil {
		return nil, err
	}
	if setDev {
		d := driver
		if addr > 0x0 {
			d = fmt.Sprintf("%s 0x%02x", driver, addr)
		}
		err = p.SetDevice(d).Err()
		if err != nil {
			p.SetMode("none")
			return nil, err
		}
	}

	s, err := ev3dev.SensorFor(fmt.Sprintf("%s:i2c%d", port, addr), driver)
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
