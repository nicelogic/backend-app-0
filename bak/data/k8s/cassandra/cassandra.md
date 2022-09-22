
# cassandra

## cmd


helm --kubeconfig tecent-cloud/base.luojm.com/token/admin.conf uninstall k8ssandra -n cassandra

registry.cn-shanghai.aliyuncs.com/logic-base/kube-webhook-certgen

kubectl edit sts k8ssandra-dc1-default-sts

## namespace problem

(
NAMESPACE=cassandra
kubectl proxy &
kubectl get namespace $NAMESPACE -o json |jq '.spec = {"finalizers":[]}' >temp.json
curl -k -H "Content-Type: application/json" -X PUT --data-binary @temp.json 127.0.0.1:8001/api/v1/namespaces/$NAMESPACE/finalize
)