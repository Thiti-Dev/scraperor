@echo off
for /f "delims== tokens=1*" %%A in (.env) do (
    if %%A equ ELEPHANTSQL_URL (
        set ELEPHANTSQL_URL=%%B
        goto process
    )
)
:process
if [%ELEPHANTSQL_URL%] == [] (
    echo No .env file found with the ELEPHANTSQL_URL variable
    goto:eof
)


migrate -path "postgres/migrations" -database %ELEPHANTSQL_URL% %*