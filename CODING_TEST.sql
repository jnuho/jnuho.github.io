CREATE DATABASE `CODING_TEST`;
USE `CODING_TEST`;

-- DROP TABLE ANIMAL_INS;
-- CREATE TABLE `ANIMAL_INS` (
--     age_upon_intake	VARCHAR(100),
--     animal_id	VARCHAR(10),
--     animal_type VARCHAR(20),
-- 	breed	VARCHAR(100),
--     color	VARCHAR(100),
--     datetime	DATETIME,
--     datetime2	DATETIME,
--     found_location	VARCHAR(200),
--     intake_condition	VARCHAR(200),
--     intake_type	VARCHAR(100),
--     name	VARCHAR(200),
--     sex_upon_intake VARCHAR(30)
-- );

-- DROP TABLE ANIMAL_OUTS;
-- CREATE TABLE `ANIMAL_OUTS` (
--     age_upon_outcome	VARCHAR(100),
--     animal_id	VARCHAR(10),
--     animal_type VARCHAR(20),
-- 	breed	VARCHAR(100),
--     color	VARCHAR(100),
--     date_of_birth	DATETIME,
--     datetime	DATETIME,
-- 	monthyear	DATETIME,
--     name	VARCHAR(200),
--     outcome_subtype	VARCHAR(20),
--     outcome_type	VARCHAR(20),
--     sex_upon_intake VARCHAR(30)
-- );
select * from ANIMAL_INS;
-- SHOW VARIABLES LIKE "secure_file_priv";
-- SHOW GLOBAL VARIABLES LIKE 'local_infile';
-- SET GLOBAL local_infile=1;

WITH RECURSIVE OT AS
(
    SELECT 0 AS N
    UNION
    SELECT N + 1 FROM OT WHERE N < 23
)
SELECT OT.N HOUR, COUNT(HOUR(OUTS.DATETIME)) COUNT
FROM OT
LEFT JOIN ANIMAL_OUTS AS OUTS ON OT.N = HOUR(OUTS.DATETIME)
GROUP BY OT.N
ORDER BY HOUR;

-- LOAD DATA INFILE 'aac_intakes.csv'
-- INTO TABLE `ANIMAL_INS`
-- FIELDS TERMINATED BY ',' 
-- ENCLOSED BY '"'
-- LINES TERMINATED BY '\n'
-- IGNORE 1 ROWS;

-- LOAD DATA INFILE 'aac_outcomes.csv'
-- INTO TABLE `ANIMAL_OUTS`
-- FIELDS TERMINATED BY ',' 
-- ENCLOSED BY '"'
-- LINES TERMINATED BY '\n'
-- IGNORE 1 ROWS;

-- LOAD DATA INFILE 'aac_intakes_outcomes.csv'
-- INTO TABLE ``
-- FIELDS TERMINATED BY ',' 
-- ENCLOSED BY '"'
-- LINES TERMINATED BY '\n'
-- IGNORE 1 ROWS;

