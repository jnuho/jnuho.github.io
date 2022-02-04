package com.spring.book.springboot.web;

import com.spring.book.springboot.config.auth.LoginUser;
import com.spring.book.springboot.config.auth.dto.SessionUser;
import com.spring.book.springboot.service.posts.PostsService;
import com.spring.book.springboot.web.domain.posts.PostsResponseDto;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;

import javax.servlet.http.HttpSession;

@RequiredArgsConstructor
@Controller
public class IndexController {

  private final PostsService postsService;
  private final HttpSession httpSession;

  @GetMapping("/")
  public String index(Model model, @LoginUser SessionUser user) { // 서버 템플릿 엔진에서 사용하 수 있는 객체 저장 가능
    model.addAttribute("posts", postsService.findAllDesc());
    // @LoginUser는 어느 컨트롤러에서든 어노테이션만 사용하면 세션정보를 가져올 수 있음

    // 앞서 작성된 CustomOAuth2UserService에서 로그인 성공시에 세션에 SessionUser가 저장되도록 구성
    // 즉, 로그인 성공 시 httpSession의 어트리뷰트에서 "user" 값을 가져올 수 있음
    // 컨트롤러나 메소드에서 세션값이 필요할 때 마다 호출 해야함
    // 메소드인자로 세션 값을 바로 받을 수 있도록 개선: @LoginUser 정의하여 파라미터로 받음
    // 어노테이션 정의 및 WebConfig @Congiguration을 통해 파라미터롤 받을 수있게 Resolver 설정?
//    SessionUser user = (SessionUser) httpSession.getAttribute("user");

    if (user != null) {
      model.addAttribute("userName", user.getName());
    }

    return "index";
  }

  @GetMapping("/posts/save")
  public String postsSave() {
    return "posts-save";
  }

  @GetMapping("/posts/update/{id}")
  public String postsUpdate(@PathVariable Long id, Model model) {
    PostsResponseDto dto  = postsService.findById(id);
    model.addAttribute("post", dto);
    return "posts-update";
  }

}
