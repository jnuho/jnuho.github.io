
- Backend Development<b>2019.11 ~ 2022.03</b>
	- Git
		- Code Review : PR, Branch merge, <i>Code conflict resolution</i>
		- Branch : 'master< develop< personal branch'
	- Java - Spring
		- API development for voting platform of SBS inkigayo TV program
			- APIs included : "Vote, Count votes (Batch Schedule) and Update user info, comments"
			- SMS authentication (AWS Java SDK)
			- 'Sign in with Apple' SDK for oauth login
		- Code optimization for handling user requests
		- Code Refactoring for code readability and reusability


```java
/**
	* The Following is a snippet of what I've been writing
	*		as a Java backend developer
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
		//	retrieves data for update
		//	(data is locked during Lock Wait Time)
		VoteDetailUserVO detailUser = selectVoteDetailUserForUpdate(vote);

		// Validate remaining user points
		validateVoteInfo(vote, vote.getVote_use_count());

		// Insert vote tran_no, hist_seq
		insertVoteHistory(vote);

		// We process point data based on use request that includes point use type parameter
		String voteUsePointType = param.getVote_use_point_type();
		if (VoteUsePointType.SILVER.equals(voteUsePointType)) {
			useSilverPoint(vote);
			useGoldPoint(vote);
		}
		else if (VoteUsePointType.GOLD.equals(voteUsePointType)) {
			useGoldPoint(vote);
		}

		// Insert another vote history.
		//	this history data is used for statistical purpose
		//	e.g. vote candidate recommendation or viewing target candidate history
		insertVoteHist(vote);

		// Update Vote aggregate data
		detailUser.setVote_use_point(vote.getVote_use_point());
		detailUser.setVote_count(vote.getVote_use_count());
		updateVoteDetailUser(detailUser);

		// Return remaining vote count for each user after the transaction is complete
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
