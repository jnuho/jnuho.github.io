
**✻ What I learned :**
```
Define Table
Define Indexes for table columns
	- indices for a table is preferable if there are a number of data rows
	- index for a column is preferable if the column is NOT NULL
	- 3-4 indices for a table are ok
	- updating or deleting performance decreases for tables with many indices
	- choose index column with High Cardinality
		- high : uncommon and unique. e.g. autoincrement seq, timestamp
		- high cardinality means one can opt out large amount of data by using index
		- low : with few unique values, typically status flags, boolean values.
	- Multi-column index: second column relies on first, third relies on second, and so on.
		- Order: Cardinality decreasing (High to Low)
	- WHERE clause must include first index column in order to use index
		- e.g. Given 1,2,3, must use (1 and 3) or (1 and 2)
		- the order in which columns are used in the where clause is trivial

	- `BETWEEN`, `LIKE`, `>`, `<` will skip index use
	- `=`, `IN`(same is using `=` multiple times)는 uses index
	- `IN` with Subquery will decrease performance
	- `OR` : FULL table scan
	- using arithmetic operation on index column will skip index `WHERE SALARY * 10 = 100`

Define primary, foreign, unique keys for columns

Optimize query with index columns and row limit

Make use of EXPLAIN statement to get detailed info about how statements are executed
```

<br>

**✻ Example**
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

-- Cardinality : LOGIN_ID > STAR_TYPE
ALTER TABLE VOTE_HISTORY ADD INDEX `idx_vote_history_01` (`LOGIN_ID`, `STAR_TYPE`);

SELECT
	SEQ
	, LOGIN_ID
	, STAR_CD
	, GRP_CD
WHERE
	1=1
	AND LOGIN_ID = 'OOO'
	AND STAR_TYPE = '1';
```
