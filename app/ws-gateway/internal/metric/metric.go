package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	// GatewayOnlineTotals 在线人数
	GatewayOnlineTotals prometheus.Gauge
	// GatewayInputBytes 流入流量
	GatewayInputBytes prometheus.Counter
	// GatewayOutputBytes 流出流量
	GatewayOutputBytes prometheus.Counter
)

func init() {
	GatewayOnlineTotals = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "gateway",
		Subsystem: "websocket",
		Name:      "online_totals",
		Help:      "the totals of online websocket",
	})
	prometheus.MustRegister(GatewayOnlineTotals)

	GatewayInputBytes = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "gateway",
		Subsystem: "websocket",
		Name:      "input_bytes",
		Help:      "the bytes of websocket input",
	})
	prometheus.MustRegister(GatewayInputBytes)

	GatewayOutputBytes = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "gateway",
		Subsystem: "websocket",
		Name:      "output_bytes",
		Help:      "the bytes of websocket output",
	})
	prometheus.MustRegister(GatewayOutputBytes)
}
