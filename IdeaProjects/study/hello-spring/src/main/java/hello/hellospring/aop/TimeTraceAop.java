package hello.hellospring.aop;

import org.aspectj.lang.ProceedingJoinPoint;
import org.aspectj.lang.annotation.Around;
import org.aspectj.lang.annotation.Aspect;
import org.springframework.stereotype.Component;

@Aspect
//@Component 또는 SpringConfig.java 에 등록
@Component
public class TimeTraceAop {

  // 타겟팅지정
//  @Around("execution(* hello.hellospring..*(..))")
  @Around("execution(* hello.hellospring.service..*(..))")
  public Object execute(ProceedingJoinPoint joinPoint) throws Throwable {
    long start = System.currentTimeMillis();
    System.out.println("START: " + joinPoint.toString());
    try {
      return joinPoint.proceed();
    } finally {
      long finish = System.currentTimeMillis();
      long timeMs = finish-start;
      System.out.println("END: " + joinPoint.toString() + " " + timeMs + " ms");
    }
  }
}
