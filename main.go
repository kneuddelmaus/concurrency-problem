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

// predefined method - do not modify
func fetchCPUMetric() float64 {
	//queries system and returns cpu usage
	//dummy return val
	time.Sleep(time.Millisecond * 1000)
	return float64(2)
}

// predefined method - do not modify
func fetchMemMetric() float64 {
	//queries system and returns memory usage
	//dummy return val
	time.Sleep(time.Millisecond * 500)
	return float64(5)
}

// predefined method - do not modify
func fetchDiskMetric() float64 {
	//queries system and returns disk usage
	//dummy return val
	return float64(7)
}

// returns a slice containing 3 elements in this order :
// 0 - cpu metric
// 1 - mem metric
// 2 - disk metric
func fetchPayload() []float64 {
	//TODO complete this. write the most optimized implementation to have the fastest total execution time possible.

	//My solution:
	// 1. create a slice of float64 for storing metrics
	metrics := make([]float64, 3)
	// 2. create a slice of the functions
	metricFns := []func() float64{
		fetchCPUMetric, fetchMemMetric, fetchDiskMetric,
	}
	// 3. create a wait group
	wg := &sync.WaitGroup{}
	wg.Add(len(metricFns))
	// 4. Loop over metricFns and use the index to populate right index with right metric as asked
	for i, metricFunc := range metricFns {
		go func(idx int, metric func() float64) {
			defer wg.Done()

			metrics[idx] = metric()
		}(i, metricFunc)
	}

	// 5. Wait for all the go routines to finish
	wg.Wait()

	// 6. return populated metrics
	return metrics
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
