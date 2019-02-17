Hướng dẫn đóng gói một ứng dụng Go vào trong 1 Docker Image
```
docker build -t dockergo .
docker run -d -p 9090:9090 dockergo
curl http://localhost:9090
```
