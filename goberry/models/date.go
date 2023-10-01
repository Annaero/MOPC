package models

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
)

type Date struct {
	time.Time
}

func (t Date) MarshalJSON() ([]byte, error) {
	date := t.Time.Format(time.DateOnly)
	date = fmt.Sprintf(`"%s"`, date)
	return []byte(date), nil
}

func (t *Date) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")

	date, err := time.Parse(time.DateOnly, s)
	if err != nil {
		return err
	}
	t.Time = date
	return
}

func RegisterDate() *bsoncodec.Registry {
	dateType := reflect.TypeOf(Date{Time: time.Now()})

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

		if valueReader.Type() != bson.TypeDateTime {
			return fmt.Errorf(
				"received invalid BSON type to decode into Date: %s",
				valueReader.Type())
		}
		readResult, err := valueReader.ReadDateTime()
		if err != nil {
			return err
		}

		readTime := time.UnixMilli(readResult)
		value.Field(0).Set(reflect.ValueOf(readTime)) //Date has the only field
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

		i, _ := value.Field(0).Interface().(time.Time)
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
