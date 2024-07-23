package endpoint_metric

const (
	// Labels
	endpointUrl         string = "endpoint"
	verb                string = "verb"
	pattern             string = "pattern"
	failed              string = "failed"
	error               string = "error"
	responseCode        string = "response_code"
	isAvailabilityError string = "is_availability_error"
	isReliabilityError  string = "is_reliability_error"

	// Names
	endpointRequestCounter string = "endpoint_request_counter"
	endpointRequestLatency string = "endpoint_request_latency"
)

type EndpointMetrics struct {
	Latency              float64
	Endpoint             string
	Verb                 string
	Pattern              string
	ResponseCode         int
	Failed               bool
	Error                string
	HasAvailabilityError bool
	HasReliabilityError  bool
}
