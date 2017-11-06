# Run Benchmark (minikube)

```
$ kubectl apply -f mysql.yaml --validate=false
$ kubectl run -it --rm --image=bazel/setup:devel --restart=Never dbsetup
$ eval $(minikube docker-env)
$ bazel run //app:devel
$ kubectl run -it --rm --image=bazel/app:devel --restart=Never benchmark
```

## Execution Results

```
bulk select from mem_tbl:
min=0.967750, max=13.576060, avg=1.671941, median=1.330782
bulk select from disk_tbl
min=1.020313, max=8.635318, avg=1.385587, median=1.196277
random select 10 rows by id from disk_tbl
min=3.620647, max=7.893772, avg=4.149065, median=3.991366
random select 100 rows by id from disk_tbl
min=36.957552, max=60.893856, avg=41.744088, median=40.520797
random select 10 rows by id from disk_tbl
min=3.151550, max=7.162153, avg=3.494655, median=3.359065
random select 100 rows by id from disk_tbl
min=30.343917, max=48.088546, avg=33.843704, median=32.829733
```

CPU: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
