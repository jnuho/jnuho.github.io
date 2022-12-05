package springbook.user.dao;


import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

// 애플리케이션 컨텍스트 또는 빈 팩토리가 사용할 설정정보 라는 표시
@Configuration
public class DaoFactory {

	@Bean
	public UserDao userDao() {
		return new UserDao(connectionMaker());
	}

//	public AccountDao userDao() {
//		return new AccountDao(connectionMaker());
//	}
//	public MessageDao userDao() {
//		return new MessageDao(connectionMaker());
//	}

	@Bean
	public ConnectionMaker connectionMaker() {
		return new DConnectionMaker();
	}
}
