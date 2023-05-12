dockerrun: 
	docker run -e "ACCEPT_EULA=Y" -e "MSSQL_SA_PASSWORD=kursPswd" -p 1433:1433 -d mcr.microsoft.com/mssql/server:2022-latest

dockerexec:
	docker exec -it magical_dhawan /opt/mssql-tools/bin/sqlcmd -S localhost -U sa -P kursPswd


docker exec -it sqlserver /opt/mssql-tools/bin/sqlcmd -S localhost -U SA -P kursPswd


docker run -e 'ACCEPT_EULA=Y' -e 'SA_PASSWORD=kursPswd' -p 1433:1433 --name sqlserver -d mcr.microsoft.com/mssql/server:2022-latest
