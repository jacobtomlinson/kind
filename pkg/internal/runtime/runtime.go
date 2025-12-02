package runtime

import (
	"os"

	"github.com/jacobtomlinson/kind/pkg/cluster"
	"github.com/jacobtomlinson/kind/pkg/log"
)

// GetDefault selected the default runtime from the environment override
func GetDefault(logger log.Logger) cluster.ProviderOption {
	switch p := os.Getenv("KIND_EXPERIMENTAL_PROVIDER"); p {
	case "":
		return nil
	case "podman":
		logger.Warn("using podman due to KIND_EXPERIMENTAL_PROVIDER")
		return cluster.ProviderWithPodman()
	case "docker":
		logger.Warn("using docker due to KIND_EXPERIMENTAL_PROVIDER")
		return cluster.ProviderWithDocker()
	case "nerdctl", "finch", "nerdctl.lima":
		logger.Warnf("using %s due to KIND_EXPERIMENTAL_PROVIDER", p)
		return cluster.ProviderWithNerdctl(p)
	default:
		logger.Warnf("ignoring unknown value %q for KIND_EXPERIMENTAL_PROVIDER", p)
		return nil
	}
}
