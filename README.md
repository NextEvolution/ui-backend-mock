# Backend Mock

## Instructions
1. Clone this repo in a sidecar location to the UI source (see directory structure below)
1. `cd <location of repo>`
1. `./mock <path to config file>` ie `./mock config.json`

The mock will open up on `localhost:8080`

## Directory Structure
The mock app requires the `./responses/` directory to run properly.

Here is the directory structure
```
.
ui-backend-mock
├── mock
└── responses
    ├── GET_alive.txt
    ├── GET_api_sales.json
    ├── POST_api_login.json
    └── ...
```
## Config File
```json
{
  "port": 8080,
  "static_file_path": "../ui-frontend/dist"
}
```

## Mock endpoints
* `/api/login` - facebook login
* `/api/sales` - List sales
* `/alive` - used to test if the mock is running
* `/` - root directory hosting UI source