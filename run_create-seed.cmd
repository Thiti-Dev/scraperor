@echo off
for /f "delims== tokens=1*" %%A in (.env) do (
    if %%A equ POSTGRESQL_URL (
        set POSTGRESQL_URL=%%B
        goto process
    )
)
:process
if [%POSTGRESQL_URL%] == [] (
    echo No .env file found with the POSTGRESQL_URL variable
    goto:eof
)

psql -d %POSTGRESQL_URL% -f .\postgres\seed.sql