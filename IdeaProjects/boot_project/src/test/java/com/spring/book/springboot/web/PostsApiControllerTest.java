package com.spring.book.springboot.web;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.spring.book.springboot.web.domain.posts.Posts;
import com.spring.book.springboot.web.domain.posts.PostsRepository;
import com.spring.book.springboot.web.domain.posts.PostsSaveRequestDto;
import com.spring.book.springboot.web.domain.posts.PostsUpdateRequestDto;
import org.junit.After;
import org.junit.Before;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.web.client.TestRestTemplate;
import org.springframework.boot.web.server.LocalServerPort;
import org.springframework.http.HttpEntity;
import org.springframework.http.MediaType;
import org.springframework.security.test.context.support.WithMockUser;
import org.springframework.test.context.junit4.SpringRunner;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.setup.MockMvcBuilders;
import org.springframework.web.context.WebApplicationContext;

import java.util.List;

import static org.hamcrest.Matchers.is;
import static org.junit.Assert.assertThat;
import static org.springframework.security.test.web.servlet.setup.SecurityMockMvcConfigurers.springSecurity;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.post;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.put;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@RunWith(SpringRunner.class)
@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
public class PostsApiControllerTest {

  @LocalServerPort
  private int port;

  @Autowired
  private TestRestTemplate restTemplate;

  @Autowired
  private PostsRepository postsRepository;

  @Autowired
  private WebApplicationContext context;

  private MockMvc mvc;

  @Before
  public void setUp() {
    mvc = MockMvcBuilders
        .webAppContextSetup(context)
        .apply(springSecurity())
        .build();
  }

  @After
  public void tearDown() throws Exception {
    postsRepository.deleteAll();
  }


  // 모의사용자로 테스트
  // ROLE_USER 권한을 가진 사용자가 API요청 하는것과 같은 효과
  // 단 MockMvc에서만 작동하기 때문에 @Before에서 매번 테스트 시작하기 전에
  // MockMvc 인스턴스를 생성해 줌
  @Test
  @WithMockUser(roles="USER")
  public void posts_등록() throws Exception {
    //given
    String title = "title";
    String content = "content";
    PostsSaveRequestDto requestDto = PostsSaveRequestDto
        .builder()
        .title(title)
        .content(content)
        .author("author")
        .build();
   String url  = "http://localhost:" + port + "/api/v1/posts";

    //when
//    ResponseEntity<Long> responseEntity = restTemplate
//        .postForEntity(url, requestDto, Long.class);
    mvc.perform(post(url)
        .contentType(MediaType.APPLICATION_JSON_UTF8)
        .content(new ObjectMapper().writeValueAsString(requestDto)))
        .andExpect(status().isOk());

    //then
//    assertThat(responseEntity.getStatusCode(), is(HttpStatus.OK));
//    assertThat(responseEntity.getBody(), greaterThan(0L));

    List<Posts> all = postsRepository.findAll();
    assertThat(all.get(0).getTitle(), is(title));
    assertThat(all.get(0).getContent(), is(content));
  }

  @Test
  @WithMockUser(roles="USER")
  public void posts_수정() throws Exception {
    //given
    Posts savedPosts = postsRepository.save(Posts.builder()
        .title("title")
        .content("content")
        .author("author")
        .build()
    );
    
    Long updateId = savedPosts.getId();
    String expectedTitle = "title2";
    String expectedContent  = "content2";

    PostsUpdateRequestDto requestDto =
        PostsUpdateRequestDto.builder()
            .title(expectedTitle)
            .content(expectedContent)
            .build();

    String url = "http://localhost:" + port + "/api/v1/posts/" + updateId;
    HttpEntity<PostsUpdateRequestDto> requestEntity = new HttpEntity<>(requestDto);

    //when
//    ResponseEntity<Long> responseEntity = restTemplate
//        .exchange(url, HttpMethod.PUT, requestEntity, Long.class);
    mvc.perform(put(url)
        .contentType(MediaType.APPLICATION_JSON_UTF8)
        .content(new ObjectMapper().writeValueAsString(requestDto)))
        .andExpect(status().isOk());

    //then
//    assertThat(responseEntity.getStatusCode(), is(HttpStatus.OK));
//    assertThat(responseEntity.getBody(), greaterThan(0L));

    List<Posts> all = postsRepository.findAll();
    assertThat(all.get(0).getTitle(), is(expectedTitle));
    assertThat(all.get(0).getContent(), is(expectedContent));
  }
}