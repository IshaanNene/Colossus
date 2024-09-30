// timing.go
package internal

import "time"

// Clock represents the timing for instruction execution
type Clock struct {
	CycleTime time.Duration
	LastCycle time.Time
}

// NewClock creates a new Clock instance with a specified cycle time
func NewClock(cycleTime time.Duration) *Clock {
	return &Clock{
		CycleTime: cycleTime,
		LastCycle: time.Now(),
	}
}

// WaitForNextCycle waits until the next clock cycle
func (clock *Clock) WaitForNextCycle() {
	now := time.Now()
	elapsed := now.Sub(clock.LastCycle)
	if elapsed < clock.CycleTime {
		time.Sleep(clock.CycleTime - elapsed)
	}
	clock.LastCycle = time.Now()
}
