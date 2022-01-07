package hellojpa;

import hellojpa.RoleType;

import javax.persistence.*;
import java.util.Date;

@Entity
//@Table(UniqueConstraint = )
//@Entity(name="Member")
//@Table(name="USER") // 테이블명 명시
public class Member {
  @Id
  private Long id;

//  @Column(name = "name", updatable = false, unique = true)
  @Column(name = "name", nullable = false, columnDefinition = "varchar(100) default 'EMPTY'")
  private String username;

  @Column
  private int age;

//  @Enumerated(EnumType.STRING)
  @Enumerated // 기본 ORDINAL
  private RoleType roleType;

  @Temporal(TemporalType.TIMESTAMP)
  private Date createdDate;

  @Temporal(TemporalType.TIMESTAMP)
  private Date lastModifiedDate;

  @Lob
  private String description;

  // DB 맵핑 없이, 메모리에 저장
  @Transient
  private int temp;

  public Member() {}

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
}
