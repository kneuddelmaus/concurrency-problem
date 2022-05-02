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

type MetricFnAndChan struct {
	metricFn   func() float64
	metricChan chan float64
}

//predefined method - do not modify
func fetchCPUMetric() float64 {
	//queries system and returns cpu usage
	//dummy return val
	//time.Sleep(2 * time.Second)
	return float64(2)
}

//predefined method - do not modify
func fetchMemMetric() float64 {
	//queries system and returns memory usage
	//dummy return val
	//time.Sleep(2 * time.Second)
	return float64(5)
}

//predefined method - do not modify
func fetchDiskMetric() float64 {
	//queries system and returns disk usage
	//dummy return val
	//time.Sleep(2 * time.Second)
	return float64(7)
}

func createMetricFetchersAndResultChans(metricFetchers ...func() float64) []MetricFnAndChan {
	var metricFetchersAndResultChans []MetricFnAndChan
	for _, metricFetcher := range metricFetchers {
		metricFetchersAndResultChans = append(metricFetchersAndResultChans, MetricFnAndChan{metricFetcher, make(chan float64, 1)})
	}
	return metricFetchersAndResultChans
}

//returns a slice containing 3 elements in this order :
// 0 - cpu metric
// 1 - mem metric
// 2 - disk metric
func fetchPayload() []float64 {
	var metrics []float64
	metricFetchersAndResultChans := createMetricFetchersAndResultChans(fetchCPUMetric, fetchMemMetric, fetchDiskMetric)
	c := make(chan struct{})

	wg := sync.WaitGroup{}
	for _, metricFetcherAndChan := range metricFetchersAndResultChans {
		wg.Add(1)
		go func(m MetricFnAndChan) {
			defer close(m.metricChan)
			defer wg.Done()
			m.metricChan <- m.metricFn()
		}(metricFetcherAndChan)
	}

	go func(wg *sync.WaitGroup) {
		defer close(c)
		defer wg.Wait()
		for _, metricFetcherAndResultChan := range metricFetchersAndResultChans {
			for metricChan := range metricFetcherAndResultChan.metricChan {
				metrics = append(metrics, metricChan)
			}
		}
	}(&wg)

	select {
	case <-c:
		return metrics
	case <-time.After(1 * time.Second):
		return metrics
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
