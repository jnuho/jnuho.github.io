
- Git
	- Code Review : PR, Merge
	- Branch : 'main < develop < personal branch'
- Java - Spring
	- API development for SBS inkigayo voting platform
		- APIs : "Vote, Count votes, Update user info, Add user comments"
		- SMS authentication (AWS Java SDK)
		- 'Sign in with Apple' SDK for oauth login
	- Code optimization for handling thousands of user requests
	- Refactoring for code readability and reusability


```java
/**
 * a snippet of what I've been writing as a backend developer
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
