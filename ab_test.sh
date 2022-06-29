ab -n 10000 -c 100 \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjEsImV4cCI6MTY1NjQ5NDE2NiwiaXNzIjoicGFuZG9yYSJ9.Tdb1sDmG1hlJHuNpLtYSpjRgrHxs9PRoDq_NFIg2CbI" \
http://127.0.0.1:5001/stocks/daily\?page\=1\&pageSize\=10\&searchVal\=\&startDate\=2021-01-01\&endDate\=2022-01-01