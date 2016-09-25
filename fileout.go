package main

import (
    "time"
    "encoding/json"
    "io/ioutil"
)

type outputStruct struct {
    Timings []float64           `json:"timings"`
    Min float64                 `json:"min"`
    Max float64                 `json:"max"`
    Mean float64                `json:"mean"`
    P25 float64                 `json:"p25"`
    P50 float64                 `json:"p50"`
    P75 float64                 `json:"p75"`
    P90 float64                 `json:"p90"`
    P95 float64                 `json:"p95"`
    OutputTimestamp time.Time   `json:"output_timestamp"`
}

/*
WriteOutputData takes the given stats and writes the raw data and some
summarized statistics to the given file.
*/
func WriteOutputData(stats StatBucket, filename string) error {

    data := outputStruct{}
    data.Timings = stats.Elements
    data.Min = stats.Min()
    data.Max = stats.Max()
    data.Mean = stats.Mean()
    data.P25 = stats.P25()
    data.P50 = stats.P50()
    data.P75 = stats.P75()
    data.P90 = stats.P90()
    data.P95 = stats.P95()
    data.OutputTimestamp = time.Now()

    b, err := json.MarshalIndent(data, "", "  ")
    if err != nil {
        return err
    }

    return ioutil.WriteFile(filename, b, 0644)
}
