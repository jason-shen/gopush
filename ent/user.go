// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/jason-shen/gopush/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// FirstName holds the value of the "first_name" field.
	FirstName string `json:"first_name,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// Apikey holds the value of the "apikey" field.
	Apikey uuid.UUID `json:"apikey,omitempty"`
	// Jwttoken holds the value of the "jwttoken" field.
	Jwttoken *string `json:"jwttoken,omitempty"`
	// ActivateCode holds the value of the "activate_code" field.
	ActivateCode int8 `json:"activate_code,omitempty"`
	// Activated holds the value of the "activated" field.
	Activated bool `json:"activated,omitempty"`
	// Locked holds the value of the "locked" field.
	Locked bool `json:"locked,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues() []interface{} {
	return []interface{}{
		&uuid.UUID{},      // id
		&sql.NullString{}, // first_name
		&sql.NullString{}, // last_name
		&sql.NullString{}, // email
		&sql.NullString{}, // password
		&uuid.UUID{},      // apikey
		&sql.NullString{}, // jwttoken
		&sql.NullInt64{},  // activate_code
		&sql.NullBool{},   // activated
		&sql.NullBool{},   // locked
		&sql.NullTime{},   // updated_at
		&sql.NullTime{},   // created_at
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(values ...interface{}) error {
	if m, n := len(values), len(user.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value != nil {
		u.ID = *value
	}
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field first_name", values[0])
	} else if value.Valid {
		u.FirstName = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field last_name", values[1])
	} else if value.Valid {
		u.LastName = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field email", values[2])
	} else if value.Valid {
		u.Email = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field password", values[3])
	} else if value.Valid {
		u.Password = value.String
	}
	if value, ok := values[4].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field apikey", values[4])
	} else if value != nil {
		u.Apikey = *value
	}
	if value, ok := values[5].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field jwttoken", values[5])
	} else if value.Valid {
		u.Jwttoken = new(string)
		*u.Jwttoken = value.String
	}
	if value, ok := values[6].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field activate_code", values[6])
	} else if value.Valid {
		u.ActivateCode = int8(value.Int64)
	}
	if value, ok := values[7].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field activated", values[7])
	} else if value.Valid {
		u.Activated = value.Bool
	}
	if value, ok := values[8].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field locked", values[8])
	} else if value.Valid {
		u.Locked = value.Bool
	}
	if value, ok := values[9].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updated_at", values[9])
	} else if value.Valid {
		u.UpdatedAt = value.Time
	}
	if value, ok := values[10].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[10])
	} else if value.Valid {
		u.CreatedAt = value.Time
	}
	return nil
}

// Update returns a builder for updating this User.
// Note that, you need to call User.Unwrap() before calling this method, if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", first_name=")
	builder.WriteString(u.FirstName)
	builder.WriteString(", last_name=")
	builder.WriteString(u.LastName)
	builder.WriteString(", email=")
	builder.WriteString(u.Email)
	builder.WriteString(", password=")
	builder.WriteString(u.Password)
	builder.WriteString(", apikey=")
	builder.WriteString(fmt.Sprintf("%v", u.Apikey))
	if v := u.Jwttoken; v != nil {
		builder.WriteString(", jwttoken=")
		builder.WriteString(*v)
	}
	builder.WriteString(", activate_code=")
	builder.WriteString(fmt.Sprintf("%v", u.ActivateCode))
	builder.WriteString(", activated=")
	builder.WriteString(fmt.Sprintf("%v", u.Activated))
	builder.WriteString(", locked=")
	builder.WriteString(fmt.Sprintf("%v", u.Locked))
	builder.WriteString(", updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
