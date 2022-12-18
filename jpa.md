

- 객체답게 모델링 할수록 매핑작업만 늘어남 (MyBatis) -> JPA도입
  - 객체를 자바 컬렉션에 저장하듯이 DB에 저장하는 JPA
  - Java Persistence API (자바 진영의 ORM 기술 표준)
  - ORM: 객체와 관계형데이터베이스 RDB 맵핑
  - EJB 엔티디 빈(자바표준) -> 하이버네이트(오픈소스) -> JPA (자바표준)
  - SQL 중심 개발에서 객체중심 개발로 변환
- JPA와 CRUD
  - `jpa.persist(member);`
  - `Member member = jpa.find(memmberId);`
  - `member.setName("변경힐 이름");`
  - `jpa.remove(member);`

- New Project > Java, Maven
- GroupId : jpa-basic
- Artifact : ex1-hello-jpa
- Dependencies (pom.xml)
  - JPA 하이버네이트
  - H2 데이터베이스
- src > main > resources > META-INF 디렉토리 > persistence.xml 생성

```xml
<dependencies>
    <!-- JPA 하이버네이트 -->
    <dependency>
        <groupId>org.hibernate</groupId>
        <artifactId>hibernate-entitymanager</artifactId>
        <version>5.3.10.Final</version>
    </dependency>
    <!-- H2 데이터베이스 -->
    <dependency>
        <groupId>com.h2database</groupId>
        <artifactId>h2</artifactId>
        <version>1.4.199</version>
    </dependency>
</dependencies>
```

- h2 200 버전 다운로드 https://h2database.com/h2-2019-10-14.zip
  - jdbc:h2:~/test 커넥트
  - jdbc:h2:tcp://localhost/~/test 커넥트
  - 테이블생성

```sql
CREATE TABLE MEMBER (
  id bigint not null,
  name varchar(255),
  primary key (id)
);
SELECT * FROM MEMBER;
```

```java
package hello.jpa;

import javax.persistence.Entity;
import javax.persistence.Id;

@Entity
public class Member {

	@Id
	private Long id;
	private String name;
	// getter, setter
```

- Persistence 설정정보 -> EntityManagerFactory -> EntityManager
  - emf: 웹서버 올라오는 시점에 1개 객체 생성됨 (싱글톤)
  - em: 생성/삭제 반복: thread공유 하면 안됨!
  - JPA 모든 변경은 트랜젝션 안에서

- JPQL
  - 단순조회 EntityManager.find() 또는 객체그래프 탐색 a.getB().getC()
  - 복잡한 세부조건 where절 사용하기 위해 JPQL 사용
  - JPQL은 엔티티 객체를대상으로 쿼리
  - SQL은 데이터베이스 테이블을 대상으로 쿼리

- 핵심 2가지
  - 객체와 관계형 데이터베이스 매핑하기 (Object Relational Mapping)
  - 영속성 컨텍스트

- 영속성 컨텍스트
  - 엔티티를 (영속성 컨텍스트에) 영구 저장하는 환경 EntityManager.persist(entity);
  - 영속성 컨텍스트는 논리적인 개념. 엔티티매니저를 통해서 영속성 컨텍스트에 접근
  - EntityManager - PersistentContext 1:1
  - 엔티티 생성주기
    - 비영속 new/transient
    - 영속 managed : em.persist(member);
    - 준영속 detached : em.detach(member);
    - 삭제 removed : em.remove(member);
  - 애플리케이션과 DB 사이 중간에 영속성 컨텍스트 : (버퍼링, 캐싱 등 이점 있음)
  - 엔티티 조회, 1차캐시 key,value = (@Id, Entity 객체) 에서 조회, 반환
  - 1차캐시에 없으면 -> DB 조회 후 -> 1차 캐시에 저장 -> 객체 반환
  - DB 트랜젝션끝나면 entityManager는 사라지기때문에 1차캐시의 큰 성능적 이점은 없음

```java
package hello.jpa;

		import javax.persistence.EntityManager;
		import javax.persistence.EntityManagerFactory; import javax.persistence.EntityTransaction;
		import javax.persistence.Persistence;
		import java.util.List;

public class JpaMain {
	public static void main(String[] args) {
		// persistence.xml에 정의된 hello 설정정보 가져옴
		EntityManagerFactory emf = Persistence.createEntityManagerFactory("hello");
		EntityManager em = emf.createEntityManager();

		EntityTransaction tx = em.getTransaction();
		tx.begin();
		try {
			// 1. Create 등록

			// 비영속
//			Member member = new Member();
//			member.setId(1L);
//			member.setName("HelloA");
			//	영속 DB쿼리는 아직-> commit(); 시점에 INSERT 됨
			// 1차캐시에 저장됨
//			em.persist(member);
			// 비슷한 방식으로 select 를 할때도 DB 조회후에 1차캐시에 저장!
			//  -> 2번째 조회시에는 1차캐시에서 가져옴

			// 2. Read 조회
			// 1차캐시에서 조회함 : 쿼리 select 쿼리 안나감(1차캐시)
//			Member findMember = em.find(Member.class, 1L);
//			System.out.println("findMember.Id=" + findMember.getId());
//			System.out.println("findMember.Name=" + findMember.getName());
			List<Member> result = em.createQuery("select m from Member as m", Member.class)
					.setFirstResult(5) // 5번 부터
					.setMaxResults(8)  // 8개 데이터
					.getResultList();
			for (Member m : result) {
				System.out.println("member.name = " + m.getName());
			}

			// 3. Update 수정
			Member findMember = em.find(Member.class, 1L);
			findMember.setName("HelloJPA");
			// em.persist(member); 필요없음! 마치 컬렉션에 저장하는것 같음
			// -> UPDATE쿼리 tx.commit(); 시점에 발생!

			// 4. Delete 삭제
//			em.remove(findMmeber);

			// 커밋!
			tx.commit();

		} catch(Exception e) {
			tx.rollback();
		} finally {
			em.close();
		}

		emf.close();
	}
}

```

- 영속 엔티티의 동일성 보장
  - Member a = em.find(Member.class, "member1");
  - Member b = em.find(Member.class, "member1");
  - a==b (1차캐시에서 가져 옴. 같은 트랜젝션에서만 가능)

- 엔티티등록: 트랜젝션을 지원하는 쓰기지연 (버퍼링)
  - em.persist(member); 이때는 INSERT 안하고, tx.commit();시에 INSERT 쿼리 실행
  - tx.commit(); -> flush
  - JDBC 배치 : persistence.xml "`hibernate.jdbc.batch_size" 버퍼사이즈 설정

- 엔티티 수정 변경 감지
  - 1차캐시에 @Id, Entity, 스냅샷 있음.
  - tx.commit(); 데이터 업데이트 시 앤티티와 스냅샷 비교 후 쓰기지연 업데이트 SQL -> DB실행

```java
  // 3. Update 수정
  Member findMember = em.find(Member.class, 1L);
  findMember.setName("HelloJPA");
  // em.persist(member); 필요없음! 마치 컬렉션에 저장하는것 같음
  // -> UPDATE쿼리 tx.commit(); 시점에 발생!
```

- 플러시 : 영속성 컨텍스트의 변경내용을 데이터베이스에 반영
  - 변경감지-> 수정된 엔티티 쓰기지연 SQL저장소에 등록
  - 쓰기지연 SQL 저장소의 쿼리를 데이터베이스에 전송 (등록, 수정, 삭제 쿼리)

- 영속성 컨텍스트를 플러시 하는방법
  - em.flush() 직접 플러시 호출 (tx.commit(); 하기전에 강제로 DB반영)
  - tx.commit() 자동 플러시 호출
  - JPQL 쿼리실행 자동 플러시 호출. em.persist()하고 tx.commit()전에 JPQL(select)하는 경우. flush() 호출

- 플러시 모드 옵션
  - em.setFlushMode(FlushModeType.COMMIT); // 커밋할때만 플러시
  - em.setFlushMode(FlushModeType.AUTO); // 커밋이나 쿼리를 실행할 때 플러시 (기본값)


- 준영속 상태
  - 영속(1차캐시에 들어가서 JPA가 관리하는 상태) -> 준영속
  - 영속상태 엔티티가 영속성 컨텍스트에서 분리(detached)
  - 영속성 컨텍스트가 제공하는 기능을 사용하지 못함
  - em.detach(entity) 특정 엔티티만 준영속 상태로 전환
  - em.clear() 영속성 컨텍스트를 완전히 초기화
  - em.close() 영속성 컨텍스트를 종료

```java
Member m = em.find(Member.class, 150L);
m.setName("AAA");
em.detach(m);

// 아무일도 일어나지 않음(detach되어서 JPA가 관리하지 않기 때문)
// select쿼리만 나오고 insert는 실행안됨
tx.commit();
```



- 엔티티 매핑
  - 객체와 테이블 매핑 @Entity @Table
    - @Table(name="MBR") // 테이블을 다른이름으로 매핑하고 싶을때 (INSERT INTO MBR...)
  - 데이터베이스 스키마 자동생성
  - 필드와 컬럼 매핑 @Column
  - 기본키 매핑 @Id
  - 연관관계 매핑 @ManyToOne, @JoinColumn

- @Entity
  - JPA로 테이블 매핑할 클래스는 필수
  - '기본생성자' 필수 (public, protected)
  - final 클래스, enum, interface, inner클래스 사용 X
  - 저장할 필드에 final 사용 X

- 데이터베이스 스키마 자동 생성
  - '개발서버'에서만 사용 권유 (create, update)
  - '테스트서버'는 (update, validate)
  - '스테이징'/'운영서버'는 다듬은 후 사용. 주의! (validate, none)
  - name="hibernate.hbm2ddl.auto" value="create"
    - 애플리케이션 로딩 시점에 create table
    - value: create, create-drop, update, validate, none
    - update: 컬럼 지우는건 안됨. 컬럼 필드 Member에 추가하고 Drop 이 아닌 Alter로 DB반영
    - validate: 엔티티와 테이블이 정상 매핑되었는지 확인
    - none: 주석처리하거나, "none"
    - @Column(unique=true, length=10) DDL 생성기능 unique 제약조건
    - DDL 생성기능은 DDL을 자동 생성할 때만 사용되고, JPA 실행로직에는 영향 X

- 필드와 컬럼 매핑 @Column
  - @Column 컬럼매핑
    - name
    - insertable,updatable: TRUE디폴트. 컬럼수정 시 DB에 반영할지여부
    - nullable (DDL) : true디폴트
    - unique (DDL) : 제약조건명이 랜덤으로 생성되므로 잘안쓰고 직접 unique 제약 DDL 따로 실행
    - columnDefinition (DDL) 컬럼정의 직접하고싶을때
      - (..., columnDefinition="varchar(100) default `EMPTY`")
    - length (DDL)
    - precision, scale (DDL)
      - BigDecimal 타입에서 사용
  - @Temporal 날짜 타입 매핑: 요즘은 필요 X
    - LocalDate, LocalDateTime으로 쓰면됨 : 어노테이션없어도 JPA에서 생성해줌
    - private LocalDate testLocalDate;
    - private LocalDateTime testLocalDateTime;
  - @Enumerated enum 타입 매핑
    - EnumType.ORDINAL 기본값 enum 순서를 데이터베이스에 저장
    - EnumType.STRING 기본값 enum 이름을 데이터베이스에 저장
    - @Enumerated로 바꾸면 기본 디폴트 (EnumType.ORDINAL)
    - EnumType.ORDINAL로 하면 0,1,2 저장되는데 나중에 GUEST, USER, ADMIN 으로
    - 바뀌는경우 숫자구분이 꼬임-> EnumType.STRING 으로 쓰기!
  - @Lob BLOB, CLOB 매핑: 문자면 CLOB, 나머지는 BLOB 매핑
  - @Transient 특정필드를 컬럼에 매핑: 메모리에서만 씀-> DB컬럼 반영안됨

- 기본키 매핑 @Id
  - 직접할당 : @Id만 사용
  - 자동할당 : @GeneratedValue
    - IDENTITY 기본키 생성을 데이터베이스에 위임 (e.g. MySQL auto_increment)
    - SEQUENCE 오라클에서 주로 사용 (private Long id; Integer대신 Long 권장)
      - @SequenceGenerator(...)
      - call next value for MEMBER_SEQ로 seq 값을 받아와서 INSERT할때 MEMBER.ID에 넣어줌
      - allocationSize=50: DB에 50->100->.. 단위로 생성 하고, 메모리에서는 1씩
    - TABLE 키생성 전용 테이블 하나 만듦. 성능적으로 안좋음. 운영서버에서 쓰기 부담
      - @TableGenerator(name = "...", table = "MY_SEQUENCES", pkColumnValue = “MEMBER_SEQ", allocationSize = 1)
    - AUTO
  - 권장 식별자 전략
    - 기본키 제약조건: not null, unique, 변하면안됨(설계가 힘듦)
      - 주로 not null, unique 조건 까지
      - 대리키(대체키)로 사용 권장
      - *권장 : Long형 + 대체키 + 키 생성전략 사용
      - auto_increment 또는 uuid/랜덤 값을 기본키 값으로 권장

- IDENTITY 전략 특징
  - 기본키 생성을 데이터베이스에 위임 (e.g. MySQL auto_increment)
  - PK값이 DB에 들어가야 값을 알 수 있음
  - IDENTITY 전략에서만 예외적으로 tx.commit();이 아닌,
  - em.persist(member); 호출 시점에 바로 INSERT 쿼리 날림
  - AUTO_INCREMENT는 INSERT 실행 한 후에야 ID 알 수 있음
 
```java
  System.out.println("===");
  em.persist(m);
  System.out.println("id== " + m.getId()); // select없이 INSERT실행후 받아 옴
  System.out.println("===");
```

```java
@Entity
public class Member {
	@Id
	private Long id;

	@Column(name = "name")
	private String username;

	private Integer age;

	@Enumerated(EnumType.STRING)
	private RoleType roleType;

	@Temporal(TemporalType.TIMESTAMP)
	private Date createdDate;

	@Temporal(TemporalType.TIMESTAMP)
	private Date lastModifiedDate;

	@Lob
	private String description;

	public Member() {}
}

```

```sql
create table Member (
  id bigint not null,
  age integer,
  createdDate timestamp,
  description clob,
  lastModifiedDate timestamp,
  roleType varchar(255),
  name varchar(255),
  primary key (id)
)
```




- 실전예제1 - 요구사항 분석과 기본 매핑
  - jdbc:h2:~/jpashop 커넥트
  - jdbc:h2:tcp://localhost/~/jpashop 커넥트
  - @Entity의 테이블 및 컬럼 속성은 explicit 하는것이 좋음
    - @Table(indexes = @Index...)
    - @Column(length = 10)
    - boot에서 JPA 사용시, 기본설정 camelCase -> under_bar_name


- 관계형 DB에 맞춘 설계 (Order 등 클래스 필드에 Member가 아닌 memberId로 연결)

```java
  Order order = em.find(Order.class, 1L);
  Long memberId = order.getMemberId();
  Member member = em.find(Member.class, memberId);
```

- 객체지향적 설계 (Order 등 클래스 필드로 Member 정의)

```java
  // Order 클래스 멤버로 Member 필드 정의함
  Order order = em.find(Order.class, 1L);
  Member member = order.getMember();
```




- 주인관계 사용시 주의사항
  - Member, Team
  - 무한루프 가능: toString(), lombok, JSON 생성 라이브러리
    - lombok toString() 사용자제
    - 컨트롤러에서 return 타입 엔티티 사용하지 말기 => DTO로 변환해서 반환
  - 단방향 객체설계 후에 추후 필요시 수정 `Team 클래스에 private List<Member> members;`
  - 연관관계의 주인은 외래 키의 위치를 기준으로 정해야함 (e.g. `Member`가 외래키 들고있음)

- 1:1
  - 주 테이블 외래키
    - Member-Locker
    - Member에 FK(LOCKER_ID) 정의하는 방법 선호
    - Member 조회하는 경우가 많고, 조인없이 Locker에 대한 정보 얻을 수 있음
  - 대상 테이블 외래키

- 다:다
  - 실무 사용지양

