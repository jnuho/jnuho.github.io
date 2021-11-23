package springbook.user.dao;

public class DaoFactory {

  public UserDao userDao() {
    UserDao userDao = new UserDao(new NConnectionMaker());
    return userDao;
  }
}
