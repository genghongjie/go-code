v=1.0
build:
	echo "Begin build yinjianxia/httpserver:${v} "
	docker build -t yinjianxia/httpserver:${v} .

push:
	docker push yinjianxia/httpserver:${v}

run_image:
	docker run -d -p 8080:8080 httpserver yinjianxia/httpserver:${v}

test_ns_ip:
	ContainerID=$(docker ps|grep httpserver |  awk '{print $1}')
	PID=$(docker inspect --format "{{ .State.Pid }}" ContainerID)
	#nsenter --target $PID  --net --pidiner_id
	nsenter -t $PID -net ip a
	nsenter -t $PID -net ip r
