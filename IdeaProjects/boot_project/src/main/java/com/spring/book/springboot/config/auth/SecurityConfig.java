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
          .antMatchers("/api/v1/**").hasRole(Role.USER.name())
          .anyRequest().authenticated() // 설정된 값들 이외의 나머지 URL들은 인증된 사용자(로그인된 사용자)에게만 허용
        .and()
          .logout()
            .logoutSuccessUrl("/") // 로그아웃시 이동 url
        .and()
          .oauth2Login()
            .userInfoEndpoint()
              .userService(customOAuth2UserService);

    super.configure(http);
  }
}
