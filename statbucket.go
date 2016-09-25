package main

import (
	"sort"
)

/*
StatBucket is a useful struct for determining various statistics regarding
a set of float elements.
*/
type StatBucket struct {
	// public
	Elements []float64

	// private
	isSorted bool
	hasMin   bool
	minimum  float64
	hasMax   bool
	maximum  float64
	hasMean  bool
	mean     float64
	hasp25   bool
	p25      float64
	hasp50   bool
	p50      float64
	hasp75   bool
	p75      float64
	hasp90   bool
	p90      float64
	hasp95   bool
	p95      float64
}

func (s *StatBucket) Count() int {
	return len(s.Elements)
}

func (s *StatBucket) Sum() (total float64) {
	for _, e := range s.Elements {
		total += e
	}
	return total
}

func (s *StatBucket) CheckSize() {
	if s.Count() == 0 {
		panic("StatBucket contains 0 items")
	}
}

func (s *StatBucket) Min() float64 {
	if s.hasMin {
		return s.minimum
	}
	s.CheckSize()
	minimum := s.Elements[0]
	for _, e := range s.Elements {
		if e < minimum {
			minimum = e
		}
	}
	s.minimum = minimum
	s.hasMin = true
	return minimum
}

func (s *StatBucket) Max() float64 {
	if s.hasMax {
		return s.maximum
	}
	s.CheckSize()
	maximum := s.Elements[0]
	for _, e := range s.Elements {
		if e > maximum {
			maximum = e
		}
	}
	s.maximum = maximum
	s.hasMax = true
	return maximum
}

func (s *StatBucket) Mean() float64 {
	if !s.hasMean {
		s.CheckSize()
		s.mean = s.Sum() / float64(s.Count())
		s.hasMean = true
	}
	return s.mean
}

func (s *StatBucket) Percentile(percentile float64) float64 {
	s.CheckSize()

	// make sure we are sorted
	if !s.isSorted {
		sort.Float64s(s.Elements)
		s.isSorted = true
	}

	// identify key for this percentile entry
	k := float64(s.Count()-1) * percentile
	// floor and ceiling
	f := float64(int(k))
	c := float64(int(k + 0.5))
	// exact item
	if f == c {
		return s.Elements[int(k)]
	}
	// interpolate between 2 adjacent items
	d0 := s.Elements[int(f)] * (c - k)
	d1 := s.Elements[int(c)] * (k - f)
	return d0 + d1
}

func (s *StatBucket) P25() float64 {
	if s.hasp25 {
		return s.p25
	}
	v := s.Percentile(0.25)
	s.p25 = v
	s.hasp25 = true
	return v
}

func (s *StatBucket) P50() float64 {
	if s.hasp50 {
		return s.p50
	}
	v := s.Percentile(0.5)
	s.p50 = v
	s.hasp50 = true
	return v
}

func (s *StatBucket) P75() float64 {
	if s.hasp75 {
		return s.p75
	}
	v := s.Percentile(0.75)
	s.p75 = v
	s.hasp75 = true
	return v
}

func (s *StatBucket) P90() float64 {
	if s.hasp90 {
		return s.p90
	}
	v := s.Percentile(0.90)
	s.p90 = v
	s.hasp90 = true
	return v
}

func (s *StatBucket) P95() float64 {
	if s.hasp95 {
		return s.p95
	}
	v := s.Percentile(0.95)
	s.p95 = v
	s.hasp95 = true
	return v
}
