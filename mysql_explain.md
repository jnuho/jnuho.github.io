




























EXPLAIN 결과 컬럼

- id: 행이 속한 SELECT 구분자
	UNION이나 서브쿼리가 없으면 하나의 SELECT만 존재하여 모든 행의 id=1
	Derived Table or Subquery will have different 'id' values

- select_type: SELECT 쿼리타입 구분 Simple or Complex
	-Simple 서브쿼리나 UNION 없음
	-Complex
		1. Simple subqueries
		2. Derived tables
		3. UNIONs(id = null)

EXPLAIN SELECT (SELECT 1 FROM sakila.actor LIMIT 1) FROM sakila.film;

id | select_type | table
---|---|---
1 | PRIMARY | film
2 | SUBQUERY | actor

EXPLAIN SELECT film_id FROM (SELECT film_id FROM sakila.film) AS der;

id | select_type | table
---|---|---
1 | PRIMARY | ```<derived2>```
2 | DERIVED | film

EXPLAIN SELECT 1 UNION ALL SELECT 1;

id | select_type | table
---|---|---
1 | PRIMARY | NULL
2 | UNION | NULL
NULL | UNION RESULT | ```<union1,2>```


