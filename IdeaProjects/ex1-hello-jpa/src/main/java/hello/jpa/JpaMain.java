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

			Movie movie = new Movie();
			movie.setActor("aaaa");
			movie.setDirector("bbbb");
			movie.setName("No Country for old men");
			movie.setPrice(10000);
			em.persist(movie);

			em.flush(); // DB에 먼저 반영
			em.clear();

			Movie findMovie = em.find(Movie.class, movie.getId());
			System.out.println("findMovie = " + findMovie);

			tx.commit();
		} catch(Exception e) {
			tx.rollback();
		} finally {
			em.close();
		}

		emf.close();
	}
}
