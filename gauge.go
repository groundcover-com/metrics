package metrics

import (
	"sync/atomic"

	"github.com/groundcover-com/metrics/pkg/options"
	"github.com/groundcover-com/metrics/pkg/set"
)

type Gauge = set.Gauge

func CreateGauge(name string, labels map[string]string, f func() float64, opts options.Options) *Gauge {
	return defaultSet.CreateGauge(name, labels, f, opts)
}

func GetOrCreateGauge(
	name string,
	labels map[string]string,
	f func() float64,
	opts options.Options,
) *Gauge {
	return defaultSet.GetOrCreateGauge(name, labels, f, opts)
}

func CreateErrorGauge(name string, labels map[string]string, f func() float64) *Gauge {
	return CreateGauge(name, labels, f, options.Error)
}

func CreateWarningGauge(name string, labels map[string]string, f func() float64) *Gauge {
	return CreateGauge(name, labels, f, options.Warning)
}

func CreateInfoGauge(name string, labels map[string]string, f func() float64) *Gauge {
	return CreateGauge(name, labels, f, options.Info)
}

func GetOrCreateErrorGauge(name string, labels map[string]string, f func() float64) *Gauge {
	return GetOrCreateGauge(name, labels, f, options.Error)
}

func GetOrCreateWarningGauge(name string, labels map[string]string, f func() float64) *Gauge {
	return GetOrCreateGauge(name, labels, f, options.Warning)
}

func GetOrCreateInfoGauge(name string, labels map[string]string, f func() float64) *Gauge {
	return GetOrCreateGauge(name, labels, f, options.Info)
}

// Subtract from atomic uint variable with 64 bits.
// Useful for unsigned gauges.
func SubtractAtomicUint64(variable *atomic.Uint64, delta uint64) {
	variable.Add(^uint64(delta - 1))
}
