//Assume you are given three functions that simulate querying `/proc` directory to obtain and return metrics representing
//system cpu, mem, and disk usage.
//Remember that querying the file system involves making system calls that are blocking in nature.
//Your goal is to fill in just the method `fetchPayload` to return those metrics as a slice.
//Optimize your implementation to have the fastest total execution time possible.
//Don't spend more than 20 mins to complete the exercise.

package main

import (
	"fmt"
	"sync"
	"time"
)

//predefined method - do not modify
func fetchCPUMetric() float64 {
	//queries system and returns cpu usage
	//dummy return val
	return float64(2)
}

//predefined method - do not modify
func fetchMemMetric() float64 {
	//queries system and returns memory usage
	//dummy return val
	return float64(5)
}

//predefined method - do not modify
func fetchDiskMetric() float64 {
	//queries system and returns disk usage
	//dummy return val
	return float64(7)
}

//returns a slice containing 3 elements in this order :
// 0 - cpu metric
// 1 - mem metric
// 2 - disk metric
func fetchPayload() []float64 {
	metric := []float64{}
	cpuMetricChan := make(chan float64, 1)
	memMetricChan := make(chan float64, 1)
	diskMetricChan := make(chan float64, 1)
	c := make(chan struct{})

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()
		cpuMetricChan <- fetchCPUMetric()
	}()
	go func() {
		defer wg.Done()
		memMetricChan <- fetchMemMetric()
	}()
	go func() {
		defer wg.Done()
		diskMetricChan <- fetchDiskMetric()
	}()

	go func() {
		defer close(c)
		var chanSlice = []chan float64{
			cpuMetricChan, memMetricChan, diskMetricChan,
		}
		for _, metricChan := range chanSlice {
			metric = append(metric, <-metricChan)
		}
		wg.Wait()
	}()

	select {
	case <-c:
		return metric
	case <-time.After(1 * time.Second):
		return metric
	}
}

func main() {
	p := fetchPayload()
	fmt.Println(p)

	if len(p) == 3 {
		if p[0] == 2 && p[1] == 5 && p[2] == 7 {
			fmt.Println("Correct")
			return
		}
	}

	fmt.Println("Fail")
}
