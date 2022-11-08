
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
public class User {
	private String id;
	private String name;
	private String passwordk;
}
```

```java
public class UserDao {
	public void add(User user) throws ClassNotFoundException, SQLException {
		Connection c = getConnection();
		PreparedStatement ps = c.preparedStatement(
			"insert into users(id, name, password) values(?,?,?)");
		ps.setString(1, user.getId());
		ps.setString(2,user.getName());
		ps.setString(3, user.getPassword());
		ps.executeUpdate();

		ps.close();
		c.close();
	}

	public User get(String id) throws ClassNotFoundException, SQLException {
		Connection c = getConnection();
		PreparedStatement ps = c.preparedStatement(
			"select * from users where id=?");
		ps.setString(1, id);

		ResultSet rs = ps.executeQuery();
		rs.next();
		User user = new User();
		user.setId(rs.getString("id"));
		user.setName(rs.getString("name"));
		user.setPassword(rs.getString("password"));

		rs.close();
		ps.close();
		c.close();
	}

	public Connection getConnection throws ClassNotFoundException, SQLException {
		Class.forName("com.mysql.jdbc.Driver");
		Connection c = DriverManager.getConnection(
			"jdbc:mysql://localhost/springbook", "spring", "book"
		);
		return c;
	}
}

```




- 커넥션 만들기의 독립
	- 템플릿 메소드 패턴: 슈퍼클래스에서 기본적인 로직흐름만들고
	- 그 기능의 일부를 추상메소드나 오버라이딩가능한 protected 메소드로 만듦
	- 팩토리 메소드 패턴 : 서브클래스에서 구체적인 오브젝트 생성방법 결정
	- "UserDao에 팩토리 메소드 패턴을 적용해서 getConnection()을 분리"
	- "디자인패턴": 소프트웨어 설계시 특정 상황에서 자주만나는
		- 문제를 해결하기 위한 재사용 가능한 솔루션


```
UserDao [add(), get(), getConnection()]
	ㄴNUserDao [getConnection()]
	ㄴDUserDao [getConnection()]
```

```java
public class UserDao {
	public void add(User user) throws ClassNotFoundException, SQLException {
		Connection c = getConnection();
		PreparedStatement ps = c.preparedStatement(
			"insert into users(id, name, password) values(?,?,?)");
		ps.setString(1, user.getId());
		ps.setString(2,user.getName());
		ps.setString(3, user.getPassword());
		ps.executeUpdate();

		ps.close();
		c.close();
	}

	public User get(String id) throws ClassNotFoundException, SQLException {
		Connection c = getConnection();
		PreparedStatement ps = c.preparedStatement(
			"select * from users where id=?");
		ps.setString(1, id);

		ResultSet rs = ps.executeQuery();
		rs.next();
		User user = new User();
		user.setId(rs.getString("id"));
		user.setName(rs.getString("name"));
		user.setPassword(rs.getString("password"));

		rs.close();
		ps.close();
		c.close();
	}

	// 추상메소드-> 메소드구현은 서브클래스가 담당
	// 또는 오버라이딩이 가능한 'protected' 메소드 정의
	public abstract Connection getConnection throws ClassNotFoundException, SQLException;

}
```

```java
public class NUserDao extends UserDao{
	// add, get은 상속됨

	// N사의implementation
	public Connection getConnection throws ClassNotFoundException, SQLException {
		Class.forName("com.mysql.jdbc.Driver");
		Connection c = DriverManager.getConnection(
			"jdbc:mysql://localhost/springbook", "spring", "book"
		);
		return c;
	}
}
```
```java
public class DUserDao extends UserDao{
	// add, get은 상속됨

	// D사의implementation
	public Connection getConnection throws ClassNotFoundException, SQLException {
		...
	}

}
```

- 상속의 한계점
	- 다중상속 문제발생
		- 만약 UserDao가 다른목적으로 다른 상속을 이미 사용하고 있다면?
	- 상속 상하위 클래스 관계는 생각보다 너무 밀접
		- 서브클래스가 슈퍼클래스 기능 여전히 사용 가능
		- 슈퍼클래스 수정 시, 서브 클래스도 코드 수정소요


