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
//      Member member1 = new Member();
//      member1.setUsername("A");
//
//      Member member2 = new Member();
//      member2.setUsername("B");
//
//      Member member3 = new Member();
//      member3.setUsername("C");
//
//      em.persist(member1);
//      em.persist(member2);
//      em.persist(member3);

      // identity전략: DB insert후 JPA로 id에 세팅. 버퍼링 없이 insert
      // sequence전략: DB next 시퀀스 가져와서 id에 세팅. 버퍼링을 통해 insert는 commit()에 실행
//      System.out.println("==========");
//      System.out.println("member1.id " + member1.getId());
//      System.out.println("member2.id " + member2.getId());
//      System.out.println("member3.id " + member3.getId());
//      System.out.println("==========");
      //DB SEQ - 1  | WAS 1
      //DB SEQ - 51 | WAS 2
      //DB SEQ - 51 | WAS 3

      Team team = new Team();
      team.setName("TeamA");
      em.persist(team);

      Member member = new Member();
      member.setUsername("member1");
//      member.setTeam(team); // 주인은 Member이므로 DB반영 됨
      member.changeTeam(team); // 주인은 Member이므로 DB반영 됨
      em.persist(member);

      // DB반영 안됨. 주인은 Member이므로
      // 다음 라인 실행 안하고 em.flush()안할 경우 DB반영 commit()전에 없기때문에 'm=' 출력이 안됨
      // 객체지향적인 관점으로 다음 라인을 실행하도록 함 (테스트시에도 유용)
      // 양방향 연관관계의경우 양쪽에 세팅하는 걸로
      // 양방향 세팅 human error 있을 수 있으므로, 연관관계 편의 메소드 생성:
      // Member의 setTeam메소드를 changeTeam으로 수정
      // 또는 Team의 addMember메소드 정의하여 그 안에서 member.setTeam()
//      public void setTeam(Team team) {
//
//        this.team = team;
//        team.getMembers().add(this);
//      }

//      team.getMembers().add(member);

      em.flush();
      em.clear(); // 영속성컨텍스트 초기화

      Team findTeam = em.find(Team.class, team.getId());
      List<Member> members = findTeam.getMembers();
      for (Member m : members) {
        System.out.println("m = " + m.getUsername());
        System.out.println("m.team = " + findTeam);
      }

      tx.commit();

    } catch(Exception e) {
      tx.rollback();
    } finally {
      em.close();
    }

    emf.close();
  }
}
