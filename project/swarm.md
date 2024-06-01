# Docker Swarm

## push image to docker hub 

### 1. built image
```bash
cd /<service directory>
docker build -f <path to docker file>.dockerfile -t <username>/<service_name>:<version> .
```
### 2. push to docker hub
```bash
docker push <username>/<service_name>:<version>
```

### 3. swarm init
```bash
docker swarn init
```
if want to add more worker & manager
`docker swarm join-token worker` : add worker
`docker swarm join-token manager` : add manager

### 4. deploy to swarm
```bash
docker stack deploy -c <path to swarm.yml file> <app_name>
docker stack ps <app_name>
```

### 5. Scailing up&down service
```bash
docker service scale <service_name>=<#replicas>
```

### 6. Update service to new/old version
* build new version of image
* push to docker hub
```bash
docker service update --image <username>/<service_name>:<new_version> <app_service_name>
```
`app_service_name` : example myapp_logger-service

### 7. Stop service
solution_1
```bash
docker service scale <app_service_name>=0
```
solution_2
```bash
docker stack rm <app_name>
```
if u want to use it any more 
```bash
docker swarm leave --force
```


