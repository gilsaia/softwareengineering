cd ..
docker rm -f car_port
docker build -f car_port/Dockerfile -t car_port .
docker run -d --network project --network-alias car_port --name car_port car_port