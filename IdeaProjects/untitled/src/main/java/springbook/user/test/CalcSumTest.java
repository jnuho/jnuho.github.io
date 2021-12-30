package springbook.user.test;

import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.TestInstance;
import org.junit.jupiter.api.extension.ExtendWith;
import org.springframework.test.context.ContextConfiguration;
import org.springframework.test.context.junit.jupiter.SpringExtension;
import springbook.learningtest.template.Calculator;

import java.io.IOException;

import static org.hamcrest.CoreMatchers.is;
import static org.hamcrest.MatcherAssert.assertThat;

// JUnit 4
//@RunWith(SpringJUnit4ClassRunner.class)
// JUnit 5
// 스프링의 테스트 컨텍스트 프레임워크의 JUnit 확장기능 지점
@ExtendWith(SpringExtension.class)
// 테스트 컨텍스트가 자동으로 만들어줄 애플리케이션 컨텍스트의 위치 지정
// 모든 테스트 클레스에서 같은 애플리케이션 컨텍스트 객체가 사용됨
@ContextConfiguration({"/applicationContext.xml"})
//@BeforeAll 의 non-static 메소드 사용을 위해
//클래스에 @TestInstance(TestInstance.Lifecycle.PER_CLASS) 추가
@TestInstance(TestInstance.Lifecycle.PER_CLASS)
public class CalcSumTest {

  Calculator calculator;
  String numFilePath;

  @BeforeAll
  public void setUp() {
    this.calculator = new Calculator();
    this.numFilePath = getClass().getResource("numbers.txt").getPath();
  }

  @Test
  public void sumOfNumbers() throws IOException {
    assertThat(calculator.calcSum(this.numFilePath), is(10));
  }

  @Test
  public void multipleOfNumbers() throws IOException {
    assertThat(calculator.calcMultiply(this.numFilePath), is(24));
  }

}
