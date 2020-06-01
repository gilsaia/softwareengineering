cd ..
docker rm -f account
docker build -f account/Dockerfile -t account .
docker run -d --network project --network-alias account --name account account