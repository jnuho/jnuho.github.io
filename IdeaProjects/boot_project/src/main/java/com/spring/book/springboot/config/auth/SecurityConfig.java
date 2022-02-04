package com.spring.book.springboot.config.auth;

import com.spring.book.springboot.web.domain.user.Role;
import lombok.RequiredArgsConstructor;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.config.annotation.web.configuration.EnableWebSecurity;
import org.springframework.security.config.annotation.web.configuration.WebSecurityConfigurerAdapter;

@RequiredArgsConstructor
@EnableWebSecurity // Spring Security 설정들을 활성화
public class SecurityConfig extends WebSecurityConfigurerAdapter {

  private final CustomOAuth2UserService customOAuth2UserService;

  @Override
  protected void configure(HttpSecurity http) throws Exception {
    http.csrf().disable() // h2-console화면 사용하기 위해 옵션 disable
        .headers().frameOptions().disable() // h2-console화면 사용하기 위해 옵션 disable
        .and()
          .authorizeRequests() // URL별 권한권리 설정 옵션의 시작점
          .antMatchers("/","/css/**", "/images/**", "/js/**", "/h2-console/**").permitAll()
          .antMatchers("/api/v1/**").hasRole(Role.USER.name()) // 해당 URL은 USER권한 가진 사람들만 허용
          .anyRequest().authenticated() // 설정된 값들 이외의 나머지 URL(anyRequest()): 인증된 사용자(authenticated/로그인된 사용자)에게만 허용
        .and()
          .logout()
            .logoutSuccessUrl("/") // 로그아웃시 이동 url
        .and()
          .oauth2Login() // OAuth2 로그인 기능에 대한 설정
            .userInfoEndpoint() // 로그인 성공 이후 사용자들의 정보를 가져올 때의 설정
            // 로그인 성공 후 후속조치 진행할 구현체
            // 리소스 서버(소셜 서비스)에서 사용자 정보를 가져온 상태에서, 추가로 진행 하고자 하는 기능 명시
            .userService(customOAuth2UserService);
  }
}
