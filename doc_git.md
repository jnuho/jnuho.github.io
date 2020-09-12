
# 로컬 환경에서 브랜치 변경하기
- 로컬에 master 브랜치가 있고, 원격(리모트) origin/master에 연결된 상태
- 로컬에 develop 브랜치를 만들고, 원격(리모트) origin/develop 브랜치에 연결하기.

``` bash
# 로컬 master 브랜치 확인
git branch
  * master
# 원격(remote) 브랜치 확인 - remote-tracking 로컬 브랜치로 확인
git branch -r
  origin/HEAD -> origin/master
  origin/master

git fetch --all

# 로컬 develop 브랜치 생성 및 원격 origin/develop에 연결
git checkout --track origin/develop

# 원격 develop 브랜치에서 소스 받기(pull)
# git pull REMOTE-NAME BRANCH-NAME
git pull origin develop

git branch
  * develop
    master
git branch -r
  origin/HEAD -> origin/master
  origin/develop
  origin/master
```


- 현재 안섭매니저 master로 푸시된 부분을 develop으로 merge시켜서 develop을 최신화
```bash
git checkout develop
git merge master
git push -u origin develop
```

- master branch는 안섭매니저님 푸시하기 전상태도 원복
```bash
git revert --no-commit d8b12c90ba..HEAD
# d8b12c90ba커밋이 마지막 커밋이 되도록 원복
git commit -m 'd8b12c90ba이후 부터 HEAD 까지의 커밋을 롤백'
git push -u origin master
```
