
- Git
	- PR, Code Review, Merge
	- Branch : 'main < test < develop'
- Java
	- RESTful API development for SBS inkigayo voting platform
	- SDKs used : AWS SNS, AWS S3
	- APIs used : Payletter, Apple login, 
	- Scheduling Tasks (Spring Framework) for Batch Job
	- Refactoring for readability and reusability of codes

### for-loop optimization

```java
// ...
		// fetch size (query limit count) for each sql query execution
		final int fetchSize = AppConstants.DEFAULT_FETCH_SIZE;

		// total list count
		final int listCount = batchApiMapper.getVoteListCnt();

		// limit for iteration
		final int limit = listCount + fetchSize;
		BatchPlayvoteVO param = new BatchPlayvoteVO();
		for (int ii = 0; ii < limit;) {
			param.setPage_index(ii);
			param.setRecord_count_per_page(fetchSize);

			List<BatchVoteVO> mainList = batchApiMapper.getPlayVoteInfoList(param);
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
 * a snippet of what I've been working as a backend developer
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

