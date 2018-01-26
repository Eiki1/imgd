package main

import "github.com/prometheus/client_golang/prometheus"

const namespace = "imgd"

var (
	inFlightGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Subsystem: "http",
		Name:      "in_flight_requests",
		Help:      "A gauge of requests currently being served.",
	})

	requestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: namespace,
		Subsystem: "http",
		Name:      "request_duration_seconds",
		Help:      "Histogram of the time (in seconds) each HTTP request took.",
		Buckets:   []float64{.001, .005, 0.0075, .01, .025, .1, .5, 1, 5},
	}, []string{"code"})

	responseSize = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: namespace,
		Subsystem: "http",
		Name:      "response_size_bytes",
		Help:      "A histogram of response sizes (in bytes) for requests.",
		Buckets:   []float64{100, 500, 1000, 2500, 5000},
	}, []string{})

	processingDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: namespace,
		Subsystem: "image",
		Name:      "processing_duration_seconds",
		Help:      "Histogram of the time (in seconds) image processing took.",
		Buckets:   []float64{.00025, .0005, 0.001, 0.0025, .005},
	}, []string{"resource"})

	getDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: namespace,
		Subsystem: "texture",
		Name:      "get_duration_seconds",
		Help:      "Histogram of the time (in seconds) each texture GET took.",
		Buckets:   []float64{.1, .25, .5, 1},
	}, []string{"source"})

	cacheDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: namespace,
		Subsystem: "cache",
		Name:      "operation_duration_seconds",
		Help:      "Histogram of the time (in seconds) each cache operation took.",
		Buckets:   []float64{.00005, .0001, .0005, 0.0025, .005, 0.0075, 0.01, 0.1},
	}, []string{"operation"})

	errorCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: "status",
			Name:      "errors",
			Help:      "Error events",
		},
		[]string{"event"},
	)

	cacheCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: "status",
			Name:      "cache",
			Help:      "Cache status",
		},
		[]string{"status"},
	)

	requestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: "status",
			Name:      "requests",
			Help:      "Resource requests",
		},
		[]string{"resource"},
	)

	// Latency on Get (source of skin) :tick:
	// Total latency for HTTP request (response code) :tick:
	// Latency on cache  (has, puul or add) :tick:
	// Gauge for cache hit, miss :tick:
	// Gauge for request (type) :tick:
	// Latency for processing (type) :tick:
)

func init() {
	prometheus.MustRegister(inFlightGauge)
	prometheus.MustRegister(requestDuration)
	prometheus.MustRegister(responseSize)
	prometheus.MustRegister(processingDuration)
	prometheus.MustRegister(getDuration)
	prometheus.MustRegister(cacheDuration)
	prometheus.MustRegister(errorCounter)
	prometheus.MustRegister(cacheCounter)
	prometheus.MustRegister(requestCounter)
}
