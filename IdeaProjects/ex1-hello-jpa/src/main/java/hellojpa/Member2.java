package hellojpa;

import javax.persistence.*;
import java.time.LocalDate;
import java.time.LocalDateTime;
import java.util.Date;

//@Entity
//@Table(UniqueConstraint = )
//@Entity(name="Member")
//@Table(name="USER") // 테이블명 명시
public class Member2 {

  @Id // PK컬럼
  private Long id;

//  @Column(name = "name", updatable = false, nullable= false, unique = true)
  // 유니크 constraint 임의의 명칭으로 생성 되므로 unique 사용 자제
  // 클래스의 @Table(UniqueConstaint=)로 이름 정하여 사용)
  @Column(name = "name", nullable = false, columnDefinition = "varchar(100) default 'EMPTY'")
  private String username;

  private int age;

  @Enumerated(EnumType.STRING)
//  @Enumerated // 디폴트 ORDINAL: 순서
  // ORDINAL 로 하고 RoleType에 추가하면 ORDINAL 꼬일 수 있음
  private RoleType roleType;

  @Temporal(TemporalType.TIMESTAMP)
  private Date createdDate;

  @Temporal(TemporalType.TIMESTAMP)
  private Date lastModifiedDate;

  // 다음 라이브러리를 쓰면 @Temporal 생략해도 DB에 Date, Timestamp 타입으로 생성됨
  private LocalDate testLocalDate;
  private LocalDateTime testLocalDateTime;

  // 문자면 Clob, 나머지는 Blob으로 맵핑
  @Lob
  private String description;

  // DB 맵핑 없이, 메모리 저장 용도 필드
  @Transient
  private int temp;

  public Member2() {}

  public Long getId() {
    return id;
  }

  public void setId(Long id) {
    this.id = id;
  }

  public String getUsername() {
    return username;
  }

  public void setUsername(String username) {
    this.username = username;
  }

  public int getAge() {
    return age;
  }

  public void setAge(int age) {
    this.age = age;
  }

  public RoleType getRoleType() {
    return roleType;
  }

  public void setRoleType(RoleType roleType) {
    this.roleType = roleType;
  }
}
