// package main

// import (
// 	"fmt"
// )

// // Metric represents the structure of a single metric entry.
// type Metric struct {
// 	Name   string            `json:"name"`
// 	Labels map[string]string `json:"labels"`
// 	Values []int             `json:"values"`
// }

// // AggregateMetricsByMethod aggregates the values of metrics with the same name and method.
// func AggregateMetricsByMethod(metrics []Metric) []Metric {
// 	aggregated := make(map[string]Metric)

// 	// Iterate through metrics and aggregate values based on name and method.
// 	for _, metric := range metrics {
// 		method := metric.Labels["method"]
// 		key := metric.Name + "|" + method // Unique key for grouping

// 		if _, exists := aggregated[key]; !exists {
// 			// Initialize if not already present
// 			aggregated[key] = Metric{
// 				Name:   metric.Name,
// 				Labels: map[string]string{"method": method},
// 				Values: make([]int, len(metric.Values)),
// 			}
// 		}

// 		// Aggregate values
// 		aggMetric := aggregated[key]
// 		for i, value := range metric.Values {
// 			aggMetric.Values[i] += value
// 		}
// 		aggregated[key] = aggMetric
// 	}

// 	// Convert map to slice
// 	var result []Metric
// 	for _, metric := range aggregated {
// 		result = append(result, metric)
// 	}
// 	return result
// }

// func main() {
// 	// Example metrics data
// 	metrics := []Metric{
// 		{
// 			Name:   "http_requests",
// 			Labels: map[string]string{"path": "/user", "method": "GET"},
// 			Values: []int{1, 2, 3},
// 		},
// 		{
// 			Name:   "http_requests",
// 			Labels: map[string]string{"path": "/group", "method": "GET"},
// 			Values: []int{5, 1, 2},
// 		},
// 		{
// 			Name:   "http_requests",
// 			Labels: map[string]string{"path": "/group", "method": "POST"},
// 			Values: []int{1, 2, 3},
// 		},
// 	}

// 	// Aggregate metrics by method
// 	aggregatedMetrics := AggregateMetricsByMethod(metrics)

// 	// Print results
// 	for _, metric := range aggregatedMetrics {
// 		fmt.Printf("Name: %s, Method: %s, Values: %v\n",
// 			metric.Name, metric.Labels["method"], metric.Values)
// 	}
// }

package main

import (
	"fmt"
)

// Metric represents the structure of a single metric entry.
type Metric struct {
	Name   string            `json:"name"`
	Labels map[string]string `json:"labels"`
	Values []int             `json:"values"`
}

// SumByNameAndMethods calculates the sum of the values for metrics grouped by name and method.
func SumByNameAndMethods(metrics []Metric, methods []string) []Metric {
	// Map to store the aggregated results for each (name, method) pair
	resultMap := make(map[string]map[string][]int)

	// Initialize resultMap
	for _, metric := range metrics {
		if _, exists := resultMap[metric.Name]; !exists {
			resultMap[metric.Name] = make(map[string][]int)
		}
	}

	// Aggregate values for each (name, method) pair
	for _, metric := range metrics {
		method := metric.Labels["method"]
		if contains(methods, method) {
			if _, exists := resultMap[metric.Name][method]; !exists {
				resultMap[metric.Name][method] = make([]int, len(metric.Values))
			}
			for i, value := range metric.Values {
				resultMap[metric.Name][method][i] += value
			}
		}
	}

	// Build the resulting list of metrics
	var result []Metric
	for name, methodsMap := range resultMap {
		for method, values := range methodsMap {
			result = append(result, Metric{
				Name:   name,
				Labels: map[string]string{"method": method},
				Values: values,
			})
		}
	}

	return result
}

// Helper function to check if a string is in a slice
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func main() {
	// Example metrics data
	metrics := []Metric{
		{
			Name:   "http_requests",
			Labels: map[string]string{"path": "/user", "method": "GET"},
			Values: []int{1, 2, 3},
		},
		{
			Name:   "http_requests",
			Labels: map[string]string{"path": "/group", "method": "GET"},
			Values: []int{5, 1, 2},
		},
		{
			Name:   "http_requests",
			Labels: map[string]string{"path": "/group", "method": "POST"},
			Values: []int{1, 2, 3},
		},
		{
			Name:   "https_requests",
			Labels: map[string]string{"path": "/user", "method": "GET"},
			Values: []int{1, 2, 3},
		},
	}

	// List of methods to aggregate
	methods := []string{"GET", "POST"}

	// Compute the sums for the specified methods
	result := SumByNameAndMethods(metrics, methods)

	// Print the result
	for _, metric := range result {
		fmt.Printf("Name: %s, Method: %s, Values: %v\n", metric.Name, metric.Labels["method"], metric.Values)
	}
}