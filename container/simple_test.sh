docker build -t simple_test .
docker run -d --network backend --network-alias simple_test --name simple_test simple_test