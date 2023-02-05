package jpabook.jpashop;


import org.assertj.core.api.Assertions;
import org.junit.jupiter.api.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.annotation.Rollback;
import org.springframework.test.context.junit4.SpringRunner;
import org.springframework.transaction.annotation.Transactional;

@RunWith(SpringRunner.class)
@SpringBootTest
public class MemberRepositoryTest {

	@Autowired
	MemberRepository memberRepository;

	@Test
	@Transactional
	@Rollback(false)
	public void testMember() throws Exception {
		//given
		Member member = new Member();
		member.setUsername("memberA");

		//when
		Long saveId = memberRepository.save(member);
		Member findMember = memberRepository.find(saveId);

	  //then
		Assertions.assertThat(findMember.getId()).isEqualTo(member.getId());
		Assertions.assertThat(findMember.getUsername()).isEqualTo(member.getUsername());

		// 같은 tx 안에서, 엔티티 같음! (1차캐시에서 가져오므로 select 쿼리 실행 X)
		Assertions.assertThat(findMember).isEqualTo(member);
	}

}