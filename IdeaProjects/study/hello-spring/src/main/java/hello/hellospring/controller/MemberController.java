package hello.hellospring.controller;

import hello.hellospring.domain.Member;
import hello.hellospring.service.MemberService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;

import java.util.List;

@Controller
public class MemberController {

  // 1.필드주입
//  @Autowired
//  private MemberService memberService;

  private final MemberService memberService;

  // 2. 생성자주입
  @Autowired
  public MemberController(MemberService memberService) {
    this.memberService = memberService;
    // 진짜 Service 실행전 프록시 Service 생성되어 AOP 적용 및 실행 후 joinPointProceed() 하게 되므로
    // 여기서는 CGLIB가 Service를 프록시가 생성되어 주입되고 있음을 확인
    // DI를 사용하기 때문에 프록시 주입이 가능!
    System.out.println("[AOP확인]memberService = " + memberService.getClass());
  }

  @GetMapping("/members/new")
  public String createForm() {
    return "members/createMemberForm";
  }

  @PostMapping("/members/new")
  public String create(MemberForm form) {
    Member member = new Member();
    member.setName(form.getName());

    memberService.join(member);
    return "redirect:/";
  }

  @GetMapping("/members")
  public String list(Model model) {
    List<Member> members = memberService.findMembers();
    model.addAttribute("members", members);
    return "members/memberList";
  }
}
