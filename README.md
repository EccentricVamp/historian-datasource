# Proficy Historian Grafana Datasource

Tested to work with GE Historian 7.1 and up. Uses the [REST API](https://www.ge.com/digital/documentation/historian/version71/IOTcwM2Y5YzctNGZhMy00M2IzLWFlZmUtNjcxODkwMzNlM2Zh.html)

## Prerequisites

1. [Visual Studio Code](https://code.visualstudio.com/)

   [Go language extension](https://marketplace.visualstudio.com/items?itemName=golang.Go)

2. [Node.js](https://nodejs.dev/)

3. [Go](https://go.dev/)

   [Mage](https://magefile.org/)

## Install dependencies

Backend: `go mod download`

Frontend: `npm install`

## Build

Backend: `mage`

Frontend: `npm run build`
