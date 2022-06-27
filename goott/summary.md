
# 교육 내용
- Containr 이해와 Kubernetes 구축
  - Why Docker & Kubernetes
  - Docker review: 다양한 Container Service 구성
  - Docker review: IaC를 위한 Dockerfile & docker-compose
  - Kubernetes cluster 구축1: Host(CentOS7) & Dashboard 구성
- K8s 아키텍쳐
  - k8s cluster 구축2: GCP & Dashboard 구성

# 실습 진행

- CentOS 베어메탈 서버 기반 K8S -> Master, Worker node x2
- 클라우드 기반의 GCP -> GKE -> Worker node * 3

### Why Docker?

- container 근간 : 이미지 hub.docker.com
- 가상화 보다 훨씬 가벼운 경량화 기술

- 모니터링 툴 : prometheous, grafana 또는 EFK

- Dockerfile -> docker build : image생성 -> private registry 등록/업로드 또는 github
  코드, 구성, 종속성(환경 일관성), 런타임엔진


### Why Kubernetes

- 컨테이너화(pod)된 워크로드와 서비스를 관리하기 위한 이식성이 있고, 확장가능한 오픈소스 플랫폼
  - (pod <- container <- application)
  - 쿠버네티스 플랫폼을 사용: GKE(fluentd 저장서버), EKS, AKS

```
# 마스터노드 + 워커노드 w1, w2, w3

<표준화된 모델>

---
Request -> HAProxy ->
  Master, w1, w2, w3
  Master, w1, w2, w3
  Master, w1, w2, w3

etcd server (저장 서버) 노드들의 DB정보
---
```

- CNCF(Cloud Native Computing Foundation) www.cncf.io


- 선언적구성 -> YAML code -> object 정의, 활동 <-> 실행적 구성
- 자동회- > CI/CD, helm, chart...
- json vs yaml(상위언어, # 주석가능!)
  - 들여쓰기 검증
  - http://www.yamllint.com/
  - https://codebeautify.org/yaml-validator
  - https://onlineyamltools.com/validate-yaml


```sh
kubectl apply -f pod.yaml
```

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: webapp-ingress
```

- Kubernetes는 자동 restart 디폴트
- 제한없이 컨테이너에 핵심 기능을 제공하여 인프라 종속을 제거하고,
- pod(컨테이너를 담는 그릇) 및 service(load balancer, 포트)를 포함한
  kubernetes 플랫폼 내의 기능 조합을 통해 이를 달성
  컨테이너 안에 application
- 모듈화를 통한 더 나은 관리
service가 pod여러개에 분산(load balancing)
kubernetes service는 유사한 기능을 수행하는 pod모음을 함께 그룹화하는데 사용
검색 및 관찰 (service discovery: DNS)가능성, 수평적 확장
및 로드밸런싱을 위해 서비스를 쉽게 구성

- k8s목적: "Desired State Management"
  사용자 정의 구성 yaml코드에 맞춰 사용자가 기대하는 상태로 동작하도록 안정적으로 유지해주는 것
  (auto scaling, 리소스분배, 자동화, 자동복구, 동일한 환경구성, 안정적 서비스)


- 자격증 : CKA, CKS(시큐리티), 개발분야

https://www.docker.com/play-with-docker/
> Lab Environment > Get started
docker pull nginx:1.21-alpine
docker run -d -p 8001:80 --name=webserver nginx:1.21-alpine
curl localhost:8001
> Open port (8001)

```sh

### 1
# --publish
# -p 호스트포트:컨테이너포트
docker run -d -it -p 9000:8000 --name=nodejs_test node:16.6-slim
docker inspect nodejs_test | grep IPAddress
docker cp app.js  nodejs_test:/app.js
docker exec -it nodejs_test ls
docker exec -it nodejs_test node app.js # node service start


docker build -t mynode:1.0 .
docker images
docker run -d -it -p 9001:8000 --name=nodejs_test2 mynode:1.0
docker login
docker image tag mynode:1.0 jnuho/mynode:1.0
docker push jnuho/mynode:1.0
docker image rm jnuho/mynode:1.0
docker pull jnuho/mynode:1.0

---> 윈도우 크롬에서 PWD주소 http://ip172--...



### 2
git clone https://github.com/brayanlee/docker-phpserver.git
cd docker-phpserver/
ls
cat index.php3
vi Dockerfile

docker build -t phpserver:1.0 .
docker images | grep php
docker run -it -d -p 8004:80 -h phpserver --name=phpserver phpserver:1.0
curl localhost:8004

# upload
docker login
docker image tag phpserver:1.0 jnuho/phpserver:3.0
docker push jnuho/phpserver:3.0
docker image rm jnuho/phpserver:3.0
docker pull jnuho/phpserver:3.0



### 3

# 빌드속도 향상 : docker buildkit

# 작은 MSA 구조:
mkdir py_flask && cd $_
vi Dockerfile
mkdir app && cd $_
vi py_app.py
vi requirements.txt
Flast==2.1.0
cd ..
vi .dockerignore
Dockerfile
tree
DOCKER_BUILDKIT=1 docker build -t py_flask:1.0 .
docker run -it -p 9000:9000 -d py_flask:1.0

# upload
docker login
docker image tag py_flask:1.0 jnuho/py_flask:1.0
docker push jnuho/py_flask:1.0
docker image rm jnuho/py_flask:1.0
docker pull jnuho/py_flask:1.0

```


k8s-master   -> Control plane -> 운영, 관리
                ㄴ api-server,etcd,scheduler,controller,...
                    ㄴ Pod임
                  -> kubeadm init 작업 ->  주요 구성요소
            Q) Master node에는 application이 포함한 Pod가 실행되지 못한다? X 그렇지 않다!
            A) 실행될 수 는 있지만, 막아놓음! ("taint"->NoSchedule) <-> toleration
                master/worker노드 둘다 container 실행가능
ㄴ worker node  -> Data plane -> application 포함한 container -> Pod실행
                                 ㄴ container runtime
                  -> kubeadm join : 마스터노드에연결 -> worker node 신규 및 확장 지원
                              <-> drain
                  -> kubeadm upgrade : worker node3 추가 시 사용

VM 최소 CPU-2core, Memory-4G 이상
  새로만들기 -> 
  - 오디오 USB 체크제거
  - 네트워크 카드 2개 :
    - Adaptor1 enp0s3 -> NAT       -> external(외부) -> 10.2.15.0
    - Adaptor2 enp0s8 -> Host-Only -> internal(infra) 고정 -> 192.168.56.xxx
      (2번체크 호스트전용: 윈도우에 있음)


1. VirtualBox 구성
  메모리 2048
  프로세스 2
2. Master CM 구성
3. Master Node 구성
4. Worker Node 복제
5. Worker Node 연결
6. Plugin 추가: 
  네트워크 plugin Calico, BGP 프로토콜 (OSI7계층 증 3계층 서비스)
    3계층 대표 프로토콜 IP
    인증서-> 윈도우 클라이언트에 심어줌
    Cognito 창 열어서


- 로그인 정보 : root/k8spass#
```sh
hostname
ifconfig
Ethernet(enp0s8) -> wired setting
host only
IPV4 >
Address: 
  192.168.56.100
  255.255.255.0
  192.168.56.2
DNS
  8.8.8.8
ping 192.168.56.100

로컬 푸티
  192.168.56.100
  root/k8spass#
```
  
2) 방화벽 설정 보안정책 낮춤
k8s 서비스가 많은데 전부 방화벽에 등록하는건 불가능: 일단 보안 낮추고 나중에 설정
setenforce 0
vi /etc/selinux/config
  # SELINUX=enforcing
  # 강제하지 않는 대신에 warning만 출력
  SELINUX=permissive

sestatus
systemctl stop firewalld && systemctl disable firewalld
systemctl list-unit-files | grep firewall

컨테이너 실체 -> process -> memory에서 작업 수행
k8s: SWAP에 가지 않도록 함 : swap device 해제 후 k8s 설치
vi /etc/fstab
  # /dev/... swap
swapoff -a

- 네트워크 플러그인 calico 사용 -> BGP 프로토콜 사용 -> L3 -> IP 프로토콜사용
  -> packet의 올바른 포워딩 설정 요구
- 클러스터에 네트워크 트래픽 전송
- 노드간 트래픽을 정상적으로 잘 전달하기 위한 ip forwarding 작업

# 1로 되어 있어야함
cat /proc/sys/net/ipv4/ip_forward
# echo '0' > /proc/sys/net/ipv4/ip_forward
# 0인경우 ip클러스터상에 문제가 될 수 있음
# ERROR FileContent--proc-sys-net-.. ... not set to 1
# 복제시에 이런 문제 발생 가능
# 해결책! :
echo 1 > /proc/sys/net/ipv4/ip_forward

--iptables rule 설정 하여 트래픽 전달
Q) docker 기본 네트워크는? bridge (L2) (OSI 네트워크 2계층)

modprobe br_netfilter
vi /etc/modules-load.d/k8.conf
  br_netfilter

# ipv4,ipv8 브릿지 네트워크 허용
vi /etc/sysctl.d/k8s.conf

  net.bridge.bridge-nf-call-ip6tables = 1
  net.bridge.bridge-nf-call-iptables = 1

sysctl --system


3) docker 구성

- 최소단위 POD : 컨테이너 돌림

yum -y install yum-utils device-mapper-persistent-data lvm2
cd /etc/yum.repos.d/

yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
cat docker-ce.repo
yum list | grep docker-ce

# 요청 docker 버전 설치
# docker-ce 버전 전부 조회
yum list docker-ce --showduplicates
yum -y install docker-ce
  docker-ce
  containerd.io
  fuse
  container-selinux
  ... 등 설치 작업

도커 cgroup 드라이버: croupfs(기본값)로 하면 안됨!
  -> K8S 쓰기위해서 다른 드라이버로 변경
docker version
docker info

# 드라이버 cgroupfs -> systemd로 변경하는 작업
mkdir /etc/docker
vi /etc/docker/daemon.json

{
  "exec-opts": ["native.cgroupdriver=systemd"],
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "100m"
  },
  "storage-driver": "overlay2",
  "storage-opts": [
    "overlay2.override_kernel_check=true"
  ]
}

# 서비스등록
mkdir -p /etc/systemd/system/docker.service.d
systemctl daemon-reload
systemctl enable --now docker
systemctl start docker
systemctl status docker

docker info | grep -i cgroup
docker pull nginx:1.21-alpine

docker login
docker info | grep Username

docker run -d 8001:80 --name=myweb nginx:1.21-alpine
  Message fron syslogd@k8s-master at .,.,
  Kernal NMI watchdog: BUG: soft lockup CPU#1 stuck for 22s [docker-app:2760]

vi /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=https://packages.cloud.google.com/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=0
repo_gpgcheck=0
gpgkey=https://packages.cloud.google.com/yum/doc/yum-key.gpg https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg

yum repolist
yum -y update


# 최신버전 1.23.8
# 1.23.5에 맞춰서 설치
yum list | grep kubernetes
yum list kubeadm --showduplicates | grep 1.23

yum -y install kubeadm-1.23.5 kubectl-1.23.5 kubelet-1.23.5 --disableexcludes=kubernetes
yum list installed | grep kubernetes

# kubelet은 데몬 (리부팅되도 항상 켜있어야함), kubeadm, kubectl은 명령어
systemctl daemon-reload
systemctl enable --now kubelet

vi /etc/hosts
  192.168.56.100 k8s-master
  192.168.56.101 k8s-node1
  192.168.56.102 k8s-node2
  192.168.56.103 k8s-node3

shutdown -h now

VM > kmaster노드 > 복제
  이름 k8s-node1, k8s-node2
  경로 H:/k8s/k8s-node1
  MAX주소정책 : 모든 네트워크 어배터의 새 MAC 주소 생성
  다음 > "완전한 복제"
각각 os 접속해서 hostname, ip변경하기

# node1 booting > 제어판 enp0s8 > 192.168.56.100->101
# 작업후putty 연결
hostnamectl set-hostname k8s-node1
ifconfig
ping 192.168.56.1
<!-- reboot -->
shutdown -h now

# node2 booting > 제어판 enp0s8 > 192.168.56.100->102 설정후 커넥션 on/off
# 작업후putty 연결
hostnamectl set-hostname k8s-node2
ifconfig
ping 192.168.56.1
<!-- reboot -->
shutdown -h now

systemctl restart network

# node1,2 작업후 master노드 켜기 (putty 3개)



- kubeadm init, kubeadm join 실습

네이버카페 Kubernetes 게시판 > installation


-- cluster 노드 3개
  ㄴ 시간동기화
  ㄴ 1회 이상 상호 SSH 접근 작업

# ntpd 올려놓아야함
rdate -s time.bora.net
date

# 각각 노드(master,node,2)에서 known_hosts등록 작업
cat /etc/hosts

ssh k8s-master
  yes
ssh k8s-node1
  yes
ssh k8s-node2
  yes


ssh root@192.168.56.101


# 밑에 파일에 node1,2 known 호스트 등록 되어야함
cat .ssh/known_hosts



### kubeadm init 작업

- 3대의 노드를 묶는작업!

- Master노드를 Control Plane으로 등록

--pod-network-cide=IP/bit
  3개를 네트워크로 묶을거기 때문에
  Pod를 어디에 배포하든
  그 Pod에 접근 가능
--service-cidr=IP/bit
  Cluster는 네트워크
  K8s는 클라우드 네이티브
  IP대역을 잡는 방법
NAT range

# 버전 체크해서 k8s가 업데이트 되버린걸 확인!
yum list installed | grep kubernetes
yum -y remove kubeadm kubectl kubelet --disableexcludes=kubenetes
yum -y install kubeadm-1.23.5 kubectl-1.23.5 kubelet-1.23.5 --disableexcludes=kubenetes
history


# 마스터노드에서 작업
#   apiserver-advertise-address=마스터노드 ip
kubeadm init --pod-network-cidr=10.96.0.0/12 --apiserver-advertise-address=192.168.56.100
  조인키 복사해서 -> node1,2에 join 명령어 실행
  => Run 'kubectl get nodes'! 메시지 확인.

- 로그: 

```sh
[root@k8s-master ~]# kubeadm init --pod-network-cidr=10.96.0.0/12 --apiserver-advertise-address=192.168.56.100
I0625 09:28:34.945745   12406 version.go:255] remote version is much newer: v1.24.2; falling back to: stable-1.23
[init] Using Kubernetes version: v1.23.8
[preflight] Running pre-flight checks
[preflight] Pulling images required for setting up a Kubernetes cluster
[preflight] This might take a minute or two, depending on the speed of your internet connection
[preflight] You can also perform this action in beforehand using 'kubeadm config images pull'



[certs] Using certificateDir folder "/etc/kubernetes/pki"
[certs] Generating "ca" certificate and key
[certs] Generating "apiserver" certificate and key
[certs] apiserver serving cert is signed for DNS names [k8s-master kubernetes kubernetes.default kubernetes.default.svc kubernetes.default.svc.cluster.local] and IPs [10.96.0.1 192.168.56.100]
[certs] Generating "apiserver-kubelet-client" certificate and key
[certs] Generating "front-proxy-ca" certificate and key
[certs] Generating "front-proxy-client" certificate and key
[certs] Generating "etcd/ca" certificate and key
[certs] Generating "etcd/server" certificate and key
[certs] etcd/server serving cert is signed for DNS names [k8s-master localhost] and IPs [192.168.56.100 127.0.0.1 ::1]
[certs] Generating "etcd/peer" certificate and key
[certs] etcd/peer serving cert is signed for DNS names [k8s-master localhost] and IPs [192.168.56.100 127.0.0.1 ::1]
[certs] Generating "etcd/healthcheck-client" certificate and key
[certs] Generating "apiserver-etcd-client" certificate and key
[certs] Generating "sa" key and public key
[kubeconfig] Using kubeconfig folder "/etc/kubernetes"
[kubeconfig] Writing "admin.conf" kubeconfig file
[kubeconfig] Writing "kubelet.conf" kubeconfig file
[kubeconfig] Writing "controller-manager.conf" kubeconfig file
[kubeconfig] Writing "scheduler.conf" kubeconfig file
[kubelet-start] Writing kubelet environment file with flags to file "/var/lib/kubelet/kubeadm-flags.env"
[kubelet-start] Writing kubelet configuration to file "/var/lib/kubelet/config.yaml"
[kubelet-start] Starting the kubelet
[control-plane] Using manifest folder "/etc/kubernetes/manifests"
[control-plane] Creating static Pod manifest for "kube-apiserver"
[control-plane] Creating static Pod manifest for "kube-controller-manager"
[control-plane] Creating static Pod manifest for "kube-scheduler"
[etcd] Creating static Pod manifest for local etcd in "/etc/kubernetes/manifests"
[wait-control-plane] Waiting for the kubelet to boot up the control plane as static Pods from directory "/etc/kubernetes/manifests". This can take up to 4m0s
[kubelet-check] Initial timeout of 40s passed.
[apiclient] All control plane components are healthy after 84.017203 seconds
[upload-config] Storing the configuration used in ConfigMap "kubeadm-config" in the "kube-system" Namespace
[kubelet] Creating a ConfigMap "kubelet-config-1.23" in namespace kube-system with the configuration for the kubelets in the cluster
NOTE: The "kubelet-config-1.23" naming of the kubelet ConfigMap is deprecated. Once the UnversionedKubeletConfigMap feature gate graduates to Beta the default name will become just "kubelet-config". Kubeadm upgrade will handle this transition transparently.
[upload-certs] Skipping phase. Please see --upload-certs
[mark-control-plane] Marking the node k8s-master as control-plane by adding the labels: [node-role.kubernetes.io/master(deprecated) node-role.kubernetes.io/control-plane node.kubernetes.io/exclude-from-external-load-balancers]
[mark-control-plane] Marking the node k8s-master as control-plane by adding the taints [node-role.kubernetes.io/master:NoSchedule]
[bootstrap-token] Using token: 49f5gs.7v4rquxe006nxsw5
[bootstrap-token] Configuring bootstrap tokens, cluster-info ConfigMap, RBAC Roles
[bootstrap-token] configured RBAC rules to allow Node Bootstrap tokens to get nodes
[bootstrap-token] configured RBAC rules to allow Node Bootstrap tokens to post CSRs in order for nodes to get long term certificate credentials
[bootstrap-token] configured RBAC rules to allow the csrapprover controller automatically approve CSRs from a Node Bootstrap Token
[bootstrap-token] configured RBAC rules to allow certificate rotation for all node client certificates in the cluster
[bootstrap-token] Creating the "cluster-info" ConfigMap in the "kube-public" namespace
[kubelet-finalize] Updating "/etc/kubernetes/kubelet.conf" to point to a rotatable kubelet client certificate and key
[addons] Applied essential addon: CoreDNS
[addons] Applied essential addon: kube-proxy

Your Kubernetes control-plane has initialized successfully!

To start using your cluster, you need to run the following as a regular user:

  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

Alternatively, if you are the root user, you can run:

  export KUBECONFIG=/etc/kubernetes/admin.conf

You should now deploy a pod network to the cluster.
Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
  https://kubernetes.io/docs/concepts/cluster-administration/addons/

Then you can join any number of worker nodes by running the following on each as root:

kubeadm join 192.168.56.100:6443 --token 49f5gs.7v4rquxe006nxsw5 \
        --discovery-token-ca-cert-hash sha256:5d5781cbacaac33bf0de53d107d62a8e8959b31799bbf1f4f3c7fffea42e1bec
```

동일 네트워크 클러스터에 연결된 노드간에 존재하는 Pod는 모든 노드에서 조회가 가능하다
-> Pod를 생성
-> Pod IP 할당

# 1.23.5로 진행하기로 했는데 yum update한후에 1.24.2로 버전업됬는데 그냥 그대로 진행하는건가요?

kube-apiserver
kube-controller-manager
kube-scheduler

... control-plnae has initialized successfully!
To start using your cluster, you need to run the following as a regular user:
Alterna.. if root userA;


mkdir -p $HOME/.kube
cd $HOME/.kube
cp -i /etc/kub
rnetes/admin.conf $HOME/.kube/config
chown $(id -u):$(id -g) $HOME/.kube/config
export KUBECONFIG=/etc/kubernetes/admin.conf

source <(kubectl completion bash)
echo "source <(kubectl completion bash)" >> .bashrc

-> 나머지 2개 노드에도 같은 작업 반복

네트워크 플러그인 : kubernetes.io/docs/concepts/cluster-administration/addons/

# 마스터노드
kubectl get node


# pending -> running 바뀌는거는 캘리코 설치 후
kubectl get pod --all-namespaces

netstat -nltp | grep LISTEN

nc 127.0.0.1 6443

# 조인키 카피해서  각 노드1,2에 붙여넣기: 작업 후 
kubectl get pod --all-namespaces
  kube-system coredns-... Pending!!
  => calico 설치!

# 캘리코 돌리기! 마스터노드:
kubectl apply -f https://docs.projectcalico.org/manifests/calico.yaml

# 또는
curl -O https://docs.projectcalico.org/manifests/calico.yaml
kubectl apply -f  calico.yaml


 

# 캘리코 Running 확인하기!
kubectl get pod --all-namespaces

watch kubectl get pods --all-namespaces

# 메이저 버전 확인
kubeadm version
kubectl version
kubelet --version

# 이제 Dashboard 올려야함!
# v2.5.0 버전 체크: 1.23 체크되어 있음!
# https://github.com/kubernetes/dashboard/releases
# 마스터노드에서 실행:
kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.5.0/aio/deploy/recommended.yaml
  namespace/kubernetes-dashboard created
  serviceaccount/kubernetes-dashboard created
  service/kubernetes-dashboard created

  I0625 10:26:43.091782    8154 request.go:665] Waited for 1.085917571s due to client-side throttling, not priority and fairness, request: GET:https://192.168.56.100:6443/apis/rbac.authorization.k8s.io/v1/namespaces/kubernetes-dashboard/roles/kubernetes-dashboard


kubectl get pod --all-namespaces
kubectl cluster-info
  Kubernetes control plane is running at ...
  CoreDNS is running at ...


# dashboard 기동 방법
  ㄴ proxy        -> 8001 port -> 보안성 X -> 브라우저 접근은 인증서가 필요
      클라우드에서 쓰는 기법 (proxy)
  ㄴ NodePort
  ㄴ apiserver    -> 6443 port -> 접근 시 -> token -> openssl로 인증서 생성 -> 윈도우에 배포 (certmgr.msc > 인증서)
      온프레미스 시 보안을 위한 인증서 추가 설치 (클라우드에서는 안씀)

vi .bashrc
  alias pall=='kubectl get pod --all-namespaces'
. .bashrc

# Pod안에 컨테이너 리스트
# Pod 장애가 생기면 RESTARTS 개수 늘어남!
pall

kubectl get pod
# 디폴트 네임스페이스 이외 조회
# -n {네임스페이스 이름}
kubectl get pod -n kubernetes-dashboard

kubectl get namespaces
kubectl get ns


# kubernetes security -> RBAC(Role), ABAC(Attribute)

kubectl get no

dashboard: 쿠버네티스가 사용되는 리소스를 스크래핑 dashboard-metrics-scraper
하려면 dashboard에 cluster admin 권한 줘야함

# 쿠버네티스의 기본적인 권한자
kubectl get cluster-admin
kubectl get clusterrole cluster-admin
kubectl describe clusterrole cluster-admin
  Verbs: 권한 이름 (create,select...)

# 서비스어카운트
# 대시보드의 이름이 sa
kubectl get sa
kubectl get sa -n kubernetes-dashboard kubernetes-dashboard

# cluster-admin과 sa(default)가 보유한 super 권한을 "kubernetes-dashboard"에 연결
  -> ClusterRoleBinding

kubectl describe sa -n kubernetes-dashboard kubernetes-dashboard
# 마스터노드:
mkdir dashboard_token && cd &_

# 또는 vi ClusterRoleBinding.yml
vi ClusterRoleBinding.yaml

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: kubernetes-dashboard2
  labels:
    k8s-app: kubernetes-dashboard
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
- kind: ServiceAccount
  name: kubernetes-dashboard
  namespace: kubernetes-dashboard

kubectl create -f ClusterRoleBinding.yaml
  kubectl {apply | create | delete | get | describe} [-f] *.yaml
    apply     -> object 생성, 생성 후 재적용(update) 가능 (일부 리소스에 한함)
    create    -> object 생성, 생성 후 재적용 불가
    delete    -> object 삭제
    get       -> object 조회
    describe  -> object 세부조회

# Dashboard 토큰값 입력해야함 (입장권)
# 매번할 때 마다 토큰 다름! 생성:
kubectl -n kubernetes-dashboard get secrets
kubectl -n kubernetes-dashboard describe secrets kubernetes-dashboard-token-7x2r4
  'token : ..."이 dashboard 입장권

grep 'client-certificate-data' ~/.kube/config | head -n 1 | aws '{print $2}' | base64 -d >> kubecfg.crt
grep 'client-key-data' ~/.kube/config | head -n 1 | aws '{print $2}' | base64 -d >> kubecfg.key
cd dashboard_token
mv ../kubecfg.crt .
mv ../kubecfg.key .
ls
  ClusterRoleBinding.yaml kubecfg.crt kubecfg.key

# .crt, .key 로 p12 생성
openssl pkcs12 -export -clcerts -inkey kubecfg.key -in kubecfg.crt -out kubecfg.12 --name "kubernetes-admin"
cp /etc/kubernetes/pki/ca.crt .
ls
  ca.crt ClusterRoleBinding.yaml kubecfg.crt kubecfg.key kubecfg.p12
--windows로이전 (winscp) -> certmg.msc -> 여기에 등록 -> powershell(관리자로)
scp로 위 값을 H/k8s/dashboard_token 윈도우에 생성
파워쉘 관리자로 실행
cd H/k8s/dashboard_token

certutil.exe --addstore "Root" H:\k8s\dashboard_token\ca.crt
certutil.exe -p k8spass# -user -importPFX H:\k8s\dashboard_token\kubecfg.p12

window + r = 실행창 -> certmgr.msc
  (certmgr.msc > 인증서로 생성 확인)

크롬 cofnito브라우저 :
https://192.168.56.100:6443/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/#/login
  select certificate 창 떠야함
  토큰입력 > Cluster > Nodes에 master+node1,2보여야함



GCP기반 GKE

실제 kubectl 명령어 등 작업 어떻게 수행?
Kubernetes Engine올리고,
Google Cloud SDK Shell 설치하여 인증후 접속!

- GCP 구성
- Google SDK Shell 설치
- 인증
- GKE Dashboard -> proxy -> 8001 port -> token (입장권/ 로그인 암호를 대체)
  클라우드레벨 보안 제공되기 때문에 별도 인증서 없이 프록시로만 접근




프로젝트 10자이자
k8s-jnuho
apiservice -> library -> kubernetes검색
Kubernetes engine api > 사용 버튼 누르면 -> 관리로 바뀜

사이드바 Kubernetes Engine
-> Cluster -> GKE Standard 구성
  클러스터이름 cluster-k8s-jnuho
  영역 asia-northeast3-a
  버전 1.22.8-gke.202 (기본값)
-> 노드 풀 
  노드 풀
    k8s-pool 노드 수 3개
  노드
    이미지유형 containerd를 포함한 (기본값)
    머신구성 (워커노드 3대)
      일반용도 시리즈 E2, e2-medium(vCPU2개, 4GB메모리)
      CPU,GPU: SSD, 100GB

SHell> gcloud init
Default compute Region and Zone? Y
  [50] asia-northeast3-a

gcloud components update
gcloud components install kubectl

대시보드 > 클러스터 > 연결 명령어
gcloud container clusters get-credentials k8s-junho --zone asia-northeast3-a --project k8s-jnuho
gcloud container clusters list
kubectl get node
  3개 노드 확인
kubectl get pod --all-namespaces
  fluentd도 설치 확인

22.8버전에 맞는 대시보드 버전
  https://github.com/kubernetes/dashboard/releases

kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.4.0/aio/deploy/recommended.yaml

GKE dashboard 데몬 띄워놓기 proxy 8001 > 별도 터미널로
kubectl proxy

크롬 cognito 브라우저 >
  http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/

토큰 필요함 > RBAG작업! > H드라이버 이동
h:k8s\dashboard_token
dashboard_token
dir
  ClusterRoleBinding.yaml
kubectl apply -f ClusterRoleBinding.yaml
  또는
kubectl create -f ClusterRoleBinding.yaml

kubectl -n kubernetes-dashboard get secret
kubectl -n kubernetes-dashboard describe secret kubernetes-dashboard-token-84f5n


H:\k8s\dashboard_token>kubectl -n kubernetes-dashboard describe secret kubernetes-dashboard-token-84f5n
Name:         kubernetes-dashboard-token-84f5n
Namespace:    kubernetes-dashboard
Labels:       <none>
Annotations:  kubernetes.io/service-account.name: kubernetes-dashboard
              kubernetes.io/service-account.uid: 6b24f495-11a5-4ab9-ae0a-5870aa8c763f

Type:  kubernetes.io/service-account-token

Data
====
token:      eyJhbGciOiJSUzI1NiIsImtpZCI6InU1WjB6UmN5TlZ5MWU5cEhfa0ctUEJ3OXh4N0ZBaDduVURod2FsbVViOHcifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJrdWJlcm5ldGVzLWRhc2hib2FyZCIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJrdWJlcm5ldGVzLWRhc2hib2FyZC10b2tlbi04NGY1biIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50Lm5hbWUiOiJrdWJlcm5ldGVzLWRhc2hib2FyZCIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50LnVpZCI6IjZiMjRmNDk1LTExYTUtNGFiOS1hZTBhLTU4NzBhYThjNzYzZiIsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDprdWJlcm5ldGVzLWRhc2hib2FyZDprdWJlcm5ldGVzLWRhc2hib2FyZCJ9.csp2bL-ukXJh1peI3_hz6NAV20BiN9GG_ylM9ufohLEaUq3WcLqgpo4s7eIYgcuGIkDeR0MoJILgCnJ4AYU1aRGGSxmdO96akHb2NBo0GUsV30TJ2BYZ-35O0h8Zwhh64yC0jA3pnE3vN6NPBtIWIvM-5Q4WouUlo665Zk_uvkMR-yrYvze37RtI51E-wRC7FE_U_R81wDXVCsWy03XjLsRuoirOS1m5AamrsGM5Jpr9zM39fVowPhKFY0nMsV9YosdN5tF1UC2oPuOLWEJp1yQyKzWFx6-1cGZOpF5pTlsQm_nmVv2rvIgWaL4DHqhKE-D2AxXEhvhZZgAoETqRfg
ca.crt:     1509 bytes
namespace:  20 bytes

대시보드 만료되면 신규토큰 발급
  https://cafe.naver.com/ocmkorea/19052

kubectl -n kubernetes-dashboard edit deployments kubernetes-dashboard
  41 --token-ttl=0
kubectl get pod --all-namespaces




CENTOS 마스터노드 (VM):
object 이름 조회, alias, KIND 등
두개이상 단어조합시 CamelCase로 표현

kubectl api-resources
kubectl api-resources | grep deploy


yaml vs. kubectl 명령어
명령어들을 yaml로 작성하여 간편화
콜론: 스페이스에러 조심

mkdir LABs && cd $_

# deployment
kubectl api-resources | grep -i {대소문자무관 리소스명}
k get componentsstatuses

# 노드 내렸다가 올리기
https://discuss.kubernetes.io/t/how-to-drain-a-node-and-restart-it/16331/2

kubectl drain gke-k8s-junho-k8s-pool-92768e66-9s5n --ignore-daemonsets
kubectl get node
  gke-k8s-junho-k8s-pool-92768e66-9s5n   Ready,SchedulingDisabled   <none>   57m   v1.22.8-gke.202
kubectl uncordon gke-k8s-junho-k8s-pool-92768e66-9s5n


설치 후 쿠버네티스 아키텍쳐 이해
https://www.eginnovations.com/documentation/Kubernetes/The-Kubernetes-Architecture.htm

  - 구성요소들 흐름
  - 노드-마스터 네트워크 정보 주고받는 방법
  - kubernetes architecture


# Master node (=Control plane)

# kubectl get pod --all-namespaces
pall

# yaml 4개 (Pod 껍데기 안에 container 스펙이 있음)
ls /etc/kubernetes/manifests
etcd.yaml kube-apiserver.yaml kube-controller-manager.yaml kube-scheduler.yaml

  kube-apiserver-k8s-master (6443포트)
    - pod로 구성
    - 전체 클러스터에 대한 작업 정의, 배포 상태 확인 등의 역할
    - REST API 기법을 이용한 통신
  etcd-k8s-master
    - k8s의 DB 역할 -> key: value 구조로 저장 (<-> apiserver) -> B*Tree INDEX 구조 (root-branch-leaf)
    - snapshot save & restore로 관리 -> CKA
    - 단일 object ~ 전체 cluster의 상태 정보 표현, 암호화 기능 제공

  kube-scheduler-k8s-master
    - kubectl get node -o wide
    - kubectl describe node k8s-node1
    - kubectl describe nodes | grep -i taint
    - 알고리즘 : 각노드 상태값 확인, 배치할 노드 결정

  kube-controller-manager-k8s-master
    - Pod를 다루는, 사용하는 상위 object 관리
    - deployment, daemonset, cronjob, replication, replication controller, replicaset


- Kubernetes components(architecture, 동작방식) 흐름:

1. 사용자가 실행 kubectl apply -f pod.yaml
2. API 서버가 받기전에 TLS 암호화 수행 (security check)
  1. Authentication 계정인증
  2. Authorization 권한
  3. Admission controller 승인제어 시스템(1,2값을 넣어서 확인)
3. API Server에게 전달 (REST API)
4. API Server는 etcd 저장하고 다시 돌려받음!
  3-way handshake
  request하면 sync와 Acknoledge 받음
5. Scheduler에게 어떤 노드에 배포할지 (Bind Pod to Node) 확인후 받음
6. API Server는 etcd에 기록후 다시 회신 받음
7. kubelet(통신병역할?)에 yaml 전달-> node1이 받음
8. kube-proxy > Kubelet
9. kube-proxy > Container Runtime 실행
    Docker의 containerd or cri-o
    도커엔진 구조 => dockerd --> containerd --> runC(OCI)
  kube-proxy > OS커널에 씀
10. API Server에 Pod 생성을 보고함
11. 사용자 kubectl get pod로 생성 확인



kubectl apply -f mynode-pod.yaml
kubectl api-resources | grep -i po
  Pod
kubectl get po
  ... ContainerCreating
kubectl get po -wide
  IP: 10.111.156.67,  NODE: k8s-node1
docker images | grep mynode

curl 10.111.156.67:8000



docker-proxy -> docker run -p 8001:80 ~
  # publish -p

kube-proxy 모든 노드에 다 걸려있음
  kube-proxy 모든 노드에 다 걸려있음(노드개수마다 있음)
  kube-proxy ---> Pod --- Service

kubectl describe pod kube-proxy-44g4c -n kube-system

calico (OSI 3계층 서비스)들이 BIRD로 연결되어 있음
모든 노드는 BIRD(데몬, calico의 모듈)으로 연결되어 있음
Pod의 IP는 calico, coredns 도움으로 할당

Pod -> pod ip 보유, 내부 컨테이너들의 호스트와 같다.
  Pause container: 
    -> infra container
    -> 내부 컨테이너에 IP, Mac등의 네트워크를 부여
    -> 할당 이후 sleep
    docker ps | grep pause
  container1
  container2


# 로컬에서 접근 X, 클러스터에서만 접근하는 private ip
# 서비스 object를 통해 외부연결 활성화
curl 10.111.156.67:8000


kube-proxy ---> Pod --- Service
application - container - Pod <- Service <- 외부연결
                                    ㄴ endpoints=podIP
서비스와 pod 연결하기 위해서 .yaml의 "labels:" 지정
=> pod는 "labels:", service는 "selector:"



# service 만들기!

kubectl apply -f mynode_svc.yaml
kubectl get po,svc -o wide
curl 192.168.56.101:9000

kubectl describe svc mynode_svc


mkdir nginx-lab && cd $_

kubectl cp index.html k8s-nginx-pod:/usr/share/nginx/html/index.html


kubectl exec -it k8s-nginx-pod -- bash
cat /usr/share/nginx/html/index.html


## GKE에서 mynode 이미지 실습
# pod,svc만들기, External IP 로드밸런서 연결, 외부 접속확인
YAML:
svc의 spec > ports > port: 호스트포트
svc의 spec > ports > targetPort = pod의 containerPort 컨테이너포트
svc의 spec > type: LoadBalancer

kubectl apply -f mynode.yaml
kubectl get po,svc -o wide
curl 34.64.45.164:9000
kubectl delete -f mynode.yaml


## php서버 테스트
kubectl apply -f devphp.yaml
kubectl get pods
kubectl delete -f devphp.yaml

# mongodb 테스트 (클라우드 아닌 VM에서)
pod apply 한다음에 어떤 노드에 deploy된지 확인해서
svc.yaml ExternalIPs에 192.168.56.102
  externalIPs:
  - 192.168.56.102
mkdir mongodb && cd $_
kubectl apply -f mongodb-pod.yaml
kubectl apply -f mongodb-svc.yaml

대시보드 > POD >  mongodb
use k8sdb
db
db.k8scollection.insrt({name:"junho.lee", job:"SOFTWARE ENGINEER"})
노드2번에

netstat -nlp | grep 27017

예제 실행 후 pod, svc 삭제
kubectl get pods
kubectl get svc
kubectl get po,svc -o wide


## mysql 예제
mkdir mysql && cd $_
# 에러! MYSQL_ROOT_PASSWORD 필수값 넣어야함!
# dockerhub > mysql 설명
kubectl apply -f mysql_pod.yaml
kubectl logs pod/mysql-pod
vi mysql_pod.yaml
      env:
    - name: MYSQL_ROOT_PASSWORD
      value: k8spass#
vi mysql_svc.yaml
  ports:
  - port: 13306
    targetPort: 3306
  type: LoadBalancer

MYSQL 워크벤치에서 접근가능
  SSH host 192.168.56.102:22
  SSH user root
  MYSQL host 192.168.56.102:22
  MYSQL port 13306

대시보드 > POD >  mysql로 접근가능



## GO 언어 활용 Pod & Service 배포 테스트 => GKE에서 테스트!
로컬/CENTOS에서 :
git clone https://github.com/brayanlee/goapp.git
cd goapp
ls
  Dockerfile goapp.go
docker build -t goapp:1.0 .
docker images | grep goapp
# 해당 Dockerfile로 이미지 빌드 되는지 테스트!!!
docker run --name goapp-deploy -p 9090:9090 -d \
  -h goapp-container goapp:1.0
# 테스트 완료!
curl localhost:9090
  hostname: goapp-container
  IP: [172.17.0.13]
# 해당 이미지 dockerhub에 올리기!
docker image tag goapp:1.0 jnuho/goapp:1.0
docker login
docker push jnuho/goapp:1.0

# GKE에서 yaml로 pod, svc 생성하기!
vi go_pod.yaml
  apiVersion: v1
  kind: Pod
  metadata:
    name: goapp-pod
    labels:
      app: hi-goapp
  spec:
    containers:
    - name: goapp-container
      image: jnuho/goapp:1.0
      ports:
      - containerPort: 9090

vi go_svc.yaml
  apiVersion: v1
  kind: Service
  metadata:
    name: goapp-svc
  spec:
    selector:
      app: hi-goapp
    ports:
    - port: 9090
      targetPort: 9090
    type: LoadBalancer

# GKE에서 실행!
kubectl apply -f go_pod.yaml
kubectl apply -f go_svc.yaml

# 아무데서나 external-ip로 접근가능!
curl {SVC-EXTERNAL-IP}:9090

# 실습후 pod, svc 자원 삭제!
kubectl delete -f go_pod.yaml
kubectl delete -f go_svc.yaml



# POD 파드 (object = api resource의 하나)

- 작은무리 (컨테이너 모음)
- 하나의 파드안에 여러개 컨테이너
- k8s의 기본적인 application 배포단위 (container를 포함하는)
- 하나의 Pod는 하나 이상의 container를 포함할 수 있다.
  - Pod는 컨테이너의 호스트와 같음
  - 컨테이너1 포트 8080, 컨테이너2 포트 8000
  - 컨테이너 1,2 는  서로 통신할 수 있음 (localhost로 통신)
    localhost:8080, localhost:8000
    두개 컨테이너가 같은 포트 쓸수 없음!
    같은 호스트밑에서 두개 어플리케이션 포트 중복 불가!
- container in Pod 종류
    ㄴ 일반 application을 수행하는 container -> Runtime container
    ㄴ 조건 설정을 통해 main container(runtime)의 기동 시점을 정하는 container -> init container
    ㄴ Pod 내에서 특정 container의 로그수집, 모니터링 등의 목적으로 사용되는 container -> sidecar container
    * CKA: runtime/init/sidecar 컨테이너 구분 중요
- Pod의 상위 object (kubernetes object relationship) ->
  - Deployment, DaemonSet, CronJob(분 시 일 월 요일), Job, ReplicaSet, Replication controller

Deployment (replicas:2)
  ㄴ ReplicaSet (replicas:2)
  Pod 2개 중 1개 장애 발생 시 Desired State Management일환으로 한개 신규생성
  Deployment의 무중단 배포 (rolling update) 실행



# multi-container 실습
mkdir multi-pod && cd $_
kubectl apply -f multi-pod.yaml
kubectl exec -it multi-pod -c one-container -- bash
curl localhost:8000
kubectl exec -it multi-pod -c two-container -- bash
curl localhost:8080


# pod는 container의 호스트와 같으므로
# curl localhost대신 curl {pod_ip} 가능
C:\Users\user\Downloads\goott\LABs\multi-pod>kubectl get po,svc -o wide
NAME            READY   STATUS    RESTARTS   AGE     IP           NODE                                   NOMINATED NODE   READINESS GATES
pod/multi-pod   2/2     Running   0          3m52s   10.44.0.16   gke-k8s-junho-k8s-pool-92768e66-2zq6   <none>           <none>

kubectl exec -it multi-pod -c one-container -- curl {pod_ip}:8080
kubectl exec -it multi-pod -c two-container -- curl {pod_ip}:8000

kubectl exec -it multi-pod -c one-container -- curl 10.44.0.16:8080
kubectl exec -it multi-pod -c two-container -- curl 10.44.0.16:8000


# Pod내의 컨테이너들은 localhost 통신이 가능하다 ---> 'namespace'에 의해 제공됨
  컨테이너 본연의 기술 namespace! <- pause container가 제공함!


# 포트 둘다 8080으로 하면 충돌나서 컨테이너 1개만 실행됨!
cp multi-pod.yaml multi-pod2.yaml
kubectl apply -f multi-pod2.yaml
kubectl get po,svc -o wide
kubectl logs multi-pod2 -c one-container
kubectl logs multi-pod2 -c two-container


# 실습 자원 지우기
kubectl delete pod multi-pod multi-pod2
kubectl delete svc mysql-svc




kubectl apply -f multi-pod.yaml

kubectl get pod multi-pod -o wide
kubectl get pod multi-pod -o yaml
kubectl get pod multi-pod -o yaml > multi-pod-yaml.txt
kubectl get pod multi-pod -o json


# 30초 타임아웃 graceful (정상종료)
kubectl delete pod multi-pod

# 강제 종료 (주의 warning 표시됨)
kubectl delete pod multi-pod --grace-period=0 --force




# Pod - Label: service와 pod 연결 접점!
    ㄴ key: value 구조로 여러개 지정 가능
    ㄴ pod(labels)와 service(selector)간의 연결(접점)성을 제공
    ㄴ 모듈화
        pod1
        pod2 -> 공통 -> labels: app: myweb [kube-dns] ---> 서비스의 selector에 지정 시 그룹화
        pod3
        ㄴ 집합구성
      서비스와 컨테이너간에 LoadBalancer 생성됨
    ㄴ 객체를 식별하고 그룹화!
    ㄴ 확인? kubectl get pod xxx --show-labels


cd LABs/nginx-lab
kubectl apply -f nginx_test_pod.yml
kubectl apply -f nginx_test_svc.yml
kubectl get po,svc --show-labels

# pod - label 실습!
# 시각화 툴 : RealOpInsight

kubectl create namespace infra-team-ns1
kubectl get namespaces

mkdir infra-team && cd $_

# Pod3개 (같은 label) + svc (label과 같은 selector)
vi label-1.yaml

apiVersion: v1
kind: Pod
metadata:
  name: label-pod-a
  namespace: infra-team-ns1
  ## LABEL
  labels:
    type: infra1
spec:
  containers:
  - name: pod-a-container
    image: dbgurum/k8s-lab:initial
---
apiVersion: v1
kind: Pod
metadata:
  name: label-pod-b
  namespace: infra-team-ns1
  ## LABEL
  labels:
    type: infra1
spec:
  containers:
  - name: pod-b-container
    image: dbgurum/k8s-lab:initial
---
apiVersion: v1
kind: Pod
metadata:
  name: label-pod-c
  namespace: infra-team-ns1
  ## LABEL
  labels:
    type: infra1
spec:
  containers:
  - name: pod-c-container
    image: dbgurum/k8s-lab:initial
---
apiVersion: v1
kind: Service
metadata:
  name: infra-svc1
  namespace: infra-team-ns1
spec:
  # SELECTOR
  selector:
    type: infra1
  ports:
  - port: 7777


kubectl create namespace infra-team-ns1
kubectl get namespaces

kubectl apply -f label-1.yaml
kubectl get po,svc -o wide -n infra-team-ns1
kubectl get all -n infra-team-ns1
kubectl describe service/infra-svc1 -n infra-team-ns1
kubectl get pods --show-labels -n infra-team-ns1
kubectl get pods --selector="type" -n infra-team-ns1


kubectl delete -f label-1.yaml
kubectl delete namespace infra-team-ns1


# label2 실습

- namespace 필수아님 : default namespace!

mkdir label2 && cd $_
vi label-2.yaml
kubectl apply -f label-2.yaml

kubectl describe svc/front-web-svc
kubectl describe svc/infra-svc

C:\Users\user\Downloads\goott\LABs\label2>kubectl get pods
NAME       READY   STATUS    RESTARTS   AGE
label-p1   1/1     Running   0          15s
label-p2   1/1     Running   0          15s
label-p3   1/1     Running   0          15s
label-p4   1/1     Running   0          15s
label-p5   1/1     Running   0          15s
label-p6   1/1     Running   0          15s

C:\Users\user\Downloads\goott\LABs\label2>kubectl get svc
NAME            TYPE        CLUSTER-IP    EXTERNAL-IP   PORT(S)    AGE
front-web-svc   ClusterIP   10.48.4.166   <none>        8081/TCP   18s
infra-svc       ClusterIP   10.48.9.10    <none>        8082/TCP   18s
kubernetes      ClusterIP   10.48.0.1     <none>        443/TCP    6h52m

kubectl delete -f label-2.yaml



### POD를 노드에 자동/수동 지정

- pod먼저 apply 하고나서 Scheduler가 노드1/2에
  '자동' 지정(노드 metric을 바탕으로 algorithm을 통해)하여 ExternalIP 생성
- nodeSelector / nodeName 으로 '수동' 지정 가능함!

# 노드 2개의 라벨이 같음, 그중 호스트이름만 다름-> 이걸 nodeSelector로 지정!
#   kubernetes.io/hostname=k8s-node1
#   kubernetes.io/hostname=k8s-node2
kubectl get no --show-labeols

- nodeSelector는 라벨에 설정된 이름을 써야함
- nodeName은 단순히 k8s-node2 이름만 써줘도 됨!

Q1) 현재 운영중인 worker node는 100대다. 이중 HDD를 사용하려는 node와
SSD를 사용하는 node를 label로 분류된다. 새로 배포하려는 Pod는
신속한 데이터처리가 요구되므로 SSD로 구성된 node에 배치될 필요가 있다.
이를 만족하도록 Pod를 배치하려면?
  nodeSelector:
    disktype: ssd
    이렇게 쓰려면 노드에도 라벨을 붙여줘야함!

Q2) 식별 Label을  gpu 사용 유무로 지정하고자 한다?

* 에러발생! -> error -> Boolena type -> "true"
Labels:
  gpu: true

* 에러해결!
Labels:
  gpu: "true"


# kubectl label nodes k8s-node2 disktype=ssd
kubectl label nodes gke-k8s-junho-k8s-pool-92768e66-9s5n distype=ssd
kubectl get no --show-labels
  ... disktype=ssd ,...

vi nodeselect.yaml

apiVersion: v1
kind: Pod
metadata:
  name: nginx-gpu
spec:
  nodeSelector:
    disktype: ssd
  containers:
  - image: nginx
    name: nginx-container

kubectl get pod -o wide
# kubectl label nodes k8s-node2 distype-
kubectl label nodes gke-k8s-junho-k8s-pool-92768e66-9s5n distype-



kubectl describe no | grep -i taint
cd LABs/nodeselector



아래 옵션없이 실행시 pod "Pending"!
toleration옵션을 통해 마스터 노드에도 가능!!

vi taint-toleration-pod.yaml

apiVersion: v1
kind: Pod
metadata:
  name: taint-pod
spec:
  containers:
  - image: nginx:1.21-alpine
    name: taint-container
  nodeSelector:
    node-role.kubernetes.io/master: ""
  tolerations:
  - effect: NoSchedule
    key: node-role.kubernetes.io/master


kubectl describe no k8s-master
  node-role.kubernetes.io/master

kubectl apply -f taint-toleration-pod.yaml
kubectl  get po -o wide


# 자원정리

kubectl delete pod {p1} {p2} ...



gcloud container clusters resize $CLUSTER_NAME --num-nodes=0
gcloud container clusters resize k8s-junho --num-nodes=0
gcloud container clusters resize k8s-junho --num-nodes=3



- Service Object(resource)

kubectl apply -f mynode-pod.yaml
k get pod -o wide
    IP 10.111.156.72:8000
# 외부연결이 안됨 -> Cluster IP (클러스터 내부에서만 사용가능) 할당
# Cluster IP 대역 지정가능
curl 10.111.156.72:8000

# Cluster 내에서만 사용가능한 IP 대역을 지정
=>
kubeadm init --pod-network-cidr=10.96.0.0/12 --service-cidr=10.10.0.0/16...
                                          ㄴ 12bit=약 100만개          ㄴ 65536
# Service object(resource)
    ㄴ --service-cidr=10.10.0.0/16 사용하지 않았기 때문에 pod network의 IP가 할당
    ㄴ ClusterIP (기본값): cluster 내에서만! 사용가능한 IP 할당!
      => 같은 클러스터의 노드들에서는 접근가능 하지만 외부에서는 접근 안됨!!!

- Pod IP(EXTERNAL-IP) vs Cluster IP?
  - Why service? 서비스를 사용하는 이유?
    - 클러스터 아이피는 엔드포인트 여러곳으로, 외부에서 연결하기 위해? Load Balancer?
    1. Pod 장애로 IP가 변경되도 Service는 Label(이름)로 연결하기 때문에 application code 수정이 불필요!
                                        ㄴ How? kube-dns
      - 서비스는 IP상관없이 selecter > label로 붙음!
    2. 외부 연결을 위해서
    3. 로드밸런싱

    pod & service 생성 시 kubernetes master의 kube-dns에 자동 등록 -> 이름:IP
                                              ㄴ kube-proxy
    외부연결 --- Service --- Pod
                         ㄴ NAT

mkdir clusterip_test && cd $_
vi clustip-test.yaml

apiVersion: v1
kind: Pod
metadata:
  name: clusterip-pod
  labels:
    app: backend
spec:
  nodeSelector:
    kubernetes.io/hostname: k8s-node1
  containers:
  - name: container
    image: dbgurum/k8s-lab:v1.0
    ports:
    - containerPort: 8080
---
apiVersion: v1
kind: Service
metadata:
  name: clusterip-svc
spec:
  selector:
    app: backend
  ports:
  - port: 9000
    targetPort: 8080
  type: ClusterIP

kubectl apply -f clustip-test.yaml

kubectl get pod/clusterip-pod -o wide
kubectl get service/clusterip-svc -o yaml grep IP
kubectl get services
kubectl describe svc clusterip-svc | grep -i endpoint
kubectl get endpoints clusterip-svc

# cluster의 모든 node에서 cluster-IP로 조회 가능
node1# curl 10.100.1.153:9000/hostname

# 외부 (windows) 에서는 조회 안됨
C:\ curl 10.100.1.153:9000/hostname


# IP 외부로연결 방식!

1. externalIPS -> 특정 노드에 직접 연결하는 방식으로 외부 연결을 수행
  -> 서비스 오브젝트의 타입은 아님!
2. NodePort
  ㄴ 외부 트래픽을 서비스로 직접 가져오는 방식 제공
  ㄴ [NodeIP]:[NodePort] -> 서비스 생성 시 모든~노드에 지정 port open
                ㄴ 30000~32767 -> netstat -nlp grep 30000
                  -> nodePort 옵션 미지정 시 자동 할당!
                  -> nodePort 옵션 지정 시 지정된 port open!
                ㄴ 외부 트래픽 -> all node:port -> Service 연결 -> Pods에 전달 -> container -> application
                                  kube-proxy (traffic forward)
                                        ㄴ reoute table 이 작성
                ㄴ type: NodePort -> [nodePort]
                -> netstat -nlp grep 30000 -> ................[kube-proxy]



cat > mynode.yaml

apiVersion: v1
kind: Pod
metadata:
  name: mynode-pod
  labels:
    app: hi-mynode
spec:
  nodeSelector:
    kubernetes.io/hostname: k8s-node1
  containers:
  - name: mynode-container
    image: dbgurum/k8s-lab:v1.0
    ports:
    - containerPort: 8080
---
apiVersion: v1
kind: Service
metadata:
  name: mynode-svc
spec:
  selector:
    app: hi-mynode
type: NodePort
ports:
  - port: 8899
    targetPort:8000
    nodePort: 30000

kubectl apply -f mynode.yaml


kubectl get pods -o wide
  k8s-node1 확인
kubectl get services
  NodePort 8899:30267/TCP 확인
kubectl get endpoints mynode-svc
  mynode-svc 10.111.156.90:8000 확인



# 전달순서!
#   nodePort (외부) -> port (서비스) -> targetPort (Pod)

# node1
netstat -nlp grep 30267

# node2
netstat -nlp grep 30267

# master
netstat -nlp grep 30267

# 각 노드IP:30267로 모두 접속 가능
curl 192.168.56.101:30267


# TIME_WAIT 걸려있음...
netstat anlp | grep 30267



- 노드포트도 Load Balancer 있음!

# 파드 2개 만들기

cat > np-lb.yaml

apiVersion: v1
kind: Pod
metadata:
  name: pod-1
  labels:
    app: test
spec:
  nodeSelector:
    kubernetes.io/hostname: k8s-node1
  containers:
  - image: dbgurum/k8s-lab:v1.0
    name: container1
    ports:
    - containerPort: 8080
---
apiVersion: v1
kind: Pod
metadata:
  name: pod-2
  labels:
    app: test
spec:
  nodeSelector:
    kubernetes.io/hostname: k8s-node2
  containers:
  - image: dbgurum/k8s-lab:v1.0
    name: container2
    ports:
    - containerPort: 8080

--- 
apiVersion: v1
kind: Service
metadata:
  name: np-lb-svc
spec:
  selector:
    app: test
  ports:
  - port: 9000
    targetPort: 8080
    nodePort: 31111
  type: NodePort

kubectl apply -f np-lb.yaml

# 랜덤으로 로드밸런싱하여 노드접근 (노드포트 기법중 하나)
curl 192.168.56.101:31111/hostname

# 실제 돌고 있는 Pod/Service 수정 가능
# - Local : 처음 들어간 노드만 나옴!
# - Cluster : 랜덤으로 노드 접근
kubectl edit service/np-lb-svc
  20 externalTrafficPolicy: Cluster <-> Local
  21 internalTrafficPolicy: Cluster <-> Local

단, 동일 노드에 pod가 2개이상 있다면, Local로 해도 자동 LB 발생

3) Loadaancer
  - 클라우드기반
  - vm

cat > lb-test1.yaml

apiVersion: v1
kind: Pod
metadata:
  name: pod-lb
  labels:
    app: lb-pod
spec:
  containers:
  - image: dbgurum/k8s-lab:v1.0
    name: container1
    ports:
    - containerPort: 8080
--- 
apiVersion: v1
kind: Service
metadata:
  name: lb-svc
spec:
  selector:
    app : lb-pod
  ports:
  - port: 9000
    targetPort: 8080
  type: NodeBalancer

kubectl apply -f np-lb.yaml

kubectl get po,svc -o wide

curl 34.64.252.135:9000/hostname

kubectl delete pod pod-lb
kubectl delete svc lb-svc



# Deployment내의 파드 3개를 의무적으로 유지!!!

cat > lb-deploy.yaml

apiVersion: apps/v1
kind: Deployment
metadata:
  name: dev-deploy
  labels:
    app: devdeploy
spec:
  replicas: 3
  selector: 
    matchLabels:
      app: devdeploy
  template:
    metadata:
      labels:
        app: devdeploy
    spec:
      containers:
      - image: dbgurum/k8s-lab:v1.0
        name: dev-container
        ports:
        - containerPort: 8080

---
apiVersion: v1
kind: Service
metadata:
  name: dev-svc
  labels:
spec:
  selector:
    app: devdeploy
  ports:
  - port: 80
    targetPort: 8080
  type: LoadBalancer



kubectl get svc
curl 34.64.252.135:80/hostname


kubectl get deploy,po -o wide
kubectl get deploy,svc -o wide

# 다시 띄워짐 2->3
kubectl delete pod/dev-deploy-765758bcb6-kkx48

# 완전 3개 지우려면
kubectl delete deploy dev-deploy


- Server 오브젝트
4) Ingress -> L7 -> http(80), https(443)  -> 라우팅 서비스
      ㄴ www.example.com
      ㄴ 외부 트래픽을 연결해주는 서비스

                          (route rule)      (labels)
  www.example.com/          -> main-svc      -> pods -> container -> app
  www.example.com/customer  -> customer-svc  -> pods
  www.example.com/order     -> order-svc     -> pods
  www.example.com/payment   -> payment-svc   -> pods



- Ingress controller 설치
https://kubernetes.github.io/ingress-nginx/deploy/

- bare-metal 또는 GKE용

kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.2.0/deploy/static/provider/baremetal/deploy.yaml
  Ingress-nginx 3개 생성됨: 컨트롤서 생성완료!

kubectl get svc -n ingress-nginx
  80:31526/TCP, 443:32073/TCP



- 파드만들기

# 서비스1개 파드1개 <- ingress X 없이!

mkdir ingress1 && cd $_
cat > hi-app.yaml

apiVersion: v1
kind: Pod
metadata:
  name: hi-pod
  labels:
    app: hi
spec:
  containers:
  - image: dbgurum/ingress:hi
    name: hi-container
    args:
    - "-text=Hi! Kubernetes~"
---
apiVersion: v1
kind: Service
metadata:
  name: hi-svc
spec:
  selector:
    app: hi
  ports:
  - port: 5678


kubectl get po,svc -o wide
service/hi-svc ClusterIp 10.110.149.98

curl 10.110.149.98:5678
curl http://10.110.149.98:5678

kubectl api-resources | grep ingr


# 서비스1개 파드1개 <- ingress O
cat > hi-ingress.yaml

apiVersion: v1
kind: Pod
metadata:
  name: hi-pod
  labels:
    app: hi
spec:
  containers:
  - image: dbgurum/ingress:hi
    name: hi-container
    args:
    - "-text=Hi! Kubernetes~"
---
apiVersion: v1
kind: Service
metadata:
  name: hi-svc
spec:
  selector:
    app: hi
  ports:
  - port: 5678
---
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: http-hi
  annotations:
    kubernetes.io/ingress.class: "nginx"
    ingress.kubernetes.io/rewrite-target: /
spec:   # http를 이용하여 /hi로 들어오면?
  rules:
  - http:
      paths:
      - path: /hi
        pathType: Prefix
        backend:
          service:
            name: hi-svc
            port:
              number: 5678

kubectl apply -f hi-ingress.yaml

kubectl get svc -n ingress-nginx
kubectl describe ingress http-hi
kubectl get ing,po,svc -o wide

curl http://192.168.56.100:31016/hi
curl http://192.168.56.101:31016/hi
curl http://192.168.56.102:31016/hi

- ingress -> webhook error 에러!!
  - https://cafe.naver.com/ocmkorea/19060

Error from server (InternalError): error when creating "hi-ingress.yaml": Internal error occurred: failed calling webhook "validate.nginx.ingress.kubernetes.io": failed to call webhook: Post "https://ingress-nginx-controller-admission.ingress-nginx.svc:443/networking/v1/ingresses?timeout=10s": dial tcp 10.99.64.181:443: connect: connection refused

nginx controller에
kubectl delete validatingwebhookconfiguration ingress-nginx-admission



- 크래쉬루프 crashloop : 익명의 사용자의 경우 에러가능...





- 서비스3개 + 파드6개!
  - 파드2개 서비스1개 * 3 (/, /bookstore, /payment)


cat > hi-ingress-2.yaml

apiVersion: v1
kind: Pod
metadata:
  name: main-pod-1
  labels:
    app: main
spec:
  containers:
  - image: dbgurum/ingress:hi
    name: main-container-1
    args:
    - "-text=This is the MAIN-page-1."
---
apiVersion: v1
kind: Pod
metadata:
  name: main-pod-2
  labels:
    app: main
spec:
  containers:
  - image: dbgurum/ingress:hi
    name: main-container-2
    args:
    - "-text=This is the MAIN-page-1."
---
apiVersion: v1
kind: Service
metadata:
  name: main-svc
spec:
  selector:
    app: main
  ports:
  - port: 5678
---

apiVersion: v1
kind: Pod
metadata:
  name: bookstore-pod-1
  labels:
    app: bookstore
spec:
  containers:
  - image: dbgurum/ingress:hi
    name: bookstore-container-1
    args:
    - "-text=This is the BOOKSTORE-page-1."
---
apiVersion: v1
kind: Pod
metadata:
  name: bookstore-pod-2
  labels:
    app: bookstore
spec:
  containers:
  - image: dbgurum/ingress:hi
    name: bookstore-container-2
    args:
    - "-text=This is the BOOKSTORE-page-1."
---
apiVersion: v1
kind: Service
metadata:
  name: bookstore-svc
spec:
  selector:
    app: bookstore
  ports:
  - port: 5678
---

apiVersion: v1
kind: Pod
metadata:
  name: payment-pod-1
  labels:
    app: payment
spec:
  containers:
  - image: dbgurum/ingress:hi
    name: payment-container-1
    args:
    - "-text=This is the PAYMENT-page-1."
---
apiVersion: v1
kind: Pod
metadata:
  name: payment-pod-2
  labels:
    app: payment
spec:
  containers:
  - image: dbgurum/ingress:hi
    name: payment-container-2
    args:
    - "-text=This is the PAYMENT-page-1."
---
apiVersion: v1
kind: Service
metadata:
  name: payment-svc
spec:
  selector:
    app: payment
  ports:
  - port: 5678
---

apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: http-main-bookstore-payment
  annotations:
    kubernetes.io/ingress.class: "nginx"
    ingress.kubernetes.io/rewrite-target: /
spec:   # http -> /
  rules:
  - http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: main-svc
            port:
              number: 5678
      - path: /bookstore
        pathType: Prefix
        backend:
          service:
            name: bookstore-svc
            port:
              number: 5678
      - path: /payment
        pathType: Prefix
        backend:
          service:
            name: payment-svc
            port:
              number: 5678


kubectl get svc -n ingress-nginx
kubectl describe ingress http-hi
kubectl get ing,po,svc -o wide

curl 192.168.56.101:31016/
  This is the MAIN-page-1.
curl 192.168.56.101:31295/bookstore
  This is the BOOKSTORE-page-1.
curl 192.168.56.101:31295/payment
  This is the PAYMENT-page-1.


- ERROR! 400 The plain HTTP request was sent to HTTPS port
- pod ip, cluster ip 로 연결 테스트 다 된다면, ingress 설정 문제로 보시면 되요. 80:31295 이렇게 연결된 정보인지도 다시 확인해주시구요.


- https 의 경우 secret적용!
  - 시크릿 object

```yaml
tls:
- hosts
```



- Volume and data

  -> 공유 -> container to host / container to container
  -> database container + volume -> db container 삭제 후 다 시 run -> 기존 DB-Table-data 그대로 이전
        ㄴ why? ??? 통신

1) emptyDir (in Pod, multi-container)
                      ㄴ container1: /mount1 <---> container2: /mount2

- 별도 공간 volume 있는게 아니라 컨테이너 1,2의 /mount1, /mount2를 연결!
  ㄴ emptyDir 기법의 볼륨 데이터는 Pod lifecycle과 같다.

2) hostPath -> container to host
- pod내의 컨테이너 볼륨은 pod 삭제시 없어지는 단점 보완


cat > emptydir1.yaml

apiVersion: v1
kind: Pod
metadata:
  name: pod-vol1
spec:
  containers:
  - name: container1
    image: dbgurum/k8s-lab:initial
    volumeMounts:
    - name: empty-dir
      mountPath: /mount1
  - name: container2
    image: dbgurum/k8s-lab:initial
    volumeMounts:
    - name: empty-dir
      mountPath: /mount2
  volumes:
  - name: empty-dir
    emptyDir: {}

kubectl apply -f emptydir1.yaml

-- container1: # kubectl exec -it pod-vol1 -c container1 -- bash
mount | grep /mount1
cd /mount1/
echo "i love k8s~" >> k8s.txt
ls
  k8s.txt
cat k8s.txt

-- container2: # kubectl exec -it pod-vol1 -c container2 -- bash
cd /mount2/
ls
  k8s.txt
cat k8s.txt
mount | grep /mount2


-- host
kubectl cp pod-vol1:mount1/k8s.txt -c container1 k8s.txt
ls
  anaconda-ks.cfg Desktop Downloads k8s.txt


[질문] 쿠버 igress, pod, service, volume 설계구성시
igress->service->pod, volume 이런식으로 구성 후 추가확장 시에는 service, pod 추가하면 되죠?
전체 구성시 어떤식으로 구성하는지 궁금해서요?

[답변] 전체 설계에 기본적으론 순서는 없지만, 말씀하신대로 전체 설계를 해 놓고
top-down, bottom-up 방식으로 해 나가시면 됩니다. 확장이야 그 후 문제니 어려울건 없구요~



(주의) DirectoryOrCreate가 아닌 Directory를 사용하면 Pod가 실행될 노드의 해당 디렉토리 미리 생성해야 됨!

node1, node2의 dataDir를 nfs로 붙임!!!

cat > emptydir2.yaml

apiVersion: v1
kind: Pod
metadata:
  name: pod-vol2
spec:
  nodeSelector:
    kubernetes.io/hostname: k8s-node1
  containers:
  - name: container
    image: dbgurum/k8s-lab:initial
    volumeMounts:
    - name: host-path
      mountPath: /mount1
  volumes:
  - name: host-path
    hostPath:
      path: /data_dir
      type: DirectoryOrCreate

k get po,svc -o wide
(k8s-master) kubectl exec -it pod-vol2 -c container -- bash
mount | grep /mount1
echo "i love you" >> k8s-2.txt
cat k8s-2.txt
kubectl get po -o wide

# pod가 저장된 node에서 확인
(k8s-node1) cd /data_dir/
(k8s-node1) ls
  k8s-2.txt



cd LABs/
mkdir hostpath && cd $_

cat > weblog-pod.yaml

# nginx -> access_log, error_log
#   host-path 기법으로 로그파일을 호스트로 뺄수 있음!!
#   /data_dir
apiVersion: v1
kind: Pod
metadata:
  name: weblog-pod
spec:
  nodeSelector:
    kubernetes.io/hostname: k8s-node1
  containers:
  - name: nginx-container
    image: nginx:1.21-alpine
    volumeMounts:
    - name: host-path
      mountPath: /var/log/nginx
  volumes:
  - name: host-path
    hostPath:
      path: /data_dir/web-log
      type: DirectoryOrCreate


k apply -f weblog-pod.yaml

k get po -o wide
# (node1)
cd /data_dir/web-log
# (node1)
tail -f access.log

# (master)
k get po -o wide | grep weblog-pod
# (master)
curl 10.111.156.71

각노드에서 실행! : 
rdate -s time.bora.net

-> pattern -> awk $1=IP, $4=time

# (node1)
awk '$4>"[25/Jun/2022:05:42:53]" && $4"[25/Jun/2022:05:44:40]"' \
  access.log | awk '{ print $1}'| \
  sort | uniq -c | sort -r | more





 - mysql volume 실습

cat > mysql-pod.sql

apiVersion: v1
kind: Pod
metadata:
  name: mysql-pod
spec:
  nodeSelector:
    kubernetes.io/hostname: k8s-node2
  containers:
  - name: db-container
    image: mysql:8.0
    volumeMounts:
    - name: host-path
      mountPath: /var/lib/mysql
    env:
    - name: MYSQL_ROOT_PASSWORD
      value: "k8spass#"
  volumes:
  - name: host-path
    hostPath:
      path: /data_dir/mysql-data
      type: DirectoryOrCreate


k get po -o wide

# master
kubectl exec -it mysql-pod -- bash
mysql -uroot
show databases;
create database k8sdb;
use k8sdb;
create table product (pord_id int, prod_name varchar(30));
insert into product values (1, 'kubernetes');
select * from product;

k delete -f mysql-pod.yaml

# (node2)
# 위에서 pod 지웠어도 데이터 남아있음
cd /data_dir/mysql-data


# master
# 여전히 데이터 남아있음 : 데이터 영속성!
k apply -f mysql-pod.yaml

kubectl exec -it mysql-pod -- bash
mysql - uroot -p
show databases;





# Linux 커멘드 사용
# 2개의 노드가 파일시스템을 공유 테스트: /data_dir
# node1의 /data_dir을 node1과 node2에서 공유
# node1
vi /etc/exports
  /data_dir *(rw,sync,all_squash,anonuid=0,anongid=0)

systemctl restart nfs
exportfs -v
  /data_dir       <world>(sync,wdelay,hide,no_subtree_check,anonuid=0,anongid=0,sec=sys,rw,secure,root_squash,all_squash)

# node2
vi /etc/fstab
  k8s-node1:/data_dir   /data_dir   nfs   defaults   0 0
mount /data_dir
df -h | grep node1
mount | grep k8s-node1

# 모든 노드에서 확인
showmount -e 192.168.56.101
  Export list for 192.168.56.101:
  /data_dir *

# node1
cd /data_dir
touch testfile.txt
# node2
ls /data_dir
  testfile.txt k8s-2.txt k8s-3.txt



- GKE
  - 컴퓨팅 > Compute Engine > 디스크 > 디스크 만들기

# gcloud에 디스크추가! 
# mongodb 전영 디스크 스토리지!

gcloud container clusters list
gcloud compute disks create mongo-storage --size=10GiB --zone=asia-northeast3-a
gcloud compute disks list
  NAME                                  LOCATION           LOCATION_SCOPE  SIZE_GB  TYPE         STATUS
  gke-k8s-junho-k8s-pool-92768e66-3bvv  asia-northeast3-a  zone            100      pd-ssd       READY
  gke-k8s-junho-k8s-pool-92768e66-g61l  asia-northeast3-a  zone            100      pd-ssd       READY
  gke-k8s-junho-k8s-pool-92768e66-zrgq  asia-northeast3-a  zone            100      pd-ssd       READY
  mongo-storage                         asia-northeast3-a  zone            10       pd-standard  READY

# 해당 스토리지로 접속
cat > gce-mongo.yaml

apiVersion: v1
kind: Pod
metadata:
  name: gce-mongo-pod
spec:
  containers:
  - image: mongo
    name: mongo-container
    volumeMounts:
    - name: mongodb-data
      mountPath: /data/db
    ports:
    - containerPort: 27017
      protocol: TCP
  volumes:
  - name: mongodb-data
    gcePersistentDisk:
      pdName: mongo-storage
      fsType: ext4

kubectl apply -f gce-mongo.yaml
kubectl get po -o wide

# 영속적인 볼륨
kubectl describe pod gce-mongo-pod
  Volumes:
  mongodb-data:
    Type:       GCEPersistentDisk (a Persistent Disk resource in Google Compute Engine)
    PDName:     mongo-storage
    FSType:     ext4
    Partition:  0
    ReadOnly:   false
  kube-api-access-pn4cq:
    Type:                    Projected (a volume that contains injected data from multiple sources)
    TokenExpirationSeconds:  3607
    ConfigMapName:           kube-root-ca.crt
    ConfigMapOptional:       <nil>
    DownwardAPI:             true

kubectl exec -it gce-mongo-pod -- mongo
# 또는 대시보드 > 파드
#   mongo

```
use mystore
db
db.customer.insert({"name": "oooo"})
db.customer.find()
```


# 파드 삭제후 재기동해도 데이터 영속성 유지!
kubectl delete -f gce-mongo.yaml
kubectl apply -f gce-mongo.yaml
kubectl exec -it gce-mongo-pod -- mongo




# PV (persistent Volume) -> 지속적 볼륨
        ㄴ k8s cluster 내에서 추가되는 가상 스토리지 인스턴스 -> 관리자 생성
        ㄴ host 영역의 공간을 프로비저닝 (claim)

  --- bind (연결) -> how? PVC 생성 시 지정한 용량 & 접근방식(accessModes)이 매칭되는 PV가 자동 선택!

  PVC (Persistent Volume Claim) -> 지속적 볼륨 요청 <- Pod 가 사용 -> 개발자 생성
  -> Pod는 볼륨에 PVC를 지정

    accessModes:
      - ReadWriteOnce(RWO)    -> 하나의 노드가 볼륨을 Read/Write 가능하도록 mount
      - ReadOnlyMany(ROX)     -> 여러 노드가 읽기 전용으로 사용 가능하도록 mount
      - ReadWriteMany(RWX)    -> 여러 노드가 Read/Write 가능하도록 mount

    persistentVolumeReclaimPolicy: Retain
      - Retain(기본값) -> PV의 데이터는 그대로 보존
      - Recycle -> 재사용 시 PV 데이터를 모두 삭제 후 재사용
      - Delete -> 사용이 종료되면 볼륨 삭제

(질문) retain 값이면 파드 종류후 다시 생성되는 파드가 재활용이 가능한 부분인가요?
(답변) O


[LAB]
-- pv1 : storage 1G ReadWriteOnce (RWO)
  L 하나의 노드가 볼륨을 Read/Write 가능하도록 마운트
-- pv2 : storage 1G ReadOnlyMany (ROX)
  L 여러개 노드가 Read전용으로 사용하도록 마운트
-- pv3 : storage 2G ReadWriteMany (RWX)
  L 여러개의 노드가 Read/Write 가능하도록 마운트


# PV 생성!

cat > persistent-vol.yaml

apiVersion: v1
kind: PersistentVolume
metadata:
  name: pv1
spec:
  capacity:
    storage: 1Gi
  accessModes:
  - ReadWriteOnce
  # Retain이 디폴트라서 생략가능
  # persistentVolumeReclaimPolicy: Retain
  local:
    path: /data_dir
  nodeAffinity:
    required:
      nodeSelectorTerms:
      - matchExpressions:
        - {key: kubernetes.io/hostname, operator: In, values: [k8s-node1]}
---
apiVersion: v1
kind: PersistentVolume
metadata:
  name: pv2
spec:
  capacity:
    storage: 1Gi
  accessModes:
  - ReadOnlyMany
  persistentVolumeReclaimPolicy: Retain
  local:
    path: /data_dir
  nodeAffinity:
    required:
      nodeSelectorTerms:
      - matchExpressions:
        - {key: kubernetes.io/hostname, operator: In, values: [k8s-node1]}
---
apiVersion: v1
kind: PersistentVolume
metadata:
  name: pv3
spec:
  capacity:
    storage: 2Gi
  accessModes:
  - ReadWriteMany
  persistentVolumeReclaimPolicy: Retain
  local:
    path: /data_dir
  nodeAffinity:
    required:
      nodeSelectorTerms:
      - matchExpressions:
        - {key: kubernetes.io/hostname, operator: In, values: [k8s-node1]}


# pv, pvc 조회
kubectl get pv,pvc -o wide
  NAME                   CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS      CLAIM   STORAGECLASS   REASON   AGE   VOLUMEMODE
  persistentvolume/pv1   1Gi        RWO            Retain           Available                                   8s    Filesystem
  persistentvolume/pv2   1Gi        ROX            Retain           Available                                   8s    Filesystem
  persistentvolume/pv3   2Gi        RWX            Retain           Available                                   8s    Filesystem


(질문) 화면의 yaml 코드로 pv 를 생성하면 nodeAffinity 에 의해서 k8s-node1의 실제 디스크를 기반으로 pv 를 생성하게끔 되는건가요?
(답변) 네 아까도 말씀드렸었죠~ 호스트(local)의 /data_dir에 프로비저닝된다는~~ 대신 호스트에는 어떤 흔적이나 공간 차지는 남지 않고 예약만 되는겁니다.


(질문) 노드1,2에서 kubectl 커멘드시 localhost:8080 refused나오면
kubeadm join을 다시 해줘야하는것인가요?
(답변) 아뇨 원래 kubectl로 조회는 마스터(컨트롤플레인) 에서만 되는겁니다.  노드에서는 일체 그런 작업 안합니다.


--- bind (연결) -> how? PVC 생성 시 지정한 용량 & 접근방식(accessModes)이 매칭되는 PV가 자동 선택!


# PVC 생성!

cat > pvc.yaml

apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: pvc1
spec:
  accessModes:
  - ReadOnlyMany
  resources:
    requests:
      storage: 1G
  storageClassName: ""
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: pvc2
spec:
  accessModes:
  - ReadWriteOnce
  resources:
    requests:
      storage: 1G
  storageClassName: ""
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: pvc3
spec:
  accessModes:
  - ReadWriteOnce
  resources:
    requests:
      storage: 5G
  storageClassName: ""
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: pvc4
spec:
  accessModes:
  - ReadWriteMany
  resources:
    requests:
      storage: 1G
  storageClassName: ""


k apply pvc.yaml
k get pv,pvc


# 삭제시 pvc 지우고 -> pv 삭제

pvc를 파드에 붙이면 노드1에 /data_dir





- Kubernetes Volume PV & PVC 이용한 Pod

cat > pod-pvc.yaml

apiVersion: v1
kind: Pod
metadata:
  name: mynode-pod
spec:
  containers:
  - image: dbgurum/mynode:1.0
    name: mynode-container
    ports:
    - containerPort: 8000
    volumeMounts:
    - name: mynode-path
      mountPath: /mynode
  volumes:
  - name: mynode-path
    persistentVolumeClaim:
      claimName: pvc1


kubectl apply -f pod-pvc.yaml
k get po
k exec -it mynode-pod -- bash
ls
  mynode
cd /mynode/
touch k8s-5.txt
ls
  k8s-3.txt k8s-4.txt k8s-5.txt

# node1
ls
  k8s-5.txt





- pvc5번에 pv4에 붙음
- PV에 라벨 명시

cat > pv4-pvc5.yaml

apiVersion: v1
kind: PersistentVolume
metadata:
  name: pv4
  labels:
    name: pv4
spec:
  capacity:
    storage: 1G
  accessModes:
  - ReadWriteOnce
  local:
    path: /data_dir
  nodeAffinity:
    required:
      nodeSelectorTerms:
      - matchExpressions:
        - {key: kubernetes.io/hostname, operator: In, values: [k8s-node1]}
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: pvc5
spec:
  accessModes:
  - ReadWriteOnce
  resources:
    requests:
      storage: 1G
  storageClassName: ""
  selector:
    matchLabels:
      name: pv4


k get pv,pvc -o wide
  NAME                   CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM          STORAGECLASS   REASON   AGE   VOLUMEMODE
  persistentvolume/pv1   1Gi        RWO            Retain           Bound    default/pvc2                           34m   Filesystem
  persistentvolume/pv2   1Gi        ROX            Retain           Bound    default/pvc1                           34m   Filesystem
  persistentvolume/pv3   2Gi        RWX            Retain           Bound    default/pvc4                           34m   Filesystem
  persistentvolume/pv4   1G         RWO            Retain           Bound    default/pvc5                           25s   Filesystem

  NAME                         STATUS    VOLUME   CAPACITY   ACCESS MODES   STORAGECLASS   AGE   VOLUMEMODE
  persistentvolumeclaim/pvc1   Bound     pv2      1Gi        ROX                           18m   Filesystem
  persistentvolumeclaim/pvc2   Bound     pv1      1Gi        RWO                           18m   Filesystem
  persistentvolumeclaim/pvc3   Pending                                                     18m   Filesystem
  persistentvolumeclaim/pvc4   Bound     pv3      2Gi        RWX                           18m   Filesystem
  persistentvolumeclaim/pvc5   Bound     pv4      1G         RWO                           25s   Filesystem



cat > pvc-pv-pod.yaml

apiVersion: v1
kind: PersistentVolume
metadata:
  name: pv-nfs
spec:
  capacity:
    storage: 1G
  volumeMode: Filesystem
  accessModes:
  - ReadWriteOnce
  persistentVolumeReclaimPolicy: Recycle
  mountOptions:
    - hard
  nfs:
    path: /data_dir
    server: 192.168.56.101
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: pvc-nfs
spec:
  accessModes:
  - ReadWriteOnce
  resources:
    requests:
      storage: 1Gi
---
apiVersion: v1
kind: Pod
metadata:
  name: nodejs-pvc-pod
spec:
  containers:
  - image: dbgurum/mynode:1.0
    name: nodejs-container
    ports:
    - containerPort: 8000
    volumeMounts:
    - name: testpath
      mountPath: /mount1
  volumes:
  - name: testpath
    persistentVolumeClaim:
      claimName: pvc-nfs


k get po,pv,pvc -o wide
k describe pod nodejs-pvc-pod
k exec -it nodejs-pvc-pod -- bash
cd /monut1
touch kubernetes_pvc.txt

# node1
ls
  kubernetes_pvc.txt






# day04
# 네임스페이스를 따로 만들어서 그안에 작업하는 경우 제외하고
# 전체삭제는 지양. 디폴트 네임스페이스에서 작업했으면
# po,svc 등 직접 하나하나 삭제
k get po,svc -o wide
k delete po p1 p2 p3
 


# worker node <- Not Ready 리소스 부족!
  ㄴ 리소스 부족
  ㄴ 리눅스 네트워크
  ㄴ kubelet 장애 -> 
  해당 노드의 systemctl restart kubelet

kubectl edit pv pv-volume

# PV -> 물리적 용량(node1 -> /data_dir) 에 대한 provisioning
  ㄴ PV -> PVC -> Pod
    ㄴ 용량 동적 변경




k get pv, pvc
k edit pv pv-volume
  - ReadWriteOnce
  capacity:
    storage: 20Gi

k get pv,pvc


cat > pv-test.yaml

apiVersion: v1
kind: PersistentVolume
metadata:
  name: pv-volume
  labels:
    type: local
spec:
  storageClassName: manual
  capacity:
    storage: 10Gi
  accessModes:
  - ReadWriteOnce
  hostPath:
    path: "/mnt/data"
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: pvc-volume
spec:
  storageClassName: manual
  accessModes:
  - ReadWriteOnce
  resources:
    requests:
      storage: 3Gi
---
apiVersion: v1
kind: Pod
metadata:
  name: web-pod
spec:
  volumes:
  - name: web-pv-storage
    persistentVolumeClaim:
      claimName: pvc-volume
  containers:
  - name: web-container
    image: nginx
    ports:
    - containerPort: 80
      name: "web-server"
    volumeMounts:
    - mountPath: "/usr/share/nginx/html"
      name: web-pv-storage

k apply -f pv-test.yaml
k get pv, svc

k delete -f pv-test.yaml


# 파일 기반의 인프라 구축 -> IaC -> 코드형 인프라




# GKE실습 pv,pvc

cat > gce-pv-mongo2.yaml

apiVersion: v1
kind: PersistentVolume
metadata:
  name: mongodb-pv
spec:
  capacity:
    storage: 1Gi
  accessModes:
  - ReadWriteOnce
  - ReadOnlyMany
  persistentVolumeReclaimPolicy: Retain
  gcePersistentDisk:
    pdName: mongo-stroage # GCE로 생성
    fsType: ext4

cat > gce-pvc-mongo2.yaml

apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: mongodb-pvc
spec:
  resources:
    requests:
      storage: 1Gi
  accessModes:
  - ReadWriteOnce
  storageClassName: ""


kubectl apply -f gce-pv-mongo2.yaml
kubectl apply -f gce-pvc-mongo2.yaml
kubectl get pv,pvc -o wide



- 로보3t로 접근하려면 외부접근이 필요하므로 svc 생성해야함!
- Kubernetes API Resource -> Volume

cat > gce-mongo2-pod.yaml

apiVersion: v1
kind: Pod
metadata:
  name: mongodb-pod
spec:
  containers:
  - image: mongo
    name: mongodb
    volumeMounts:
    - name: mongodb-data
      mountPath: /data/db
    ports:
    - containerPort: 27017
      protocol: TCP
  volumes:
  - name: mongodb-data
    persistentVolumeClaim:
      claimName: mongodb-pvc

kubectl apply -f gce-mongo2-pod.yaml

kubectl exec -it mongodb-pod -- mongo
use mystore
db.customer.find()





- PV에 PVC가 붙어서 사용가능한 공간이 되고 Pod등 실행
- PV(1G) -> PVC (1G) -> Pod(최대 1G)
  - StorageClass -> '동적할당'이 가능한 볼륨 기법 -> PV 처럼 사용가능 -> PVC에 매칭가능 -> Pod
    : 동적 볼륨!


# GCE 실습

kubectl get pv,pvc -o wide


kubectl api-resources | grep -i storageclass
  storageclasses                    sc           storage.k8s.io/v1                      false        StorageClass
gcloud container clusters list
  NAME       LOCATION           MASTER_VERSION  MASTER_IP      MACHINE_TYPE  NODE_VERSION    NUM_NODES  STATUS
  k8s-junho  asia-northeast3-a  1.22.8-gke.202  34.64.186.244  e2-medium     1.22.8-gke.202  3          RUNNING


# 용량이 없음: StorageClass 동적할당이므로!!
# 1. StorageClass 만들기!!
cat > gce-sclass.yaml

apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: fast
provisioner: kubernetes.io/gce-pd
parameters:
  type: pd-ssd
  zone: asia-northeast3-a

kubectl apply -f gce-sclass.yaml
kubectl get sc -o wide

# 2. PVC 만들기 -> 위의 StorageClass 사용

cat > mongo-pvc-sclass.yaml
apiVersion: storage.k8s.io/v1
kind: PersistentVolumeClaim
metadata:
  name: mongo-pvc-sclass
spec:
  storageClassName: fast
  resources:
    requests:
      storage: 1Gi
  accessModes:
  - ReadWriteOnce


kubectl apply -f mongo-pvc-sclass.yaml
kubectl get sc -o wide



# 3. Pod
cat > gce-pod-sclass.yaml
apiVersion: v1
kind: Pod
metadata:
  name: mongodb-pvc-sclass
spec:
  containers:
  - image: mongo
    name: mongodb
    volumeMounts:
    - name: mongodb-data
      mountPath: /data/db
    ports:
    - containerPort: 27017
      protocol: TCP
  volumes:
  - name: mongodb-data
    persistentVolumeClaim:
      claimName: mongodb-pvc-sclass

kubectl apply -f gce-pod-sclass.yaml
kubectl get po,svc,deploy


gcloud container clusters list
gcloud compute disks delete mongo-stroage --zone asia-northeast3-a



# configMap
  ㄴ 환경변수, 특정 인수값, * 구성파일(*.conf)*
                              ㄴ nginx.conf -> 웹서비스를 수행할 수 있는 구성
                                    ㄴ proxy로 재구성
  - 키-값 구조, 기밀 데이터가 아닌 데이터 저장
  - Pod에 참조값으로 지정
  - 생성
      pod -> kubectl run | create -> CKA
      kubectl create configmap ...
                                    --from-literal mypwd=k8spass#
                                    --from-literal=mypwd=k8spass#
                                    --from-literal a=b --from-literal a2=b2 여러번 지정가능

      -> 여러개의 환경값을 지정해야 되는 경우 -> 환경값을 저장한 파일 생성
                                    --from-file=파일명

  - Pod에 참조값으로 지정


# secret (암호키 -> 아닌 -> 인코딩 -> 'generic')
  - 키-값 구조, 기밀 데이터 저장 시 사용
  - Pod에 참조값으로 지정
  - 암호화된 값으로 저장 -> "base64" encoding (가볍고 빠름) -> aGVsbG8ga3ViZXJuZXRlcwo=
  - secret은 인코딩 값을 저장
      ㄴ Pod에 참조값으로 지정 (Pod내에 참조값으로 쓸때는 디코딩 됨)
                ㄴ 자동 디코딩

# 인코딩값

echo 'hello kubernetes' | base645
  aGVsbG8ga3ViZXJuZXRlcwo=
echo aGVsbG8ga3ViZXJuZXRlcwo= | base64 --decode



cd LABs
mkdir configmap && cd $_
kubectl create configmap log-mod-cm --from-literal=LOG_LEVEL=DEBUG
k get cm
k describe cm log-mod-cm
k get cm log-mod-cm -o yaml > log-mod-cm.yaml
scp root@192.168.56.100:/root/log-mod-cm.yaml .

apiVersion: v1
data:
  LOG_LEVEL: DEBUG
kind: ConfigMap
metadata:
  creationTimestamp: "2022-06-26T02:08:10Z"
  name: log-mod-cm
  namespace: default
  resourceVersion: "41861"
  uid: 913e11b1-c7be-48f8-9c6c-0be21764bfa1


# 위의 configMap 참조하는 Pod 만들기!

cat > log-mod-pod1.yaml
apiVersion: v1
kind: Pod
metadata:
  name: log-mod-pod1
spec:
  containers:
  - image: nginx:1.21-alpine
    name: log-container
    # configMap을 참조!
    envFrom:
    - configMapRef:
        name: log-mod-cm

scp log-mod-pod1.yaml root@192.168.56.100:/root

k apply -f log-mod-pod1.yaml
k exec -it log-mod-pod1 -- env
  PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
  HOSTNAME=log-mod-pod1
  TERM=xterm
  LOG_LEVEL=DEBUG
  KUBERNETES_SERVICE_HOST=10.96.0.1
  KUBERNETES_SERVICE_PORT=443
  KUBERNETES_SERVICE_PORT_HTTPS=443
  KUBERNETES_PORT=tcp://10.96.0.1:443
  KUBERNETES_PORT_443_TCP=tcp://10.96.0.1:443
  KUBERNETES_PORT_443_TCP_PROTO=tcp
  KUBERNETES_PORT_443_TCP_PORT=443
  KUBERNETES_PORT_443_TCP_ADDR=10.96.0.1
  NGINX_VERSION=1.21.6
  NJS_VERSION=0.7.3
  PKG_RELEASE=1
  HOME=/root

# 컨피그맵을 Pod에 참조시키는 기법 #2

k describe log-mod-cm

cat > log-mod-pod2.yaml
apiVersion: v1
kind: Pod
metadata:
  name: log-mod-pod2
spec:
  containers:
  - image: nginx:1.21-alpine
    name: log-container
    # configMap을 참조-#2! 볼륨마운트 방식!
    # 로그레벨 파일이 /etc/config에 생성이 되고 LOG_LVEL=DEBUG 지정됨
    volumeMounts:
    - name: cm-volume
      mountPath: /etc/config
  volumes:
  - name: cm-volume
    configMap:
      name: log-mode-cm

scp log-mod-pod2.yaml root@192.168.56.100:/root

k apply -f log-mod-pod2.yaml
k get pod
k exec -it log-mod-pod2 -- ls /etc/config
k exec -it log-mod-pod2 -- ls /etc/config/LOG_LEVEL
  DEBUG



vi redis.conf
key1=docker
key2=kubernetes
key3=pending

k create configmap redis-config --from-file=redis.conf
k get cm
k describe cm redis-config

cat > redis-pod.yaml
apiVersion: v1
kind: Pod
metadata:
  name: redis-pod
spec:
  containers:
  - image: redis
    name: redis-container
    volumeMounts:
    - name: redis-volume
      mountPath: /opt/redis-config
  volumes:
  - name: redis-volume
    configMap:
      name: redis-config

k apply -f redis-pod.yaml



k create configmap my-cm --from-literal=k8skey=k8svalue \
  --dry-run=client -o yaml > my-cm.yaml

cat my-cm.yaml





# 실습

- 3개 오브젝트 필요 pod, configMap, svc
- Pod의 멀티컨테이너 정의하기!
  - 컨테이너통신: localhost:9000
- 외부연결 수행 192.168.56.101:31111
  서비스생성 (NodePort)
  - flaskapp-container
  - nginx-container(웹서버)
        ㄴ proxy거치도록
        ㄴ nginx웹서버를 버리고 nginx proxy로 재구성!!!
        ㄴ proxy로 구성한 nginx.conf를 만들어 컨테이너에 넣어줌
        ㄴ 해당 conf 가지고 있는 configMap 만들어서 Pod에 ref시키기
          (방법1 envFrom, 방법2-실습에 사용: mount-볼륨기법사용할거임 !! )
      => nginx proxy에 의해 flask 9000번으로

      flaskapp-conf.yaml > proxy_pass
        -> localhost:9000 보내기!


cat > flaskapp-conf.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: nginx-conf
data:
  nginx.conf: |-
    user  nginx;
    worker_processes  1;

    error_log  /var/log/nginx/error.log warn;
    pid        /var/run/nginx.pid;

    events {
        worker_connections  1024;
    }

    http {
        include       /etc/nginx/mime.types;
        default_type  application/octet-stream;

        sendfile        on;
        keepalive_timeout  65;

        upstream flaskapp {
            server 127.0.0.1:9000;
        }

        server {
            listen 80;

            location / {
                proxy_pass         http://flaskapp;
                proxy_redirect     off;
            }
        }
    }
 
k apply -f flaskapp-conf.yaml
k get cm


cat > flaskapp-pod.yaml
apiVersion: v1
kind: Pod
metadata:
  name: flaskapp-pod
  labels:
    app: flaskapp
spec:
  containers:
  - image: nginx:1.21-alpine
    name: proxy-container
    volumeMounts:
    - name: nginx-proxy-config
      mountPath: /etc/nginx/nginx.conf
      subPath: nginx.conf
    ports:
    - containerPort: 80
      protocol: TCP
  # - image: jnuho/py_flask:1.0
  - image: dbgurum/py_flask:1.0
    name: flaskapp-container
    ports:
    - containerPort: 9000
      protocol: TCP
  volumes:
  - name: nginx-proxy-config
    configMap:
      name: nginx-conf


k apply -f flaskapp-pod.yaml
k get cm,po,svc





cat > flaskapp-svc.yaml
apiVersion: v1
kind: Service
metadata:
  name: flaskapp-svc
spec:
  selector:
    app: flaskapp
  ports:
  - protocol: TCP
    port: 80
    targetPort: 80
    nodePort: 31111
  type: NodePort

k apply -f flaskapp-svc.yaml
k get cm,po,svc




# 시크릿이용

k create secret generic my-pwd --from-literal mypwd=k8spass#
k get secrets
k get secrets -o yaml
  - apiVersion: v1
    data:
      mypwd: azhzcGFzcyM=
echo azhzcGFzcyM= | base64 --decode
k delete secrets my-pwd



k create secret generic my-pwd --from-literal mypwd=k8spass# --dry-run=client -o yaml > my-pwd.yaml
  apiVersion: v1
  data:
    mypwd: azhzcGFzcyM=
  kind: Secret
  metadata:
    creationTimestamp: null
    name: my-pwd

k apply -f my-pwd.yaml

cat > secret-pod.yaml

apiVersion: v1
kind: Pod
metadata:
  name: secret-pod
spec:
  containers:
  - image: nginx:1.21-alpine
    name: sec-container
    envFrom:
    - secretRef:
        name: my-pwd

k apply -f secret-pod.yaml






# 파드에 마운트 기법으로 시크릿 넣기!
k create secret generic super-secret --from-literal=superpwd=k8spass#

k get secrets

k describe secrets super-secret

cat > super-pod.yaml
apiVersion: v1
kind: Pod
metadata:
  name: super-pod
spec:
  containers:
  - image: redis
    name: redis-container
    volumeMounts:
    - name: sec-volume
      mountPath: /secret
  volumes:
  - name: sec-volume
    secret:
      secretName: super-secret

k apply -f super-pod.yaml
k describe pod super-pod

k exec -it super-pod -- ls /secret
k exec -it super-pod -- cat /secret/superpwd





# 실습2

cat > cm-lab.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: cm-dev
data:
  SSH: 'false'
  User: kevin

k apply -f cm-lab.yaml

cat > sec-lab.yaml
apiVersion: v1
kind: Secret
metadata:
  name: sec-dev
data:
  Key: YnJhdm9teWxpZmU=

k apply -f sec-lab.yaml



cat > cm-sec-pod.yaml
apiVersion: v1
kind: Pod
metadata:
  name: cm-sec-pod
spec:
  containers:
  - name: cm-sec-container
    image: dbgurum/k8s-lab:initial
    envFrom:
    - configMapRef:
        name: cm-dev
    - secretRef:
        name: sec-dev

k apply -f cm-sec-pod.yaml


k exec -it cm-sec-pod -- env
HOSTNAME=cm-sec-pod
TERM=xterm
User=kevin
Key=bravomylife
SSH=false







# 마지막실습!

- Deployment 등 실습 (Deployment = ReplicaSet + Pod)
      ㄴ image
      ㄴ multi Pod 유지
      ㄴ rolling update -> 무중단 업데이트 -> rollback ->  이전 버전 (old version) 관리 = undo

- Ingress + secret + Deployment

- HPA Horizontal Pod Autoscaling
  -> 수평적 Pod 자동 확장기
             ㄴ 기준? 자원 상태 -> CPU utilization(활용률)

- Pilot project -> wordpress + mysql -> 9개 object -> semi MSA
        ㄴ Centos
        ㄴ GKE -> public IP










- Deployment 등 실습 (Deployment = ReplicaSet + Pod)

# CentOS -----
scp What\ is\ your\ motto.zip root@192.168.56.100:/root

mkdir dev_deploy && cd $_
unzip What\ is\ your\ motto.zip
vi Dockerfile
FROM nginx:1.21
RUN mkdir -p /usr/share/nginx/html/assets
RUN mkdir -p /usr/share/nginx/html/css
RUN mkdir -p /usr/share/nginx/html/js
COPY index.html /usr/share/nginx/html/index.html
COPY assets /usr/share/nginx/html/assets
COPY css /usr/share/nginx/html/css
COPY js /usr/share/nginx/html/js
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]

docker build -t jnuho/devweb:1.0 .
docker images | grep devweb
docker push jnuho/devweb:1.0


# GKE -------
#   Deployment(replicas=3), Service(LoadBalancer) 생성
#   GKE 제공 ExternalIP:8090 접속 결과 확인

cat > devweb.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: devweb-deployment
spec:
  selector: 
    matchLabels:
      app: my-devweb
  replicas: 3
  template:
    metadata:
      labels:
        app: my-devweb
    spec:
      containers:
      - name: devweb-container
        image: jnuho/devweb:1.0
        ports:
        - containerPort: 80
---
apiVersion: v1
kind: Service
metadata:
  name: devweb-svc
spec:
  selector:
    app: my-devweb
  ports:
  - port: 8090
    targetPort: 80
  type: LoadBalancer



kubectl apply -f devweb.yaml
kubectl get po,svc,deployment
  NAME                                     READY   STATUS    RESTARTS   AGE
  pod/devweb-deployment-64c477f7bc-49cmq   1/1     Running   0          3m53s
  pod/devweb-deployment-64c477f7bc-8qrmn   1/1     Running   0          3m53s
  pod/devweb-deployment-64c477f7bc-x8qq8   1/1     Running   0          3m53s

  NAME                 TYPE           CLUSTER-IP    EXTERNAL-IP     PORT(S)          AGE
  service/devweb-svc   LoadBalancer   10.48.12.92   34.64.252.135   8090:32565/TCP   3m53s
  service/kubernetes   ClusterIP      10.48.0.1     <none>          443/TCP          7d4h

  NAME                                READY   UP-TO-DATE   AVAILABLE   AGE
  deployment.apps/devweb-deployme

- 결과 확인: 34.64.252.135:8090

kubectl delete -f devweb.yaml









- Ingress + secret + Deployment


# Deployemnt 생성

- nginx / goapp deployment 생성

cat > nginx-go-deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx-container
        image: nginx:1.21
        ports:
        - containerPort: 80
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: goapp-deployment
  labels:
    app: goapp
spec:
  replicas: 3
  selector:
    matchLabels:
      app: goapp
  template:
    metadata:
      labels:
        app: goapp
    spec:
      containers:
      - name: goapp-container
        image: 본인ID/goapp:1.0
        ports:
        - containerPort: 9090

kubectl apply -f nginx-go-deployment.yaml
kubectl get deployment,po,svc

# Service 생성

cat > nginx-go.svc.yaml
apiVersion: v1
kind: Service
metadata:
  name: nginx-lb
spec:
  type: NodePort
  ports:
  - port: 80
    targetPort: 80
  selector:
    app: nginx
---
apiVersion: v1
kind: Service
metadata:
  name: goapp-lb
spec:
  type: NodePort
  ports:
  - port: 80
    targetPort: 9090
  selector:
    app: goapp

kubectl apply -f nginx-go.svc.yaml

# secret 생성!

kubectl get secrets
openssl genrsa -out server.key 2048

openssl req -new -x509 -key server.key -out server.cert -days 360 -subj /CN=nginx.k8s.com,goapp.k8s.com
kubectl create secret tls k8s-secret --cert=server.cert --key=server.key
kubectl get secrets
  NAME                  TYPE                                  DATA   AGE
  k8s-secret            kubernetes.io/tls                     2      2s

# ingress 생성!
  - nginx.k8s.com -> nginx-lb
  - goapp.k8s.com -> goapp-lb


질문) goapp 이미지이름 잘못기업 -> yaml 수정후 apply -f 해도되는지?
답변) 이미지 변경은 deploy에 구조가 변견되기때문에 update가 안되고, delete를 하고 다시 apply하셔야 됩니다.


cat > nginx-go-ingress.yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: nginx-goapp-ingress
spec:
  tls:
  - hosts:
    - nginx.k8s.com
    - goapp.k8s.com
    secretName: k8s-secret
  rules:
  - host: nginx.k8s.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: nginx-lb
            port:
              number: 80
  - host: goapp.k8s.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: goapp-lb
            port:
              number: 80

k apply -f nginx-go-ingress.yaml

# Ingress 조회

kubectl get ingress
kubectl get deploy,po,svc,ingress -o wide
  pod/go..deployment..-mj9iw
  pod/go..deployment..-zg4lc
  pod/go..deployment..-zz29f
  pod/nginx..deployment..-
  pod/nginx..deployment..-
  pod/nginx..deployment..-
  ...
  service/goapp-lb  .. 80:30120
  service/nginx-lb  .. 80:31446


# host 접근을 위한 추가 설정

vi /etc/hosts
192.168.56.100 k8s-master nginx.k8s.com goapp.k8s.com
192.168.56.101 k8s-node1 nginx.k8s.com goapp.k8s.com
192.168.56.102 k8s-node2 nginx.k8s.com goapp.k8s.com
192.168.56.103 k8s-node3 nginx.k8s.com goapp.k8s.com


# goapp 접근 테스트: 각 pod로 접근하는 것을 확인
curl goapp.k8s.com:30120
  hostname: goapp-deployment-7745ddfdb4-zg4lc
curl goapp.k8s.com:30120
  hostname: goapp-deployment-7745ddfdb4-dsads
curl goapp.k8s.com:30120
  hostname: goapp-deployment-7745ddfdb4-eiwoq


# nginx 접근 테스트: Deployment로 배포된 각 노드의 Pod를 식별하기 위해 index.html 생성 후 적용
# (master) -> node1
ssh k8s-node1
docker ps
  CONTAINER ID
  215fc... dbgurum/nginx ...

vi index.html
<h1> Great! kubernetes ingress 1. </h1>
docker cp index.html 215fc..:/usr/share/nginx/html/

# 다른 node2,3도 각 container ID로 진행
<h1> Great! kubernetes ingress 2. </h1>
<h1> Great! kubernetes ingress 3. </h1>


# 접근 확인
#   master 노드에서:
curl nginx.k8s.com:31446
  <h1> Great! kubernetes ingress 3. </h1>
curl nginx.k8s.com:31446
  <h1> Great! kubernetes ingress 2. </h1>
curl nginx.k8s.com:31446
  <h1> Great! kubernetes ingress 1. </h1>

# Traffic check: 각 노드에 iptraf-ng 설치 후 패킷 유입확인
#  node1:
#  node2:
yum -y install iptraf-ng

# 반복시도
#   master노드:
curl nginx.k8s.com:31446














# 마지막 실습!!!

- Pilot project -> wordpress + mysql -> 9개 object -> semi MSA

# 마지막 실습!!! - CentOs

cd LABs/
mkdir web-db && cd $_
cd LABs/web-db
scp * root@192.168.56.100:/root/web-db

cat > web-db-pv1.yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: web-db-pv1
  labels:
    type: local
spec:
  capacity:
    storage: 10Gi
  accessModes:
  - ReadWriteOnce
  hostPath:
    path: "/web-db-storage/pv01"


cat > web-db-pv2.yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: web-db-pv2
  labels:
    type: local
spec:
  capacity:
    storage: 10Gi
  accessModes:
  - ReadWriteOnce
  hostPath:
    path: "/web-db-storage/pv02"


kubectl apply -f web-db-pv1.yaml
kubectl apply -f web-db-pv2.yaml

kubectl get pv
  NAME         CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS      CLAIM   STORAGECLASS   REASON   AGE
  web-db-pv1   10Gi       RWO            Retain           Available                                   5s
  web-db-pv2   10Gi       RWO            Retain           Available                                   3s


-- 생성한 PV를 사용할 PVC 생성: 용량과 접근권한을 비교하여 자동 지정
cat > wordpress-pvc.yaml
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: wordpress-pvc
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 5Gi

cat > mysql-pvc.yaml
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: mysql-pvc
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 5Gi


kubectl apply -f wordpress-pvc.yaml
kubectl apply -f mysql-pvc.yaml

kubectl get pv,pvc


-- secret object 생성: mysql 암호 저장용

kubectl create secret generic mysql-pwd --from-literal=password=Passw0rD
  secret/mysql-pwd created
kubectl describe secret mysql-pwd
  Name: mysql-pwd
  Namespace: default
  Labels: <none>
  Annotations: <none>
  Type: Opaque
  Data
  ====
  password: 8 bytes



-- mysql pod를 배포할 Deployment 생성

cat > mysql-deploy.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: mysql
  labels:
    app: mysql
spec:
  replicas: 1
  selector:
    matchLabels:
      app: mysql
    template:
      metadata:
        labels:
          app: mysql
      spec:
        containers:
          - image: mysql:5.6
            name: mysql
            env:
              - name: MYSQL_ROOT_PASSWORD
                valueFrom:
                  secretKeyRef:
                    name: mysql-pwd
                    key: password
              - name: MYSQL_DATABASE
                value: kube-db
              - name: MYSQL_USER
                value: kube-user
              - name: MYSQL_ROOT_HOST
                value: '%'
              - name: MYSQL_PASSWORD
                value: Passw0rD
            ports:
              - containerPort: 3306
                name: mysql
            volumeMounts:
              - name: mysql-persistent-storage
                mountPath: /var/lib/mysql
        volumes:
          - name: mysql-persistent-storage
            persistentVolumeClaim:
              claimName: mysql-pvc


kubectl apply -f mysql-deploy.yaml
kubectl get deploy,pods -o wide



-- mysql service 생성

cat > mysql-service.yaml
apiVersion: v1
kind: Service
metadata:
  name: mysql
  labels:
    app: mysql
spec:
  type: ClusterIP
  ports:
  - port: 3306
  selector:
    app: mysql

kubectl apply -f mysql-service.yaml
kubectl get svc -o wide


-- wordpress를 배포할 Deployment 생성

cat > wordpress-deploy.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: wordpress
  labels:
    app: wordpress
spec:
  replicas: 1
  selector:
    matchLabels:
      app: wordpress
  template:
    metadata:
      labels:
        app: wordpress
    spec:
      containers:
        - image: wordpress
          name: wordpress
          env:
          - name: WORDPRESS_DB_HOST
            value: mysql:3306
          - name: WORDPRESS_DB_NAME
            value: kube-db
          - name: WORDPRESS_DB_USER
            value: kube-user
          - name: WORDPRESS_DB_PASSWORD
            value: Passw0rD
          ports:
            - containerPort: 80
              name: wordpress
          volumeMounts:
            - name: wordpress-persistent-storage
              mountPath: /var/www/html
      volumes:
        - name: wordpress-persistent-storage
          persistentVolumeClaim:
            claimName: wordpress-pvc

kubectl apply -f wordpress-deploy.yaml
  deployment.apps/wordpress created
kubectl get deploy,pods -o wide



-- wordpress service 생성

cat > wordpress-service.yaml
apiVersion: v1
kind: Service
metadata:
  name: wordpress
  labels:
    app: wordpress
spec:
  type: NodePort
  ports:
    - port: 80
      targetPort: 80
      protocol: TCP
  selector:
    app: wordpress

kubectl apply -f wordpress-service.yaml
kubectl get svc -o wide
kubectl get pod -o wide
kubectl get pod/wordpress-778cb9d9fb-xk44j -o wide
kubectl describe pod/wordpress-778cb9d9fb-xk44j
kubectl get po,svc,deploy

# node1에서:
netstat -nlp | grep 30548
  tcp 0 0 0.0.0.0:30548 0.0.0.0:* LISTEN 3407/kube-proxy

# 마스터에서
kubectl exec -it mysql-7b54d57fcc-v9rzd -- bash
env | grep MYSQL
  MYSQL_PASSWORD=Passw0rD
  MYSQL_DATABASE=kube-db
  MYSQL_ROOT_PASSWORD=Passw0rD
  MYSQL_MAJOR=5.6
  MYSQL_USER=kube-user
  MYSQL_VERSION=5.6.50-1debian9
  MYSQL_ROOT_HOST=%

mysql -uroot -p
  Enter Password: Passw0rD
show databases;
use kube-db
show tables;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| kube-db            |
| mysql              |
| performance_schema |
+--------------------+


[wordpress Pod가 생성된 node의 IP와 NodePort가 지정한 port로 접속]

-- wordpress 사이트 install 이후 테이블 생성됨
192.168.56.100:30548

show tables;
+-----------------------+
| Tables_in_kube-db     |
+-----------------------+
| wp_commentmeta        |
| wp_comments           |
| wp_links              |
| wp_options            |
| wp_postmeta           |
| wp_posts              |
| wp_term_relationships |
| wp_term_taxonomy      |
| wp_termmeta           |
| wp_terms              |
| wp_usermeta           |
| wp_users              |
+-----------------------+

select * from wp_users;

# node1:
cd /
ls
  web-db-storage
cd web-db-storage
ls
  pv01 pv02










# 마지막 실습!!! - GKE

cd LABs/
mkdir web-db-GKE && cd $_
cd LABs/web-db




# mysql을 위한 서비스 구성: mysql-svc.yaml

cat > mysql-svc.yaml
apiVersion: v1
kind: Service
metadata:
  name: wordpress-mysql
  labels:
    app: wordpress
spec:
  ports:
  - port: 3306
  selector:
    app: wordpress
    tier: mysql
  clusterIP: None

kubectl apply -f mysql-svc.yaml
kubectl get svc -o wide


# mysql을 위한 서비스 구성: mysql-pvc.yaml
-- PVC 생성: 용량과 접근권한을 비교하여 자동 지정

cat > mysql-pvc.yaml
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: mysql-pv-claim
  labels:
    app: wordpress
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 20Gi

kubectl apply -f mysql-pvc.yaml
kubectl get pvc


# mysql Pod를 위한 Deployment 구성: mysql-deploy.yaml

-- mysql pod를 배포할 Deployment 생성

cat > mysql-deploy.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: wordpress-mysql
  labels:
    app: wordpress
spec:
  selector:
    matchLabels:
      app: wordpress
      tier: mysql
  strategy:
    type: Recreate
  template:
    metadata:
      labels:
        app: wordpress
        tier: mysql
    spec:
      containers:
        - image: mysql:5.6
          name: mysql
          env:
            - name: MYSQL_ROOT_PASSWORD
              valueFrom:
                secretKeyRef:
                  name: mysql-pass
                  key: password
          ports:
            - containerPort: 3306
              name: mysql
          volumeMounts:
            - name: mysql-persistent-storage
              mountPath: /var/lib/mysql
      volumes:
        - name: mysql-persistent-storage
          persistentVolumeClaim:
            claimName: mysql-pv-claim


kubectl apply -f mysql-deploy.yaml
kubectl get deploy,pods -o wide

# 외부 서비스를 위한 LoadBalancer 구성

cat > wordpress-service.yaml
apiVersion: v1
kind: Service
metadata:
  name: wordpress
  labels:
    app: wordpress
spec:
  ports:
    - port: 80
  selector:
    app: wordpress
    tier: frontend
  type: LoadBalancer

kubectl apply -f wordpress-service.yaml
kubectl get svc -o wide
kubectl get pod -o wide
kubectl get pod/wordpress-778cb9d9fb-xk44j -o wide
kubectl describe pod/wordpress-778cb9d9fb-xk44j
kubectl get po,svc,deploy



# mysql 관리자 암호 저장을 위한 secret 생성

kubectl create secret generic mysql-pass --from-literal=password=Passw0rD
  secret/mysql-pwd created
kubectl describe secret mysql-pass
  Name: mysql-pass
  Namespace: default
  Labels: <none>
  Annotations: <none>
  Type: Opaque
  Data
  ====
  password: 8 bytes



# wordpress를 위한 PVC 구성

cat > wordpress-pvc.yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: wp-pv-claim
  labels:
    app: wordpress
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 20Gi

kubectl apply -f wordpress-pvc.yaml


# WordPress Pod를 위한 Deployment 구성


cat > wordpress-deploy.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: wordpress
  labels:
    app: wordpress
spec:
  selector:
    matchLabels:
      app: wordpress
      tier: frontend
  template:
    metadata:
      labels:
        app: wordpress
        tier: frontend
    spec:
      containers:
        - image: wordpress:4.8-apache
          name: wordpress
          env:
          - name: WORDPRESS_DB_HOST
            value: wordpress-mysql
          - name: WORDPRESS_DB_PASSWORD
            valueFrom:
              secretKeyRef:
                name: mysql-pass
                key: password
          ports:
            - containerPort: 80
              name: wordpress
          volumeMounts:
            - name: wordpress-persistent-storage
              mountPath: /var/www/html
      volumes:
        - name: wordpress-persistent-storage
          persistentVolumeClaim:
            claimName: wp-pv-claim

kubectl apply -f wordpress-deploy.yaml
  deployment.apps/wordpress created
kubectl get deploy,pods -o wide



# 정보조회
kubectl get deploy,pod,svc -o wide

# Wordpress WEB 접속:
http://34.64.252.135:80

# mysql 접속
kubectl exec -it wordpress-mysql-6c479567b-hfghj -- bash
env | grep MYSQL
  MYSQL_ROOT_PASSWORD=Passw0rD
  MYSQL_MAJOR=5.6
  MYSQL_VERSION=5.6.51-1debian9

mysql -uroot -p
  Enter Password: Passw0rD
show databases;
+---------------------+
| Database            |
+---------------------+
| information_schema  |
| #mysql50#lost+found |
| mysql               |
| performance_schema  |
| wordpress           |
+---------------------+

use wordpress
show tables;
Empty set (0.00 sec)

# 워드프레스 사이트 install 후 
+--------------------+
| Database           |
+--------------------+
| information_schema |
| kube-db            |
| mysql              |
| performance_schema |
+--------------------+


# wordpress 사이트 install 이후 테이블 생성됨
192.168.56.100:30548

mysql> show tables;
+-----------------------+
| Tables_in_wordpress   |
+-----------------------+
| wp_commentmeta        |
| wp_comments           |
| wp_links              |
| wp_options            |
| wp_postmeta           |
| wp_posts              |
| wp_term_relationships |
| wp_term_taxonomy      |
| wp_termmeta           |
| wp_terms              |
| wp_usermeta           |
| wp_users              |
+-----------------------+
12 rows in set (0.00 sec)

select * from wp_users;
+----+------------+------------------------------------+---------------+-------------------+----------+---------------------+---------------------+-------------+--------------+
| ID | user_login | user_pass                          | user_nicename | user_email        | user_url | user_registered     | user_activation_key | user_status | display_name |
+----+------------+------------------------------------+---------------+-------------------+----------+---------------------+---------------------+-------------+--------------+
|  1 | jnuho      | $P$Bq9uB4jTJz9G10ZoTeKh4sTf27wr310 | jnuho         | jnuho@outlook.com |          | 2022-06-26 08:39:23 |                     |           0 | jnuho        |
+----+------------+------------------------------------+---------------+-------------------+----------+---------------------+---------------------+-------------+--------------+
1 row in set (0.00 sec)


# PVC 할당 공간 확인

df -h
  Filesystem      Size  Used Avail Use% Mounted on
  overlay          95G  4.3G   90G   5% /
  tmpfs            64M     0   64M   0% /dev
  tmpfs           2.0G     0  2.0G   0% /sys/fs/cgroup
  shm              64M     0   64M   0% /dev/shm
  /dev/sda1        95G  4.3G   90G   5% /etc/hosts
  /dev/sdb         20G  235M   20G   2% /var/lib/mysql ★★★★★
  tmpfs           2.8G   12K  2.8G   1% /run/secrets/kubernetes.io/serviceaccount
  tmpfs           2.0G     0  2.0G   0% /proc/acpi
  tmpfs           2.0G     0  2.0G   0% /proc/scsi
  tmpfs           2.0G     0  2.0G   0% /sys/firmware


# 글등록 후

mysql> use wordpress;
mysql> select * from wp_posts;