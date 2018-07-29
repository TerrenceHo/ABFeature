# ABFeature
### Feature management made easy.

[![Build Status](https://travis-ci.org/TerrenceHo/ABFeature.svg?branch=master "Travis CI
status")](https://travis-ci.org/TerrenceHo/ABFeature)

## Dependencies
1. Golang 1.9 or higher
2. Postgres
    - SQLite is possible, but not recommended
3. Go Deb
4. GPG
5. Mockery
    - To regenerate mocking code
6. Docker (optional)

## How to run

1. Run dep ensure
```bash
dep ensure
```

2. Unencrypt the config file
```bash
make config
```

3. Run the server
```bash
make run
```
## Running Tests
1. Make mocks (Optional)
```bash
make gen-mocks
```

2. Run tests
```bash
make tests
```
