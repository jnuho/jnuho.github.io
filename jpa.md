





















- Maven Project
- Spring
- GroupId : jpa-basic
- Artifact : ex1-hello-jpa
- Dependencies (pom.xml)
  - JPA 하이버네이트
  - H2 데이터베이스


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

Intellij > new > maven
Intellij > open > pom.xml

- 테이블생성
- Member.java 클래스 @Entity 생성


- JPQL : 검색시 세부 조건 가능 age >= 18
- 영속성 컨텍스트
  - 엔티티 조회, 1차캐시 key,value = (@Id, Entity객체)
  - 1차캐시에 없으면 DB조회 후 1차캐시에 저장

- 배치
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

