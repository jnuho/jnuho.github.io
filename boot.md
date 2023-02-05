
- 스프링 부트 스타터(https://start.spring.io/)
  - Spring boot 2.7.8
  - Group : `jpabook`
  - Artifact : `jpashop`
  - 디펜던시 추가 :  Spring Web, Thymeleaf, Spring Data JPA, H2 Database, Lombok
  - 프로젝트로드후 Enable Annotation Processing



- 인텔리제이 shortcut
  - option + command + V : Create return variable
  - command + shift + T : Create new test
  - control + shift + R : Run Tests
  - Editor > Live Template > custom ('tdd')

```java
@Test
public void $NAME$() throws Exception {
	//given
	$END$
  //when
  //then
}
```
 

- 라이브러리 확인
  - cli 또는 인텔리제이 Gradle 탭

```shell
cd jpashop
./gradlew dependencies
```

- HTML 수정 후 서비 재시작 필요 X
  - 라이브러리 : spring-boot-devtools
  - build.gradle > `implementation 'org.springframework.boot:spring-boot-devtools'`
  - build > recompile xxx.html
 
- h2 실행
- application.properties -> application.yaml
  - 스프링부트 매뉴얼 'spring.io/projects/spring-boot#learn'


- 테스트
  - @Tranactional 은 테스트의 경우 롤백
  - @Rollback(false) 옵션으로 롤백 막을 수 있긴 함

- jar 동작 확인

```shell
cd japshop

# 빌드실패 (tools.jar 등 jdk에 대한 에러)
# https://stackoverflow.com/a/24657630
# macbook에서 default jdk 변경
/usr/libexec/java_home -V        
# Matching Java Virtual Machines (2):
# 1.8.131.11 (x86_64) "Oracle Corporation" - "Java" /Library/Internet Plug-Ins/JavaAppletPlugin.plugin/Contents/Home
# 1.8.0_131 (x86_64) "Oracle Corporation" - "Java SE 8" /Library/Java/JavaVirtualMachines/jdk1.8.0_131.jdk/Contents/Home
# 또는 ~/.bash_profile 에 아래 내용 저장 후 source ~/.bash_profile
export JAVA_HOME=`/usr/libexec/java_home -v 1.8.0_131`

./gradlew clean build
cd build/libs
ls 
# jpashop-0.0.1-SNAPSHOT-plain.jar        jpashop-0.0.1-SNAPSHOT.jar
# 스프링 시작
java -jar jpashop-0.0.1-SNAPSHOT.jar

```

- 쿼리 파라미터 로그 남기기
  - org.hibernate.type: trace
  - 외부 라이브러리 추가 (spring-boot-data-source-decorator)
  - build.gradle > `implementation 'com.github.gavlyukovskiy:p6spy-spring-boot-starter:1.5.6'`
  - 운영 환경에서는 성능 저하 가능 (개발환경에서만 권장)

