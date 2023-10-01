package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
)

type Date struct {
	Time    time.Time
	NotNull bool
}

func (t Date) GetTime() (time.Time, error) {
	if !t.NotNull {
		return t.Time, errors.New("time is null")
	}
	return t.Time, nil
}

func NewDate(time time.Time) *Date {
	return &Date{Time: time, NotNull: true}
}

func (t Date) MarshalJSON() ([]byte, error) {
	if !t.NotNull {
		return []byte(`""`), nil
	}
	date := t.Time.Format(time.DateOnly)
	date = fmt.Sprintf(`"%s"`, date)
	return []byte(date), nil
}

func (t *Date) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")

	if s == "" {
		t.NotNull = false
		return
	}

	date, err := time.Parse(time.DateOnly, s)
	if err != nil {
		return err
	}
	t.Time = date
	return
}

func RegisterDate() *bsoncodec.Registry {
	dateType := reflect.TypeOf(Date{})

	dateDecoderFunc := func(
		decodeContext bsoncodec.DecodeContext,
		valueReader bsonrw.ValueReader,
		value reflect.Value,
	) error {
		// All decoder implementations should check that val is valid, settable,
		// and is of the correct kind before proceeding.
		if !value.IsValid() || !value.CanSet() || value.Type() != dateType {
			return bsoncodec.ValueDecoderError{
				Name:     "dateDecoder",
				Types:    []reflect.Type{dateType},
				Received: value,
			}
		}

		switch valueReader.Type() {
		case bson.TypeNull:
			if err := valueReader.ReadNull(); err != nil {
				return err
			}
			value.FieldByName("NotNull").SetBool(false)
		case bson.TypeDateTime:
			readResult, err := valueReader.ReadDateTime()
			if err != nil {
				return err
			}
			readTime := time.UnixMilli(readResult)
			value.FieldByName("Time").Set(reflect.ValueOf(readTime)) //Date has the only field
			if readResult == 0 {
				value.FieldByName("NotNull").SetBool(false)
			} else {
				value.FieldByName("NotNull").SetBool(true)
			}
		default:
			return fmt.Errorf(
				"received invalid BSON type to decode into Date: %s",
				valueReader.Type())
		}

		return nil
	}

	dateEncoderFunc := func(
		encodeContext bsoncodec.EncodeContext,
		valueWriter bsonrw.ValueWriter,
		value reflect.Value,
	) error {
		if !value.IsValid() || value.Type() != dateType {
			return bsoncodec.ValueEncoderError{
				Name:     "received invalid type to encode: %s",
				Types:    []reflect.Type{dateType},
				Received: value}
		}

		if !value.FieldByName("NotNull").Interface().(bool) {
			return valueWriter.WriteNull()
		}

		i, _ := value.FieldByName("Time").Interface().(time.Time)
		return valueWriter.WriteDateTime(i.UnixMilli()) //Date has the only field
	}

	reg := bson.NewRegistry()
	reg.RegisterTypeDecoder(
		dateType,
		bsoncodec.ValueDecoderFunc(dateDecoderFunc))
	reg.RegisterTypeEncoder(
		dateType,
		bsoncodec.ValueEncoderFunc(dateEncoderFunc))
	return reg
}
