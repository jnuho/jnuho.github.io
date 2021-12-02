package springbook.user.dao;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class DaoFactory {

  @Bean
  public UserDao userDao() {
    return new UserDao(connectionMaker());
  }

  @Bean
  public ConnectionMaker connectionMaker() {
    return new NConnectionMaker();
    // 필요에 따라 어떤 ConnectionMaker 사용할 지 수정
//    return new LocalDBConnectionMaker();
//    return new ProductionlDBConnectionMaker();
  }
}
