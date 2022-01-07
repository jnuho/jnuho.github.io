



- Maven Project
- Spring
- GroupId : jpa-basic
- Artifact : ex1-hello-jpa
- Dependencies
  - JPA 하이버네이트
  - H2 데이터베이스

Intellij > new > maven
Intellij > open > pom.xml

- 테이블생성
- Member.java 클래스 @Entity 생성


- JPQL : 검색시 세부 조건 가능 age >=18

- 영속성 컨텍스트
  - 엔티티 조회, 1차캐시 key,value = (@Id, Entity객체)
  - 1차캐시에 없으면 DB조회 후 1차캐시에 저장


- 배치

- 데이터베이스 스키마 자동 생성
  - 개발서버 에서만 사용
  - 운영은 다듬은 후 사용
  - hibernate.hbm2ddl.auto
    - update: 컬럼 지우는건 안됨
