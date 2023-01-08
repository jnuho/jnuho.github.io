package hello.jpa;

import javax.persistence.*;
import java.time.LocalDateTime;

@Entity
public class Member extends BaseEntity {
	@Id
	@GeneratedValue(strategy = GenerationType.AUTO)
	@Column(name = "MEMBER_ID")
	private Long id;

	@Column(name = "USERNAME")
	private String username;

//	@Column(name = "TEAM_ID")
//	private Long teamId;

	// 객체 team과 DB 조인컬럼 TEAM_ID 명시
	@ManyToOne
	@JoinColumn(name = "TEAM_ID")
	private Team team;

	private String createdBy;
	private LocalDateTime createdDate;
	private String lastModifiedBy;
	private LocalDateTime lastModifiedDate;

	public Team getTeam() {
		return team;
	}

	public void changeTeam(Team team) {
		this.team = team;
//		team.getMembers().add(this);
	}

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
