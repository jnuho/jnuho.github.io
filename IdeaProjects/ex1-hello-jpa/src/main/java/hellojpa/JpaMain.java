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
      Member member1 = new Member();
      member1.setUsername("A");

      Member member2 = new Member();
      member2.setUsername("B");

      Member member3 = new Member();
      member3.setUsername("C");

      em.persist(member1);
      em.persist(member2);
      em.persist(member3);

      // identity전략: DB insert후 JPA로 id에 세팅. 버퍼링 없이 insert
      // sequence전략: DB next 시퀀스 가져와서 id에 세팅. 버퍼링을 통해 insert는 commit()에 실행
      System.out.println("==========");
      System.out.println("member1.id " + member1.getId());
      System.out.println("member2.id " + member2.getId());
      System.out.println("member3.id " + member3.getId());
      System.out.println("==========");
      //DB SEQ - 1  | WAS 1
      //DB SEQ - 51 | WAS 2
      //DB SEQ - 51 | WAS 3

      tx.commit();

    } catch(Exception e) {
      tx.rollback();
    } finally {
      em.close();
    }

    emf.close();
  }
}
