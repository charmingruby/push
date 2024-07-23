package endpoint_metric

// const (
// 	// Labels
// 	endpointUrl         string = "endpointUrl"
// 	verb                string = "verb"
// 	pattern             string = "pattern"
// 	failed              string = "failed"
// 	error               string = "error"
// 	responseCode        string = "response_code"
// 	isAvailabilityError string = "is_availability_error"
// 	isInfraError        string = "is_infra_error"

// 	// Names
// 	endpointRequestCounter string = "endpoint_request_counter"
// 	endpointRequestLatency string = "endpoint_request_latency"
// )

// type EndpointMetrics struct {
// 	// Metric
// 	Latency float64

// 	// Labels
// 	Endpoint     string
// 	Verb         string
// 	Pattern      string
// 	ResponseCode int
// 	Failed       bool
// 	Error        string
// 	IsInfraError bool
// }

// func Send(metrics EndpointMetrics) {
// 	labels := map[string]string{
// 		endpointUrl:  metrics.Endpoint,
// 		verb:         metrics.Verb,
// 		pattern:      metrics.Pattern,
// 		responseCode: fmt.Sprintf("%d", metrics.ResponseCode),
// 		failed:       fmt.Sprintf("%v", metrics.Failed),
// 		error:        metrics.Error,
// 		isInfraError: fmt.Sprintf("%v", metrics.IsInfraError),
// 	}
// }
