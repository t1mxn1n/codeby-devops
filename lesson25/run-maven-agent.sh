docker run -d \
  --name jenkins-maven-agent \
  --network jenkins \
  --restart=on-failure \
  -e JENKINS_URL=http://jenkins-blueocean:8080/ \
  -e JENKINS_AGENT_NAME=maven-agent \
  -e JENKINS_AGENT_WORKDIR=/home/jenkins/agent \
  -e JENKINS_SECRET=17b5c4ce883883d78b89f5ff8dbab8a59b825159097f2dc39a1f3c790bc097e4 \
  -v jenkins_agent_workdir:/home/jenkins/agent \
  jenkins-inbound-agent-maven