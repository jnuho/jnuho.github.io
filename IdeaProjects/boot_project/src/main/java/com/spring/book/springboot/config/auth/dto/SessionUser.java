package com.spring.book.springboot.config.auth.dto;

import com.spring.book.springboot.web.domain.user.User;
import lombok.Getter;

// SessionUser에는 인증된 사용자 정보만 필요
// 그외에 필요한 정보들은 없으니 name, email, picture만 필드로 선언
@Getter
public class SessionUser {
  private String name;
  private String email;
  private String picture;

  public SessionUser(User user) {
    this.name = user.getName();
    this.email = user.getEmail();
    this.picture = user.getPicture();
  }
}
