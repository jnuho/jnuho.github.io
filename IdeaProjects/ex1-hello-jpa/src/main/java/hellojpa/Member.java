package hellojpa;

import javax.persistence.Entity;
import javax.persistence.Id;

@Entity
//@Table(name="USER") // 테이블명 명시
public class Member {

  @Id
  private Long id;

  public Member() {}
  public Member(Long id, String name) {
    this.id = id;
    this.name = name;
  }

  //  @Column(name="username") // 컬럼명 명시
  private String name;

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
