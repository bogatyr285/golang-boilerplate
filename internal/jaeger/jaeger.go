package jaeger

import (
	"os"
	"strconv"

	"contrib.go.opencensus.io/exporter/jaeger"
	"github.com/bogatyr285/golang-boilerplate/internal/buildinfo"
	"go.opencensus.io/trace"
)

type Jaeger struct {
	jexporter *jaeger.Exporter
}

func NewExporter(tracingEndpoint, appName string, bi buildinfo.BuildInfo) (*Jaeger, error) {
	hostname, _ := os.Hostname()

	jexporter, err := jaeger.NewExporter(jaeger.Options{
		AgentEndpoint: tracingEndpoint,
		Process: jaeger.Process{
			ServiceName: appName,
			Tags: []jaeger.Tag{
				jaeger.StringTag("version", bi.Version+` (`+bi.CommitHash+`@`+bi.BuildDate+`)`),
				jaeger.StringTag("process_id", strconv.Itoa(os.Getpid())),
				jaeger.StringTag("hostname", hostname),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	trace.RegisterExporter(jexporter)
	trace.ApplyConfig(trace.Config{
		DefaultSampler: trace.AlwaysSample(),
	})

	return &Jaeger{jexporter: jexporter}, nil
}

func (j *Jaeger) Stop() error {
	j.jexporter.Flush()

	return nil
}
