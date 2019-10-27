# developMental
developer forum / community application created with java spring framework

# Spring environment setup

## Set encoding(UTF-8)
Preferences - General - Editors - Spelling
Workspace
JSON Files - UTF-8
Web - css, html, jsp
XML Files

## Add apache tomcat server v8.5
check "serve modules without publishing"

## Download sts and maven
1. Spring Tool Suite IDE (STS 3.9.10)
2. Apache Maven (apache-maven-3.6.2-bin.zip)
```
https://spring.io/tools/sts/all
http://maven.apache.org/
```

## maven : setup directory for storing .jar files
In order to download spring ```.jar``` files from MAVEN repositories specified in ```pom.xml``` into C:\maven\repository

## Create Spring Legacy Project
Spring MVC Project as 'template', specify top-level package (e.g. com.kh.spring); the last package name i.e. 'spring' becomes root context (http://localhost:9090/spring/)

## maven
```C:\Program Files\apache-maven-3.6.2\conf\settings.xml```
```xml
<settings xmlns="http://maven.apache.org/SETTINGS/1.0.0"
          xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
          xsi:schemaLocation="http://maven.apache.org/SETTINGS/1.0.0 http://maven.apache.org/xsd/settings-1.0.0.xsd">
  <!-- localRepository
   | The path to the local repository maven will use to store artifacts.
   |
   | Default: ${user.home}/.m2/repository
  -->
  <localRepository>C:\maven\repository</localRepository>
```

Go to preferences - maven -user settings - change
    from ```C:\Program Files\apache-maven-3.6.2\conf\settings.xml```
    to ```C:\maven\repository from C:\Users\foo\.m2\repository```

- \src\main\java - java srouce files
- \src\test\java - source files for test
- src.main.webapp.resrouces - js, css files
- serverlet-context.xml - .xml for servlet
- root-context.xml - .xml for whole project
- Note the view folder is inside the WEB-INF folder
- pom.xml - specifying dependencies (mvnrepository url) allows sts to automatically download .jar files into Maven Dependencies(C:\maven\repository)

## UTF-8 inside .jsp
```jsp
<%@ page language="java" contentType="text/html; charset=UTF-8" page Encoding="UTF-8"%> 
```

## pom.xml
```xml
<properties>
  <java-version>1.8</java-version>
  <org.springframework-version>5.0.6.RELEASE</org.springframework-version>
  <org.aspectj-version>1.6.10</org.aspectj-version>
  <org.slf4j-version>1.6.6</org.slf4j-version>
</properties>

<!-- Servlet -->
<dependency>
  <groupId>javax.servlet</groupId>
  <artifactId>javax.servlet-api</artifactId>
  <version>3.1.0</version>
  <scope>provided</scope>
</dependency>
<dependency>
  <groupId>javax.servlet.jsp</groupId>
  <artifactId>javax.servlet.jsp-api</artifactId>
  <version>2.3.1</version>
  <scope>provided</scope>
</dependency>

<!-- DB설정 jar dependency -->
<!-- https://mvnrepository.com/artifact/org.mybatis/mybatis -->
<dependency>
  <groupId>org.mybatis</groupId>
  <artifactId>mybatis</artifactId>
  <version>3.4.6</version>
</dependency>
  
<!-- https://mvnrepository.com/artifact/org.mybatis/mybatis-spring -->
<dependency>
  <groupId>org.mybatis</groupId>
  <artifactId>mybatis-spring</artifactId>
  <version>1.3.2</version>
</dependency>
<!-- /mybatis 적용 파일 끝 -->

<!-- common-dbcp DB연결을 위한 connection-pool 라이브러리 -->
<!-- https://mvnrepository.com/artifact/commons-dbcp/commons-dbcp -->
<dependency>
  <groupId>commons-dbcp</groupId>
  <artifactId>commons-dbcp</artifactId>
  <version>1.4</version>
</dependency>

<!-- https://mvnrepository.com/artifact/org.springframework/spring-jdbc -->
<dependency>
<groupId>org.springframework</groupId>
<artifactId>spring-jdbc</artifactId>
<version>${org.springframework-version}</version>
</dependency>

<!-- spring-loaded 설정하기 : tomcat 서버 재시작 안해도 자동 load 되도록-->
<!-- https://mvnrepository.com/artifact/org.springframework/springloaded -->
<dependency>
  <groupId>org.springframework</groupId>
  <artifactId>springloaded</artifactId>
  <version>1.2.8.RELEASE</version>
  <scope>provided</scope>
</dependency>
  
<!-- lombok library받기 -->
<!-- https://mvnrepository.com/artifact/org.projectlombok/lombok -->
<dependency>
  <groupId>org.projectlombok</groupId>
  <artifactId>lombok</artifactId>
  <version>1.18.10</version>
  <scope>provided</scope>
</dependency>

<!-- Spring Security 스프링 비밀번호 암호화 적용하기 feat security (search: spring-security)-->
<!-- https://mvnrepository.com/artifact/org.springframework.security/spring-security-core -->
<dependency>
  <groupId>org.springframework.security</groupId>
  <artifactId>spring-security-core</artifactId>
  <version>${org.springframework-version}</version>
</dependency>
<!-- https://mvnrepository.com/artifact/org.springframework.security/spring-security-web -->
<dependency>
  <groupId>org.springframework.security</groupId>
  <artifactId>spring-security-web</artifactId>
    <version>${org.springframework-version}</version>
</dependency>
<!-- https://mvnrepository.com/artifact/org.springframework.security/spring-security-config -->
<dependency>
  <groupId>org.springframework.security</groupId>
  <artifactId>spring-security-config</artifactId>
  <version>${org.springframework-version}</version>
</dependency>

<!-- Test -->
<plugin>
  <groupId>org.apache.maven.plugins</groupId>
  <artifactId>maven-compiler-plugin</artifactId>
  <version>3.6.1</version>
  <configuration>
    <source>1.8</source>
    <target>1.8</target>
    <compilerArgument>-Xlint:all</compilerArgument>
    <showWarnings>true</showWarnings>
    <showDeprecation>true</showDeprecation>
  </configuration>
</plugin>
```

## Project Facets
```
Dynamic Web Module 3.1
java 1.8
```

## Build path
```
project r_click - congigure build path - java 1.8, apache tomcat, maven
      Deployment Assembly - add - java build path entries - Maven dependency
project r_click - maven - update project
```

## xml catalog
preferences - xml - xml catalog - Add
```
http://mybatis.org/dtd/mybatis-3-config.dtd
-//mybatis.org//DTD Config 3.0//EN

http://mybatis.org/dtd/mybatis-3-mapper.dtd
-//mybatis.org//DTD Mapper 3.0//EN
```

## 1.lombok 2.springreloaded 3.spring-security
1. lombok -> vo 자동 생성 : mvnrepository lombok
```
* C:\maven\repository\org\projectlombok\lombok\1.18.10
* open powershell (shitft+r_click)> java -jar .\lombok-1.18.10.jar
  - specify location(sts.exe)
```

2. tomcat 재시작 안해도 자동 reload되는 dependency 추가 mvnrepository-> springloaded
```xml
<!-- pom.xml -->
<!-- https://mvnrepository.com/artifact/org.springframework/springloaded
  Maven Dependencies에 .jar파일 받아짐
  Tomcat server double click : Module Auto Reload 해제 uncheck
 > module 탭에서 edit web module: auto reload (uncheck)
 > overview탭에서 publishing> Automatically publish when resources change
 > overview탭에서 open launch configuarton> arguments
 -javaagent:C:\maven\repository\org\springframework\springloaded\1.2.8.RELEASE\springloaded-1.2.8.RELEASE.jar -noverify
-->
```

3. 암호화 처리 spring security
* pom.xml : mvnrepository spring-security &lt;dependency&gt; 등록 - .jar파일 다운로드(web/config/core 3개)
```xml
<!-- https://mvnrepository.com/artifact/org.springframework.security/spring-security-core -->
<!-- https://mvnrepository.com/artifact/org.springframework.security/spring-security-web -->
<!-- https://mvnrepository.com/artifact/org.springframework.security/spring-security-config -->
```

* security-context.xml : spring folder - new - spring bean configuration -> skip next -> finish
```xml
<!-- src/main/resources/security-context.xml -->
<!-- No such Bean Definition : spring security -->
<bean id="bcryptPasswordEncoder"
  class="org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder" />
```

* web.xml : security-context.xml -> tomcat에 설정해야함
```xml
  <!-- web.xml -->
  <!-- The definition of the Root Spring Container shared by all Servlets and Filters -->
  <context-param>
    <param-name>contextConfigLocation</param-name>
    <param-value>
      /WEB-INF/spring/root-context.xml
      /WEB-INF/spring/security-context.xml
    </param-value>
  </context-param>
```
## web.xml
```xml
<!-- Server web.xml -->
<web-app version="3.1" xmlns="http://xmlns.jcp.org/xml/ns/javaee"
  xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
  xsi:schemaLocation="http://xmlns.jcp.org/xml/ns/javaee http://xmlns.jcp.org/xml/ns/javaee/web-app_3_1.xsd">

<!-- comment the following from my web app web.xml -->
<!-- and replace it with the one from Server web.xml -->
<!-- <web-app version="2.5" xmlns="http://java.sun.com/xml/ns/javaee"
xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
xsi:schemaLocation="http://java.sun.com/xml/ns/javaee https://java.sun.com/xml/ns/javaee/web-app_2_5.xsd"> -->

<!-- spring has pre-defined filter! -->
<filter>
  <filter-name>encodingFilter</filter-name>
  <filter-class>org.springframework.web.filter.CharacterEncodingFilter</filter-class>
  <init-param>
    <param-name>encoding</param-name>
    <param-value>UTF-8</param-value>
  </init-param>
</filter>
<filter-mapping>
  <filter-name>encodingFilter</filter-name>
  <url-pattern>/*</url-pattern>
</filter-mapping>

<!-- add welcome-file-list -->
<welcome-file-list>
  <welcome-file>index.jsp</welcome-file>
</welcome-file-list>
```

## create index.jsp in webapp folder

## Spring settings

```xml
<!-- root-context.xml
  defines shared resources visible to all other web components.
  root-context.xml 은 스프링 돌아가는데 필요한거 다 담고 있다
  스프링은 객체들에 의해 돌아간다. 여기에 디비에 대한 설정이 들어간다.
-->
<!-- DB JDBC MYBATIS 설정 (pom.xml에서 필요한 resource 저장해놓음) jar file -> dependency 등록 -->
<!-- mybatis-config.xml을 spring에 맞게 수정.
  Spring이 시작하면서 이용할 수 있는 객체로 등록됨
  Spring 설정파일: BasicDataSource를 spring/spring 계정에게 부여. spring bean 등록
  close 생성자를 이용해서 destroy. Spring xml 파일을 보고 생성자를 생성. 연결내용.
  DB에 필요한 데이터를 BasicDataSource객체에 값을 넣음 setter(=property)
-->
<bean id="dataSource" class="org.apache.commons.dbcp.BasicDataSource" destroy-method="close">
  <property name="driverClassName" value="oracle.jdbc.driver.OracleDriver"/>
  <property name="url" value="jdbc:oracle:thin:@localhost:1521:xe"/>
  <property name="username" value="spring"/>
  <property name="password" value="spring"/>
</bean>

<!-- Mybatis에서 SqlSession 객체를 생성하는 FactoryBean 을 springbean으로 등록 
  Mybatis Spring jar파일에 있는 객체를 씀 객체생성 멤버변수로 BasicDataSource -->
<bean id="sqlSessionFactoryBean" class="org.mybatis.spring.SqlSessionFactoryBean">
  <property name="dataSource" ref="dataSource"/>
  <property name="mapperLocations" value="classpath*:mapper/**/*.xml"/>
  <property name="configLocation" value="classpath:mybatis-config.xml"/>
</bean>

<!-- SessionTemplate을 springbean으로 등록하여 활용 -->
<bean id="sqlSessionTemplate" class="org.mybatis.spring.SqlSessionTemplate">
  <constructor-arg index="0" ref="sqlSessionFactoryBean"/>
</bean>

```

## mybatis-config.xml

## ArrayTypeHandler implements TypeHandler<String[]>
```xml
<!-- member-mapper.xml -->
<!-- resultType은 select에서만 필요
resultType update insert delete 은 디폴트로 _int이므로 생략
  -->
<insert id="insertMember" parameterType="member">
  INSERT INTO MEMBER VALUES(
                  #{userId},#{password},#{userName},#{gender},
                  #{age},#{email},#{phone},
                  #{address},#{hobby, typeHandler=arrType},SYSDATE)
<!--   private String[] hobby; private Date enrollDate; -->
</insert>
```

## mapper/emp/emp-mapper.xml
```xml
<mapper namespace="dev"></mapper>
```

## SPRING LOG : 등록된 bean과 autowiring된 것들을 생성 및 bean으로 등록된 controller load?
pom.xml log4j mvnrepository is included in default

```
log4j 구조
```
1. Appender 태그 : log출력에 대한 환경설정하는 태그
  * 로그를 어디로 출력을 할지 : (콘솔/파일/DB)
2. Logger 태그 : 어디 부분에서 Appender를 실행 할 지
  * 어떤 단계에서 실행할지
3. Root 태그 : 기본적용 Logger (부모; 최상위 객체)
```
기본 패턴설정 : %-5p: %c - %m%n
    이벤트명, 카테고리명, 로그전달메세지 개행
```
* %c : 카테고리명(logger이름)을 표시
* 카테고리명이 a.b.c일때, %c{2}는 b.c를 출력
* %C : 클래스명을 표시함.	
* 풀 클래스 명이 com.kh.logger일때, %C{2}는 kh.logger를 출력
* %d : 로그 시간을 출력한다. java.text.SimpleDateFormat에서 적절한 출력 포맷을 지정할 수 있다. 
* %d{HH:mm:ss, SSS}
* %d{yyyy MMM dd HH:mm:ss, SSS}
* %d{ABSOLUTE} 
* %d{DATE} 
* %d{ISO8601}
* %F : 파일명을 출력. 로그시 수행한 메소드, 라인번호가 함께 출력된다.
* %l : 로깅이 발생한 caller의 위치정보. 자바파일명:라인번호(링크제공) 
* %L : 라인 번호만 출력한다(링크없음)
* %m : 로그로 전달된 메시지를 출력한다.
* %M : 로그를 수행한 메소드명을 출력한다. 
* %n : 플랫폼 종속적인 개행문자가 출력. rn 또는 n
* %p : 로그 이벤트명등의 priority 가 출력(debug, info, warn, error, fatal )
* %r : 로그 처리시간 (milliseconds)
* %t : 로그이벤트가 발생된 쓰레드의 이름을 출력
* %% : % 표시를 출력. escaping
* %r : 어플리케이션 시작 이후 부터 로깅이 발생한 시점의 시간(milliseconds)
* %X : 로깅이 발생한 thread와 관련된 MDC(mapped diagnostic context)를 출력합니다. %X{key} 형태.

```xml
<!-- pom.xml -->
<!-- Logging -->
<dependency>
  <groupId>org.slf4j</groupId>
  <artifactId>slf4j-api</artifactId>
  <version>${org.slf4j-version}</version>
</dependency>

<dependency>
  <groupId>log4j</groupId>
  <artifactId>log4j</artifactId>
  <version>1.2.15</version>
  <exclusions> ... </exclusions>
  <scope>runtime</scope>
</dependency>
```

src/main/resources -> log4j.xml
```xml
<!-- 1. Appender : console에 찍음-->
<appender name="console" class="org.apache.log4j.ConsoleAppender">
  <param name="Target" value="System.out" />
  <!-- 로그가 찍히는 형식 -->
  <layout class="org.apache.log4j.PatternLayout">
    <!-- m : message n: newline -->
    <param name="ConversionPattern" value="[%d{yyyy-MM-dd HH:mm:ss}] %-5p: %l - %m%n" />
  </layout>
  <!-- <layout class="org.apache.log4j.HTMLLayout"></layout> -->
  <!-- <layout class="org.apache.log4j.xml.XMLLayout"></layout> -->
</appender>

<!-- 2. Appender : 파일에 로그 찍기 다른Append를 이용하면됨 -> DailyRollingFileAppender -->
<appender name="filelogger" class="org.apache.log4j.DailyRollingFileAppender">
  <param name="file" value="C://logs//spring//spring.log" />
  <!-- determine log append / replace -->
  <param name="Append" value="true" />
  <param name="dataPattern" value=".yyyy-MM-dd" />
  <layout class="org.apache.log4j.PatternLayout">
      <param name="ConversionPattern" value="[%d{yyyy-MM-dd HH:mm:ss}] %-5p: %l - %m%n" />
  </layout>
</appender>
<!-- 3. sql구문 로그 남기기  -->
<appender name="sqlLogger" class="org.apache.log4j.ConsoleAppender">
  <layout class="org.apache.log4j.PatternLayout">
    <param name="ConversionPattern" value="%-5p: %m%n" />
  </layout>
</appender>

<!-- Application Loggers -->
<logger name="jdbc.sqlonly" additivity="false">
  <level value="INFO" />
  <appender-ref ref="sqlLogger" />
</logger>
<logger name="jdbc.resultsettable" additivity="false">
  <level value="INFO" />
  <appender-ref ref="sqlLogger" />
</logger>

<!-- Application Loggers -->
<logger name="com.kh.dm">
  <!-- <level value="info" /> -->
  <level value="debug" />
  <!-- name에 해당 하는 패키지의 모든 로그가 level INFO이상의 로그를 찍음
          TRACE
          DEBUG(개발시 사용하는 로그들)
          INFO(RUNTIME 중 상태변경, 정보성 메시지를 담을때)
          WARN(프로그램 실행시 문제가 없지만, 향후 시스템에서 error의 원인이 될 수 있다는 경고메시지)
          ERROR(어떤 요청을 처리할때 발생한 문제, 프로그램 동작안함)
          FATAL(심각한 에러, 메모리에 대한 손상, 운영체제 손상)
          : 개발자가 직접 에러 레벨 정함
    -->
</logger>

<!-- 3rdparty Loggers -->
<logger name="org.springframework.core">
  <level value="info" />
</logger>

<logger name="org.springframework.beans">
  <level value="info" />
</logger>

<logger name="org.springframework.context">
  <level value="info" />
</logger>

<logger name="org.springframework.web">
  <level value="info" />
</logger>

<!-- Root Logger -->
<root>
  <priority value="warn" />
  <appender-ref ref="console" />
  <appender-ref ref="filelogger" />
</root>
```

## add log4jdbc-remix to pom.xml
```xml
<!--  DB쿼리문, 결과를 log로 출력하는 라이브러리 추가 log4jdbc-remix -->
<!-- https://mvnrepository.com/artifact/org.lazyluke/log4jdbc-remix -->
<dependency>
  <groupId>org.lazyluke</groupId>
  <artifactId>log4jdbc-remix</artifactId>
  <version>0.2.7</version>
</dependency>
```

com.kh.spring.logger.LoggerInterceptor(superclass as HandlerInterceptorAdapter)
```
method mapping 실행전,후, 응답완료후,  -> filter대신에 interception 사용

interceptor
구현 : HandlerInterceptorAdapter를 상속해서 구현

전처리(preHandler) : dispatcherServlet이 컨트롤러가 매핑된 메소드 호출전에 실행되는 로직구현
후처리(postHandler) : 컨트롤러가 매핑된 메소드 실행 후 실행되는 로직 구현
뷰처리후(afterCompletion) : 응답까지 완료된 후 실행되는 매소드 구현

1. intercepter클래스를 등록을 해야함
  -> servlet-context.xml에 등록!
```

```text
10월 08, 2019 9:10:58 오전 org.apache.catalina.core.ApplicationContext log
정보: Initializing Spring FrameworkServlet 'appServlet'
INFO : org.springframework.web.context.support.XmlWebApplicationContext - Refreshing WebApplicationContext for namespace 'appServlet-servlet': startup date [Tue Oct 08 09:10:58 KST 2019]; parent: Root WebApplicationContext
INFO : org.springframework.beans.factory.xml.XmlBeanDefinitionReader - Loading XML bean definitions from ServletContext resource [/WEB-INF/spring/appServlet/servlet-context.xml]
INFO : org.springframework.beans.factory.annotation.AutowiredAnnotationBeanPostProcessor - JSR-330 'javax.inject.Inject' annotation found and supported for autowiring
INFO : org.springframework.web.servlet.mvc.method.annotation.RequestMappingHandlerMapping - Mapped "{[/demo/demo3.do]}" onto public java.lang.String com.kh.spring.demo.controller.DemoController.demo3(java.util.Map,org.springframework.ui.Model)
INFO : org.springframework.web.servlet.mvc.method.annotation.RequestMappingHandlerMapping - Mapped "{[/demo/demo4.do]}" onto public java.lang.String com.kh.spring.demo.controller.DemoController.demo4(com.kh.spring.demo.model.vo.Dev,org.springframework.ui.Model)
INFO : org.springframework.web.servlet.mvc.method.annotation.RequestMappingHandlerMapping - Mapped "{[/demo/demo.do]}" onto public java.lang.String com.kh.spring.demo.controller.DemoController.demo()
INFO : org.springframework.web.servlet.mvc.method.annotation.RequestMappingHandlerMapping - Mapped "{[/demo/demo1.do]}" onto public java.lang.String com.kh.spring.demo.controller.DemoController.demo1(javax.servlet.http.HttpServletRequest,javax.servlet.http.HttpServletResponse)
INFO : org.springframework.web.servlet.mvc.method.annotation.RequestMappingHandlerMapping - Mapped "{[/demo/demo2.do]}" onto public java.lang.String com.kh.spring.demo.controller.DemoController.demo2(java.lang.String,int,java.lang.String,java.lang.String,java.lang.String[],org.springframework.ui.Model)
INFO : org.springframework.web.servlet.mvc.method.annotation.RequestMappingHandlerMapping - Mapped "{[/demo/insertDev.do]}"
```

## (AOP) Aspect Oriented Prorgramming
```
관점지향 프로그래밍 controller로 매칭 X method 명칭으로 mapping
insert update delete (Service)메소드가 실행될떄, aop(관점지향)로 권한 판단
소규모 프로젝트에서는 interceptor로 구현. interceptor는 controller 접근일때만 
service나 dao도 쓸수 있는 interceptor.

AOP는 oop를 더 oop스럽게 사용할 수 있게 하는 기술
객체를 재사용함으로서 개발자들은 반복되는 코드를 많이 줄였지만,
매요청마다 로그, 권한체크, 인증, 예외처리 등 필수요소는 반복될 수 밖에 없었음.
AOP를 통해 비즈니스로직(주업무)과 공통모듈(보조업무)로 구분한 후 비지니스 로직
코드외부에서 필요한 시점에 공통모듈을 삽입하여 실행하는 것.
e.g. 자동 transaction 처리
e.g. delete update메소드 실행 할때, aop를 통해 session에 있는
     로그인 유저 정보를 테이블에 저장 및 로그 남김, 또는 session으로 권한 체크 등

주요개념.
a. 관점(aspect)
  구현하고자 하는 횡단의 관심사의 기능, 클래스 단위
b. Join point (적용가능한 지점)
  관점을 삽입하여 Advice(적용될 코드/메소드 /기능)가 적용될 수 있는 위치 (before, after, around)
c. Pointcut 어느클래스의 어느 메소드 전,후,전/후에 적용할지 패턴형식으로 작성하는 것을 말함
d. Advice
   공통 모듈 로직
```

```
AOP적용
1. 선언적방식 : XML에서 설정
```
e.g.
```xml
<!-- spring/appServlet/aspect-context.xml -->
<aop:config>
  <aop:aspect id="test" ref="loggerAspect" >
    <aop:pointcut
        expression="execution(접근제한자 클래스명(패키지포함) method param)" id="pc" />
    <aop:after|before|around method="loggerAspect의 메소드" pointcut="pc" />
  </aop:aspect>
</aop:config>

```
2. Annotation : 클래스 내부에서 Annotation 구현
annotation을 검색할수 있게 설정 *xml파일에 반드시 있어야함
```xml
<aop:aspectj-autoproxy/>
```
```java
aop설정 메소드에 어노테이션 표시
@Pointcut("execution(public**(..))")
메소드명
```

## aop -> pom.xml
```xml
<!-- mvnrepository : search for "AspectJ Weaver" -->
<!-- https://mvnrepository.com/artifact/org.aspectj/aspectjweaver -->
<dependency>
  <groupId>org.aspectj</groupId>
  <artifactId>aspectjweaver</artifactId>
  <version>1.6.10</version>
</dependency>
```

## appServlet 밑에 spring bean configuration file (xml)
```aspect-context-xml```


## paging
```xml
<!-- board들어갈떄 로그인유저만 servlet-context.xml -->
  <!-- 인터셉터 등록하기 -->
  <interceptors>
    <interceptor>
      <!-- ** 전체에 interceptor 적용하여 log 찍음 -->
      <!-- <mapping path="/**" /> -->
      <!-- member만 interceptor 적용하여 log 찍도록 -->
      <!-- [2019-10-14 09:15:22] DEBUG: com.kh.spring.common.interceptor.LoggerInterceptor.preHandle(LoggerInterceptor.java:19) - ======================start====================== -->
      <!-- [2019-10-14 09:15:22] DEBUG: com.kh.spring.common.interceptor.LoggerInterceptor.preHandle(LoggerInterceptor.java:20) - /spring/member/memberEnroll.do -->

      <!-- exclude mapping -->
      <!-- <exclude-mapping path="/member/*" /> -->

      <!-- <mapping path="/member/*" /> -->
      <mapping path="/demo/*" />
      <mapping path="/board/*" />
      <beans:bean id="loggerInterceptor" class="com.kh.spring.common.interceptor.LoggerInterceptor" />
    </interceptor>
  </interceptors>

```

## file upload
mvnrepository 


## transaction manager 
insert into attachmet : tablename ('attachmet') exception :
board is still inserted & committed

```
transaction시에는 aop와 연결해서 구현 -> 에러 발생시 rollback!
처리하는 객체를 spring 이 만들어 놓음 : transaction manager
```

transaction:
```
spring에서는 트랜젝션 관리자가 처리함(transaction manager)
spring에서는 DB에 접근해서 처리할 수 있는 방법이 다양함
jsp, 하이버네이트, mybatis 여러가지 라이브러리, 프레임워크를
적용할 수 있기 때문에, 이를 통합적으로 관리 할 수있는 관리자를 둔것임
```

트렌젝션 처리하는 방법:
```xml
<!--
1. 선언적 방법: -> xml설정
 - transaction 관리하는 xml을 만들어서 처리
  root-context.xml에 태그를 추가해서
-->
<tx:advice transaction-Manager="transactionManager">
  <tx:attributes>
    <!-- 해당하는 메소드 지정 -->
    <!-- => 트랜젝션 매니져를 통해 관리되야할 메소드 -->
    <!-- insert로 시작되는 메소드를 매니져 관리 안으로 집어 넣음 -->
    <tx:method name="insert*" rollback-for="Exception" />
  </tx:attributes>
</tx:advice>
<!-- runtime exception 발생했을때 처리 기본적으로 기능 있음
syntax 에러 및 다른 에러 처리도 범위 확장/축소 가능 -->
<bean id="transactionManager" class="a123asdf" >

<!--
2. 어노테이션 방법(programatic) - 소스코드상에 @으로 처리
-->
1) root-context.xml : 어노테이션 살펴봐. servlet-context : 어노테이션 driven
 -> annotation표시를 검색할 수 있게 설정
<tx:annotation-driven transaction-manager="" />
2) Service 객체 트렌젝션 처리
메소드 마다 -> insert, update, delete 메소드에 트랜젝션
```
```java
@Service
class EmpService{

  @Override
  @Transactional
  // @Transactional([옵션])
  // @Transactional(rollback-for="Exception.class")
  // @Transactional(rollback-for="타임아웃:일정시간 지나면 rollback")
  public int insertEmp(){

  }
}
```

* 옵션
```
propatation 트랜젝션.
  기존에 트렌젝션이 발생한 뒤 다른 트렌젝션이 참여하는 방법을 설정

  REQUIRED = DEFAULT //기본적으로 그냥 이걸로 하세요!
    이미 시작된 트랜젝션이 있으면 참여하고 아니면 새로 생성

  SUPPORT
    이미 시작된 트랜젝션이 있으면 참여, 없으면 트랜젝션 없이 처리

  MANDATORY
    REQUIRED와 비슷. 이미 시작된 트렌젝션이 있으면 참여, 없으면 예외발생 시킴

  REQUIRED_NEW
    항상 새로운 트렌젝션을 수행. 이미 사자된 트렌젝션은 보류
  NEVER, NOT_SUPPORTED

  ISOLATION 격리수준 (commit 된것을 select또는 이전 것 을 select 한 것인지 등 설정)

  DEFAULT : DBMS가 가지고 있는 트렌젝션 설정을 따르는 것

  READ_COMMITTED
    기본적으로 가장 많이 쓰는 방법 COMMIT 까지 SELECT 대기

  READ_UNCOMMITTED 가장 낮은 격리수준
    COMMIT이 되기 전에 SELECT 됨. DIRTYREAD발생

  REPEATABLE_READ
    트랜젝션이 읽은 row를 다른 트렌젝션이 수정하는 것을 막는 것

  SERIALIZABLE
    격리수준 가장 강력(READ 못하게), 성능을 떨어뜨림... 자주 사용 안함

  기타: TIMEOUT 커밋하는 기간을 정해놓음, READONLY
 ```
 ```xml
    <!-- https://mvnrepository.com/artifact/org.aspectj/aspectjweaver -->
    <dependency>
      <groupId>org.aspectj</groupId>
      <artifactId>aspectjweaver</artifactId>
      <version>1.6.10</version>
    </dependency>
기존에 추가한 aspectjweaver 사용함 

 ```
 namespace탭에서 tx 체크하기
 ```java
 // boardServiceImpl.java
  //  @Transactional //트랜젝션의 기준이 되는 것 : RunTimeException 발생시! Exception으로 하면 안됨
  //  @Transactional(rollbackFor = Exception.class) //트랜젝션의 기준이 되는 것 : 모든 에러 Exception (RunTimeException 포함) 발생시!
  //annotation 방식 말고 root-context.xml에 등록가능
  public int insertBoard(Map<String, String> param, List<Attachment> attachList) {
    //세션 트랜젝션 관리(by spring)
    int result = 0;
    result = dao.insertBoard(sqlSession, param); //board테이블에 데이터 입력!

    if(result ==0)
      throw new RuntimeException(); //TransactionManager가 자동으로 롤백 처리함

    if(attachList.size() > 0) {
      for(Attachment a : attachList) {
        a.setBoardNo(Integer.parseInt(param.get("boardNo")));
        result = dao.insertAttachment(sqlSession, a);
      }
    }
    
    return result;
  }
  

 ```

 ```xml
 <!-- root-context.xml -->
   <!-- annotation 방식으로 처리하기 -->
  <!-- tx는 namespace에 있는 것: namespace 등록하기 -->
  <!-- <tx:annotation-driven transaction-manager="transactionManager" /> -->
  
  <!-- 선언적 방식으로 Transaction 처리 -->
  <tx:advice id="txAdvice" transaction-manager="transactionManager">
    <tx:attributes>
      <tx:method name="insert*" rollback-for="Exception" />
    </tx:attributes>
  </tx:advice>

  <!-- aop와 연결하여 트랜젝션을 적용 -->
  <aop:config>
    <aop:pointcut expression="execution(* com.kh.spring..*ServiceImpl.*(..))" id="serviceMethod" />
    <aop:advisor advice-ref="txAdvice" pointcut-ref="serviceMethod" />
  </aop:config>


 ```

 ## 스프링 Ajax 처리하기
 3가지 방법 있음
 ```
 1. ServletOutStream 이용하는 방법
 2. ModelAndView 객체를 반환형으로 하고, viewResolver를 이용하는 방법
 3. @Repository 어노테이션을 이용하는 방법

  2~3 jackson ObjectMapper를 이용함
 ```


mvnrepository - Json Lib4Spring 1.0.2
```
 json-lib-ext-spring
```

```xml
  <!-- servlet-context.xml -->
  <!-- jsonView 등록하기 -->
  <beans:bean id="viewResolver" class="org.springframework.web.servlet.view.BeanNameViewResolver">
  </beans:bean>

```

## jackson-databind
mvnrepo jackson-databind
```xml
<!-- pom.xml -->
<!-- https://mvnrepository.com/artifact/com.fasterxml.jackson.core/jackson-databind -->
<dependency>
    <groupId>com.fasterxml.jackson.core</groupId>
    <artifactId>jackson-databind</artifactId>
    <version>2.9.10</version>
</dependency>

```

```xml
<!-- servlet-context.xml -->
  <!-- jackson -->
  <beans:bean id="jacksonMessageConverter" class="org.springframework.http.converter.json.MappingJackson2HttpMessageConverter">
  </beans:bean>
  
  <!-- jackson handler -->
  <beans:bean id="" class="org.springframework.web.servlet.mvc.method.annotation.RequestMappingHandlerAdapter">
    <beans:property name="messageConverters">
      <beans:list>
        <beans:ref bean="jacksonMessageConverter" />
      </beans:list>
    </beans:property>
  </beans:bean>
```

## Changing Context root
```xml
<!-- project r_click -> Web project setting -> edit context path
Servers/server.xml -->
<Context docBase="ParkingSpring" path="/parking" reloadable="true" source="org.eclipse.jst.jee.server:ParkingSpring"/></Host>
```

