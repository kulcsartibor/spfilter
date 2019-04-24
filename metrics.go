package filter

import (
	"sync"
)

// requestCount exports a prometheus metric that is incremented every time a query is seen by the example plugin.
/*var requestCount = prometheus.NewCounterVec(prometheus.CounterOpts{
	Namespace: plugin.Namespace,
	Subsystem: "spfilter",
	Name:      "request_count_total",
	Help:      "Counter of requests made.",
}, []string{"server"})
*/
var once sync.Once