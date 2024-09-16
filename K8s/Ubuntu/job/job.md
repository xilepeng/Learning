

``` go
ubuntu@master:~$ vim job-demo.yaml

ubuntu@master:~$ kubectl create -f job-demo.yaml
job.batch/job-demo created

ubuntu@master:~$ kubectl get jobs
NAME       COMPLETIONS   DURATION   AGE
job-demo   1/1           18s        3m29s
ubuntu@master:~$ kubectl get pods
NAME                       READY   STATUS      RESTARTS   AGE
hpa-demo-c9ddb6864-tldck   1/1     Running     0          163m
job-demo--1-6lgx5          0/1     Completed   0          4m4s



ubuntu@master:~$ kubectl describe job job-demo
Name:             job-demo
Namespace:        default
Selector:         controller-uid=2afed7a9-972e-41ab-8e0a-83aa59ea576c
Labels:           controller-uid=2afed7a9-972e-41ab-8e0a-83aa59ea576c
                  job-name=job-demo
Annotations:      <none>
Parallelism:      1
Completions:      1
Completion Mode:  NonIndexed
Start Time:       Mon, 13 Dec 2021 16:17:27 +0800
Completed At:     Mon, 13 Dec 2021 16:17:45 +0800
Duration:         18s
Pods Statuses:    0 Running / 1 Succeeded / 0 Failed
Pod Template:
  Labels:  controller-uid=2afed7a9-972e-41ab-8e0a-83aa59ea576c
           job-name=job-demo
  Containers:
   counter:
    Image:      busybox
    Port:       <none>
    Host Port:  <none>
    Command:
      bin/sh
      -c
      for i in 9 8 7 6 5 4 3 2 1; do echo $i; done
    Environment:  <none>
    Mounts:       <none>
  Volumes:        <none>
Events:
  Type    Reason            Age    From            Message
  ----    ------            ----   ----            -------
  Normal  SuccessfulCreate  5m14s  job-controller  Created pod: job-demo--1-6lgx5
  Normal  Completed         4m56s  job-controller  Job completed


ubuntu@master:~$ kubectl logs job-demo--1-6lgx5
9
8
7
6
5
4
3
2
1

```



``` go
ubuntu@master:~$ vim cronjob.yaml
ubuntu@master:~$ kubectl apply -f cronjob.yaml
cronjob.batch/cronjob-demo configured

ubuntu@master:~$ kubectl get cronjob
NAME           SCHEDULE      SUSPEND   ACTIVE   LAST SCHEDULE   AGE
cronjob-demo   */1 * * * *   False     1        0s              2m36s

ubuntu@master:~$ kubectl get jobs
NAME                    COMPLETIONS   DURATION   AGE
cronjob-demo-27324528   1/1           18s        2m51s
cronjob-demo-27324529   1/1           18s        111s
cronjob-demo-27324530   1/1           18s        51s
job-demo                1/1           18s        24h


ubuntu@master:~$ kubectl get pods
NAME                             READY   STATUS      RESTARTS   AGE
cronjob-demo-27324529--1-z4v58   0/1     Completed   0          2m35s
cronjob-demo-27324530--1-b5chw   0/1     Completed   0          95s
cronjob-demo-27324531--1-98cpr   0/1     Completed   0          35s
hpa-demo-c9ddb6864-tldck         1/1     Running     0          27h
job-demo--1-6lgx5                0/1     Completed   0          24h


ubuntu@master:~$ kubectl delete -f cronjob.yaml
cronjob.batch "cronjob-demo" deleted
ubuntu@master:~$ kubectl get jobs
NAME       COMPLETIONS   DURATION   AGE
job-demo   1/1           18s        25h


ubuntu@master:~$ kubectl get jobs
NAME                    COMPLETIONS   DURATION   AGE
cronjob-demo-27324579   1/1           18s        4m35s
cronjob-demo-27324580   1/1           18s        3m35s
cronjob-demo-27324581   1/1           18s        2m35s
cronjob-demo-27324582   1/1           19s        95s
cronjob-demo-27324583   1/1           19s        35s
job-demo                1/1           18s        25h
ubuntu@master:~$ kubectl get cronjob
NAME           SCHEDULE      SUSPEND   ACTIVE   LAST SCHEDULE   AGE
cronjob-demo   */1 * * * *   False     0        52s             10m
ubuntu@master:~$ kubectl get pods
NAME                             READY   STATUS              RESTARTS   AGE
cronjob-demo-27324579--1-5xsgd   0/1     Completed           0          5m9s
cronjob-demo-27324580--1-fmk2l   0/1     Completed           0          4m9s
cronjob-demo-27324581--1-98q56   0/1     Completed           0          3m9s
cronjob-demo-27324582--1-m9bg6   0/1     Completed           0          2m9s
cronjob-demo-27324583--1-gm7tm   0/1     Completed           0          69s
cronjob-demo-27324584--1-hz7rd   0/1     ContainerCreating   0          9s


ubuntu@master:~$ kubectl delete cronjob cronjob-demo
cronjob.batch "cronjob-demo" deleted
ubuntu@master:~$ kubectl get cronjob
No resources found in default namespace.
```