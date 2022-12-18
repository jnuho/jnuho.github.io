

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
import javax.persistence.EntityManagerFactory;
import javax.persistence.EntityTransaction;
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
			
			Member member = new Member();
			member.setId(1L);
			member.setName("HelloA");
			//	영속 DB쿼리는 아직-> commit(); 시점에 INSERT 됨
      // 1차캐시에 저장됨
			em.persist(member);
			// 비슷한 방식으로 select 를 할때도 DB 조회후에 1차캐시에 저장!
			//  -> 2번째 조회시에는 1차캐시에서 가져옴

			// 1차캐시에서 조회함
      // 쿼리 select 쿼리 안나감(1차캐시)
			Member findMember = em.find(Member.class, 1L);
			System.out.println("findMember.Id=" + findMember.getId());
			System.out.println("findMember.Name=" + findMember.getName());

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


- 데이터베이스 스키마 자동 생성
  - 개발서버 에서만 사용
  - 운영은 다듬은 후 사용
  - hibernate.hbm2ddl.auto
    - update: 컬럼 지우는건 안됨

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

