# Go Project

## This project using liquibase for database migrations, the migrations scripts are in the folder: `liquibase/my_app/sql` and the file `liquibase/my_app/my-app-changelog.sql` is the main file for liquibase.

## To run the the script for liquibase, you need to have the liquibase installed in your machine, you can download the liquibase in the link: https://www.liquibase.org/download

## After install the liquibase, you need to run the command: `liquibase update-sql --changelogfile=my-app-changelog.sql` in the folder `liquibase/my_app/sql` to run the migrations scripts.

## And then you can run the command: `liquibase update --changelogfile=my-app-changelog.sql` to update the database.



```
curl --location 'http://localhost:10001/country/Adam' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJodHRwOi8vc2NoZW1hcy54bWxzb2FwLm9yZy93cy8yMDA1LzA1L2lkZW50aXR5L2NsYWltcy9lbWFpbGFkZHJlc3MiOiJ0ZXN0MThAZ21haWwuY29tIiwiaHR0cDovL3NjaGVtYXMubWljcm9zb2Z0LmNvbS93cy8yMDA4LzA2L2lkZW50aXR5L2NsYWltcy9yb2xlIjoiVmlld2VyIiwiZXhwIjoxNjk0NDcyNjczLCJpc3MiOiJodHRwOi8vbG9jYWxob3N0OjUwMDEvLGh0dHA6Ly9sb2NhbGhvc3Q6NTAwMy8saHR0cDovL2xvY2FsaG9zdDo1MDA1LyIsImF1ZCI6Imh0dHA6Ly9sb2NhbGhvc3Q6NDIwMS8ifQ.jysG9r0KI_D2PLQJBOlz1TSEhWtMXBw-LYqoNMwgBD0' \
--data ''
```


```
curl --location 'http://localhost:10001/current-time?timezone=Europe%2FAmsterdam' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJodHRwOi8vc2NoZW1hcy54bWxzb2FwLm9yZy93cy8yMDA1LzA1L2lkZW50aXR5L2NsYWltcy9lbWFpbGFkZHJlc3MiOiJ0ZXN0MThAZ21haWwuY29tIiwiaHR0cDovL3NjaGVtYXMubWljcm9zb2Z0LmNvbS93cy8yMDA4LzA2L2lkZW50aXR5L2NsYWltcy9yb2xlIjoiVmlld2VyIiwiZXhwIjoxNjk0NDcyNjczLCJpc3MiOiJodHRwOi8vbG9jYWxob3N0OjUwMDEvLGh0dHA6Ly9sb2NhbGhvc3Q6NTAwMy8saHR0cDovL2xvY2FsaG9zdDo1MDA1LyIsImF1ZCI6Imh0dHA6Ly9sb2NhbGhvc3Q6NDIwMS8ifQ.jysG9r0KI_D2PLQJBOlz1TSEhWtMXBw-LYqoNMwgBD0' \
--data ''
```