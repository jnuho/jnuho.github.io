package springbook.user.test;

import org.junit.jupiter.api.Test;
import org.springframework.context.ApplicationContext;
import org.springframework.context.support.GenericXmlApplicationContext;
import springbook.user.dao.DConnectionMaker;
import springbook.user.dao.UserDao;
import springbook.user.domain.User;

import java.sql.SQLException;

import static org.hamcrest.CoreMatchers.is;
import static org.hamcrest.MatcherAssert.assertThat;

public class UserDaoTest {

  @Test
  public void addAndGet() throws SQLException, ClassNotFoundException {

    /** 의존관계 검색 */
    // 1. XML방식
    ApplicationContext context = new GenericXmlApplicationContext("applicationContext.xml");
    UserDao dao = context.getBean("userDao", UserDao.class); // 메소드 명
    // 2. 어노테이션 방식
//    ApplicationContext context = new AnnotationConfigApplicationContext(DaoFactory.class);
//    UserDao dao = new DaoFactory().userDao();

    // 커넥션 횟수
    DConnectionMaker pcm = context.getBean("connectionMaker", DConnectionMaker.class);
    System.out.println("커넥션 횟수 : "+ pcm.getCounter());

    User user = new User();
    user.setId("whiteship");
    user.setName("백기선");
    user.setPassword("married");

    dao.add(user);

    User user2 = dao.get(user.getId());


    assertThat(user2.getName(), is(user.getName()));
    assertThat(user2.getPassword(), is(user.getPassword()));

  }
}
