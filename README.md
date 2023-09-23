# Go Project

## This project using liquibase for database migrations, the migrations scripts are in the folder: `liquibase/my_app/sql` and the file `liquibase/my_app/my-app-changelog.sql` is the main file for liquibase.

## To run the the script for liquibase, you need to have the liquibase installed in your machine, you can download the liquibase in the link: https://www.liquibase.org/download

## After install the liquibase, you need to run the command: `liquibase update-sql --changelogfile=my-app-changelog.sql` in the folder `liquibase/my_app/sql` to run the migrations scripts.

## And then you can run the command: `liquibase update --changelogfile=my-app-changelog.sql` to update the database.