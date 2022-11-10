
- 스프링 컨테이너
	- '애플리케이션 컨텍스트' 스프링 런타임
	- 애플리케이션 오브젝트 생성하고 관리: 웹모듈에서 동작하는 서비스나 서블릿으로 등록해서 사용
- 공통 프로그래밍 모델
	- IoC/DI 오브젝트 생명주기와 의존관계에 대한 프로그래밍 모델
	- 서비스 추상화
	- AOP 모델

- Simplicity: EJB -> 객체지향 POJO
- Flexibility

- DAO (Data Access Object) 데이터 조회 및 조작하는 기능 전담하는 오브젝트

```java
package springbook.user.domain;

public class User {
	String id;
	String name;
	String password;

	//public Getter, Setter
}
```

```sql
CREATE TABLE USERS (
	id varchar(10) primary key,
	name varchar(20) not null,
	password varchar(10) not null
)
```

- JavaBean 조건
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
package springbook.user.dao;

public class UserDao {

	public void add(User user) throws ClassNotFoundException, SQLException {
		Class.forName("com.mysql.jdbc.Driver");
		Connection c = DriverManager.getConnection(
			"jdbc:mysql//localhost/springbook", "spring", "book");
		PreparedStatement ps = c.prepareStatement(
			"insert into users(id, name, password) values(?,?,?)");
		ps.setString(user.getId());
		ps.setString(user.geName());
		ps.setString(user.getPassword());
		ps.executeUpdate();

		ps.close();
		c.close();
	}

	public User get(String id) throws ClassNotFoundException, SQLException {
		Class.forName("com.mysql.jdbc.Driver");
		Connection c = DriverManger.getConnection(
			"jdbc:mysql//localhost/springbook", "spring", "book");
		PreparedStatement ps = c.prepareStatement(
			"select * from users where id = ?");
		ps.setString(id);
		ResultSet rs = ps.executeQuery();
		rs.next();

		User user = new User();
		user.setId(rs.getString("id"));
		user.setName(rs.getString("name"));
		user.setPassword(rs.getString("password"));

		rs.close();
		ps.close();
		c.close();

		return user;
	}

	public static void main(String[] args) {
		UserDao dao = new UserDao();

		User user = new User();
		user.setId("1");
		user.setName("Bob");
		user.setPassword("fooo");

		dao.add(User);
		System.out.Println(user.getId() + " 등록 성공");

		User user2 = dao.get(user.getId());
		System.out.Println("조회 성공: name=" + user2.getName());
		System.out.Println("조회 성공: password=" + user2.getPassword());

		System.out.Println(user.getId() + " 조회 성공");
	}
}

```

- DAO의 분리
	- 애플리케이션 설계 변경에따른 영향도 최소화
	- 관심사의 분리

- UserDao의 관심사 (1. DB 커넥션, 2. DB Statement 실행, 3. 리소스 close())

- 커넥션만들기 추출 getConnection()

```java
public class UserDao {
	public void add(User user) throws ClassNotFoundException, SQLException {
		Connection c = getConnection();
		PreparedStatement ps = c.prepareStatement(
				"insert into users(id, name, password) values(?,?,?)");
		ps.setString(user.getId());
		ps.setString(user.geName());
		ps.setString(user.getPassword());
		ps.executeUpdate();

		ps.close();
		c.close();
	}

	public User get(String id) throws ClassNotFoundException, SQLException {
		Connection c = getConnection();
		PreparedStatement ps = c.prepareStatement(
				"select * from users where id = ?");
		ps.setString(id);
		ResultSet rs = ps.executeQuery();
		rs.next();

		User user = new User();
		user.setId(rs.getString("id"));
		user.setName(rs.getString("name"));
		user.setPassword(rs.getString("password"));

		rs.close();
		ps.close();
		c.close();

		return user;
	}

	public Connection getConnection() throws ClassNotFoundException, SQLException {
		Class.forName("com.mysql.jdbc.Driver");
		Connection c = DriverManager.getConnection(
				"jdbc:mysql://localhost/springbook", "spring", "book"
		);
		return c;
	}
}
```


- 커넥션만들기 독립
	- UserDao 확장 클래스의 DB연결 인터페이스 제공: 상속(inheritance)
	- 2개 업체-> UserDao 사용하여 각각 자체 구현 DB 커넥션 코드 사용 하려함
	- UserDao 개발자는 getConnection() 코드를 공개하지 않을때 구현방법?

- 커넥션 만들기의 독립
	- 템플릿 메소드 패턴
		- 슈퍼클래스에서 기본적인 로직흐름만들고
		- 서브클래스에서 그 기능의 일부를 추상메소드나 오버라이딩가능한 protected 메소드로 만듦
	- 팩토리 메소드 패턴 : 서브클래스에서 구체적인 오브젝트 생성방법 결정
		- "UserDao에 팩토리 메소드 패턴을 적용해서 getConnection()을 분리"
		- "디자인패턴": 소프트웨어 설계시 특정 상황에서 자주만나는
			- 문제를 해결하기 위한 재사용 가능한 솔루션


- 템플릿 메소드 패턴
	- 서브 클래스에서 오버라이드
	- `protectded void hookMethod() {}`: 선택적으로 오버라이드
	- `public abstract void abstractMethod() {}`: 반드시 오버라이드

- 팩토리메소드 패턴
	- 객체 생성방법을 결정하고, 그렇게 만들어진 오브젝트를 돌려주는 오브젝트
	- UserDao, ConnectionMaker: 애플리케이션의 핵심적인 데이터 로직과 기술 로직
	- DaoFactory:  애플리케이션의 오브젝트들을 구성하고 그 관계를 정의

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
	public abstract Connection getConnection() throws ClassNotFoundException, SQLException;

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
	- 커넥션 생성 확장기능을 다른 Dao클래스에서 사용 X
		- Dao클래스 추가시 구현 코드가 계속 중복



- 클래스의 분리
	- 하지만 N, D사의 UserDao 확장 시, UserDao클래스만 제공하고,
	- 커넥션기능 각 업체별 구현 제공 불가능 해짐
		- UserDao 코드 수정 없이 D사 구현 커넥션 제공 메소드 사용 불가능 (e.g. D사의 openConnection())
		- DB커넥션을 제공하는 클래스가 뭔지 UserDao클래스가 알고 있어야 함!
			- UserDao가 DB커넥션 가져오는 구체적인 방법에 종속 되어 버림

```java
package springbook.user.dao;

public class SimpleConnectionMaker {
	public Connection makeConnection() throws ClassNotFoundException, SQLException {
		Class.forName("com.mysql.jdbc.Driver");
		Connection c = DriverManger.getConnection(
			"jdbc:mysql//localhost/springbook", "spring", "book");
		return c;
	}
}
```

```java
public class UserDao {
	private SimpleConnectionMaker simpleConnectionMaker;

	public UserDao() {
		simpleConnectionMaker = new SimpleConnectionMaker();
	}

	public void add(User user) throws ClassNotFoundException, SQLException {
		Connection c = simpleConnectionMaker.makeConnection();
		// ...
	}

	public User get(String id) throws ClassNotFoundException, SQLException {
		Connection c = simpleConnectionMaker.makeConnection();
		// ...
	}
}
```










- 제어권의 이전을 통한 제어관계 역전
  - 사용 오브젝트 결정, 생성, 오브젝트 메소드 호출
  - 이런 흐름을 역전하여 자신이 사용할 오브젝트를 스스로 선택하거나 생성하지 않음
  - 위임받은 제어 권한을 갖는 특별한 오브젝트에 의해 결정되고 만들어지도록 함


- `@RunWith`
    - 다른 프레임워크와 test 컨텍스트를 통합하여 사용하거나
    - JUnit 4 테스트케이스 실행 플로우를 변경하기 위해 사용
    - JUnit 5에서는 `@ExtendWith` 사용

- 스프링 IOC
  - 스프링 빈:
    - 스프링이 제어권을 가지고 직접 만들고 관계를 부여하는 오브젝트
    - 오브젝트 단위의 애플리케이션 컴포넌트
    - 스프링 컨테이너가 생성과 관계설정, 사용 등을 제어해주는 제어의 역전이 적용된 오브젝트
    - `@Bean` : 오브젝트 생성을 담당하는 IoC용 메소드에 추가
      - 빈의 이름 : getBean("메소드명"). @Bean 붙힌 메소드명
      - 빈의 클래스 : 빈 오브젝트를 어떤 클래스를 이용해서 만들지 정의
      - 빈의 의존 오브젝트 : 빈의 생성자나 수정자 메소드를 통해 의존 오브젝트 넣어줌
    - `@Configuration` 애플리케이션 컨텍스트 또는 빈팩토리가 사용할 설정정보
      - 빈팩토리를 위한 오브젝트 설정을 담당하는 클래스
  - 빈팩토리 : 빈의 생성과 관계설정 제어 담당하는 IoC 오브젝트
  - 빈팩토리 vs. 어플리케이션 컨텍스트(IoC방식을 따라 만들어진 일종의 빈팩토리)

- 자바 싱글톤 패턴
  - new DaoFactory().userDao() vs. 어플리케이션 컨텍스트의 의존관계 검색에 의한 UserDao객체 생성
    - 전자는 호출할때마다 새로운 오브젝트 생성 vs 후자는 동일 오브젝트 반복 호출
    - 애플리케이션 컨텍스트는 싱글톤을 저장하고 관리하는 싱글톤 레지스트리
    - 애플리케이션 안에 제한된 수, 대개 한개의 오브젝트,만 만들어 사용하는 것이 싱글톤 패턴의 원리
  - private 생성자 정의
  - 싱글톤 오브젝트 저장할 자신과 같은 타입의 스태틱 필드 정의
  - static factory method getInstance() 만들고, 호출되는 시점에서
  한번만 오브젝트가 만들어지게 함. 생성된 오브젝트는 스태틱 필드에 저장된다
  - 한번 오브젝트(싱글톤) 만들어지고 난 후에는 getInstance() 메소드로 이미 만들어져 있는 스태틱 필드에 저장해둔 오브젝트를 넘겨줌

- 싱글톤 레지스트리
  - 스태틱 메소드와 private 생성자를 사용해야하는 비정상적인 클래스가 아니라, 평범한 자배클래스를 싱글톤으로 활용하도록 해줌
  - 멀티스레드 환경에서 여러스레드가 동시에 접근해서 사용 가능
  - 서비스 형태의 오브젝트로 사용되는 경우에는 상태정보를 내부에 갖고있지않은 stateless 방식으로 만들어 져야함
  - 변수저장 하는 방식은 덮어씌워질 수 있기 때문에 위험
  
- 의존관계
  - A -> B
  런타임시 발생 하는 의존관계 (오브젝트 의존관계)
  Dependent object
  
- 의존관계 주입 (Dependency Injection)
  - 클래스 모델이나 코드에는 런타임 시점의 의존관계 X (인터페이스에만 의존)
  - 런타임 시점의 의존관계는 컨테이너나 팩토리 같은 제 3의 존재가 결정한다
  - 의존관계는 사용할 오브젝트에 대한 레퍼런스를 외부에서 제공(주입) 해줌으로써 만들어진다
  - DaoFactory : 두오브젝트 사이의 런타임 의존관계를 설정해주는 의존관계 주입 작업 주도
    IoC방식으로 오브젝트의 생성과 초기화, 제공 작업을 수행하는 컨테이너; DI 컨테이너
        
- 의존관계 검색과 주입
  - 런타임시에 의존관계를 결정
  - 의존관계를 맺는 방법이 외부로부터 주입이 아니라, 스스로 검색을 이용
  - dependency lookup
    
- 차이점
  - DI를 원하는 오브젝트는 먼저 자기 자신이 컨테이너가 관리하는 빈이 되어야 함
  - 의존관계 검색 방식에서는 검색하는 오브젝트는 자신이 스프링의 빈일 필요가 없음 e.g. UserDao
  - ConnectionMaker만 스프링빈이면 됨
  - 반면 의존성주입 시 UserDao와 ConnectionMaker 사이에 DI 적용하려면, UserDao도 컨테이너가 만드는 빈 오브젝트여야 함
  - DI방법을 좀더 선호. 다만 static 메소드인 main에서는 의존성 검색 필요
  - 사용자 요청을 받을 때 마다 이런식의 구현이 필요하지만, 스프링이 서블릿으로 제공
  - UserDao는 ConnectionMaker에만 의존

- 메소드를 이용한 DI
  - 수정자 메소드 setter
  - 일반 메소드

- XML을 이용한 설정
  - DaoFactory의 bean인 userDao와 makeConnection을 xml로 대체가능

|자바 코드 설정정보 | XML 설정정보
---|---|---
빈 설정파일 | @Configuration | `<beans>`
빈 이름 | @Bean methodName() | `<bean id="methodName">`
빈 클래스 | return new BeanClass() | `class="a.b.c...BeanClass`

- @Autowired
  - 테스트 컨텍스트 프레임워크는 변수타입과 일치하는 컨텍스트 내의 빈을 찾는다.
  - 타입일치 빈이 있으면 인스턴스 변수에 주입 (타입에 의한 자동 와이어링)
  - ApplicationContext 는 applicationContext.xml 정의여부 상관없음
    - 스프링 어플리케이션 컨텍스트는 초기화 시 자신도 빈으로 등록
    - 어플리케이션 컨텍스트에는 ApplicationContext 타입의 빈이 존재, DI도 가능
    - ApplicationContext 빈을 가져올 필요없이 UserDao를 자동와이어링 가능
    - 같은타입의 빈이 2개이상이면 자동와이어링 불가능
    
    
``` java
@Autowired
SimpleDriverDataSource simpleDataSource;

@Autowired
DataSource dataSource;
```

- DI사용 및 SimpleDriverDataSource 대신 DataSource 인터페이스 사용 이유
  1. 추후 구현체가 변경 될수 있기때문
  2. 다른 서비스기능을 도입가능 (ex. 커넥션 counter)
  3. 효율적인 테스트


- `@BeforeAll`의 non-static 메소드 사용을 위해,
  - 클래스에 `@TestInstance(TestInstance.Lifecycle.PER_CLASS)` 추가

- `@DirtiesContext` : 테스트 메소드에서 애플리케이션 컨텍스트의 구성이나 상태를 변경한다는 것을
  - 테스트 컨텍스트 프레임워크에 알려줌
- 테스트용 test-applicationContext.xml 생성 및 dataSource빈의 "url" 프로퍼티를 testdb바라보게 설정


```
docker run -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root --name mysql8 mysql:8.0 --lower_case_table_names=1
```

```sql
-- 데이터베이스 생성 
CREATE DATABASE `testdb` CHARACTER SET utf8;
USE `testdb`;

# 계정생성 (5.7 이상은 Create 후 권한부여)
# % 인 경우 remote access 허용 
# localhost 인 경우 localhost만 허용 
# IP 인 경우 해당 IP만 허용 (192.168.0.% 인 경우 Host 대역 허용)
CREATE USER 'spring'@'%' IDENTIFIED BY 'book';

# 권한부여 (계정 생성 후 부여)
# 외부 % 및 개발서버 lowem-139 접근권한 허용
GRANT ALL PRIVILEGES ON `testdb`.* TO 'spring'@'%' WITH GRANT OPTION;

# 테이블 생성
CREATE TABLE users (
  ID varchar(10) primary key,
  NAME varchar(20) not null,
  PASSWORD varchar(10) not null
);
```

- 스프링 학습테스트
  - JUnit실행 클래스 : 테스트메소드 @Test 실행할 떄 마다 새로운 테스트클래스의 오브젝트가 만들어짐
  - ApplicationContext : 테스트 개수에 상관없이 한 개 만들어지고, 모든 테스트에서 공유됨

- 디펜던시 인젝션 pom.xml 또는 @Autowired
  - JdbcContext는 인터페이스 없으므로 pom.xml대신 수동 DI 사용


- 템플릿과 콜백
  - 템플릿
  - 콜백: 실행 되는 것을 목적으로 다른 오브젝트 메소드에 전달되는 오브젝트





