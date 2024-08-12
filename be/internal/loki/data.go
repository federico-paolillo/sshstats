package loki

type SampleLabels = map[string]string

type RawSample struct {
	Labels SampleLabels
	Value  float64
}
