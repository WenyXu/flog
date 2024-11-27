package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

type TimestampGenerator struct {
	created  time.Time
	interval time.Duration
	noise    time.Duration
}

func (tg *TimestampGenerator) Next() time.Time {
	time := tg.created.Add(tg.interval + time.Duration(gofakeit.Number(0, int(tg.noise))))
	tg.created = time
	return time
}

// Generate generates the logs with given options
func Generate(option *Option) error {
	var (
		splitCount = 1
		timestamp  = &TimestampGenerator{
			created:  gofakeit.Date(),
			interval: time.Second * time.Duration(option.Interval),
			noise:    time.Second * time.Duration(option.TimestampNoise),
		}

		interval time.Duration
		delay    time.Duration
	)

	if option.Delay > 0 {
		interval = option.Delay
		delay = interval
	}

	logFileName := option.Output
	writer, err := NewWriter(option.Type, logFileName)
	if err != nil {
		return err
	}

	log := NewLogHeader(option.Format)
	if log != "" {
		_, _ = writer.Write([]byte(log + "\n"))
	}

	if option.Forever {
		for {
			time.Sleep(delay)
			log := NewLog(option.Format, timestamp.Next())
			_, _ = writer.Write([]byte(log + "\n"))
		}
	}

	if option.Bytes == 0 {
		// Generates the logs until the certain number of lines is reached
		for line := 0; line < option.Number; line++ {
			time.Sleep(delay)
			log := NewLog(option.Format, timestamp.Next())
			_, _ = writer.Write([]byte(log + "\n"))

			if (option.Type != "stdout") && (option.SplitBy > 0) && (line > option.SplitBy*splitCount) {
				_ = writer.Close()
				fmt.Println(logFileName, "is created.")

				logFileName = NewSplitFileName(option.Output, splitCount)
				writer, _ = NewWriter(option.Type, logFileName)

				splitCount++
			}
		}
	} else {
		// Generates the logs until the certain size in bytes is reached
		bytes := 0
		for bytes < option.Bytes {
			time.Sleep(delay)
			log := NewLog(option.Format, timestamp.Next())
			_, _ = writer.Write([]byte(log + "\n"))

			bytes += len(log)
			if (option.Type != "stdout") && (option.SplitBy > 0) && (bytes > option.SplitBy*splitCount+1) {
				_ = writer.Close()
				fmt.Println(logFileName, "is created.")

				logFileName = NewSplitFileName(option.Output, splitCount)
				writer, _ = NewWriter(option.Type, logFileName)

				splitCount++
			}
		}
	}

	if option.Type != "stdout" {
		_ = writer.Close()
		fmt.Println(logFileName, "is created.")
	}
	return nil
}

// NewWriter returns a closeable writer corresponding to given log type
func NewWriter(logType string, logFileName string) (io.WriteCloser, error) {
	switch logType {
	case "stdout":
		return os.Stdout, nil
	case "log":
		logFile, err := os.Create(logFileName)
		if err != nil {
			return nil, err
		}
		return logFile, nil
	case "gz":
		logFile, err := os.Create(logFileName)
		if err != nil {
			return nil, err
		}
		return gzip.NewWriter(logFile), nil
	default:
		return nil, nil
	}
}

// NewLogHeader creates a log header for given format
func NewLogHeader(format string) string {
	switch format {
	case "csv":
		return CSVLogHeader
	default:
		return ""
	}
}

// NewLog creates a log for given format
func NewLog(format string, t time.Time) string {
	switch format {
	case "apache_common":
		return NewApacheCommonLog(t)
	case "apache_combined":
		return NewApacheCombinedLog(t)
	case "apache_error":
		return NewApacheErrorLog(t)
	case "rfc3164":
		return NewRFC3164Log(t)
	case "rfc5424":
		return NewRFC5424Log(t)
	case "common_log":
		return NewCommonLogFormat(t)
	case "json":
		return NewJSONLogFormat(t)
	case "csv":
		return NewCSVLogFormat(t)
	default:
		return ""
	}
}

// NewSplitFileName creates a new file path with split count
func NewSplitFileName(path string, count int) string {
	logFileNameExt := filepath.Ext(path)
	pathWithoutExt := strings.TrimSuffix(path, logFileNameExt)
	return pathWithoutExt + strconv.Itoa(count) + logFileNameExt
}
