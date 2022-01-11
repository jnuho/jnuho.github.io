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
      Member member = new Member();
      member.setUsername("C");
      em.persist(member);
      System.out.println("member.id " + member.getId());

      tx.commit();

    } catch(Exception e) {
      tx.rollback();
    } finally {
      em.close();
    }

    emf.close();
  }
}
