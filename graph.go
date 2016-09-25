package main

import (
	"fmt"
)

/*
histogram takes an array of floats and counts the numbers into bins.
*/
func histogram(entries StatBucket, bins int) []int {
	minimum := entries.Min()
	maximum := entries.Max()

	bars := make([]int, bins)
	divisor := (maximum - minimum) / float64(bins)

	// count up in buckets
	for _, e := range entries.Elements {
		kf := (e - minimum) / divisor
		k := int(kf)
		if k >= bins {
			k = bins - 1
		}
		bars[k]++
	}
	return bars
}

/* normalize the columns down to 0.0-1.0 range */
func normalize(entries []int) []float64 {
	maximum := float64(entries[0])
	for _, e := range entries {
		if float64(e) > maximum {
			maximum = float64(e)
		}
	}

	bars := make([]float64, len(entries))
	for i, e := range entries {
		bars[i] = float64(e) / maximum
	}
	return bars
}

/* given the normalized bar hieghts, and a bar position and height
return the correct bar character */
func barChar(entries []float64, bar int, mi float64, mx float64) string {
	v := entries[bar]

	m2 := mi + (mx-mi)/3
	m3 := m2 + (mx-mi)/3

	if v > m3 {
		return "█"
	} else if v > m2 {
		return "▄"
	} else if v > mi {
		return "_"
	}

	return " "
}

func PrintGraph(entries StatBucket, graphHeight int, graphWidth int) error {
	fmt.Println("Distribution (normalized):")
	fmt.Println("-------------------------")

	normalized := normalize(histogram(entries, graphWidth))
	dv := float64(1) / float64(graphHeight)
	for i := graphHeight; i > 0; i-- {
		fi := float64(i) / float64(graphHeight)
		for b := 0; b < graphWidth; b++ {
			fmt.Printf(barChar(normalized, b, fi-dv, fi))
		}
		fmt.Println()
	}
	for i := 0; i < graphWidth; i++ {
		fmt.Printf("-")
	}
	fmt.Println()

	lp := graphWidth / 2
	lf := fmt.Sprintf("%%-%ds", lp)
	rf := fmt.Sprintf("%%%ds", graphWidth-lp)
	fmt.Printf(lf, PrettyDuration(entries.Min()))
	fmt.Printf(rf, PrettyDuration(entries.Max()))
	fmt.Println()

	return nil
}
