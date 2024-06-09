
CLUSTER="clicker-cluster"
stage?=local
STAGE=$(stage)
REGISTY?=localhost:5001

ifeq ($(STAGE), local)
	REGISTY=localhost:5001
else ifeq ($(STAGE), dev)
	REGISTY=localhost:5001
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
	@make deploy-game-redis 
	@make deploy-log-redis
	@make deploy-mysql

remove-base-all:
	@make remove-argo 
	@make remove-nats 
	@make remove-game-redis 
	@make remove-log-redis
	@make remove-mysql

deploy-app-all:
	@make deploy-login

remove-app-all:
	@make remove-login

repo-add:
	@helm repo add bitnami https://charts.bitnami.com/bitnami
	@helm repo add nats https://nats-io.github.io/k8s/helm/charts/
	@helm repo add argo https://argoproj.github.io/argo-helm
	@helm repo update

# 로그인 서버 배포
push-login:
	@docker build -t ${REGISTY}/login:latest -f cmd/login/Dockerfile .
	@docker push ${REGISTY}/login:latest

deploy-login:
	@helm upgrade --install login install/helm/login -n login -f values/login/values.${STAGE}.yaml --create-namespace

remove-login:
	@helm uninstall login -n login

# 레디스 배포 
deploy-game-redis:
	@helm upgrade --install game-redis install/helm/game-redis-19.5.2 -n redis -f values/game-redis/values.${STAGE}.yaml --create-namespace

remove-game-redis:
	@helm uninstall game-redis -n redis

deploy-log-redis:
	@helm upgrade --install log-redis install/helm/log-redis-19.5.2 -n redis -f values/log-redis/values.${STAGE}.yaml --create-namespace

remove-log-redis:
	@helm uninstall log-redis -n redis

# 레디스 로깅
push-log-client:
	@docker build -t ${REGISTY}/log-client:latest -f cmd/log-client/Dockerfile .
	@docker push ${REGISTY}/log-client:latest

deploy-log-client:
	@helm upgrade --install log-client install/helm/log-client -n redis-client -f values/log-client/values.${STAGE}.yaml --create-namespace

remove-log-client:
	@helm uninstall log-client -n redis-client

# 레디스 게임
push-game-client:
	@docker build -t ${REGISTY}game-client:latest -f cmd/game-client/Dockerfile .
	@docker push ${REGISTY}/game-client:latest

deploy-game-client:
	@helm upgrade --install game-client install/helm/game-client -n redis-client -f values/game-client/values.${STAGE}.yaml --create-namespace

remove-game-client:
	@helm uninstall game-client -n redis-client

# 게임 서버 배포

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


deploy-debug:
	kubectl apply -f https://k8s.io/examples/application/nginx-with-request.yaml