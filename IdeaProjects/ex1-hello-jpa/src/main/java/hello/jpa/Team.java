package hello.jpa;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.Id;

@Entity
public class Team {

	@Id @GeneratedValue
	@Column(name = "TEAM_ID")
	private Long id;

	private String name;

	// Member객체의 필드명 "team"을 맵핑
//	@OneToMany(mappedBy = "team")
//	private List<Member> members = new ArrayList<>();

//	public List<Member> getMembers() {
//		return members;
//	}
//
//	public void setMembers(List<Member> members) {
//		this.members = members;
//	}

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
