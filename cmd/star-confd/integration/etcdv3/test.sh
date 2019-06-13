#!/bin/bash

export HOSTNAME="localhost"
export ETCDCTL_API="3"

etcdctl put /stargate_test/key foobar --endpoints=10.129.11.12:2379
etcdctl put /stargate_test/database/host 127.0.0.1 --endpoints=10.129.11.12:2379
etcdctl put /stargate_test/database/password p@sSw0rd --endpoints=10.129.11.12:2379
etcdctl put /stargate_test/database/port 3306 --endpoints=10.129.11.12:2379
etcdctl put /stargate_test/database/username confd --endpoints=10.129.11.12:2379
etcdctl put /stargate_test/upstream/app1 10.0.1.10:8080 --endpoints=10.129.11.12:2379
etcdctl put /stargate_test/upstream/app2 10.0.1.11:8080 --endpoints=10.129.11.12:2379

# Run confd
# confd --onetime --log-level debug --confdir ./integration/confdir --backend etcdv3 --node http://127.0.0.1:2379 --watch
go run ./cmd/star-confd --onetime --confdir ./cmd/star-confd/integration/confdir --log-level debug --backend=etcdv3 --node http://10.129.11.12:2379 --watch