package hello.hellospring.controller;

import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseBody;

@Controller
public class HelloController {
  @GetMapping("hello")
  public String hello(Model model) {
    model.addAttribute("data", "Hello spring!!");
    return "hello";
  }

  @GetMapping("hello-mvc")
  public String helloMvc(@RequestParam(value = "name", required=true) String name, Model model) {
    model.addAttribute("name", name);
    return "hello-template";
  }

  @GetMapping("hello-string")
  @ResponseBody // 템플릿이 viewResolver 아닌 응답결과를 직접 결정: string 또는 객체는 디폴트 json
  public String helloString(@RequestParam(value = "name") String name) {
    return "HELLO " + name;
//    return "<html>HELLO " + name + "</html>";
  }

  @GetMapping("hello-api")
  @ResponseBody // 템플릿이 viewResolver 아닌 응답결과를 직접 결정: string 또는 객체를 디폴트 json
  public Hello helloApi(@RequestParam(value = "name") String name) {
    Hello hello = new Hello();
    hello.setName(name);
    return hello;
  }

  static class Hello {
    private String name;

    public String getName() {
      return name;
    }

    public void setName(String name) {
      this.name = name;
    }
  }
}
