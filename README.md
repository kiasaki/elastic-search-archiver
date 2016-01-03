# Elastic Search Archiver

This little command line tool is built to help you archive old logs in Elastic
Search, most likely logs Logstash put in there.

It will go and **flush indices transactions** then proceed to **close** those indices
so that they don't cost you CPU or memory, all this up to a certain duration,
the default is 2 weeks in the past. On top of that it has the ability to completely
delete (forever) indices up to a certain duration, default is not deletion.

## Usage

```
Usage of elastic-search-archiver:
  -archive=false: Pass flag to execute archival
  -archive-age=360h0m0s: Archive only indices old X time (Default is 15 days written 360h
  -delete=false: Pass flag to execute deletion
  -delete-age=720h0m0s: Delete only indices old X time (Default is 30 days written 720h
  -es-host="": Elastic Search host to delete indices from (required)
  -prefix="logstash-": Prefix behind ElasticSearch indices to delete or archive
```

## Developing

```bash
go get .
go run cmd/elastic-search-archiver/main.go
make build
```

## License

See `LICENSE` file.
