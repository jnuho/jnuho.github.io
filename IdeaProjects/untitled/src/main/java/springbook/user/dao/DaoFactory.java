package springbook.user.dao;

import org.springframework.jdbc.datasource.SimpleDriverDataSource;

import javax.sql.DataSource;

//@Configuration
public class DaoFactory {

//  @Bean
  public UserDao userDao() {

    UserDao userDao = new UserDao(connectionMaker());
//    userDao.setDataSource(dataSource());
    return userDao;
  }

  public ConnectionMaker connectionMaker() {
    return new DConnectionMaker();
  }

//  @Bean
  public DataSource dataSource() {
    SimpleDriverDataSource dataSource = new SimpleDriverDataSource();
    dataSource.setDriverClass(com.mysql.jdbc.Driver.class);
    dataSource.setUsername("spring");
    dataSource.setPassword("book");
    return dataSource;
  }
}
