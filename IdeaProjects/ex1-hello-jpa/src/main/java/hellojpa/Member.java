package hellojpa;

import javax.persistence.*;
import java.util.Date;

@Entity
//@Entity(name="Member")
//@Table(name="USER") // 테이블명 명시
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
  public Member(Long id, String name) {
    this.id = id;
    this.name = name;
  }

  public Long getId() {
    return id;
  }

  public void setId(Long id) {
    this.id = id;
  }

  public String getName() {
    return name;
  }

  public void setName(String name) {
    this.name = name;
  }
}
