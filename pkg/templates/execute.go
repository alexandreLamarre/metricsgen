package templates

import (
	"bytes"
	_ "embed"
)

func ExecuteMetrics(cfg GenConfig) ([]byte, error) {
	var b bytes.Buffer
	if err := metricsGenOtelTemplate.Execute(&b, cfg); err != nil {
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
