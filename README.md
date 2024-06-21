# GoHttpRedirecter

A simple HTTP request redirector written in go

## CLI Arguments

| Argument             | Description                                                                          |
|----------------------| -------------------------------------------------------------------------------------|
| --port               | Webserver port (default: 7878)                                                       |
| --routes <FILE_NAME> | Name of the routes configuration file inside the routes directory (default: default) |


## Routes configuration file syntax
```
<REQUEST_PATH> => <REDIRECT_URL>
<REQUEST_PATH_2> => <REDIRECT_URL_2>
```
