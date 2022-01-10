package hellojpa;

import javax.persistence.*;
import java.math.BigInteger;
import java.time.LocalDate;
import java.time.LocalDateTime;
import java.util.Date;

@Entity
//@Table(UniqueConstraint = )
//@Entity(name="Member")
//@Table(name="USER") // 테이블명 명시
public class Member {

  @Id
//  @GeneratedValue(strategy = GenerationType.IDENTITY)
  @GeneratedValue(strategy = GenerationType.SEQUENCE)
  private Long id;

  @Column(name="name", nullable= false)
  private String username;

  public Member(){}

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
}
