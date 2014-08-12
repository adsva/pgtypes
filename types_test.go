package pgtypes

import (
	"testing"
	"time"
)

func TestPoint(t *testing.T) {
	p1 := &Point{3.141592, 3.141592}

	val, err := p1.Value()
	if err != nil {
		t.Error(err)
	}

	p2 := &Point{}
	err = p2.Scan(val)
	if err != nil {
		t.Error(err)
	}

	if *p1 != *p2 {
		t.Errorf("Value output: %s does not match input: %s", p1, p2)
	}

}

func TestDateRange(t *testing.T) {
	from, err := time.Parse(dateRangeFormat, "1970-01-01")
	to, err := time.Parse(dateRangeFormat, "1970-01-02")
	r1 := &DateRange{from, to}

	val, err := r1.Value()
	if err != nil {
		t.Error(err)
	}

	r2 := &DateRange{}
	err = r2.Scan(val)
	if err != nil {
		t.Error(err)
	}

	if !(r1.From.Equal(r2.From) && r2.To.Equal(r2.To)) {
		t.Errorf(
			"Value output: %s--%s does not match input: %s--%s",
			r1.From.Format(dateRangeFormat),
			r1.To.Format(dateRangeFormat),
			r2.From.Format(dateRangeFormat),
			r2.To.Format(dateRangeFormat),
		)
	}

}
