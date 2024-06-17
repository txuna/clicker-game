
CLUSTER="clicker-cluster"
stage?=local
STAGE=$(stage)
REGISTY?=localhost:5001

ifeq ($(STAGE), local)
	REGISTY=localhost:5001
else ifeq ($(STAGE), live)
	REGISTY=txuna
endif

kind-create:
	@sh ./kind/create-cluster.sh ${CLUSTER}

kind-delete:
	@kind delete cluster -n ${CLUSTER}

gen-proto:
	protoc --go_out=./pkg --go_opt=paths=source_relative \
    		--go-grpc_out=./pkg --go-grpc_opt=paths=source_relative \
    		./apis/*.proto%

deploy-base-all:
	@make repo-add 
	@make deploy-argo 
	@make deploy-nats 
	@make deploy-redis 
	@make deploy-mysql
	@make deploy-loki-stack

remove-base-all:
	@make remove-argo 
	@make remove-nats 
	@make remove-redis 
	@make remove-mysql
	@make remove-loki-stack

deploy-app-all:
	@make deploy-login
	@make deploy-redis-client 
	@make deploy-game 

remove-app-all:
	@make remove-login
	@make remove-game
	@make remove-redis-client

repo-add:
	@helm repo add bitnami https://charts.bitnami.com/bitnami
	@helm repo add nats https://nats-io.github.io/k8s/helm/charts/
	@helm repo add argo https://argoproj.github.io/argo-helm
	@helm repo add grafana https://grafana.github.io/helm-charts
	@helm repo update

print-grafana-password:
	@kubectl get secret --namespace loki-stack loki-grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo

deploy-loki-stack:
	@helm upgrade --install loki --namespace=loki-stack grafana/loki-stack --set grafana.enabled=true --create-namespace 

remove-loki-stack:
	@helm uninstall loki -n loki-stack

# 로그인 서버 배포
push-login:
	@docker build -t ${REGISTY}/login:latest -f cmd/login/Dockerfile .
	@docker push ${REGISTY}/login:latest

deploy-login:
	@helm upgrade --install login install/helm/login -n login -f values/login/values.${STAGE}.yaml --create-namespace

remove-login:
	@helm uninstall login -n login

# 레디스 배포 
deploy-redis:
	@helm upgrade --install redis install/helm/redis-19.5.2 -n redis -f values/redis/values.${STAGE}.yaml --create-namespace

remove-redis:
	@helm uninstall redis -n redis

# 레디스 클라이언트
push-redis-client:
	@docker build -t ${REGISTY}/redis-client:latest -f cmd/redis-client/Dockerfile .
	@docker push ${REGISTY}/redis-client:latest

deploy-redis-client:
	@helm upgrade --install redis-client install/helm/redis-client -n redis -f values/redis-client/values.${STAGE}.yaml --create-namespace

remove-redis-client:
	@helm uninstall redis-client -n redis

# 게임 서버 배포
push-game: 
	@docker build -t ${REGISTY}/game:latest -f cmd/game/Dockerfile .
	@docker push ${REGISTY}/game:latest

deploy-game:
	@helm upgrade --install game install/helm/game -n game -f values/game/values.${STAGE}.yaml --create-namespace

remove-game:
	@helm uninstall game -n game

# nats 배포
deploy-nats:
	@helm upgrade --install nats install/helm/nats-1.1.12 -n nats -f values/nats/values.${STAGE}.yaml --create-namespace

remove-nats:
	@helm uninstall nats -n nats

# mysql 배포
deploy-mysql:
	@helm upgrade --install mysql install/helm/mysql-11.1.2 -n mysql -f values/mysql/values.${STAGE}.yaml --create-namespace

remove-mysql:
	@helm uninstall mysql -n mysql

# argo 배포
deploy-argo:
	@helm upgrade --install argocd argo/argo-cd -n argocd --create-namespace

remove-argo:
	@helm uninstall argocd -n argocd

init-mysql:
	@mysql -h 127.0.0.1 -P 3307 -u root -p1234qwer! clicker-game < sql/clicker-game.sql

deploy-debug:
	kubectl apply -f https://k8s.io/examples/application/nginx-with-request.yaml
