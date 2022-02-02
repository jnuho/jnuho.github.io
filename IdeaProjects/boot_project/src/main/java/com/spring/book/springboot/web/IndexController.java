package com.spring.book.springboot.web;

import com.spring.book.springboot.service.posts.PostsService;
import com.spring.book.springboot.web.domain.posts.PostsResponseDto;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;

@RequiredArgsConstructor
@Controller
public class IndexController {

  private final PostsService postsService;

  @GetMapping("/")
  public String index(Model model) { // 서버 템플릿 엔진에서 사용하 수 있는 객체 저장 가능
    model.addAttribute("posts", postsService.findAllDesc());
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
