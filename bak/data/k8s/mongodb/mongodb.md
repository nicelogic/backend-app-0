
# mongodb

## enterprise

### ops-manager

kubectl delete ns mongodb-ops-manager


kubectl create secret generic  ops-manager-admin-secret\
  --from-literal=Username="logic" \
  --from-literal=Password="ccccc123" \
  --from-literal=FirstName="nice" \
  --from-literal=LastName="logic" -n mongodb

## community


kubectl create secret generic base-mongodb-user-password -n mongodb --from-literal="password=ccccc123"



<!-- mongo "mongodb://warmth-mongo.luojm.com:27017/?replicaSet=base-mongodb" --username logic --password ccccc123 --authenticationDatabase warmth -->
<!-- mongo "mongodb://base-mongodb-svc.mongodb.svc.cluster.local:27017/?replicaSet=base-mongodb" --username logic --password ccccc123 --authenticationDatabase warmth -->
<!-- mongo "mongodb://admin:ccccc123@warmth-mongo.luojm.com:27017/warmth?readPreference=primary&directConnection=true&ssl=true" -->
mongo "mongodb://logic:ccccc123@warmth-mongo.luojm.com:27017/warmth?readPreference=secondaryPreferred&directConnection=true&ssl=true"

rs.initiate();
//可以显示当前的状态
rs.status(); 

db.createUser({user:"test",pwd:"ccccc123",roles:["clusterAdmin","readWriteAnyDatabase","dbAdminAnyDatabase","userAdminAnyDatabase"]})

db.getFreeMonitoringStatus