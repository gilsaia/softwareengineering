docker rm -f simple_test
docker build -t simple_test .
docker run -dp 8081:8081 --network project --network-alias simple_test --name simple_test simple_test