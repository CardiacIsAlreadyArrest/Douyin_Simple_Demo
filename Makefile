## 启动全部服务和 api 层
run_all:
	make run_user &
	make run_publish &
	make run_feed &
	make run_api

## 启动 api 层
run_api:
	cd cmd/api && sh ./build.sh && sh ./output/bootstrap.sh

## 启动 user 服务
run_user:
	cd cmd/user && sh ./build.sh && sh ./output/bootstrap.sh

## 启动 publish 服务
run_publish:
	cd cmd/publish && sh ./build.sh && sh ./output/bootstrap.sh

## 启动 feed 服务
run_feed:
	cd cmd/feed && sh ./build.sh && sh ./output/bootstrap.sh

## 启动 favorite 服务
run_favorite:
	cd cmd/favorite && sh ./build.sh && sh ./output/bootstrap.sh


## 启动相关服务
start:
	docker compose --profile dev up -d

## 关闭相关服务
stop:
	docker-compose --profile dev stop

## 关闭并删除
down:
	docker-compose --profile dev down