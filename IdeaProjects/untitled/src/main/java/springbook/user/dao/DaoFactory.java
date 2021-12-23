package springbook.user.dao;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.jdbc.datasource.SimpleDriverDataSource;

import javax.sql.DataSource;

@Configuration
public class DaoFactory {

  @Bean
  public UserDao userDao() {

    UserDao userDao = new UserDao();
//    userDao.setDataSource(dataSource());
//    userDao.setConnectionMaker(connectionMaker());
    return userDao;
  }

  @Bean
  public ConnectionMaker connectionMaker() {
    return new DConnectionMaker(realConnectionMaker());
  }

  @Bean
  public ConnectionMaker realConnectionMaker() {
    return new ProdConnectionMaker();
  }

  @Bean
  public DataSource dataSource() {
    SimpleDriverDataSource dataSource = new SimpleDriverDataSource();
    dataSource.setDriverClass(com.mysql.jdbc.Driver.class);
    dataSource.setUsername("spring");
    dataSource.setPassword("book");
    return dataSource;
  }
}
