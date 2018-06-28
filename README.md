# ABFeature
### Feature management made easy.

## Dependencies
1. Golang 1.9 or higher
2. Postgres
    - SQLite is possible, but not recommended
3. Go Deb
4. GPG

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
