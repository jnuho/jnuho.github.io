
- [Setting Up](#setting-up)
- [Getting started with Amazon EKS](#Getting-started-with-Amazon-EKS)
  - [eksctl](#eksctl)
  - [AWS Management Console and AWS CLI](#AWS-Management-Console-and-AWS-CLI)
- [Add-ons](#add-ons)


# Setting Up

1. aws-cli
  - create IAM user and generate access key (download .csv)
  - (optional) get security token

https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html

```sh
curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
unzip awscliv2.zip
sudo ./aws/install

# Configure Access key, Secret key, region, and etc
aws configure
aws configure --profile [PROFILE_NAME]

# To verify the user identity
aws sts get-caller-identity
```

2. eksctl

https://eksctl.io/installation/

- install-eksctl.sh

```sh
# for ARM systems, set ARCH to: `arm64`, `armv6` or `armv7`
ARCH=amd64
PLATFORM=$(uname -s)_$ARCH

curl -sLO "https://github.com/eksctl-io/eksctl/releases/latest/download/eksctl_$PLATFORM.tar.gz"

# (Optional) Verify checksum
curl -sL "https://github.com/eksctl-io/eksctl/releases/latest/download/eksctl_checksums.txt" | grep $PLATFORM | sha256sum --check

tar -xzf eksctl_$PLATFORM.tar.gz -C /tmp && rm eksctl_$PLATFORM.tar.gz

sudo mv /tmp/eksctl /usr/local/bin
```

3. kubectl


- https://docs.aws.amazon.com/eks/latest/userguide/install-kubectl.html

```sh
curl -O https://s3.us-west-2.amazonaws.com/amazon-eks/1.28.3/2023-11-14/bin/linux/amd64/kubectl
chmod +x ./kubectl
mkdir -p $HOME/bin && cp ./kubectl $HOME/bin/kubectl && export PATH=$HOME/bin:$PATH
echo 'export PATH=$HOME/bin:$PATH' >> ~/.bashrc
kubectl version --client
```

- TODO kubectl 클러스터와 통신 구성
  - After cluster creation

```sh
aws sts get-caller-identity

aws eks update-kubeconfig --region ap-northeast-2 --name testcluster-001
```



# Getting started with Amazon EKS

- Create EKS cluster using one of the followings
  - eksctl command
  - AWS Management Console and AWS CLI

## eksctl

- eksctl uses cloudformation

```sh
# The cluster configuration is created in ~/.kube/config
eksctl create cluster --name testcluster-001 --region ap-northeast-2
eksctl create cluster --name testcluster-001 --region ap-northeast-2 --profile teraform

eksctl delete cluster --name testcluster-001 --region ap-northeast-2
```

- kubectl 클러스터 통신 설정
  - Configures kubectl so that you can connect to an Amazon EKS cluster.
  - configuration file is created in `~/.kube/config`

```sh
aws sts get-caller-identity
aws eks update-kubeconfig --region ap-northeast-2 --name testcluster-001
aws eks update-kubeconfig --region ap-northeast-2 --name testcluster-001 --profile kams

kubectl get no -o wide
kubectl get pods -A -o wide
```


## AWS Management Console and AWS CLI

https://docs.aws.amazon.com/eks/latest/userguide/getting-started-console.html

1. Create an Amazon VPC
  - with public and private subnets that meets Amazon EKS requirements

```sh
aws cloudformation create-stack \
  --region region-code \
  --stack-name my-eks-vpc-stack \
  --template-url https://s3.us-west-2.amazonaws.com/amazon-eks/cloudformation/2020-10-29/amazon-eks-vpc-private-subnets.yaml
```

2. EKS cluster IAM Role creation

https://docs.aws.amazon.com/ko_kr/eks/latest/userguide/service_IAM_role.html

```sh
cat > eks-cluster-role-trust-policy.json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "eks.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}

# Create the role
aws iam create-role \
  --role-name myAmazonEKSClusterRole \
  --assume-role-policy-document file://"eks-cluster-role-trust-policy.json"

# Attach the required Amazon EKS managed IAM policy to the role
aws iam attach-role-policy \
  --policy-arn arn:aws:iam::aws:policy/AmazonEKSClusterPolicy \
  --role-name myAmazonEKSClusterRole
```

3. EKS cluser creation
  - Go to EKS AWS console
  - Add cluster -> Create
    - Enter name
    - Choose Cluster Service Role
    - Specify networking: VPC
    - Configure observability: Next
    - Add-ons: https://docs.aws.amazon.com/eks/latest/userguide/eks-add-ons.html


check EKS addon such as VPC-cni, CoreDNS Version, and kube-proxy with below URL

https://docs.aws.amazon.com/ko_kr/eks/latest/userguide/managing-vpc-cni.html
https://docs.aws.amazon.com/ko_kr/eks/latest/userguide/managing-coredns.html
https://docs.aws.amazon.com/ko_kr/eks/latest/userguide/managing-kube-proxy.html

4. EKS kubeconfig setting

```sh
aws eks update-kubeconfig --region ap-northeast-2 --name testcluster-001
kubectl get svc
```

5. Create nodes

```sh
cat > node-role-trust-policy.json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "ec2.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}


# Create the node IAM role
aws iam create-role \
  --role-name myAmazonEKSNodeRole \
  --assume-role-policy-document file://"node-role-trust-policy.json"

# Attach the required managed IAM policies to the role
aws iam attach-role-policy \
  --policy-arn arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy \
  --role-name myAmazonEKSNodeRole
aws iam attach-role-policy \
  --policy-arn arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly \
  --role-name myAmazonEKSNodeRole
aws iam attach-role-policy \
  --policy-arn arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy \
  --role-name myAmazonEKSNodeRole
```

- Go to EKS console > cluster > Compute tab
  - Add Node Group: write name and choose Node IAM role name

3. EKS cluster oidc registration

4. EKS cluster node IAM Role creation

https://docs.aws.amazon.com/ko_kr/eks/latest/userguide/create-node-role.html


# Add-ons





