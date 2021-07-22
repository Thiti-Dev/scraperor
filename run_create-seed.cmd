@echo off
for /f "delims== tokens=1*" %%A in (.env) do set POSTGRESQL_URL=%%B
psql -d %POSTGRESQL_URL% -f .\postgres\seed.sql