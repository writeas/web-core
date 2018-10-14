// Package query assists in building SQL queries.
package query

import (
	"database/sql"
)

type Update struct {
	Updates, Conditions string
	Params              []interface{}

	sep string
}

func NewUpdate() *Update {
	return &Update{}
}

func (u *Update) Set(v, property string) *Update {
	if v != "" {
		u.Updates += u.sep + property + " = ?"
		u.sep = ", "
		u.Params = append(u.Params, v)
	}
	return u
}

func (u *Update) SetBytes(v []byte, property string) *Update {
	if len(v) > 0 {
		u.Updates += u.sep + property + " = ?"
		u.sep = ", "
		u.Params = append(u.Params, v)
	}
	return u
}

func (u *Update) SetStringPtr(v *string, property string) *Update {
	if v != nil {
		u.Updates += u.sep + property + " = ?"
		u.sep = ", "
		u.Params = append(u.Params, v)
	}
	return u
}

func (u *Update) SetIntPtr(v *int, property string) *Update {
	if v != nil {
		u.Updates += u.sep + property + " = ?"
		u.sep = ", "
		u.Params = append(u.Params, v)
	}
	return u
}

func (u *Update) SetBoolPtr(v *bool, property string) *Update {
	if v != nil {
		u.Updates += u.sep + property + " = ?"
		u.sep = ", "
		u.Params = append(u.Params, v)
	}
	return u
}

func (u *Update) SetNullBool(v *sql.NullBool, property string) *Update {
	if v != nil {
		u.Updates += u.sep + property + " = ?"
		u.sep = ", "
		u.Params = append(u.Params, v)
	}
	return u
}

func (u *Update) SetNullString(v *sql.NullString, property string) *Update {
	if v != nil {
		u.Updates += u.sep + property + " = ?"
		u.sep = ", "
		u.Params = append(u.Params, v)
	}
	return u
}

func (u *Update) Append(v interface{}) {
	u.Params = append(u.Params, v)
}

func (u *Update) Where(condition string, params ...interface{}) *Update {
	u.Conditions = condition
	for _, p := range params {
		u.Append(p)
	}
	return u
}
