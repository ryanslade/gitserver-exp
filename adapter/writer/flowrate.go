package writer

import (
	"io"

	"github.com/mxk/go-flowrate/flowrate"
)

const (
	gigabit = int64(1000 * 1000 * 1000)
)

// A FlowrateWriter limits the write rate of w to 1 Gbps.
//
// We are cloning repositories from within the same network from another
// Sourcegraph service (zoekt-indexserver). This can end up being so fast that
// we harm our own network connectivity. In the case of zoekt-indexserver and
// gitserver running on the same host machine, we can even reach up to ~100
// Gbps and effectively DoS the Docker network, temporarily disrupting other
// containers running on the host.
//
// Google Compute Engine has a network bandwidth of about 1.64 Gbps
// between nodes, and AWS varies widely depending on instance type.
// We play it safe and default to 1 Gbps here (~119 MiB/s), which
// means we can fetch a 1 GiB archive in ~8.5 seconds.
func FlowrateWriter(w io.Writer) io.Writer {
	return flowrate.NewWriter(w, 1)
}