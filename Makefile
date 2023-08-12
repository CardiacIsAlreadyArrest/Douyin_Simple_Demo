## 启动全部服务和 api 层
run_all:
	make run_user &
	make run_api

## 启动 api 层
run_api:
	cd cmd/api && sh ./build.sh && sh ./output/bootstrap.sh

## 启动 user 服务
run_user:
	cd cmd/user && sh ./build.sh && sh ./output/bootstrap.sh