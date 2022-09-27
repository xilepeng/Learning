

``` s
docker pull docker.dragonflydb.io/dragonflydb/dragonfly

docker run -itd --name dragonfly --network=host --ulimit memlock=-1 -p 6379:6379 docker.dragonflydb.io/dragonflydb/dragonfly

```


