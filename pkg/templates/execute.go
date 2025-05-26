package templates

import (
	"bytes"
	_ "embed"
)

func ExecuteOtelMetrics(cfg GenConfig) ([]byte, error) {
	var b bytes.Buffer
	if err := metricsGenOtelTemplate.Execute(&b, cfg); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func ExecutePrometheusMetrics(cfg GenConfig) ([]byte, error) {
	var b bytes.Buffer
	if err := metricsGenPrometheusTemplate.Execute(&b, cfg); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func ExecuteDocs(cfg DocConfig) ([]byte, error) {
	var b bytes.Buffer
	if err := docsGenTemplate.Execute(&b, cfg); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
