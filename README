# Run Benchmark (minikube)

```
$ kubectl apply -f mysql.yaml --validate=false
$ kubectl run -it --rm --image=mysql:5.7 --restart=Never mysql-client -- mysql -h db -pmysql
mysql> create database benchdb;
mysql> GRANT ALL privileges ON benchdb.* TO 'mysql'@'%' IDENTIFIED BY 'password';
mysql> FLUSH PRIVILEGES;
$ kubectl run -it --rm --image=mysql:5.7 --restart=Never mysql-client -- mysql -h db -pmysql < test-data.sql
$ eval $(minikube docker-env)
$ bazel run //app:devel
$ kubectl run -it --rm --image=bazel/app:devel --restart=Never benchmark
```

## Execution Results

```
#1
from mem_tbl: min=1.997226, max=6.255443, avg=2.527301
from disk_tbl: min=1.730931, max=2.604391, avg=2.094179
#2
from mem_tbl: min=1.775399, max=2.292928, avg=1.927773
from disk_tbl: min=1.957583, max=2.345315, avg=2.079053
#3
from mem_tbl: min=1.796944, max=2.877918, avg=2.046922
from disk_tbl: min=1.983932, max=2.354852, avg=2.081603
#4
from mem_tbl: min=1.853423, max=2.464446, avg=1.978605
from disk_tbl: min=1.999681, max=2.602639, avg=2.266579
#5
from mem_tbl: min=1.822558, max=2.366603, avg=2.088474
from disk_tbl: min=1.990294, max=2.414813, avg=2.173238
#6
from mem_tbl: min=1.858594, max=16.209477, avg=3.534728
from disk_tbl: min=1.908793, max=2.592190, avg=2.207912
#7
from mem_tbl: min=1.795732, max=3.953448, avg=2.492175
from disk_tbl: min=1.909732, max=2.679900, avg=2.229338
#8
from mem_tbl: min=1.846960, max=3.047939, avg=2.304677
from disk_tbl: min=1.883303, max=3.125494, avg=2.174343
#9
from mem_tbl: min=1.773979, max=2.130625, avg=1.901683
from disk_tbl: min=1.912291, max=2.332345, avg=2.052006
#10
from mem_tbl: min=1.933129, max=2.409287, avg=2.090920
from disk_tbl: min=1.934000, max=2.790869, avg=2.375897
```
