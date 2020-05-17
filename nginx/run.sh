docker rm -f nginx
docker build -t nginx
docker run --name nginx -dp 80:80 --network project --network-alias nginx nginx