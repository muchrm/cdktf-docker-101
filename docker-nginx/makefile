vlan:
	docker network create pvc-01

web:
	docker run -d --network pvc-01 --name web -v $(shell pwd)/www:/usr/share/nginx/html nginx

header:
	docker run -d --network pvc-01 --name header -v $(shell pwd)/header:/usr/share/nginx/html nginx

menu:
	docker run -d --network pvc-01 --name menu -v $(shell pwd)/menu:/usr/share/nginx/html nginx

content:
	docker run -d --network pvc-01 --name content -v $(shell pwd)/content:/usr/share/nginx/html nginx

build-proxy:
	docker build -t proxy proxy

run-proxy:
	docker run -d --name proxy --network pvc-01 -p 80:80 -v $(shell pwd)/proxy/conf.d:/etc/nginx/conf.d proxy

bomb:
	docker rm -f $(shell docker ps -aq)