package springbook.user.test;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.ApplicationContext;
import org.springframework.test.context.ContextConfiguration;

import java.util.HashSet;
import java.util.Set;

import static org.hamcrest.CoreMatchers.*;
import static org.hamcrest.MatcherAssert.assertThat;
import static org.junit.jupiter.api.Assertions.assertTrue;

/**
 * JUnit은 테스트 메소드를 수행할때 마다 새로운 오브젝트를 만든다
 * 정말 새로운 오브젝트가 만들어지는지 확인
 * JUnit 자신에 대한 테스트
 */
@ContextConfiguration
public class JUnitTest {

  // 테스트 컨텍스트가 매번 주입해주는 애플레케이션 컨텍스트는 항상 같은
  // 오브젝트인지 테스트로 확인
  @Autowired
  ApplicationContext context;

//  static JUnitTest testObject;
  static Set<JUnitTest> testObjects = new HashSet<JUnitTest>();
  static ApplicationContext contextObject = null;

  @Test
  public void test1() {
    assertThat(testObjects, not(hasItem(this)));
    testObjects.add(this);

    assertThat(contextObject == null || contextObject == this.context,
        is(true));
    contextObject = this.context;
  }

  @Test
  public void test2() {
    assertThat(testObjects, not(hasItem(this)));
    testObjects.add(this);

    assertTrue(contextObject == null || contextObject == this.context);
    contextObject = this.context;
  }

  @Test
  public void test3() {
    assertThat(testObjects, not(hasItem(this)));
    testObjects.add(this);

    assertThat(contextObject,
        either(is(nullValue())).or(is(this.context)));
    contextObject = this.context;
  }
}
