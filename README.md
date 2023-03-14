# GService

Small prototype of a micro service

## Deploy

```bash
make
docker-compose up -d --build
```

## Testing

```bash
make test
```

## E2E test

### Adding new post data `POST /data`

```bash
# Passing basic auth on routes for data
curl -X POST -u "admin:admin" --header "Content-Type: application/json" --data '{"name": "Test1", "stage": 1, "score": 100}' docker_host_address:3000/post-data
```

```json
{
  "msg": "post inserted with success"
}
```

### List all post data `GET /data`

```bash
# Passing basic auth on routes for data
curl -v http://docker_host_address:3000/data
```

```json
[
  {
    "UUID4": "40ff6d6b-f55d-4203-aaa5-86fa9a36e393",
    "Name": "test2",
    "Stage": 1,
    "Score": 100,
    "Timestamp": "2020-09-06T15:46:35+00:00"
  }
]
```

### Get post data by name `GET /get-data/{name}`

```bash
curl -v http://docker_host_address:3000/get-data/test2
```

```json
{
 "UUID4": "40ff6d6b-f55d-4203-aaa5-86fa9a36e393",
  "Name": "test2",
  "Stage": 1,
  "Score": 100,
  "Timestamp": "2020-09-06T15:46:35+00:00"
}
```