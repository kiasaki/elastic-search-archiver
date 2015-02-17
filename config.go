package esa

import (
	"errors"
	"flag"
	"time"
)

type Config struct {
	Archive           bool
	ArchiveAge        time.Duration
	Delete            bool
	DeleteAge         time.Duration
	ElasticSearchHost string
	IndicePrefix      string
}

var (
	fArchive           *bool
	fArchiveAge        *time.Duration
	fDelete            *bool
	fDeleteAge         *time.Duration
	fElasticSearchHost *string
	fIndicePrefix      *string
)

func SetupConfigFlags() {
	fArchive = flag.Bool("archive", false, "Pass flag to execute archival")
	fArchiveAge = flag.Duration("archive-age", time.Hour*24*15, "Archive only indices old X time (Default is 15 days written 360h")
	fDelete = flag.Bool("delete", false, "Pass flag to execute deletion")
	fDeleteAge = flag.Duration("delete-age", time.Hour*24*30, "Delete only indices old X time (Default is 30 days written 720h")
	fElasticSearchHost = flag.String("es-host", "", "Elastic Search host to delete indices from (required)")
	fIndicePrefix = flag.String("prefix", "logstash-", "Prefix behind ElasticSearch indices to delete or archive")
}

func NewConfigFromFlags() Config {
	SetupConfigFlags()
	flag.Parse()
	return Config{
		*fArchive, *fArchiveAge, *fDelete, *fDeleteAge, *fElasticSearchHost, *fIndicePrefix,
	}
}

func (c Config) Validate() error {
	if !c.Archive && !c.Delete {
		return errors.New("Nothing to do, specify one or both of the archive and delete flags")
	}
	if c.Archive && c.ArchiveAge == 0 {
		return errors.New("Can't archive for indices old 0s, pass-in archive-age flag")
	}
	if c.Delete && c.DeleteAge == 0 {
		return errors.New("Can't archive for indices old 0s, pass-in archive-age flag")
	}
	if c.ElasticSearch == "" {
		return errors.New("ElasticSearch host flag is required")
	}
	return nil
}
