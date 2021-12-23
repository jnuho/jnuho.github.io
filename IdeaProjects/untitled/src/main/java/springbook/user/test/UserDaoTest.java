package springbook.user.test;

import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.ApplicationContext;
import org.springframework.context.support.GenericXmlApplicationContext;
import org.springframework.dao.EmptyResultDataAccessException;
import org.springframework.test.context.ContextConfiguration;
import org.springframework.test.context.junit.jupiter.SpringExtension;
import springbook.user.dao.DConnectionMaker;
import springbook.user.dao.UserDao;
import springbook.user.domain.User;

import javax.sql.DataSource;
import java.sql.SQLException;

import static org.hamcrest.CoreMatchers.is;
import static org.hamcrest.MatcherAssert.assertThat;
import static org.junit.jupiter.api.Assertions.assertThrows;

// JUnit 4
//@RunWith(SpringJUnit4ClassRunner.class)
//@ContextConfiguration({"/applicationContext.xml"})
// JUnit 5
// 스프링의 테스트 컨텍스트 프레임워크의 JUnit 확장기능 지점
@ExtendWith(SpringExtension.class)
// 테스트 컨텍스트가 자동으로 만들어줄 애플리케이션 컨텍스트의 위치 지정
// 모든 테스트 클레스에서 같은 애플리케이션 컨텍스트 객체가 사용됨
@ContextConfiguration({"/applicationContext.xml"})
public class UserDaoTest {

  // 테스트 오브젝트가 만들어지고 나면 스프링 테스트 컨텍스트에 의해 자동으로 값이 주입된다.
  // 하나의 애플리케이션 컨텍스트가 만들어져 모든 테스트 메소드에 사용됨 (같은 오브젝트 주소)
  @Autowired
  private ApplicationContext context;

  private UserDao dao;
  private User user1;
  private User user2;
  private User user3;

  private int counter;

  @BeforeAll
  public void setUp() {
    /** 의존관계 검색 */
    // 1. XML방식
//    ApplicationContext context = new GenericXmlApplicationContext("applicationContext.xml");
//    UserDao dao = context.getBean("userDao", UserDao.class); // 메소드 명
    // 2. 어노테이션 방식
//    ApplicationContext context = new AnnotationConfigApplicationContext(DaoFactory.class);
//    UserDao dao = new DaoFactory().userDao();

    // 커넥션 횟수
    DConnectionMaker pcm = context.getBean("connectionMaker", DConnectionMaker.class);
    this.counter = pcm.getCounter();

    user1 = new User("whiteship", "백기선", "married");
    user2 = new User("abc", "가나다", "123");
    user3 = new User("DEF", "라마바", "456");
  }

  @Test
  public void addAndGet() throws SQLException, ClassNotFoundException {

    // 커넥션 횟수
    System.out.println("커넥션 횟수 : "+ this.counter);

    dao.deleteAll();
    assertThat(dao.getCount(), is(0));

    dao.add(user1);
    assertThat(dao.getCount(), is(1));
    dao.add(user2);
    assertThat(dao.getCount(), is(2));
    dao.add(user3);
    assertThat(dao.getCount(), is(3));

    User userGet1 = dao.get(user1.getId());
    assertThat(userGet1.getName(), is(user1.getName()));
    assertThat(userGet1.getPassword(), is(user1.getPassword()));
  }

  @Test
  public void getUserFailure() throws SQLException {
    ApplicationContext context = new GenericXmlApplicationContext("applicationContext.xml");
    UserDao dao = context.getBean("userDao", UserDao.class);
    dao.deleteAll();
    assertThat(dao.getCount(), is(0));

    assertThrows(EmptyResultDataAccessException.class, () -> {
      dao.get("unkown_id");
    });
  }

}
