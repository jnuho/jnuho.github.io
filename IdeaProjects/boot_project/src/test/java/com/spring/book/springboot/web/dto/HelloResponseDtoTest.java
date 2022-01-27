package com.spring.book.springboot.web.dto;

import junit.framework.TestCase;
import org.junit.Test;

import static org.hamcrest.MatcherAssert.assertThat;
import static org.hamcrest.Matchers.is;


public class HelloResponseDtoTest {

  @Test
  public void 롬복_기능_테스트() {
    // given
    String name ="test";
    int amount = 1000;

    // when
    HelloResponseDto dto = new HelloResponseDto(name, amount);

    // then
    assertThat(dto.getName(), is(name));
    assertThat(dto.getAmount(), is(amount));
  }
}