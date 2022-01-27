package com.spring.book.springboot.web.dto;

import lombok.Getter;
import lombok.RequiredArgsConstructor;

@Getter
@RequiredArgsConstructor // final있는 필드 포함된 생성자
public class HelloResponseDto {

  private final String name;
  private final int amount;

}
