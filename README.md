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

### Adding new post data `POST /post-data`

```bash
curl -X POST --header "Content-Type: application/json" --data '{"title": "Test1"}' docker_host_address:3000/post-data
```

```json
{
  "msg": "post inserted with success"
}
```

### List all post data `GET /post-data`

```bash
curl -v http://docker_host_address:3000/post-data
```

```json
[
  {
    "UUID4": "40ff6d6b-f55d-4203-aaa5-86fa9a36e393",
    "Title": "test2",
    "Timestamp": "2020-09-06T15:46:35+00:00"
  }
]
```

### Get post data by title `GET /post-data/{title}`

```bash
curl -v http://docker_host_address:3000/post-data/test2
```

```json
{
 "UUID4": "40ff6d6b-f55d-4203-aaa5-86fa9a36e393",
  "Title": "test2",
  "Timestamp": "2020-09-06T15:46:35+00:00"
}
```