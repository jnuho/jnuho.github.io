package hellojpa;

import javax.persistence.*;

@Entity
//@Table(UniqueConstraint = )
//@Entity(name="Member")
//@Table(name="USER") // 테이블명 명시
@SequenceGenerator(
    name="member_seq_generator",
    sequenceName = "member_seq", // 매핑할 데이터베이스 시퀀스 이름
    initialValue=1, allocationSize = 50) // 성능최적화 파라미터. 시퀀스 한번 호출에 증가하는 수. next call할때 미리 50개 가져옴
//@TableGenerator(
//    name="member_seq_generator",
//    table = " my_sequences",
//    pkColumnValue="member_seq", allocationSize = 1)
public class Member {

  @Id
//  @GeneratedValue(strategy = GenerationType.IDENTITY) // commit()전에 persist()때 DB반영됨 PK가 DB적용되야 생성되므로
  @GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "member_seq_generator")
// DB타입별 오라클 시퀀스,  MySQL auto_increment 대신 별도 테이블로 관리 (성능문제있음)
//  @GeneratedValue(strategy = GenerationType.TABLE, generator="member_seq_generator)
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
