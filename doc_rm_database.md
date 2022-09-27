```
I had a chance to work on database design and query optimization.
Followings are what I did as a backend engineer.
```

- Define Table Entity
- Define Indexes for table columns if necessary
	- preferable if there are large number of data rows
	- preferable if the column is most of the time NOT NULL
	- 3-4 indices are preferable
	- updating or deleting performance decreases if the table has many indices defined
	- Decide which column should be chosen as an index column
		- High Cardinality (High duplication. e.g. gender, grade)
		- High Cardinality means one can opt out large amount of data by using index
	- Multi-column index: 2nd column replies on 1st, 3rd replies on 2nd, and so on.
		- Order: Cardinality decreasing (High to Low)
	- WHERE clause must include 1st index column in order to use index
		- e.g. Given 1,2,3, must use (1 and 3) Or (1 and 2)
		- the order in which columns are used in the where clause is trivial

- Define primary, foreign, unique keys for columns
- Optimize query with index columns and row limit
- Make use of EXPLAIN statement to get detailed info about how statements are executed

```sql
CREATE TABLE `VOTE_HISTORY` (
	`SEQ` BIGINT(21) NOT NULL AUTO_INCREMENT COMMENT 'seq',
	`LOGIN_ID` VARCHAR(50) NOT NULL COMMENT 'login id',
	`STAR_TYPE` CHAR(1) NOT NULL COMMENT 'type 1.single 2.group',
	`STAR_CD` INT(11) DEFAULT 0 NOT NULL COMMENT 'target candidate's star code',
	`GRP_CD` VARCHAR(10) NOT NULL COMMENT 'target candidate's group code',
	`VOTE_DT` DATETIME NOT NULL COMMENT 'date of the voting',
	PRIMARY KEY (`SEQ`),
	-- a user must have exactly 'one' vote history of a candidate
	UNIQUE KEY `udx_vote_log_01` (`login_id`, `star_type`, `star_cd`, `grp_cd`)
) ENGINE = InnoDB DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci COMMENT 'vote history';

-- Cardinality : STAR_TYPE > LOGIN_ID
ALTER TABLE VOTE_HISTORY ADD INDEX `idx_vote_history_01` (`STAR_TYPE`, `LOGIN_ID`);


SELECT
	SEQ
	, LOGIN_ID
	, STAR_CD
	, GRP_CD
WHERE
	1=1
	AND STAR_TYPE = '1'
	AND LOGIN_ID = 'OOO';
```
