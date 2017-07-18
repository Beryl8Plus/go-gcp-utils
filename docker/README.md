build docker image
```
docker build -t go-gcp-utils .
```

run
```
docker-compose run app
```

to get ip of docker container
```
docker inspect $CONTAINER_ID
```

to get container id
```
docker ps
```

access app
```
http://container_ip:8080
```

run local
```
dev_appserver.py -A $GCLOUD_PROJECTID --host=0.0.0.0 app.yaml
```