package enum

type PriorityLevel uint8

// All PriorityLevel list
const (
	// PriorityNone is default PriorityLevel
	PriorityNone     PriorityLevel = 0
	PriorityNormal   PriorityLevel = 1
	PriorityHigh     PriorityLevel = 2
	PriorityCritical PriorityLevel = 3
)
