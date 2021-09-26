
```sh
sudo passwd root

sudo yum repolist
sudo yum install epel-release
sudo yum list installed | grep epel
```

- Nodejs 설치

```sh
# nvm 설치
curl -o- https://raw.githubusercontent.com/creationix/nvm/v0.33.11/install.sh | bash
source ~/.bashrc
nvm --version

# nodejs 설치
nvm install --lts

# 설치확인
npm -v
node -v

node -e "console.log('Running Node.js ' + process.version)"
Running Node.js v14.17.6
```

- Git 배포 설정

```sh
# 코드실행에 필요한 의존성 패키지
sudo yum install curl-devel expat-devel gettext-devel openssl-devel zlib-devel

# 소스코드 배포할 디렉터리
cd /var
mkdir www

# 코드관리는 일반유저로
sudo chown centos www
cd /var/www
git clone https://github.com/deopard/aws-exercise-a.git
cd aws-exercise-a
# node package manager : nodej의존성 패키지 관리
# package.json 명시 라이브러리 설치
npm install
```

- 웹서버 (WEB)
- 웹 어플리케이션 서버 (WAS)