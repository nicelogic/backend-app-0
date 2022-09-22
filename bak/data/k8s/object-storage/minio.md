
# minio

## init

wget <https://github.com/minio/operator/releases/download/v4.2.7/kubectl-minio_4.2.7_linux_amd64> -O kubectl-minio
chmod +x kubectl-minio
mv kubectl-minio /usr/local/bin/

kubectl minio version

kubectl create namespace minio-tenant-1

kubectl minio tenant create minio-tenant-1       \
  --servers                 1                    \
  --volumes                 4                   \
  --capacity                5Gi                 \
  --storage-class           minio-local-storage \
   --namespace  minio-tenant-1

  kubectl get svc --namespace minio-tenant-1

## info

Username: admin
Password: 12a52881-750a-415a-a212-ea0213fa3176


## k8s 

kubectl minio init --namespace base --cluster-domain luojm.com
kubectl minio proxy -n base 

kubectl minio tenant delete  base --namespace base
kubectl minio tenant create  base --servers 2 --volumes 4 --capacity 10Gi --storage-class minio-local-storage --namespace base

----

kubectl minio init  
kubectl minio init  --cluster-domain 10.1.0.1
kubectl minio proxy  

kubectl create ns minio-tenant-base
kubectl minio tenant create  minio-base-tenant-1 --servers 2 --volumes 4 --capacity 10Gi --storage-class minio-local-storage  --namespace minio-tenant-base

kubectl minio delete -n minio-tenant-base
kubectl minio tenant delete  minio-base-tenant-1 -n minio-tenant-base

## monitor

kubectl minio tenant  info minio-base-tenant-1 -n minio-base

---

## error

### ERROR Unable to validate passed arguments in MINIO_ARGS:env+tls://s7T9j6cBCi7zkXoy4eUn:1HlD93hX8UOFoSRmpiBjeMa1rjoAuvw6qSde9DBn@operator.minio-base.svc.cluster.local:4222/webhook/v1/getenv/minio-base/minio-base: Get "https://operator.minio-base.svc.cluster.local:4222/webhook/v1/getenv/minio-base/minio-base?key=MINIO_ARGS": context deadline exceeded


## back

Tenant 'minio-base-tenant-1' created in 'minio-tenant-base' Namespace

  Username: admin 
  Password: 932740e2-5972-483f-9ed6-344ef05528e4 
  Note: Copy the credentials to a secure location. MinIO will not display these again.

+-------------+-----------------------------+-------------------+--------------+--------------+
| APPLICATION | SERVICE NAME                | NAMESPACE         | SERVICE TYPE | SERVICE PORT |
+-------------+-----------------------------+-------------------+--------------+--------------+
| MinIO       | minio                       | minio-tenant-base | ClusterIP    | 443          |
| Console     | minio-base-tenant-1-console | minio-tenant-base | ClusterIP    | 9443  

mc alias set tenant-base https://minio-tenant-base.luojm.com admin  932740e2-5972-483f-9ed6-344ef05528e4
mc ls tenant-base/test