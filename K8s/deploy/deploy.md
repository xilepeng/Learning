

```go
滚动升级
ubuntu@master:~$ kubectl rollout status deployment deployment-demo
deployment "deployment-demo" successfully rolled out
暂停
ubuntu@master:~$ kubectl rollout pause deployment deployment-demo
deployment.apps/deployment-demo paused
继续
ubuntu@master:~$ kubectl rollout resume deployment deployment-demo
deployment.apps/deployment-demo resumed

回滚
ubuntu@master:~$ kubectl rollout history deployment deployment-demo
deployment.apps/deployment-demo
REVISION  CHANGE-CAUSE
1         <none>
2         <none>


ubuntu@master:~$ kubectl apply -f deploy-demo.yaml --record=true
Flag --record has been deprecated, --record will be removed in the future
deployment.apps/deployment-demo configured
ubuntu@master:~$ kubectl rollout status deployment deployment-demo
Waiting for deployment "deployment-demo" rollout to finish: 2 out of 3 new replicas have been updated...
Waiting for deployment "deployment-demo" rollout to finish: 2 out of 3 new replicas have been updated...
Waiting for deployment "deployment-demo" rollout to finish: 2 out of 3 new replicas have been updated...
Waiting for deployment "deployment-demo" rollout to finish: 2 out of 3 new replicas have been updated...
Waiting for deployment "deployment-demo" rollout to finish: 2 out of 3 new replicas have been updated...
Waiting for deployment "deployment-demo" rollout to finish: 2 out of 3 new replicas have been updated...
Waiting for deployment "deployment-demo" rollout to finish: 1 old replicas are pending termination...
Waiting for deployment "deployment-demo" rollout to finish: 1 old replicas are pending termination...
Waiting for deployment "deployment-demo" rollout to finish: 1 old replicas are pending termination...
Waiting for deployment "deployment-demo" rollout to finish: 2 of 3 updated replicas are available...
deployment "deployment-demo" successfully rolled out


ubuntu@master:~$ kubectl rollout history deployment deployment-demo
deployment.apps/deployment-demo
REVISION  CHANGE-CAUSE
1         <none>
2         <none>
3         kubectl apply --filename=deploy-demo.yaml --record=true


ubuntu@master:~$ kubectl rollout undo deployment deployment-demo
deployment.apps/deployment-demo rolled back

ubuntu@master:~$ kubectl rollout status deployment deployment-demo
deployment "deployment-demo" successfully rolled out

ubuntu@master:~$ kubectl get deploy deployment-demo
NAME              READY   UP-TO-DATE   AVAILABLE   AGE
deployment-demo   3/3     3            3           14h

ubuntu@master:~$ kubectl get rs
NAME                         DESIRED   CURRENT   READY   AGE
deployment-demo-55dc4587bc   0         0         0       97m
deployment-demo-845d4d9dff   3         3         3       14h
deployment-demo-d9d8cf5c7    0         0         0       14h

ubuntu@master:~$ kubectl rollout undo deployment deployment-demo --to-revision=3
deployment.apps/deployment-demo rolled back

```