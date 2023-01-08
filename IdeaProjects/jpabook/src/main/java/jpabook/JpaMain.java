package jpabook;

import jpabook.jpashop.domain.Order;
import jpabook.jpashop.domain.OrderItem;

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
			Order order = new Order();
			em.persist(order);

			OrderItem orderItem = new OrderItem();
			orderItem.setOrder(order);
			em.persist(orderItem);

			tx.commit();
		} catch(Exception e) {
			tx.rollback();
		} finally {
			em.close();
		}

		emf.close();
	}
}
