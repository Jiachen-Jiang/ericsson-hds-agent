package diskusage

import (
	"github.com/NodePrime/ericsson-hds-agent/agent/collectors"
)

// Run returns disk usage metrics
func Run() ([]*collectors.MetricResult, error) {
	data, err := loader()
	if err != nil {
		return nil, err
	}

	return preformatter(data)
}
