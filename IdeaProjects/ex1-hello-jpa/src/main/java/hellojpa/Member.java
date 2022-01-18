package hellojpa;

import javax.persistence.*;
import java.util.ArrayList;
import java.util.List;

@Entity
//@Table(UniqueConstraint = )
//@Entity(name="Member")
//@Table(name="USER") // 테이블명 명시
//@SequenceGenerator(
//    name="member_seq_generator",
//    sequenceName = "member_seq", // 매핑할 데이터베이스 시퀀스 이름
//    initialValue=1, allocationSize = 50) // 성능최적화 파라미터. 시퀀스 한번 호출에 증가하는 수. next call할때 미리 50개 가져옴
//@TableGenerator(
//    name="member_seq_generator",
//    table = " my_sequences",
//    pkColumnValue="member_seq", allocationSize = 1)
public class Member {

  @Id
  @GeneratedValue
//  @GeneratedValue(strategy = GenerationType.IDENTITY) // commit()전에 persist()때 DB반영됨 PK가 DB적용되야 생성되므로
//  @GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "member_seq_generator")
// DB타입별 오라클 시퀀스,  MySQL auto_increment 대신 별도 테이블로 관리 (성능문제있음)
//  @GeneratedValue(strategy = GenerationType.TABLE, generator="member_seq_generator)
  @Column(name = "MEMBER_ID")
  private Long id;

//  @Column(name="name", nullable= false)
  @Column(name="USERNAME")
  private String username;

//  @Column(name = "TEAM_ID")
//  private Long teamId;

  // FK있는 곳을 '주인'으로 지정
  // team변경 시 DB에 반영
//  @ManyToOne(fetch = FetchType.LAZY)
  @ManyToOne
  @JoinColumn(name = "TEAM_ID") // 연관관계 주인인 경우
//  @JoinColumn(name = "TEAM_ID", insertable = false, updatable = false) // 1:N맵핑 시, 연관관계 주인 아닌경우(읽기전용)
  private Team team;

  @OneToOne
  @JoinColumn(name = "LOCKER_ID")
  private Locker locker;


  // 다:다 맵핑
  @ManyToMany
  @JoinTable(name = "MEMBER_PRODUCT")
  private List<Product> products = new ArrayList<>();

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

  public Team getTeam() {
    return team;
  }

  public void changeTeam(Team team) {
    this.team = team;
    team.getMembers().add(this);
  }

}
