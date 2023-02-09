 
![Cashflow Logo](./1.png)  

# Cash Flow
**Manage & Monitor Internal Money Flow**

This project is intended as a internal tool that can be easily integrated into a microservice-architectured system

## Features
- Income Management
    - Sumit Income
    - Approve Income
- Expense Management
    - Submit Expense
    - Approve Expense (TBD)
    - Schedule Expense
- Analysis/Export
    - Overall (monthly,yearly)
    - Groups (monthly, yearly)
    - Individuals (monthly, yearly)

## Architecture

This itself is implemented as microservice architecture.

### Microservices included
- Auth
- User
- REST API Gateway
- Cash
- Schedule

## Tech Stack
- GoLang/Gin
- OpenAPI/AsyncAPI
- AWS (TBD)
- CICD (Jenkins)
- MongoDB
- JWT/OTP
- Docker/K8S

## REST API specification

The live OpenAPI specification of the Cashflow REST API can be found [here](https://open-api-test.ludoempire.com/) (TODO)

## Socket.io specification

The live AsyncAPI specification of the Ludo Socket.io API can be found [here](https://async-api-test.ludoempire.com/) (TODO)

 
## Run Locally  
Clone the project  

~~~bash  
  git clone https://github.com/mastervectormaster/cashflow
~~~

Go to the project directory  

~~~bash  
  cd cashflow
~~~

Run the docker container env  

~~~bash  
npm run dev
~~~  
