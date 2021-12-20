package springbook.user.test;

import org.junit.jupiter.api.Test;
import springbook.user.dao.DaoFactory;
import springbook.user.dao.UserDao;
import springbook.user.domain.User;

import java.sql.SQLException;

import static org.hamcrest.CoreMatchers.is;
import static org.hamcrest.MatcherAssert.assertThat;

public class UserDaoTest {

  @Test
  public void addAndGet() throws SQLException, ClassNotFoundException {

    // 의존관계 검색
//    ApplicationContext context = new GenericXmlApplicationContext("applicationContext.xml");
//    UserDao dao = context.getBean("userDao", UserDao.class); // 메소드 명

    UserDao dao = new DaoFactory().userDao();
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
