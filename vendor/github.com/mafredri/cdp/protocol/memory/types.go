// Code generated by cdpgen. DO NOT EDIT.

package memory

// PressureLevel Memory pressure level.
type PressureLevel string

// PressureLevel as enums.
const (
	PressureLevelNotSet   PressureLevel = ""
	PressureLevelModerate PressureLevel = "moderate"
	PressureLevelCritical PressureLevel = "critical"
)

func (e PressureLevel) Valid() bool {
	switch e {
	case "moderate", "critical":
		return true
	default:
		return false
	}
}

func (e PressureLevel) String() string {
	return string(e)
}

// SamplingProfileNode Heap profile sample.
type SamplingProfileNode struct {
	Size  float64  `json:"size"`  // Size of the sampled allocation.
	Count float64  `json:"count"` // Number of sampled allocations of that size.
	Stack []string `json:"stack"` // Execution stack at the point of allocation.
}

// SamplingProfile Array of heap profile samples.
type SamplingProfile struct {
	Samples []SamplingProfileNode `json:"samples"` // No description.
}
