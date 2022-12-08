package springbook.user.dao;

import springbook.user.domain.User;

import java.sql.*;

public class UserDao {

	// 초기에 설정하면  사용 중에는 바뀌지 않는 읽기 전용 인스턴스 변수
	// DaoFactory에서 @Bean을 붙여 만들었기 때문에, 스프링이 관리하는 빈이 됨
	// 별다른 설정 없다면, 기본적으로 싱글톤 오브젝트 한개만 만들어져,
	// UserDao의 connectionMaker 인스턴스 필드에 저장됨!
	private ConnectionMaker connectionMaker;

	// 매번 새로운 값으로 바뀌는 정보를 담은 인스턴스 변수
	// : 멀티스레드 환경에서 사용하면 심각한 문제가 발생한다!
	//	private Connection c;
	//	private User user;

	public UserDao(ConnectionMaker connectionMaker) {
		this.connectionMaker = connectionMaker;
	}

	public void add(User user) throws ClassNotFoundException, SQLException {
		Connection c = connectionMaker.getConnection();
		PreparedStatement ps = c.prepareStatement(
				"insert into users(id, name, password) values(?,?,?)");

		System.out.printf("%s %s %s\n", user.getId(), user.getName(), user.getPassword());
		ps.setString(1, user.getId());
		ps.setString(2, user.getName());
		ps.setString(3, user.getPassword());

		ps.executeUpdate();

		ps.close();
		c.close();
	}

	public User get(String id) throws ClassNotFoundException, SQLException {
		Connection c = connectionMaker.getConnection();
		PreparedStatement ps = c.prepareStatement(
				"select * from users where id=?");
		ps.setString(1, id);

		ResultSet rs = ps.executeQuery();
		rs.next();

		User user = new User();
		user.setId(rs.getString("id"));
		user.setName(rs.getString("name"));
		user.setPassword(rs.getString("password"));

		rs.close();
		ps.close();
		c.close();

		return user;
	}
}
