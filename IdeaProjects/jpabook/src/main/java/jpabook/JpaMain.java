package jpabook;

import jpabook.jpashop.domain.Member;
import jpabook.jpashop.domain.Team;

import javax.persistence.EntityManager;
import javax.persistence.EntityManagerFactory;
import javax.persistence.EntityTransaction;
import javax.persistence.Persistence;

public class JpaMain {
	public static void main(String[] args) {
		// persistence.xml에 정의된 hello 설정정보 가져옴
		EntityManagerFactory emf = Persistence.createEntityManagerFactory("hello");
		EntityManager em = emf.createEntityManager();

		EntityTransaction tx = em.getTransaction();
		tx.begin();
		try {
			Member member1 = new Member();
			member1.setUsername("member1");
			em.persist(member1);

			Member member2 = new Member();
			member2.setUsername("member2");
			em.persist(member2);

			em.flush();
			em.clear();

			Member m1 = em.find(Member.class, member1.getId());
			Member m1Ref = em.getReference(Member.class, member1.getId());
			Member m2 = em.getReference(Member.class, member2.getId());
//			System.out.println("findMember.id = " + findMember.getId());
//			System.out.println("findMember.username = " + findMember.getUsername());
			System.out.println("m1.getClass() = " + m1.getClass());
			System.out.println("m1.getClass() = " + m1Ref.getClass()); // 영속성 컨텍스트에 있는데 굳이 proxy 사용 X
			System.out.println("m1 == m1Ref = " + (m1 == m1Ref));
			System.out.println("m2.getClass() = " + m2.getClass()); // Proxy

			// 타입 체크 시 주의
			System.out.println("m1 == m2 " + logic(m1, m2));

			tx.commit();
		} catch(Exception e) {
			tx.rollback();
		} finally {
			em.close();
		}

		emf.close();
	}

	private static boolean logic(Member m1, Member m2) {
//		return m1.getClass() == m2.getClass(); // 한개는 프록시가 넘어온다면 false
		return (m1 instanceof Member) && (m2 instanceof Member);
	}

	private static void printMember(Member member) {
		System.out.println("member = " + member);
	}

	private static void printMemberAndTeam(Member member) {
		String username = member.getUsername();
		System.out.println("username = " + username);

		Team team = member.getTeam();
		System.out.println("team = " + team);
	}
}
