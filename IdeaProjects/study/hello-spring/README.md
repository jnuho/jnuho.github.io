
- Spring Initializr
  - Gradle Project
  - Spring Boot 2.x.x
  - Dependencies: Spring Web, Thymeleaf 
    - spring-boot-devtools: html 컴파일만 해주면 서버 재시작 없이 view파일 변경 가능
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
# resources/static/index.html
# resources/static/hello-static.html

# 스프링컨테이너 컨트롤러에 맵핑 없으면 static 페이지를 찾음
localhost:8080/
localhost:8080/index.html
localhost:8080/hello-static.html
```

- MVC와 템플릿 엔진
```
# resources/templates/hello.html
# resources/templates/hello-template.html
# html파일 경로 직접 열어도 thymeleaf 동작함
http://localhost:8080/hello-mvc?name=abc
```

- API
  - Http의 Body에 문자내용을 직접반환
  - 기본 문자처리 `StringHttpMessageConverter`
  - 기본 객체처리 `MappingJackson2HttpMessageConverter`
  - byte 처리 등등 기타 여러 HttpMessageConverter가 기본으로 등록되어 있음
```
http://localhost:8080/hello-api?name=abc
컨트롤러 -> @ResponseBody
-> viewResolver 대신 HttpMessageConverter (JsonConverter / StringConverter)
```

- 비즈니스 요구사항
  - 컨트롤러: 웹 MVC의 컨트롤러 역할
  - 서비스: 핵심 비즈니스 로직 구현
  - 리포지토리: DB접근, 도메인 객체를 DB에 저장, 관리
  - 도메인: ex) 회원, 주문, 쿠폰 등 비즈니스 도메인 객체

- 테스트
  - junit테스트코드
  - `graldew test` 전체 테스트

- 서버재시작없이 반영
```
# build.gradle
```

- `@Controller` 스프링컨테이너 시작시 빈등록 관리

```java
@Controller
public class MemberController {

  private final MemberService memberService;

  // MemberSerivce 빈 찾을 수 없음 에러!
  // -> MemberServicex를 @Service 등록해야 함
  @Autowired
  public MemberController(MemberService memberService) {
    this.memberService = memberService;
  }
}
```
- 컴포넌트 스캔은 해당 패키지만 스캔

- 빈등록 @Component
- 빈등록 @Configuration 직접생성


1. 생성자주입 (*권장)
2. 필드주입
3. setter주입


- Mapping 우선순위
  - 맵핑찾음
  - static파일

- DB접근

```properties
# application.properties
spring.datasource.url=jdbc:mysql://localhost:3306/testdb
spring.datasource.driver-class-name=com.mysql.jdbc.Driver
```
