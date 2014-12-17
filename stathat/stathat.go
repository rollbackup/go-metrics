// Metrics output to StatHat.
package stathat

import (
	"github.com/rollbackup/metrics"
	"github.com/simonz05/stathat"
	"log"
	"time"
)

func Stathat(r metrics.Registry, d time.Duration, userkey string) {
	for {
		if err := sh(r, userkey); nil != err {
			log.Println(err)
		}
		time.Sleep(d)
	}
}

func sh(r metrics.Registry, userkey string) error {
	r.Each(func(name string, i interface{}) {
		switch metric := i.(type) {
		case metrics.Counter:
			stathat.PostCount(name, userkey, int(metric.Count()))
		case metrics.Gauge:
			stathat.PostValue(name, userkey, float64(metric.Value()))
		case metrics.GaugeFloat64:
			stathat.PostValue(name, userkey, float64(metric.Value()))
		case metrics.Histogram:
			h := metric.Snapshot()
			stathat.PostCount(name+".count", userkey, int(h.Count()))
		case metrics.Meter:
			m := metric.Snapshot()
			stathat.PostCount(name+".count", userkey, int(m.Count()))
		case metrics.Timer:
			t := metric.Snapshot()
			stathat.PostCount(name+".count", userkey, int(t.Count()))
		}
	})
	return nil
}
