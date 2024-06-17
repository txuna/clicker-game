# Clicker Game
화면을 클릭하여 게임내 재화를 벌고 성장하는 게임입니다. 

### Architecture
![alt text](./image/arch.png)

### Build
```bash
git clone https://github.com/txuna/clicker-game.git
cd clicker-game/server
make kind-create

make push-login stage=local
make push-game stage=local
make push-redis-client stage=local

make deploy-base-all
# wait for a minute
make deploy-app-all

# port-forward
kubectl port-forward svc/mysql -n mysql 3307:3306
kubectl port-forward svc/login -n login 9001:9001
kubectl port-forward svc/game -n game 9003:9003
kubectl port-forward svc/loki-grafana -n loki-stack 3000:80

# init mysql 
make init-mysql

```

### TEST
```bash
cd clicker-game/client/test_client
go build
./client
```

### DEMO
![alt text](./image/k9s.png)

![alt text](./image/image.png)

![alt text](./image/image-1.png)

### Cleanup 
```bash
make remove-app-all
make remove-base-all
make kind-delete
```