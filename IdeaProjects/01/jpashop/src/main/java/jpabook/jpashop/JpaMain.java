package jpabook.jpashop;

import com.sun.org.apache.xpath.internal.operations.Or;
import jpabook.jpashop.domain.Member;
import jpabook.jpashop.domain.Order;
import jpabook.jpashop.domain.OrderItem;

import javax.persistence.EntityManager;
import javax.persistence.EntityManagerFactory;
import javax.persistence.EntityTransaction;
import javax.persistence.Persistence;

public class JpaMain {
  public static void main(String[] args) {
    // 애플리케이션 전체에서 하나만 생성하여 사용
    EntityManagerFactory emf = Persistence.createEntityManagerFactory("hello");

    // 쓰렏간 공유 X. 사용하고 버림
    EntityManager em = emf.createEntityManager();

    EntityTransaction tx = em.getTransaction();
    tx.begin();

    try {

      Order order = new Order();
      em.persist(order);

      OrderItem orderItem = new OrderItem();
      orderItem.setOrder(order);

      tx.commit();

    } catch(Exception e) {
      tx.rollback();
    } finally {
      em.close();
    }

    emf.close();
  }
}
