package springbook.user.dao;

import org.springframework.context.ApplicationContext;
import org.springframework.context.support.ClassPathXmlApplicationContext;

import java.sql.SQLException;

public class UserDaoConnectionCountingTest {

  public static void main(String[] args) throws ClassNotFoundException, SQLException {
//    ApplicationContext context
//        = new AnnotationConfigApplicationContext(CountingDaoFactory.class);
    // DaoFactory 대신 XML 파일 사용하도록 설정
//    ApplicationContext context = new GnericXmlApplicationContext("springbook/user/dao/applicationConext.xml");
    ApplicationContext context = new ClassPathXmlApplicationContext("applicationConext.xml", UserDao.class);
    UserDao dao = context.getBean("userDao", UserDao.class);

    // DAO 사용 코드
    CountingConnectionMaker ccm = context.getBean("connectionMaker", CountingConnectionMaker.class);
    System.out.println("Connection count : " + ccm.getCounter());
  }
}
