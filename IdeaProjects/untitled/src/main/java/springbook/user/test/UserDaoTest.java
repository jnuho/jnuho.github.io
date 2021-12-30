package springbook.user.test;

import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.TestInstance;
import org.junit.jupiter.api.extension.ExtendWith;
import org.springframework.context.ApplicationContext;
import org.springframework.context.support.GenericXmlApplicationContext;
import org.springframework.dao.EmptyResultDataAccessException;
import org.springframework.jdbc.datasource.SingleConnectionDataSource;
import org.springframework.test.annotation.DirtiesContext;
import org.springframework.test.context.ContextConfiguration;
import org.springframework.test.context.junit.jupiter.SpringExtension;
import springbook.user.dao.UserDao;
import springbook.user.domain.User;

import javax.sql.DataSource;
import java.sql.SQLException;

import static org.hamcrest.CoreMatchers.is;
import static org.hamcrest.MatcherAssert.assertThat;
import static org.junit.jupiter.api.Assertions.assertThrows;

// JUnit 4
//@RunWith(SpringJUnit4ClassRunner.class)
// JUnit 5
// 스프링의 테스트 컨텍스트 프레임워크의 JUnit 확장기능 지점
@ExtendWith(SpringExtension.class)
// 테스트 컨텍스트가 자동으로 만들어줄 애플리케이션 컨텍스트의 위치 지정
// 모든 테스트 클레스에서 같은 애플리케이션 컨텍스트 객체가 사용됨
@ContextConfiguration({"/applicationContext.xml"})
//@BeforeAll 의 non-static 메소드 사용을 위해
//클래스에 @TestInstance(TestInstance.Lifecycle.PER_CLASS) 추가
@TestInstance(TestInstance.Lifecycle.PER_CLASS)
// 테스트 메소드에서 애플리케이션 컨텍스트의 구성이나 상태를 변경한다는 것을
// 테스트 컨텍스트 프레임워크에 알려줌
@DirtiesContext
public class UserDaoTest {

  // 테스트 오브젝트가 만들어지고 나면 스프링 테스트 컨텍스트에 의해 자동으로 값이 주입된다.
  // 하나의 애플리케이션 컨텍스트가 만들어져 모든 테스트 메소드에 사용됨 (같은 오브젝트 주소)
//  @Autowired
//  private ApplicationContext context;

  private UserDao dao;

  private User user1;
  private User user2;
  private User user3;

//  private int counter;

  @BeforeAll
  public void setUp() {
    /** 의존관계 검색 */
    // 1. XML방식
    ApplicationContext context = new GenericXmlApplicationContext("applicationContext.xml");
    dao = context.getBean("userDao", UserDao.class); // 메소드 명
    // 2. 어노테이션 방식
//    ApplicationContext context = new AnnotationConfigApplicationContext(DaoFactory.class);
//    UserDao dao = new DaoFactory().userDao();

    // 커넥션 횟수
//    DConnectionMaker pcm = context.getBean("connectionMaker", DConnectionMaker.class);
//    this.counter = pcm.getCounter();
//    DataSource dataSource = new SimpleConnectionHandle

    // @DirtiesContext 사용 시 컨텍스트 구성 변경 (운영->로컬 테스트 DB)
    DataSource dataSource = new SingleConnectionDataSource(
        "jdbc:mysql://localhost/testdb", "spring", "book", true);
    dao.setDataSource(dataSource);

    user1 = new User("whiteship", "백기선", "married");
    user2 = new User("abc", "가나다", "123");
    user3 = new User("DEF", "라마바", "456");
  }

  @Test
  public void addAndGet() throws SQLException {

    // 커넥션 횟수
//    System.out.println("커넥션 횟수 : "+ this.counter);

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
//    ApplicationContext context = new GenericXmlApplicationContext("applicationContext.xml");
//    UserDao dao = context.getBean("userDao", UserDao.class);
    dao.deleteAll();
    assertThat(dao.getCount(), is(0));

    assertThrows(EmptyResultDataAccessException.class, () -> {
      dao.get("unkown_id");
    });
  }

}
