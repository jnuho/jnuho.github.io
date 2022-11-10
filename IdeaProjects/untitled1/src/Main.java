import springbook.user.dao.UserDao;
import springbook.user.domain.User;

import java.sql.SQLException;

public class Main {

	public static void main(String[] args) throws ClassNotFoundException, SQLException {
		User user = new User();
		user.setId("10");
		user.setName("abc");
		user.setPassword("123");

		UserDao dao = new UserDao();
		dao.add(user);
		System.out.println(user.getId() + " 등록성공");

		User user2 = dao.get(user.getId());
		System.out.println(user.getId() + " 조회성공");
		System.out.println(user.getName());
		System.out.println(user.getPassword());
	}
}