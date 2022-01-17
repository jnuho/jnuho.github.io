package hellojpa;

import javax.persistence.*;
import java.util.ArrayList;
import java.util.List;

@Entity
public class Team {

  @Id @GeneratedValue
  @Column(name="TEAM_ID")
  private Long id;

  private String name;

  @OneToMany(mappedBy = "team") // 읽기만 가능. members수정해도 반영 X
  @JoinColumn(name ="TEAM_ID") // 1:N맵핑만 만드는경우추가!
  // team.getMembers().add(member); // 이떄 DB.Member테이블에 업데이트 실행됨
  private List<Member> members = new ArrayList<>();

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

  public List<Member> getMembers() {
    return members;
  }

  public void setMembers(List<Member> members) {
    this.members = members;
  }
}
