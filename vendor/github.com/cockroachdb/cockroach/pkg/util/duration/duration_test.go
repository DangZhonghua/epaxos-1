// Copyright 2016 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.
//
// Author: Daniel Harrison (daniel.harrison@gmail.com)

package duration

import (
	"math"
	"testing"
	"time"

	_ "github.com/cockroachdb/cockroach/pkg/util/log"
)

type durationTest struct {
	cmpToPrev int
	duration  Duration
	err       bool
}

// positiveDurationTests is used both to check that each duration roudtrips
// through Encode/Decode and that they sort in the expected way. They are not
// required to be listed in ascending order, but for ease of maintenance, it is
// expected that they stay ascending.
//
// The negative tests are generated by prepending everything but the 0 case and
// flipping the sign of cmpToPrev, since they will be getting bigger in abolute
// value and more negative.
//
// TODO(dan): Write more tests with a mixture of positive and negative
// components.
var positiveDurationTests = []durationTest{
	{1, Duration{Months: 0, Days: 0, Nanos: 0}, false},
	{1, Duration{Months: 0, Days: 0, Nanos: 1}, false},
	{1, Duration{Months: 0, Days: 0, Nanos: nanosInDay - 1}, false},
	{1, Duration{Months: 0, Days: 1, Nanos: 0}, false},
	{0, Duration{Months: 0, Days: 0, Nanos: nanosInDay}, false},
	{1, Duration{Months: 0, Days: 0, Nanos: nanosInDay + 1}, false},
	{1, Duration{Months: 0, Days: daysInMonth - 1, Nanos: 0}, false},
	{1, Duration{Months: 0, Days: 0, Nanos: nanosInMonth - 1}, false},
	{1, Duration{Months: 1, Days: 0, Nanos: 0}, false},
	{0, Duration{Months: 0, Days: daysInMonth, Nanos: 0}, false},
	{0, Duration{Months: 0, Days: 0, Nanos: nanosInMonth}, false},
	{1, Duration{Months: 0, Days: 0, Nanos: nanosInMonth + 1}, false},
	{1, Duration{Months: 0, Days: daysInMonth + 1, Nanos: 0}, false},
	{1, Duration{Months: 1, Days: 1, Nanos: 1}, false},
	{1, Duration{Months: 1, Days: 10, Nanos: 0}, false},
	{0, Duration{Months: 0, Days: 40, Nanos: 0}, false},
	{1, Duration{Months: 2, Days: 0, Nanos: 0}, false},
	{1, Duration{Months: math.MaxInt64 - 1, Days: daysInMonth - 1, Nanos: nanosInDay * 2}, true},
	{1, Duration{Months: math.MaxInt64 - 1, Days: daysInMonth * 2, Nanos: nanosInDay * 2}, true},
	{1, Duration{Months: math.MaxInt64, Days: math.MaxInt64, Nanos: nanosInMonth + nanosInDay}, true},
	{1, Duration{Months: math.MaxInt64, Days: math.MaxInt64, Nanos: math.MaxInt64}, true},
}

func fullDurationTests() []durationTest {
	var ret []durationTest
	for _, test := range positiveDurationTests {
		d := test.duration
		negDuration := Duration{Months: -d.Months, Days: -d.Days, Nanos: -d.Nanos}
		ret = append(ret, durationTest{cmpToPrev: -test.cmpToPrev, duration: negDuration, err: test.err})
	}
	ret = append(ret, positiveDurationTests...)
	return ret
}

func TestEncodeDecode(t *testing.T) {
	for i, test := range fullDurationTests() {
		sortNanos, months, days, err := test.duration.Encode()
		if test.err && err == nil {
			t.Errorf("%d expected error but didn't get one", i)
		} else if !test.err && err != nil {
			t.Errorf("%d expected no error but got one: %s", i, err)
		}
		if err != nil {
			continue
		}
		sortNanosBig, _, _ := test.duration.EncodeBigInt()
		if sortNanos != sortNanosBig.Int64() {
			t.Errorf("%d Encode and EncodeBig didn't match [%d] vs [%s]", i, sortNanos, sortNanosBig)
		}
		d, err := Decode(sortNanos, months, days)
		if err != nil {
			t.Fatal(err)
		}
		if test.duration != d {
			t.Errorf("%d encode/decode mismatch [%v] vs [%v]", i, test, d)
		}
	}
}

func TestCompare(t *testing.T) {
	prev := Duration{Nanos: 1} // It's expected that we start with something greater than 0.
	for i, test := range fullDurationTests() {
		cmp := test.duration.Compare(prev)
		if cmp != test.cmpToPrev {
			t.Errorf("%d did not compare correctly got %d expected %d [%s] vs [%s]", i, cmp, test.cmpToPrev, prev, test.duration)
		}
		prev = test.duration
	}
}

func TestNormalize(t *testing.T) {
	for i, test := range fullDurationTests() {
		nanos, _, _ := test.duration.EncodeBigInt()
		normalized := test.duration.normalize()
		normalizedNanos, _, _ := normalized.EncodeBigInt()
		if nanos.Cmp(normalizedNanos) != 0 {
			t.Errorf("%d effective nanos were changed [%s] [%s]", i, test.duration, normalized)
		}
		if normalized.Days > daysInMonth && normalized.Months != math.MaxInt64 ||
			normalized.Days < -daysInMonth && normalized.Months != math.MinInt64 {
			t.Errorf("%d days were not normalized [%s]", i, normalized)
		}
		if normalized.Nanos > nanosInDay && normalized.Days != math.MaxInt64 ||
			normalized.Nanos < -nanosInDay && normalized.Days != math.MinInt64 {
			t.Errorf("%d nanos were not normalized [%s]", i, normalized)
		}
	}
}

func TestDiffMicros(t *testing.T) {
	tests := []struct {
		t1, t2  time.Time
		expDiff int64
	}{
		{
			t1:      time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			t2:      time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			expDiff: 0,
		},
		{
			t1:      time.Date(1, 8, 15, 12, 30, 45, 0, time.UTC),
			t2:      time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			expDiff: -63062710155000000,
		},
		{
			t1:      time.Date(1994, 8, 15, 12, 30, 45, 0, time.UTC),
			t2:      time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			expDiff: -169730955000000,
		},
		{
			t1:      time.Date(2012, 8, 15, 12, 30, 45, 0, time.UTC),
			t2:      time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			expDiff: 398349045000000,
		},
		{
			t1:      time.Date(8012, 8, 15, 12, 30, 45, 0, time.UTC),
			t2:      time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			expDiff: 189740061045000000,
		},
		// Test if the nanoseconds round correctly.
		{
			t1:      time.Date(2000, 1, 1, 0, 0, 0, 499, time.UTC),
			t2:      time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			expDiff: 0,
		},
		{
			t1:      time.Date(1999, 12, 31, 23, 59, 59, 999999501, time.UTC),
			t2:      time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			expDiff: 0,
		},
		{
			t1:      time.Date(2000, 1, 1, 0, 0, 0, 500, time.UTC),
			t2:      time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			expDiff: 1,
		},
		{
			t1:      time.Date(1999, 12, 31, 23, 59, 59, 999999500, time.UTC),
			t2:      time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			expDiff: -1,
		},
	}

	for i, test := range tests {
		if res := DiffMicros(test.t1, test.t2); res != test.expDiff {
			t.Errorf("%d: expected DiffMicros(%v, %v) = %d, found %d",
				i, test.t1, test.t2, test.expDiff, res)
		} else {
			// Swap order and make sure the results are mirrored.
			exp := -test.expDiff
			if res := DiffMicros(test.t2, test.t1); res != exp {
				t.Errorf("%d: expected DiffMicros(%v, %v) = %d, found %d",
					i, test.t2, test.t1, exp, res)
			}
		}
	}
}

func TestAddMicros(t *testing.T) {
	tests := []struct {
		t   time.Time
		d   int64
		exp time.Time
	}{
		{
			t:   time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			d:   0,
			exp: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			t:   time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			d:   123456789,
			exp: time.Date(2000, 1, 1, 0, 2, 3, 456789000, time.UTC),
		},
		{
			t:   time.Date(2000, 1, 1, 0, 0, 0, 999, time.UTC),
			d:   123456789,
			exp: time.Date(2000, 1, 1, 0, 2, 3, 456789999, time.UTC),
		},
		{
			t:   time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			d:   12345678987654321,
			exp: time.Date(2391, 03, 21, 19, 16, 27, 654321000, time.UTC),
		},
		{
			t:   time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			d:   math.MaxInt64 / 10,
			exp: time.Date(31227, 9, 14, 2, 48, 05, 477580000, time.UTC),
		},
		{
			t:   time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			d:   -123456789,
			exp: time.Date(1999, 12, 31, 23, 57, 56, 543211000, time.UTC),
		},
		{
			t:   time.Date(2000, 1, 1, 0, 0, 0, 999, time.UTC),
			d:   -123456789,
			exp: time.Date(1999, 12, 31, 23, 57, 56, 543211999, time.UTC),
		},
		{
			t:   time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			d:   -12345678987654321,
			exp: time.Date(1608, 10, 12, 04, 43, 32, 345679000, time.UTC),
		},
		{
			t:   time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			d:   -math.MaxInt64 / 10,
			exp: time.Date(-27228, 04, 18, 21, 11, 54, 522420000, time.UTC),
		},
	}

	for i, test := range tests {
		if res := AddMicros(test.t, test.d); !test.exp.Equal(res) {
			t.Errorf("%d: expected AddMicros(%v, %d) = %v, found %v",
				i, test.t, test.d, test.exp, res)
		}
	}
}
