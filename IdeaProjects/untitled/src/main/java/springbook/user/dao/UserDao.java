package springbook.user.dao;

import org.junit.jupiter.api.BeforeAll;
import org.springframework.context.ApplicationContext;
import org.springframework.context.support.GenericXmlApplicationContext;
import org.springframework.dao.EmptyResultDataAccessException;
import springbook.user.domain.User;

import javax.sql.DataSource;
import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.SQLException;

public class UserDao {
//  private ConnectionMaker connectionMaker;
//  public void setConnectionMaker(ConnectionMaker connectionMaker) {
//    this.connectionMaker = connectionMaker;
//  }

  private DataSource dataSource;

  public void setDataSource(DataSource dataSource) {
    this.dataSource = dataSource;
  }

  public void add(User user) throws ClassNotFoundException, SQLException {
//    Connection c = getConnection();
//    Connection c = connectionMaker.makeConnection();
    Connection c = dataSource.getConnection();
    PreparedStatement ps  = c.prepareStatement("insert into users(id, name, password) values(?,?,?)");
    ps.setString(1, user.getId());
    ps.setString(2, user.getName());
    ps.setString(3, user.getPassword());
    ps.executeUpdate();

    ps.close();
    c.close();
  }

  public User get(String id) throws ClassNotFoundException, SQLException {
//    Connection c = getConnection();
//    Connection c = connectionMaker.makeConnection();
    Connection c = dataSource.getConnection();
    PreparedStatement ps  = c.prepareStatement("select from users where id=?");
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

    if (null == user) {
      throw new EmptyResultDataAccessException(1);
    }

    return user;
  }

  public void deleteAll() throws SQLException {
    Connection c = dataSource.getConnection();
    PreparedStatement ps = c.prepareStatement("delete from users");
    ps.executeUpdate();

    ps.close();
    c.close();
  }

  public int getCount() throws SQLException {
    Connection c = dataSource.getConnection();
    PreparedStatement ps = c.prepareStatement("select count(*) from users");

    ResultSet rs = ps.executeQuery();
    rs.next();
    int count = rs.getInt(1);

    rs.close();
    ps.close();
    c.close();

    return count;
  }

}
