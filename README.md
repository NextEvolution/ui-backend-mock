# Backend Mock

## Instructions
1. Clone this repo in a sidecar location to the UI source (see directory structure below)
1. `cd <location of repo>`
1. `./mock`

The mock will open up on `localhost:8080`

## Directory Structure
The mock app requires two directories to run properly: `./responses` and `../ui-frontend/src`.
The mock app is meant to run as a sidecar to the UI source, hence it is not in the same directory.
```
.
├── ui-backend-mock
│   ├── mock
│   └── responses
│       ├── GET_alive.txt
│       ├── GET_api_sales.json
│       └── POST_api_login.json
└── ui-frontend
    └── src
        └── index.html
```
* mock is the executable to run the mock. Simply run as `./mock`
* Responses is the directory containing pre-made responses.
* `ui-front-end/src` is where you will place the source for the site. Note: mock is meant to side-car the ui source.

## Mock endpoints
* `/api/login` - facebook login
* `/api/sales` - List sales
* `/alive` - used to test if the mock is running
* `/` - root directory hosting UI source