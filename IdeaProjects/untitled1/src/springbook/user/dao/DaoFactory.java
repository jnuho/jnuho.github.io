package springbook.user.dao;

public class DaoFactory {

	public UserDao userDao() {
		return new UserDao(connectionMaker());
	}

//	public AccountDao userDao() {
//		return new AccountDao(connectionMaker());
//	}
//	public MessageDao userDao() {
//		return new MessageDao(connectionMaker());
//	}

	public ConnectionMaker connectionMaker() {
		return new DConnectionMaker();
	}
}
