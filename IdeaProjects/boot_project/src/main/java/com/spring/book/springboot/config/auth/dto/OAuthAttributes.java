package com.spring.book.springboot.config.auth.dto;

import com.spring.book.springboot.web.domain.user.Role;
import com.spring.book.springboot.web.domain.user.User;
import lombok.Builder;
import lombok.Getter;
import org.springframework.security.oauth2.core.OAuth2TokenValidator;

import java.util.Map;

@Getter
public class OAuthAttributes {
  private Map<String, Object> attributes;
  private String nameAttributeKey;
  private String name;
  private String email;
  private String picture;

  @Builder
  public OAuthAttributes(Map<String, Object> attributes,
                         String nameAttributeKey, String name, String email, String picture) {
    this.attributes = attributes;
    this.nameAttributeKey = nameAttributeKey;
    this.name = name;
    this.email = email;
    this.picture = picture;
  }

  /**
   * OAuth2User 에서 반환하는 사용자정보는 Map이기 떄문에 값 하나하나를 변환해야 함
   * @param registrationId
   * @param userNameAttributeName
   * @param attributes
   * @return
   */
  public static OAuthAttributes of(String registrationId,
                                   String userNameAttributeName,
                                   Map<String, Object> attributes) {
    return ofGoogle(userNameAttributeName, attributes);
  }

  private static OAuthAttributes ofGoogle(String userNameAttributeName,
                                          Map<String, Object> attributes) {
    return OAuthAttributes.builder()
        .name((String)attributes.get("name"))
        .email((String)attributes.get("email"))
        .picture((String)attributes.get("picture"))
        .attributes(attributes)
        .nameAttributeKey(userNameAttributeName)
        .build();
  }

  /**
   * OAuthAttributes에서 엔터티를 생성하는 시점은 처음 가입할 때
   * 가입할 때 기본권한은 GUEST로 설정
   * @return
   */
  public User toEntity() {
    return User.builder()
        .name(name)
        .email(email)
        .picture(picture)
        .role(Role.GUEST)
        .build();
  }

}
