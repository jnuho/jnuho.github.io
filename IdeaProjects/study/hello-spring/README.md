
- Spring Initializr
  - Gradle Project
  - Spring Boot 2.x.x
  - Dependencies: Spring Web, Thymeleaf 
    - spring-boot-devtools: html컴파일만 해주면 서버 재시작 없이 view파일 변경 가능
  - 다운로드 unzip > Open or import "build.gradle" > Open as project

- 배포파일 빌드(gradle) 및 실행

```
# 배포를 위한 빌드 
cd study/hello-spring
.gradlew build
.gradlew clean build
cd build/libs
ls -arlth
    hello-spring-0.0.1-SNAPSHOPT.jar
scp hello-spring-0.0.1-SNAPSHOPT.jar {remotehost}:/home/ec2-user

# 서버 실행
java -jar hello-spring-0.0.1-SNAPSHOPT.jar
```


- 정적컨텐츠
```
# 스프링 컨테이너 > resources/static 폴더 순서로 반환
# resources/static/hello-static.html
localhost:8080/
localhost:8080/index.html
localhost:8080/hello-static.html
```

- MVC와 템플릿 엔진
```
# resources/templates/hello.html
# resources/templates/hello-template.html
http://localhost:8080/hello-mvc?name=abc
```

- API
    - 기본 문자처리 `StringHttpMessageConverter`
    - 기본 객체처리 `MappingJackson2HttpMessageConverter`
    - byte 처리 등등 기타 여러 HttpMessageConverter가 기본으로 등록되어 있음
```
http://localhost:8080/hello-api?name=abc
컨트롤러 -> @ResponseBody
-> viewResolver 대신 HttpMessageConverter (JsonConverter / StringConverter)
```

- 서버재시작없이 반영
```
# build.gradle
```