package hello.jpa;

import javax.persistence.EntityManager;
import javax.persistence.EntityManagerFactory;
import javax.persistence.EntityTransaction;
import javax.persistence.Persistence;
import java.util.List;

public class JpaMain {
	public static void main(String[] args) {
		// persistence.xml에 정의된 hello 설정정보 가져옴
		EntityManagerFactory emf = Persistence.createEntityManagerFactory("hello");
		EntityManager em = emf.createEntityManager();

		EntityTransaction tx = em.getTransaction();
		tx.begin();
		try {
			Team team = new Team();
			team.setName("TeamA");
			em.persist(team);

			Member member = new Member();
			member.setUsername("member1");
			member.changeTeam(team);
			em.persist(member);

			// 편의메소드 정의 Member.changeTeam(m){ .. 여기서 team.getMembers().add(this);... }
//			team.getMembers().add(member);
//			em.flush();
//			em.clear();

			// flush();clear(); 없고, team의 members에 add하지않으면, 밑에서 출력 안됨
			//1차 캐시 (team객체에 컬렉션 세팅안되어있음)
			Team findTeam = em.find(Team.class, team.getId());
			List<Member> members = findTeam.getMembers();
			for (Member m : members) {
				System.out.println("m=" + m.getUsername());
			}


			tx.commit();
		} catch(Exception e) {
			tx.rollback();
		} finally {
			em.close();
		}

		emf.close();
	}
}
