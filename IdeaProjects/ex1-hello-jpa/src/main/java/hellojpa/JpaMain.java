package hellojpa;

import javax.persistence.EntityManager;
import javax.persistence.EntityManagerFactory;
import javax.persistence.EntityTransaction;
import javax.persistence.Persistence;
import java.util.List;

public class JpaMain {
  public static void main(String[] args) {
    // 애플리케이션 전체에서 하나만 생성하여 사용
    EntityManagerFactory emf = Persistence.createEntityManagerFactory("hello");

    // 쓰렏간 공유 X. 사용하고 버림
    EntityManager em = emf.createEntityManager();

    EntityTransaction tx = em.getTransaction();
    tx.begin();

    try {
      // 등록
      // 비영속
//      Member member = new Member();
//      member.setId(101L);
//      member.setName("HelloJPA");

      // 영속
//      System.out.println("=== BEFORE ===");
//      // 1차캐시 (DB에 아직 적용안된 상태) commit() 시점에 쿼리 실행
//      em.persist(member);
//      System.out.println("=== AFTER ===");

      // 1차캐시에 없으면 DB조회-> 1차캐시 저장
      // 두번쨰 조회시에는 1차캐시에서 조회
//      Member findMember = em.find(Member.class, 101L);
//      System.out.println("findMember.id = " + findMember.getId());
//      System.out.println("findMember.name = " + findMember.getName());
//      Member findMember2 = em.find(Member.class, 101L);
//      System.out.println("findMember2.id = " + findMember2.getId());
//      System.out.println("findMember2.name = " + findMember2.getName());
//
//      System.out.println("findMember == findMember2 : " + (findMember == findMember2));
      // 준영속
      // 회원 엔티티를 영속성 컨텍스트에서 분리
//      em.detach(member);

      // 삭제
      // 객제츨 삭제한 상태
//      em.remove(member);

      // 조회
//      Member findMember = em.find(Member.class, 1L);
//      System.out.println(findMember.getId());
//      System.out.println(findMember.getName());

      // 삭제
//      Member findMember = em.find(Member.class, 1L);
//      em.remove(findMember);

      // 수정
//      Member findMember = em.find(Member.class, 1L);
//      findMember.setName("HelloJPA");

      // JPQL
//      Member findMember = em.find(Member.class, 1L);
//      List<Member> result = em.createQuery("select m from Member m", Member.class)
//              .getResultList();
//      List<Member> result = em.createQuery("select m from Member m", Member.class)
//          .setFirstResult(5)
//          .setMaxResults(8)
//          .getResultList();
//      for (Member m : result) {
//        System.out.println(m.getName());
//      }

      Member member1 = new Member(150L, "A");
      Member member2 = new Member(160L, "B");

      em.persist(member1);
      em.persist(member2);

      System.out.println("===");

      tx.commit();

    } catch(Exception e) {
      tx.rollback();
    } finally {
      em.close();
    }

    emf.close();
  }
}
