
- Go1.16버전 부터 Go 모듈 사용이 기본이 됨
  - 모듈은 패키지 종속성 관리 등 패키지 관리 시스템 역할
  - `Ver < 1.16`: 모든 Go코드는 $GOPATH/src 아래
  - `Ver >= 1.16`: 모든 Go코드는 Go모듈 아래
    - `go mod init MODULE_NAME`으로 Go모듈 생성
    - 모듈이름은 유니크해야 함; GitHub 저장소 주소를 활용하거나 URL을 모듈 이름으로 사용
