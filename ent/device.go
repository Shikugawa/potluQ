// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/Shikugawa/potluq/ent/device"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Device is the model entity for the Device schema.
type Device struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Role holds the value of the "role" field.
	Role device.Role `json:"role,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Device) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},
		&sql.NullString{},
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Device fields.
func (d *Device) assignValues(values ...interface{}) error {
	if m, n := len(values), len(device.Columns); m != n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	d.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field role", values[0])
	} else if value.Valid {
		d.Role = device.Role(value.String)
	}
	return nil
}

// QueryClub queries the club edge of the Device.
func (d *Device) QueryClub() *ClubQuery {
	return (&DeviceClient{d.config}).QueryClub(d)
}

// QueryUser queries the user edge of the Device.
func (d *Device) QueryUser() *UserQuery {
	return (&DeviceClient{d.config}).QueryUser(d)
}

// Update returns a builder for updating this Device.
// Note that, you need to call Device.Unwrap() before calling this method, if this Device
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Device) Update() *DeviceUpdateOne {
	return (&DeviceClient{d.config}).UpdateOne(d)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (d *Device) Unwrap() *Device {
	tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Device is not a transactional entity")
	}
	d.config.driver = tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Device) String() string {
	var builder strings.Builder
	builder.WriteString("Device(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteString(", role=")
	builder.WriteString(fmt.Sprintf("%v", d.Role))
	builder.WriteByte(')')
	return builder.String()
}

// Devices is a parsable slice of Device.
type Devices []*Device

func (d Devices) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}
