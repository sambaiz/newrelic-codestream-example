```shell
$ export NEWRELIC_LICENSE_KEY=***
$ export NEW_RELIC_METADATA_REPOSITORY_URL=git@github.com:sambaiz/newrelic-codestream-example.git
$ export NEW_RELIC_METADATA_COMMIT=$(git rev-parse HEAD)
$ go run main.go
```
