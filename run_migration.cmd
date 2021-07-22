for /f "delims== tokens=1*" %%A in (.env) do set POSTGRESQL_URL=%%B

migrate -path "postgres/migrations" -database %POSTGRESQL_URL% %*