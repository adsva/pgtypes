/*
Package pgtypes implements some PostgreSQL data types which
implement the sql.Scanner and sql.Valuer interfaces
*/

package pgtypes

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var (
	pointRegexp     *regexp.Regexp
	dateRangeRegexp *regexp.Regexp
	dateRangeFormat = "2006-01-02"
)

func init() {
	pointRegexp = regexp.MustCompile(`(\d+\.\d+)`)
	dateRangeRegexp = regexp.MustCompile(`(\d+-\d+-\d+)`)
}

type Point struct {
	X float64
	Y float64
}

func (point *Point) Scan(src interface{}) (err error) {
	var source string
	switch src.(type) {
	case string:
		source = src.(string)
	case []byte:
		source = string(src.([]byte))
	default:
		return errors.New("Incompatible type for Point")
	}

	match := pointRegexp.FindAllString(source, 2)

	x, err := strconv.ParseFloat(match[0], 64)
	if err != nil {
		return err
	}
	y, err := strconv.ParseFloat(match[1], 64)
	if err != nil {
		return err
	}
	point.X = x
	point.Y = y
	return nil
}

func (point *Point) Value() (driver.Value, error) {
	return fmt.Sprintf("(%f,%f)", point.X, point.Y), nil
}

type DateRange struct {
	From time.Time
	To   time.Time
}

func (dr *DateRange) Scan(src interface{}) (err error) {
	var source string
	switch src.(type) {
	case string:
		source = src.(string)
	case []byte:
		source = string(src.([]byte))
	default:
		return errors.New("Incompatible type for Point")
	}

	match := dateRangeRegexp.FindAllString(source, 2)

	from, err := time.Parse(dateRangeFormat, match[0])
	if err != nil {
		return err
	}
	to, err := time.Parse(dateRangeFormat, match[1])
	if err != nil {
		return err
	}
	dr.From = from
	dr.To = to
	return nil
}

func (dr *DateRange) Value() (driver.Value, error) {
	if dr != nil {
		return fmt.Sprint(dr), nil
	} else {
		return nil, nil
	}
}

func (dr DateRange) String() string {
	return fmt.Sprintf(
		"[%s,%s)",
		dr.From.Format(dateRangeFormat),
		dr.To.Format(dateRangeFormat),
	)
}
