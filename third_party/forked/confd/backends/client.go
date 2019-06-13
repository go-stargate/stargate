package backends

import (
	"errors"
	"strings"

	"github.com/go-stargate/stargate/third_party/forked/confd/backends/etcdv3"
	"github.com/go-stargate/stargate/third_party/forked/confd/backends/redis"
	"github.com/go-stargate/stargate/third_party/forked/confd/log"
)

// The StoreClient interface is implemented by objects that can retrieve
// key/value pairs from a backend store.
type StoreClient interface {
	GetValues(keys []string) (map[string]string, error)
	WatchPrefix(prefix string, keys []string, waitIndex uint64, stopChan chan bool) (uint64, error)
}

// New is used to create a storage client based on our configuration.
func New(config Config) (StoreClient, error) {

	if config.Backend == "" {
		config.Backend = "etcd"
	}
	backendNodes := config.BackendNodes

	if config.Backend == "file" {
		log.Info("Backend source(s) set to " + strings.Join(config.YAMLFile, ", "))
	} else {
		log.Info("Backend source(s) set to " + strings.Join(backendNodes, ", "))
	}

	switch config.Backend {
	case "etcdv3":
		return etcdv3.NewEtcdClient(backendNodes, config.ClientCert, config.ClientKey, config.ClientCaKeys, config.BasicAuth, config.Username, config.Password)

	case "redis":
		return redis.NewRedisClient(backendNodes, config.ClientKey, config.Separator)
	}
	return nil, errors.New("Invalid backend")
}
