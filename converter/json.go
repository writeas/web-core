package converter

import (
	"database/sql"
	"encoding/json"
	"reflect"
)

type NullJSONBool struct {
	sql.NullBool
}

func JSONNullBool(value string) reflect.Value {
	v := NullJSONBool{}

	if value == "on" || value == "off" {
		return reflect.ValueOf(NullJSONBool{sql.NullBool{Bool: value == "on", Valid: true}})
	}
	if err := v.Scan(value); err != nil {
		return reflect.Value{}
	}
	return reflect.ValueOf(v)
}

func (v NullJSONBool) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Bool)
	} else {
		return json.Marshal(nil)
	}
}

func (v *NullJSONBool) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *bool
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Bool = *x
	} else {
		v.Valid = false
	}
	return nil
}

type NullJSONString struct {
	sql.NullString
}

func JSONNullString(value string) reflect.Value {
	v := NullJSONString{}
	if err := v.Scan(value); err != nil {
		return reflect.Value{}
	}

	return reflect.ValueOf(v)
}

func (v NullJSONString) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.String)
	} else {
		return json.Marshal(nil)
	}
}

func (v *NullJSONString) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *string
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.String = *x
	} else {
		v.Valid = false
	}
	return nil
}
