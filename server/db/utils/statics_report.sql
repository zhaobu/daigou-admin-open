SELECT
    dt,
    count(DISTINCT userid) AS num
FROM
    denglu
WHERE
    dt BETWEEN '2018-09-09'
    AND '2018-09-13'
GROUP BY
    dt;

-- 查询每日新增
SELECT
gender
    DATE(created_at) as regist_time,
    count(DISTINCT user_id) AS num
FROM
    user
WHERE
    created_at BETWEEN '2020-01-09'
    AND '2020-08-27'
GROUP BY
    DATE(created_at);