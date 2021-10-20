v=1.1
build:
	echo "Begin build yinjianxia/httpserver:${v} "
	docker build -t yinjianxia/httpserver:${v} .

push:
	docker push yinjianxia/httpserver:${v}

run_server:
	docker run -d --name httpserver -p 8080:80 yinjianxia/httpserver:${v}

test_ns_ip:
	ContainerID=$( docker ps|grep httpserver | sed -n "1,1p" |awk '{print $1}')
	PID=$(docker inspect --format "{{ .State.Pid }}" $ContainerID)
	nsenter -t $PID -net ip a
	nsenter -t $PID -net ip r
