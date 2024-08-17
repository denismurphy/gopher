package types

import (
	"fmt"
	"time"
)

// Define custom types and structs here
type (
	Int    = int
	Double = float64
	Byte   = uint8

	DateType struct {
		Time     time.Time
		Timezone string
	}

	MagicType struct {
		Name  string
		Power int
	}

	Simple struct {
		Name string
		ID   int
	}

	MyError struct {
		Message string
	}

	Truck struct {
		Time     time.Time
		Timezone string
	}
)

// Error implements the error interface for MyError
func (e *MyError) Error() string {
	return e.Message
}

// TurnOn sets initial values for a Truck
func (t *Truck) TurnOn() {
	t.Time = time.Now()
	t.Timezone = "UTC"
}

// String implements the Stringer interface for Truck
func (t *Truck) String() string {
	return fmt.Sprintf("Time: %v, Timezone: %s", t.Time, t.Timezone)
}

// Turner interface for types that can be turned on
type Turner interface {
	TurnOn()
}

// Cars represents different car manufacturers
type Cars int

const (
	Volkswagen Cars = iota
	Ford
	Toyota
	Kia
	Hyundai
)
