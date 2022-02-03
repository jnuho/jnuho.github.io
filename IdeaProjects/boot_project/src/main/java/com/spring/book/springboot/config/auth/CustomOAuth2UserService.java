package com.spring.book.springboot.config.auth;

import com.spring.book.springboot.config.auth.dto.OAuthAttributes;
import com.spring.book.springboot.config.auth.dto.SessionUser;
import com.spring.book.springboot.web.domain.user.User;
import com.spring.book.springboot.web.domain.user.UserRepository;
import lombok.RequiredArgsConstructor;
import org.apache.catalina.manager.util.SessionUtils;
import org.springframework.security.core.authority.SimpleGrantedAuthority;
import org.springframework.security.oauth2.client.userinfo.DefaultOAuth2UserService;
import org.springframework.security.oauth2.client.userinfo.OAuth2UserRequest;
import org.springframework.security.oauth2.client.userinfo.OAuth2UserService;
import org.springframework.security.oauth2.core.OAuth2AuthenticationException;
import org.springframework.security.oauth2.core.user.DefaultOAuth2User;
import org.springframework.security.oauth2.core.user.OAuth2User;
import org.springframework.stereotype.Service;

import javax.servlet.http.HttpSession;
import java.util.Collections;

/**
 * 구글로그인에서 가져온 사용자정보(email, name, picture)들을 기반으로
 * 가입 및 정보수정, 세션저장 등의 기능을 지원
 */
@RequiredArgsConstructor
@Service
public class CustomOAuth2UserService implements OAuth2UserService<OAuth2UserRequest, OAuth2User> {

  private final UserRepository userRepository;
  private final HttpSession httpSession;


  @Override
  public OAuth2User loadUser(OAuth2UserRequest userRequest) throws OAuth2AuthenticationException {

    OAuth2UserService<OAuth2UserRequest, OAuth2User> delegate = new DefaultOAuth2UserService();
    OAuth2User oAuth2User = delegate.loadUser(userRequest);

    // 소셜로그인 서비스 구분 하기위한 Id (구글,네이버, 카카오, ...)
    String registrationId = userRequest.getClientRegistration().getRegistrationId();
    // OAuth2 로그인 진행 시 키가 되는 값들 - Primary Key
    // 구글의 경우만 기본적으로 코드지원: "sub"
    // 이후 네이버 로그인과 구글 로그인을 동시 지원할 때 사용
    String userNameAttributeName = userRequest.getClientRegistration()
        .getProviderDetails()
        .getUserInfoEndpoint()
        .getUserNameAttributeName();
    // OAuth2UserService를 통해 가져온 OAuth2User의 attribute를 담은 클래스
    OAuthAttributes attributes = OAuthAttributes
        .of(registrationId, userNameAttributeName, oAuth2User.getAttributes());

    User user = saveOrUpdate(attributes);
    // 세션에 사용자 정보를 저장하기 위한 Dto 클래스 - SessionUser
    // User 클래스를 쓰지않고 Session사용자를 따로 만들어서 쓰는 이유
    httpSession.setAttribute("user", new SessionUser(user));

    return new DefaultOAuth2User(
        Collections.singleton(new SimpleGrantedAuthority(user.getRoleKey())),
            attributes.getAttributes(),
            attributes.getNameAttributeKey());
  }

  private User saveOrUpdate(OAuthAttributes attributes) {
    User user = userRepository.findByEmail(attributes.getEmail())
        .map(entity -> entity.update(attributes.getName(), attributes.getPicture()))
        .orElse(attributes.toEntity());
    return userRepository.save(user);
  }
}
