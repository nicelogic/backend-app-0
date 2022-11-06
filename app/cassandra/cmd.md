
#cmd

## cmd

kubectl exec -it cassandra-cluster-env0-dc1-default-sts-0 -n k8ssandra-operator -- /bin/bash
cqlsh -u cassandra-cluster-env0-superuser -p znk4uVfaCLm6hppEZaJl cassandra-cluster-env0-dc1-stargate-service