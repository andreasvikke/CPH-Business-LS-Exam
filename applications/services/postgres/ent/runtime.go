// Code generated by entc, DO NOT EDIT.

package ent

import (
	"postgres_service/ent/checkin"
	"postgres_service/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	checkinFields := schema.CheckIn{}.Fields()
	_ = checkinFields
	// checkinDescAttendanceCode is the schema descriptor for attendanceCode field.
	checkinDescAttendanceCode := checkinFields[0].Descriptor()
	// checkin.AttendanceCodeValidator is a validator for the "attendanceCode" field. It is called by the builders before save.
	checkin.AttendanceCodeValidator = func() func(int64) error {
		validators := checkinDescAttendanceCode.Validators
		fns := [...]func(int64) error{
			validators[0].(func(int64) error),
			validators[1].(func(int64) error),
			validators[2].(func(int64) error),
		}
		return func(attendanceCode int64) error {
			for _, fn := range fns {
				if err := fn(attendanceCode); err != nil {
					return err
				}
			}
			return nil
		}
	}()
}
