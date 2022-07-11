# gload
A very simple load test tool

## Usage

```shell
gload -s /path/to/scenarios.yml
```

## Scenario file example

```yaml
CreateUser:
  url: https://api.gloadsys.com/users
  mehotd: POST
  timeout: 5000
  iterations: 50
  processes: 3
  headers:
    Content-Type: application/json
  body: |
    {
      "username": "gloaduser",
      "email": "gload@mail.com"
    }

GetAllGroups:
  url: https://api.gloadsys.com/groups
  mehotd: GET
  timeout: 5000
  iterations: 200
  processes: 10
  headers:
    Content-Type: application/json
```