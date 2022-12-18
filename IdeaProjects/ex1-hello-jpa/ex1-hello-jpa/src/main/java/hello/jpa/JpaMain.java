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
			// 1. Create 등록

			// 비영속
//			Member member = new Member();
//			member.setId(1L);
//			member.setName("HelloA");
			//	영속 DB쿼리는 아직-> commit(); 시점에 INSERT 됨
//			em.persist(member);

			// 2. Read 조회
//			Member findMember = em.find(Member.class, 1L);
//			System.out.println("findMember.Id=" + findMember.getId());
//			System.out.println("findMember.Name=" + findMember.getName());
			List<Member> result = em.createQuery("select m from Member as m", Member.class)
					.setFirstResult(5) // 5번 부터
					.setMaxResults(8)  // 8개 데이터
					.getResultList();
			for (Member m : result) {
				System.out.println("member.name = " + m.getName());
			}

			// 3. Update 수정
			Member findMember = em.find(Member.class, 1L);
			findMember.setName("HelloJPA");
			// em.persist(member); 필요없음! 마치 컬렉션에 저장하는것 같음
			// -> UPDATE쿼리 tx.commit(); 시점에 발생!

			// 4. Delete 삭제
//			em.remove(findMmeber);

			// 커밋!
			tx.commit();

		} catch(Exception e) {
			tx.rollback();
		} finally {
			em.close();
		}

		emf.close();
	}
}
