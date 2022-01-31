package com.spring.book.springboot.web.domain.posts;

import junit.framework.TestCase;
import org.junit.After;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;

import java.util.List;

import static org.hamcrest.MatcherAssert.assertThat;
import static org.hamcrest.Matchers.is;

@RunWith(SpringRunner.class)
@SpringBootTest
public class PostsRepositoryTest extends TestCase {

  @Autowired
  PostsRepository postsRepository;


  @After
  public void cleanup() {
    postsRepository.deleteAll();
  }

  @Test
  public void 게시글저장_불러오기() {
    //given
    String title = "테스트 게시글";
    String content = "테스트 본문";
    postsRepository.save(Posts.builder()
        .title(title)
        .content(content)
        .author("lopehih@gmail.com")
        .build()
    );

    //when
    List<Posts> postsList = postsRepository.findAll();

    //then
    Posts posts = postsList.get(0);
    assertThat(posts.getTitle(), is(title));
    assertThat(posts.getContent(), is(content));
  }
}