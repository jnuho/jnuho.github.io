package com.spring.book.springboot.web.domain.posts;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;

import java.util.List;

// DB Layer  접근자 (MyBatis에서는 흔히 Dao)
// Entity(Posts)와 Entity Repository (PostsRepository) 함께 위치
// Entity 클래스는 기본 Repository 없이 제대로 역할 X
// JpaRepository<엔터티 클래스, PK타입>를 상속하면 기본적인 CRUD 메소드가 자동으로 생성됨
//  @Repository 추가할 필요 X
public interface PostsRepository extends JpaRepository<Posts, Long> {

  @Query("SELECT p FROM Posts p ORDER BY p.id DESC")
  List<Posts> findAllDesc();

}
