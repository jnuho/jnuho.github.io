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
			Member member1 = new Member();
			member1.setUsername("hello1");
			em.persist(member1);

			em.flush();
			em.clear();

			Member m1 = em.find(Member.class, member1.getId());
			System.out.println("m1 = " + m1.getClass());
			Member ref = em.getReference(Member.class, member1.getId());
			System.out.println("ref = " + ref.getClass());
			System.out.println("m1 == ref " + (m1 == ref));

			tx.commit();
		} catch(Exception e) {
			tx.rollback();
		} finally {
			em.close();
		}

		emf.close();
	}
}
