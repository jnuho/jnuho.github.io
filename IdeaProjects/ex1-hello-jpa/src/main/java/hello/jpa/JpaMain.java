package hello.jpa;

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

			Member member = new Member();
			member.setUsername("hello");

			em.persist(member);

			em.flush();
			em.clear();

//			Member findMember = em.find(Member.class, member.getId());
			Member findMember = em.getReference(Member.class, member.getId());
			System.out.println("findMember = " + findMember.getClass());
			System.out.println("findMember.id = " + findMember.getId());
			System.out.println("findMember.username = " + findMember.getUsername());

			tx.commit();
		} catch(Exception e) {
			tx.rollback();
		} finally {
			em.close();
		}

		emf.close();
	}
}
