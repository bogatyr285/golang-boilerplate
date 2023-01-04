package grpc

// func metrics() {
// 	// prometheus
// 	promExporter, err := prometheus.NewExporter(prometheus.Options{
// 		Namespace: "golangsvc",
// 	})

// 	if err != nil {
// 		log.Fatalf("Failed to create Prometheus exporter: %v", err)
// 	}
// 	view.RegisterExporter(promExporter)

// 	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
// 	grpc_prometheus.Register(grpcSrv)

// 	mergedMux := http.NewServeMux()
// 	mergedMux.Handle("/metrics", promhttp.Handler())

// }
