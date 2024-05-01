
- 인덱스
  - Cardinality가 높은 컬럼을 인덱스 지정
    - low : 성별, 학년, 상태값
    - high : 주민번호, 계좌번호, seq(auto_increment),
  - 멀티컬럼 인덱스는 Cadidnality가 높은순에서 낮은순으로 지정
    - `CREATE INDEX IDX_SALARIES_DECREASE ON SALARIES(GROUP_NO, FROM_DATE, IS_BONUS);`
    - 여러컬럼 조회시 WHERE절에 첫번째 인덱스 컬럼을 추가해야 인덱스 사용함
  - `BETWEEN`, `LIKE`, `>`, `<`는 인덱스 사용하지 않음
  - `=`, `IN`(`=`를 여러번 사용과 동일)는 인덱스 사용
  - `IN` 의 인자로 서브쿼리 사용시 성능저하
  - `OR` : FULL 테이블스캔
  - WHERE절에 인덱스 컬럼에 연산사용하면 인덱스 안탐 `WHERE SALARY * 10 = 100`

```sql
CREATE TABLE `user`(
  `ID` INT(11) NOT NULL AUTO_INCREMENT COMMENT '아이디',
  `NAME` VARCHAR(50) NOT NULL COMMENT '이름',
  `HOBBY` VARCHAR(100) COMMENT '취미',
  `STAT` CHAR(1) DEFAULT '1' COMMENT `1.활성 2.비활성`,
  PRIMARY KEY(ID)
);

ALTER TABLE user ADD INDEX `idx_user_01` (`ID`, `NAME`, `STAT`);
select
  A.ID
  , A.NAME
  , A.STAT
FROM USER A
WHERE 1=1
  A.ID = 'ABC'
  AND STAT = '1'
```

