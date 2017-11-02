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
min=0.952639, max=14.089473, avg=1.581041, median=1.381390
bulk select from disk_tbl
min=1.040869, max=2.066162, avg=1.384743, median=1.266105
```

CPU: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
