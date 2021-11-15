package springbook.user.dao;

import java.sql.Connection;
import java.sql.SQLException;

public class NUserDao extends UserDao {

  @Override
  public Connection getConnecton() throws ClassNotFoundException, SQLException {
    return null;
  }
}
