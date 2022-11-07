
- 스프링 컨테이너
	- '애플리케이션 컨텍스트' 스프링 런타임
	- 애플리케이션 오브젝트 생성하고 관리: 웹모듈에서 동작하는 서비스나 서블릿으로 등록해서 사용
- 공통 프로그래밍 모델 - IoC/DI, 서비스 추상화, AOP
	- IoC/DI 오브젝트 생명주기와 의존관계에 대한 프로그래밍 모델
	- 서비스 추상화
	- AOP 모델
- simplicity: EJB -> 객체지향 POJO
- flexibility

- DAO (Data Access Object) 데이터 조회 및 조작하는 기능 전담하는 오브젝트

```java
package springbook.user.domain;

public class User {
	String id;
	String name;
	String password;

	//Public Getter, Setter
}
```
```sql
CREATE TABLE USERS (
	id varchar(10) primary key,
	name varchar(20) not null,
	password varchar(10) not null
)
```

- JavaBean
  - 디폴트 생성자
  - 프로퍼티 (getter, setter)

- UserDao 
  - 사용자정보를 DB에 넣고 관리할 수 있는 UserDao 클래스
  - 등록, 수정, 삭제, 조회
  - JDBC 사용
    - DB연결 connection
    - SQL 담은 Statement / PreparedStatement 생성
    - Statement 실행
    - 조회의 경우 ResultSet로 받아서 오브젝트(user)에 옮겨줌
    - Connection, Statement, ResultSet 등 리소스는 close() 해준다
    - JDBC API가 만들어내는 exception을 직접 처리하거나 메소드에 throws 선언하여 메소드 밖으로 던짐

```java

```


