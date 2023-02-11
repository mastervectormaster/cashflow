<p align="center" style="margin-bottom: 0px !important;">
  <img width="200" src="logo.png" alt="Cashflow Logo" align="center">
</p>
<h1 align="center" style="margin-top: 0px;">Cash Flow</h1>

This project is intended as a internal tool for ease of management of Cash

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

```bash
  git clone https://github.com/mastervectormaster/cashflow
```

Go to the project directory

```bash
  cd cashflow
```

Run the docker container env

```bash
npm run dev
```

Stop the docker container env

```bash
npm run dev:stop
```
