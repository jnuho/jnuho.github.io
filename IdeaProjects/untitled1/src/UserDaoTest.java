import springbook.user.dao.ConnectionMaker;
import springbook.user.dao.DConnectionMaker;
import springbook.user.dao.DaoFactory;
import springbook.user.dao.UserDao;
import springbook.user.domain.User;

import java.sql.SQLException;

public class UserDaoTest {

	public static void main(String[] args) throws ClassNotFoundException, SQLException {
		User user = new User();
		user.setId("102");
		user.setName("alice");
		user.setPassword("ururle1234");

		UserDao dao = new DaoFactory().userDao();
		dao.add(user);
		System.out.println(user.getId() + " 등록성공");

		User user2 = dao.get(user.getId());
		System.out.println(user2.getId() + " 조회성공");
		System.out.println(user2.getName());
		System.out.println(user2.getPassword());
	}
}