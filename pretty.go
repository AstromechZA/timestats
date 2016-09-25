package main

import (
	"fmt"
)

var unitsNames = []string{"nanosecond", "microsecond", "millisecond", "second", "minute", "hour", "day"}
var unitRatios = []float64{1, 1000, 1000, 1000, 60, 60, 24}

/*
PrettyDuration takes in a number of nanoseconds and formats it as a string
duration. Eg: "10.002 seconds"
*/
func PrettyDuration(nanos float64) string {
	v := nanos
	lastU := unitsNames[0]
	for i := 0; i < len(unitsNames); i++ {
		if v < unitRatios[i] {
			break
		}
		lastU = unitsNames[i]
		v = v / unitRatios[i]
	}
	plural := ""
	if v != 1.0 {
		plural = "s"
	}
	return fmt.Sprintf("%.2f %s%s", v, lastU, plural)
}
