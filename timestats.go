package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
)

const usageString = `todo
`

/*
PrintStatistics takes in a StatBucket object and prints a variety of
measures to Stdout.
*/
func PrintStatistics(entries StatBucket) {
	fmt.Println("Statistics (Nanoseconds):")
	fmt.Println("-------------------------")

	fmt.Printf("Count %d\n", entries.Count())
	fmt.Printf("Min   %f (%s)\n", entries.Min(), PrettyDuration(entries.Min()))
	fmt.Printf("Mean  %f (%s)\n", entries.Mean(), PrettyDuration(entries.Mean()))
	fmt.Printf("Max   %f (%s)\n", entries.Max(), PrettyDuration(entries.Max()))
	fmt.Println()
	fmt.Printf("P25   %f (%s)\n", entries.P25(), PrettyDuration(entries.P25()))
	fmt.Printf("P50   %f (%s)\n", entries.P50(), PrettyDuration(entries.P50()))
	fmt.Printf("P75   %f (%s)\n", entries.P75(), PrettyDuration(entries.P75()))
	fmt.Printf("P90   %f (%s)\n", entries.P90(), PrettyDuration(entries.P90()))
	fmt.Printf("P95   %f (%s)\n", entries.P95(), PrettyDuration(entries.P95()))
}

/*
runIteration will run a single iteration of the given command and will return
the elapsed time, error code, and any specific failure message.

If showOutput is enabled, it will forward stdout/stderr for the process.
*/
func runIteration(command []string, showOutput bool) (time.Duration, int, error) {
	start := time.Now()
	cmd := exec.Command(command[0], command[1:]...)
	if showOutput {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	err := cmd.Run()
	elapsed := time.Since(start)
	if err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				return elapsed, status.ExitStatus(), err
			}
		} else {
			return elapsed, 127, err
		}
	}
	return elapsed, 0, nil
}

/*
runIterations will run many iterations of the same command. It will return the
array of elapsed times.

If showOutput is true, success and failure messages will be printed to stdout/
stderr.

It will sleep `interval` seconds between each run.
*/
func runIterations(command []string, iterations int, interval float64, showOutput bool) []float64 {
	timeEntries := []float64{}
	for i := 0; i < iterations; i++ {
		elapsed, exitCode, err := runIteration(command, showOutput)
		if showOutput {
			if exitCode == 0 {
				fmt.Printf("#%d succeeded after %s\n", i, elapsed)
			} else {
				os.Stderr.WriteString(fmt.Sprintf("#%d failed with code %d after %s\n", i, exitCode, err.Error()))
			}
		}
		timeEntries = append(timeEntries, float64(elapsed.Nanoseconds()))
		time.Sleep(time.Duration(interval) * time.Second)
	}
	return timeEntries
}

func mainInner() error {
	// count and interval settings
	countFlag := flag.Int("count", 1, "Number of iterations")
	intervalFlag := flag.Float64("interval", 0.0, "Seconds to wait between iterations")

	// output settings
	quietFlag := flag.Bool("quiet", false, "Don't print results of the command")
	//outputFileFlag := flag.String("output", "", "Output json file of data gathered")

	// graph dimension settings
	graphWidth := flag.Int("graphx", 100, "Width of the distribution graph")
	graphHeight := flag.Int("graphy", 10, "Height of the distribution graph")

	// set a more verbose usage message.
	flag.Usage = func() {
		os.Stderr.WriteString(usageString)
		flag.PrintDefaults()
	}
	// parse them
	flag.Parse()

	// expect a command
	if flag.NArg() < 1 {
		return errors.New("Error: Expected a command!")
	}

	// sanitize and check flags
	count := *countFlag
	if count < 1 {
		return errors.New("Error: Count must be greater than 1!")
	}
	interval := *intervalFlag
	if interval < 0 {
		return errors.New("Error: Interval must be greater than 0!")
	}
	if *graphWidth < 20 {
		return errors.New("Error: Graph width must be at least 20!")
	}
	if *graphHeight < 5 {
		return errors.New("Error: Graph height must be at least 5!")
	}

	// the command elements
	command := flag.Args()
	quiet := *quietFlag

	// gather times into statistical struct
	timings := runIterations(command, count, interval, !quiet)
	timeStats := StatBucket{Elements: timings}

	// print output stuff
	if !quiet {
		fmt.Println()
	}
	PrintStatistics(timeStats)
	fmt.Println()
	PrintGraph(timeStats, *graphHeight, *graphWidth)

	return nil
}

func main() {
	if err := mainInner(); err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}
