
- From the book "Kubernetes in Actions, 2E."

- Declarative model to define an application
	- You describe the components that make up your application
	- Kubernetes turns this description into a running application
	- It then keeps the application healthy by restarting or recreating parts of it as needed.


- Kubernetes provies infrastructure-related mechnisms
	- (frees developers from those and allow them to focus on the business logic)
  - service discovery
  - horizontal scaling
  - self-healing
  - leader election

- Self-Service Deployment of Applications
	- Because Kubernetes presents all its worker nodes as a single deployment surface,
	- it no longer matters which node you deploy your application to.


- 쿠버네티스 아키텍쳐
  - Control Plane ('Master' nodes)
    - Kubernetes API
    - etcd
    - Scheduler
    - Controllers
  - Workerload Plane ('Worker' nodes)
    - Kublet
    - Container Runtime
    - Kubernetes Service Proxy


- 쿠버네티스가 앱을 실행하는 과정
  - DEFINE your application
    - yaml, json 파일을 통해 Object들을 정의
    - e.g. Deployment 2개, Services 2개로 manifest를 만들어 어플리케이션 배포
    - 1. Kubernetes API를 통해 manifest의 Object를 생성하여 etcd 저장
    - 2. 컨트롤러가 각 application 인스턴스에 해당하는 Object 생성
    - 3. 스케쥴러가 각각의 인스턴스에 node 배정
    - 4. Kublet이 instance가 node에 배정된 것을 notice받고, Container Runtime을 통해 application instace를 실행
    - 5. Kube Proxy는 application instance가 클라이언트로 부터 connection을 받을 준비가 되면 load balancer 설정함
    - 6. Kublet과 Controller는 system을 모니터링하고 application이 실행 상태를 유지하도록 함
  - SUBMIT application to the API
    - `kubectl` 커맨드라인으로 yaml, json파일을 오브젝트 단위로 나눠서 API에 전달

- CONTROLLER
  - API서버로 부터 Object 생성알림을 받으면, Kubernetes API  통해 Object 생성

- SCHEDULER
  - 컨트롤러 타입으로, application 인스턴스를 워크노드에 스케쥴하는 역할
  - 각 어플리케이션 인스턴스 오브젝트의 최적의 워커노드를 찾아 인스턴스에 할당 (API를 통해 object를 수정하여)

- KUBELET AND THE CONTAINER RUNTIME
  - 각 워커노드에서 실행되는 kublet은 컨트롤러 타입 중 하나.
  - 해당 워커 노드에 어플리케이션 인스턴스가 할당 될때까지 기다리다가, 어플리케이션을 실행
  - 컨테이너 runtime이 해당 어플리케이션 컨테이너를 실행하도록 함

- KUBE PROXY
  - 어플리케이션 deployment는 여러 어플리케이션 인스턴스로 구성될 수 있기 때문에, 로드밸런서가 이들을 하나의 IP로 expose 해야함
  - Kube Proxy도 컨트롤러 타입으로서 로드밸런서를 만드는 역할을 함


- KEEPING THE APPLICATION HEALTHY
  - 어플리케이션이 실행되면, kubelet은 어플리케이션이 종료되는 경우, 재시작하여 healthy한 상태를 유지함
  - 어플리케이션 인스턴스를 represent하는 object를 업데이트하여, 어플리케이션의 status를 report함





- Docker Container

```sh
ls kaida-0.1
  Dockerfile  Makefile  app.js  html/
cat Dockerfile
  FROM node:16
  COPY app.js /app.js
  COPY html /html
  ENTRYPOINT ["node", "app.js"]
docker build -t kiada:latest .
docker run --namee kiada-container -p 1234:8080 -d kiada
docker tag kiada jnuho/kiada:0.1

docker login -u jnuho docker.io
docker push jnuho/kiada:0.1

# Run the image on other Hosts
docker run --name kiada-container -p 1234:8080 -d jnuho/kiada:0.1

```

- API Resources

```sh
kubectl api-resources
kubectl explain pod
kubectl explain configmap
kubectl explain service

# 파드: pods, pod, po
kubectl get pod --all-namespaces
kubectl get po --all-namespaces
```

- 쿠버네티스는 오브젝트를 yaml기반 매니페스트 파일로 관리
  - apiVersion : 오브젝트가 어떤 API 그룹에 속하고 api버전은 몇인가
  - kind : 오브젝트가 어떤 API 리소스인가
  - metadata : 오브젝트 식별하기 위한 정보 (이름, 네임스페이스, 레이블)
    - Label (optional) : 정확한 리소스 이름 정의, 소유자, 종류 명시하여 식별하기 위한 목적
    - Annotations  : 쿠버네티스 애드온이 해당 오브젝트를 어떻게 처리할지 결정하기위한 설정 용도
      - e.g. LogAgent 애드온의 로그 수집설정
  - spec : 오브젝트가 가지고자 하는 데이터
    - API 리소스에 따라 spec 대신 data,rules,subjects 등 다른 속성 사용



- `kubectl` : 쿠버네티스 클라이언트
  - Kubernetes API Server와 통신 할 수 있는 툴

- 클러스터 생성하는 방법
  1. minikube 싱글 노드 클러스터
  2. kind 멀티 노드 클러스터
  3. GKE, EKS, kind 멀티 노드 클러스터


```sh
### METHOD 1. minikube로 싱글 node 클러스터 생성

# Create a single-node Kubernetes Cluster to test Kubernetes functionalities
# testing certain features related to managing apps on multiple nodes are limited

minikube start --driver=docker

kubectl cluster-info

### METHOD 2. kind
# 문서: https://kind.sigs.k8s.io/docs/user/quick-start/
# 다른 minikube와 같은 툴에비해 아직 unstable

curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.12.0/kind-linux-amd64
chmod +x ./kind
mv ./kind /usr/local/kind
ls -al /usr/local/kind
  drwxr-xr-x  2 root root 4096 Apr  5 15:20 kind
  alias k=kubectl
  complete -F __start_kubectl k
  source <(kubectl completion bash)
  export PATH=/usr/local/kind:$PATH

vim ~ubuntu/.bashrc
  export PATH=/usr/local/kind:$PATH
source .bashrc

cat kind-multi-node.yaml
  kind: Cluster
  apiVersion: kind.sigs.k8s.io/v1alpha3
  nodes:
  - role: control-plane
  - role: worker
  - role: worker

# 3-node cluster
kind create cluster --config kind-multi-node.yaml

kind get nodes
  kind-worker
  kind-control-plane
  kind-worker2

# Instead of using Docker to run containers, nodes created by kind use the CRI-O container runtime
crictl ps
docker ps


### METHOD 3. GKE로 멀티 node 클러스터 생성 (Google 가입 및 신용카드 등록 필요)

# Signing up for a Google account, in the unlikely case you don’t have one already.
# Creating a project in the Google Cloud Platform Console.
# Enabling billing. This does require your credit card info, but Google provides a
#   12-month free trial. And they’re nice enough to not start charging automatically after the free trial is over.)
# Enabling the Kubernetes Engine API.
# Downloading and installing Google Cloud SDK. (This includes the gcloud
#   command-line tool, which you’ll need to create a Kubernetes cluster.)
# Installing the kubectl command-line tool with `gcloud components install kubectl`

# After completing the installation, you can create a Kubernetes cluster with three
#   worker nodes using the command shown in the following listing.
# Then interact with cluster using `kubectl` which issues REST request
# to the Kubernetes API Server running on the master node

# set default zone
# region europe-west3 has three different zones
# but here we set all nodes to use the same zone 'europe-west3-c'
gcloud config set compute/zone europe-west3-c

# create kubernetes cluster
# three worker nodes in the same zone
# thus --num-nodes indicates the number of nodes per zone
# if the regions contains three zones and you only want three nodes, set --num-nodes=1
gcloud container clusters create kiada --num-nodes 3

gcloud compute instances list

# scale the number of nodes
gcloud container clusters resize kiada --size 0

### 클러스터 정보 확인

# explore what's running inside a node; you can check logs inside it
gcloud compute ssh <node-name>

# check if cluster is up by listing cluster nodes
#   minikube : 1 node in the cluster
#   GKE : 3 nodes in the cluster



### METHOD 4. EKS
# Install `eksctl` command-line tool
# https://docs.aws.amazon.com/eks/latest/userguide/getting-started-eksctl.html

# creates a three-node cluster in the eu-central-1 region
eksctl create cluster --name kiada --region eu-central-1 --nodes 3 --ssh-access
```

- 쿠버네티스 툴 `kubectl`

```
kubectl get nodes
kubectl describe nodes
kubectl describe node <node-name>
```

- kubectl 이 특정 Kubernetes cluster를 사용 하도록 설정
  - point kubectl to it by setting the KUBECONFIG environment variable as follows:

```sh
cat ~/.kube/config
export KUBECONFIG=/path/to/custom/kubeconfig
```

- Kubernetes Dashboard

```sh
# https://github.com/kubernetes/dashboard
# add a configuration to a resource by file name or stdin
# k apply [options]

kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.5.1/aio/deploy/recommended.yaml

# This command runs a local proxy to the API server, allowing you to access the services through it. Let the proxy process run and use the browser to open the dashboard at the following URL:

kubectl proxy

# 대시보드 URL:
# http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/

# Token 생성 해야함:
# https://github.com/kubernetes/dashboard/blob/master/docs/user/access-control/creating-sample-user.md#getting-a-bearer-token

kubectl -n kubernetes-dashboard get secret $(kubectl -n kubernetes-dashboard get sa/admin-user -o jsonpath="{.secrets[0].name}") -o go-template="{{.data.token | base64decode}}"


kubectl -n kubernetes-dashboard describe secret $(kubectl -n kubernetes-dashboard get secret | sls admin-user | ForEach-Object { $_ -Split '\s+' } | Select -First 1)


# Using Helm to install dashboard?  Install helm :
# https://medium.com/@munza/local-kubernetes-with-kind-helm-dashboard-41152e4b3b3d
curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
chmod 700 get_helm.sh
./get_helm.sh

helm repo add kubernetes-dashboard https://kubernetes.github.io/dashboard/
helm install dashboard kubernetes-dashboard/kubernetes-dashboard -n kubernetes-dashboard --create-namespace

```


- 클러스터에 app 배포하기 (Deployment)
  - 보통은 yaml, json 파일에 Object 정의 하여 배포
    - (처음에는 일단 manifest 파일없이 Deployment 만들어보기)
  - `kubectl create` -> Kubernetes API에 전달되어 -> Deployment 생성 -> Pod 오브젝트 생성
    - 워커 노드에 assigned/scheduled 됨
    - 워커 노드의 Kublet (쿠버네티스 Agent)는 Pod가 새로 해당 node에 스케쥴된것을 확인
    - 도커 registry에서 이미지를 받아, 컨테이너로 실행
    - Deployment 오브젝트는 클러스터에 배포된 어플리케이션을 의미
    - 기존에 이미지로 만들어 컨테이너에 배포했던 kiada를 쿠버네티스 클러스터에 적용해보기
  - pod 생성
    - 쿠버네티스는 컨테이너 단위가 아닌 Pods 단위 관리
    - Pod: multiple co-located 컨테이너; 컨테이너 그룹
    - 같은 worker노드와 리눅스 namespace위에서 실행되는 밀접하게 연결된 컨테이너 그룹
    - 각각의 Pod은 separate logical machine with same IP, hostname, process, and so on.
    - 하나의 어플리케이션을 실행 함
  - service 생성
    - 각 pods들은 IP가 할당되지만 클러스터로의 internal ip라서 외부에서 접근 불가능
    - 서비스 오브젝트를 생성하여 external ip 할당하여 외부에 노출
    - LoadBalancer 타입의 서비스를 생성-> external load balancer를 생성되고
    - 이 로드밸런서의 퍼블릭 IP를 통해 pod에 접근 가능
    - pod는 ephemeral함. pod가 실패하거나 다운되고 다시 생성되면 ip변동-> service를 통해 static ip 할당

```sh
### POD 생성

# OLD VERSION :
# --generator=run/v1 옵션은 Deployment대신 ReplicationController를 create
# kubectl run kubia --image=jnuho/kubia --port=8080

# kubectl로 dokcerhub 이미지를 쿠버네티스 클러스터에서 실행
# create all the necessary components without having to deal with JSON or YAML
# Create a deployment object, called 'kiada'
# deployment to use the container image jnuho/kiada:0.1
# the Deployment object is stored in Kubernetes API
# now the jnuho/kiada container is run in the Cluster

kubectl create deployment kiada --image=jnuho/kiada:0.1
  deployment.apps/kiada created

# 생성된 Deployment 리스트 : NOT READY (컨테이너가 준비되지 않음)
# 다만 컨테이너 조회 하는 커멘드는 없음 e.g. k get containers 는 없음
# 컨테이너가 아닌 PODS가 쿠버네티스에서 가장 작은 단위임!
# POD 내의 컨테이너는 네트워크와 UTS 네임스페이스를 공유
# 각 POD은 논리적인 단위의 컴퓨터로서 어플리케이션을 실행
# 하나의 노드에 여러개의 POD가 있더라도, 각 POD들은 서로 다른 POD들의 프로세스를 볼 수 없음
k get deployments


# 생성된 pods 정보 (Pending/ContainerCreating -> Running)
kubectl get pods
kubectl describe pods

# 생성된 pod 상세 정보에 worker노드에서 실행 중 확인
kubectl describe pod <pod-name>
  Name:         kiada-7fc9cfcf4b-t8fcg
  Namespace:    default
  Priority:     0
  Node:         kind-worker/172.18.0.2


### SERVICE 생성
# Exposing your application to the world

# Create Service Object
#   expose all pods that belong to the kiada Deployment as a new service.
#   pods to be accessed from outside the cluster via a load balancer.
#   application listens on port 8080, so you want to access it via that port.
#   로드밸런서 타입 서비스는 cluster내에만 expose 시키거나 퍼블릭IP로 외부로 expose 가능

kubectl expose deployment kiada --type=LoadBalancer --name kiada-http --port 8080

# load balancer가 생성되고 EXTERNAL-IP 할당됨
kubectl get services
kubectl get svc

# List all available resources - Kubernetes Objects
kubectl api-resources


# 32232 is the node port on the worker node
# that forwards connections to your service.
# 이 포트로 모든 워커 노드에서 서비스에서 접근가능
# (minikube, kind, 등등 클러스터 종류 상관 없이)

regardless of whether you’re using Minikube or any other Kubernetes cluster.
kubectl get svc kiada-http
  NAME         TYPE           CLUSTER-IP     EXTERNAL-IP   PORT(S)          AGE
  kiada-http   LoadBalancer   10.96.59.164   <pending>     8080:32232/TCP   83m
  kubernetes   ClusterIP      10.96.0.1      <none>        443/TCP          6h50m


# Kubernetes 클러스터가 클라우드 (AWS, Gcloud)에 deploy 된다면
# 해당 클라우드 인프라에 로드밸런서를 통해
# 클러스터로 오는 트래픽을 실행 중인 컨테이너로 forward 함
# 해당 인프라는 Kubernetes에게 로드밸런서의 Ip를 알려주고 서비스의 외부주소가 됨


# Accessing your application through Load Balancer
curl <EXTERNAL-IP>:8080
curl 104.155.74.57:8080


# Accessing your application without Load Balancer
#   not all kubernetes clusters provide Load Balancer
#   minikube shows where to access the services
#     prints the url of the service
minikube service kiada --url
curl <ip:port>


# you can access this service locally using the Kubectl proxy tool.
kubectl port-forward kiada-http 8080:3002

# Deploying your application REQUIRES only two steps
# 'Deployment' is a representation of an application
# 'Service' exposes that deployment

kubectl create deployment [options]
kubectl expose deployment [options]

```


- Horizontally Scaling the application
  - increase (scale-out) / decrease the number of running application instances
  - scale the deployment object



```sh
# the result of 'Scale Out'
k get deploy
k get deployments.apps
k get deployment

  NAME    READY   UP-TO-DATE   AVAILABLE   AGE
  kiada   3/3     3            3           31h

k get pods

  NAME                     READY   STATUS    RESTARTS       AGE
  kiada-7fc9cfcf4b-64qvl   1/1     Running   0              3m3s
  kiada-7fc9cfcf4b-mf8ls   1/1     Running   0              3m3s
  kiada-7fc9cfcf4b-t8fcg   1/1     Running   1 (148m ago)   31h
```


- Displaying the pods' host node when listing pods
  - To see which nodes the pods were scheduled to,
  - you can use the -o wide option to display a more detailed pod list

- Object created
  - Deployment object you created,
  - Pod objects that were automatically created based on the Deployment
  - Service object you created manually
    - Instead of connecting directly to the pod,
    - clients should connect to the IP of the service.
    - This ensures that their connections are always routed to a healthy pod,
    - Each service provides internal load balancing in the cluster,
    - but if you set the type of service to LoadBalancer,
    - Kubernetes will ask the cloud infrastructure it runs in for an additional
    - load balancer to make your application available at a publicly accessible address


![image](https://drek4537l1klr.cloudfront.net/luksa3/v-12/Figures/03image014.png)


```sh
k get pods -o wide

# observing requests hitting all three pods when using the service
k get svc
  NAME         TYPE           CLUSTER-IP     EXTERNAL-IP   PORT(S)          AGE
  kiada-http   LoadBalancer   10.96.59.164   <pending>     8080:32232/TCP   43h
  kubernetes   ClusterIP      10.96.0.1      <none>        443/TCP          2d

k describe pods | grep Node:
Node:         kind-worker2/172.18.0.3
Node:         kind-worker/172.18.0.4
Node:         kind-worker/172.18.0.4

# Each request arrives at a different pod in random order.
# This is what services in Kubernetes do when more than one pod instance is behind them.
# They act as load balancers in front of the pods.

curl 172.18.0.3:32232
curl 172.18.0.4:32232

```



- Kubernetes API

- Understanding the structure of an object manifest
  - manifest of most Kubernetes API objects consists of the following four sections:
  - Type Metadata : type of a resource
  - Object Metadata : name and identifying information
  - Spec Metadata :  desired state
  - Status Metadata : current actual state

- Object Controller controls each object type
  - controllr monitors the desired state (Spec) constantly
  - controller updates the current status (Status) constantly

```sh

# get info about a node in a yaml format
kubectl get node <node-name> -o yaml


# one can access the API directly through proxy using plain HTTP

kubectl proxy


kubectl explain node.spec
kubectl get node kind-control-plane -o yaml

```






