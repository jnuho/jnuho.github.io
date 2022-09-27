
- Git
	- PR, Code Review, Merge
	- Branch : 'main < test < develop'
- Java
	- RESTful API development for SBS inkigayo voting platform
	- SDKs used : AWS SNS, AWS S3
	- APIs used : Payletter, Apple login
	- Scheduling (Spring Framework) for Batch Job
	- Refactoring for readability and reusability of codes
- Test
	- Junit unit test

### For-loop optimization

```java
// ...
// Limit fetch size (query limit count) for each sql query execution
//	so that each request does not put 
final int fetchSize = AppConstants.DEFAULT_FETCH_SIZE;

// Total list count
final int listCount = batchApiMapper.getVoteListCnt();

// Limit for iteration
final int limit = listCount + fetchSize;
BatchVO param = new BatchVO();
for (int ii = 0; ii < limit;) {
	param.setPage_index(ii);
	param.setRecord_count_per_page(fetchSize);

	List<BatchVoteVO> mainList = batchApiMapper.getVoteInfoList(param);
	for (BatchPlayvoteVO m : mainList) {
		batchApiMapper.updateVote(m);
	}
	ii += fetchSize;
}
// ...
```


### Refactoring

```java
/**
 * API for executing voting process includes updating and retrieving from database
 *	refactoring helped to improve readability
 */

@Override
@Transactional
public ResultVO insertVote(VoteInsertVO param) throws Exception {

	// Get vote info including target candidate that a user is requesting
	VoteInsertVO vote = getVoteUserInfo(param);

	// Calculate total vote points
	//	internal point system include point conversion!
	vote = setTotalVotePointInfo(vote);

	// Operation for SELECT FOR UPDATE
	//	retrieves data for update (update is locked during Lock Wait Time)
	VoteDetailUserVO detailUser = selectVoteDetailUserForUpdate(vote);

	// Validate remaining user points
	validateVoteInfo(vote, vote.getVote_use_count());

	// Insert vote result data : tran_no, hist_seq
	insertVoteHistory(vote);

	// Process point deduction based on the user request which includes use point type
	String voteUsePointType = param.getVote_use_point_type();
	if (VoteUsePointType.SILVER.equals(voteUsePointType)) {
		useSilverPoint(vote);
		useGoldPoint(vote);
	}
	else if (VoteUsePointType.GOLD.equals(voteUsePointType)) {
		useGoldPoint(vote);
	}

	// Insert another vote history data
	//	which is used for statistical purpose
	//	e.g. for future vote candidate recommendation or viewing voting history
	insertVoteHist(vote);

	// Update Vote aggregate data
	detailUser.setVote_use_point(vote.getVote_use_point());
	detailUser.setVote_count(vote.getVote_use_count());
	updateVoteDetailUser(detailUser);

	// Return remaining vote count for a user after the transaction is complete
	//	in case of a vote platform that restrict the vote counts for each user
	if (!VoteLimitCountType.NOLIMIT.equals(vote.getVote_limit_cnt_type())) {
		int remain_cnt = selectRemainCnt(vote);
		ResultDoPlayVoteVO result = new ResultDoPlayVoteVO();
		result.setVote_remain_cnt(remain_cnt);
		result.setCode(ErrorCode.SUCCESS.getCode());
		result.setMessage(getMessage(ErrorCode.SUCCESS.getKey()));
		return result;
	} else {
		ResultVO result = new ResultVO();
		result.setCode(ErrorCode.SUCCESS.getCode());
		result.setMessage(getMessage(ErrorCode.SUCCESS.getKey()));
		return result;
	}
}
```


### Unit Test

```java
package com.oo.ooo.home;

import static org.junit.Assert.assertEquals;
import java.util.List;
import org.junit.Test;
import org.springframework.beans.factory.annotation.Autowired;
import com.rowem.passicon.SpringTestBase;


/**
 * OO 조회 테스트
 *
 * @author junho
 *
 */
public class DataApiTest extends SpringTestBase {

	@Autowired
	private HomeApiService homeApiService;

	@Autowired
	private HomeApiMapper homeApiMapper;

	@Autowired
	private DataService dataService;

	/**
	 * Delete data
	 *
	 * @throws Exception
	 */
	private void deleteData() throws Exception {

		DataVO data = null;
		do {
			MyObj param = new MyObj();
			param.setLogin_id("ooo")

			// Find data to be deleted
			data = dataApiMapper.selectMyData(param);
			if (data != null) {
				ResultVO result = dataService.deleteData(data);
				assertEquals(ErrorCode.SUCCESS.getCode(), result.getCode());
			}
		} while(star_info != null);
	}

	/**
	 * Insert data
	 *
	 * @throws Exception
	 */
	private void insertData() throws Exception {

		MyObj s = new MyObj();
		s.setLogin_id("ooo")
		ResultVO res = dataService.insertMyData(s);
		assertEquals(ErrorCode.SUCCESS.getCode(), res.getCode());
	}

	/**
	 * 0-1 홈화면 테스트
	 *
	 * 케이스1. 마이스타 X 투표이력 X, 백오피스 > 추천스타 리스트 0개(전부 비활성 stat=2)
	 *
	 * @throws Exception
	 */
	@Test
	public void testGetStarInfo_1() throws Exception {
		String login_id = "ooo";
		DataParamVO param = new DataParamVO();
		param.setLogin_id(login_id);

		// 마이스타 X
		deleteData();

		// 투표이력 X
		List<StarInfoVO> recom_votehist_list = homeApiMapper.selectStarVoteLogList(param);
		assertEquals(0, recom_votehist_list.size());

		DataResultVO result = (DataResultVO) homeApiService.getStarInfo(param);
		assertEquals(null, result.getStar_info());
		assertEquals(0, result.getRecom_votehist_list().size());
		assertEquals(null, result.getRefresh_size());
		assertEquals(0, result.getRecom_admin_list().size());
		assertEquals(null, result.getVote_playvote_list());
		assertEquals(null, result.getVote_indiv_rank());
		assertEquals(null, result.getVote_battle_rank());
	}

	/**
	 * 케이스2. 마이스타 X 투표이력 X, 백오피스 > 추천스타 리스트 1개 이상 (활성 stat=1)
	 *
	 * @throws Exception
	 */
	@Test
	public void testGetStarInfo_2() throws Exception {
		String login_id = "kimansub87@gmail.com";
		DataParamVO param = new DataParamVO();
		param.setLogin_id(login_id);

		// 마이스타 X
		deleteData();

		// 투표이력 X
		List<StarInfoVO> recom_votehist_list = homeApiMapper.selectStarVoteLogList(param);
		assertEquals(0, recom_votehist_list.size());

		DataResultVO result = (DataResultVO) homeApiService.getStarInfo(param);
		assertEquals(null, result.getStar_info());
		assertEquals(0, result.getRecom_votehist_list().size());
		assertEquals(null, result.getRefresh_size());
		assertEquals(Boolean.TRUE, result.getRecom_admin_list().size() > 0);
		assertEquals(null, result.getVote_playvote_list());
		assertEquals(null, result.getVote_indiv_rank());
		assertEquals(null, result.getVote_battle_rank());
	}
}


```

```java
package com.oo.ooo;

import javax.servlet.ServletContext;

import org.junit.Assert;
import org.junit.Before;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.mock.web.MockServletContext;
import org.springframework.test.context.ActiveProfiles;
import org.springframework.test.context.ContextConfiguration;
import org.springframework.test.context.junit4.SpringJUnit4ClassRunner;
import org.springframework.test.context.web.WebAppConfiguration;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.setup.MockMvcBuilders;
import org.springframework.web.context.WebApplicationContext;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.core.util.DefaultPrettyPrinter;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.SerializationFeature;

/**
 * 스프링 테스트 베이스
 *
 * @author junho
 *
 */
@RunWith(SpringJUnit4ClassRunner.class)
@ContextConfiguration(locations = { "file:src/main/webapp/WEB-INF/spring/appServlet/servlet-context.xml",
		"file:src/main/webapp/WEB-INF/spring/root-context.xml",
		"file:src/test/resources/spring/context-test-mapper.xml", })
@WebAppConfiguration
@ActiveProfiles("local")
public class SpringTestBase {

	protected Logger logger = LoggerFactory.getLogger(getClass());

	@Autowired
	protected WebApplicationContext wac;

	protected MockMvc mockMvc;

	protected ObjectMapper objectMapper;

	@Before
	public void setUp() {
		mockMvc = MockMvcBuilders.webAppContextSetup(wac).build();
		objectMapper = new ObjectMapper();
		objectMapper.enable(SerializationFeature.INDENT_OUTPUT);
		objectMapper.setDefaultPrettyPrinter(new DefaultPrettyPrinter());
	}

	/**
	 * 객체를 JSON으로 로그 출력
	 *
	 * @param name
	 * @param value
	 */
	protected void printJson(String name, Object value) {
		try {
			logger.info("{}:\n{}", name, objectMapper.writeValueAsString(value));
		} catch (JsonProcessingException e) {
			logger.error("error :", e);
		}
	}

	@Test
	public void isWebAppContextLoadingProperly() {
		ServletContext servletContext = wac.getServletContext();
		Assert.assertNotNull(servletContext);
		Assert.assertTrue(servletContext instanceof MockServletContext);
	}
}
```
