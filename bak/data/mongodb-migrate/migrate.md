
# mongodb migrate

## rename db

* mongodump --archive='mongodump-account-db' --db=account
* mongorestore --archive='mongodump-account-db' --nsFrom='account.*' --nsTo='user.*'

### 如何保持服务

* 所有用到该数据库名字的服务，数据库名字需要调整成新数据库
* 两个数据库共存，并且需要保证新数据库重命名之后，两边数据一致
* 

## rename collection

* db.accounts.renameCollection('users');

