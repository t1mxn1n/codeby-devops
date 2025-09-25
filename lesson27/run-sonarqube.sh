docker run -d --name sonarqube \
  --network jenkins \
  -p 9000:9000 \
  sonarqube:community