package com.spring.book.springboot.service.posts;

import com.spring.book.springboot.web.domain.posts.*;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

import javax.transaction.Transactional;
import java.beans.Transient;

@RequiredArgsConstructor
@Service
public class PostsService {

  private final PostsRepository postsRepository;

  @Transactional
  public Long save(PostsSaveRequestDto requestDto) {
    return postsRepository.save(requestDto.toEntity()).getId();
  }

  @Transactional
  public Long update(Long id, PostsUpdateRequestDto requestDto) {
    Posts posts = postsRepository.findById(id)
        .orElseThrow(
            () -> new IllegalArgumentException("해당 게시글이 없습니다. id=" + id)
        );
    posts.update(requestDto.getTitle(), requestDto.getContent());
    return id;
  }

  public PostsResponseDto findById(Long id) {
    Posts entity = postsRepository.findById(id)
        .orElseThrow(
            () -> new IllegalArgumentException("해당 게시글이 없습니다. id=" + id)
        );
    return new PostsResponseDto(entity);
  }

//  @Transactional(readOnly = true)
//  // 트랜젝션 범위는 유지하되, 조회 기능만 남겨두어 조회속도 개선
//  // .map(posts -> new PostsListResponseDto(posts)) 와 같은 표현식
//  public List<PostsListResponseDto> findAllDesc() {
//    return postsRepository.findAllDesc().stream()
//      .map(PostsListResponseDto::new)
//      .collect(Collectiors.toList());
//  }

}
