package com.spring.book.springboot.config.auth;

import java.lang.annotation.ElementType;
import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;

@Target(ElementType.PARAMETER) // 이 어노테이션이 생성될 수 있는 위치: 파라미터로 선언된 객체에서만 사용 (e.g 클래스어노테이션도 가능)
@Retention(RetentionPolicy.RUNTIME)
public @interface LoginUser { // @interface: 클래스를 어노테이션으로 정의
}
