# ClickHouse

## Install ClickHouse
```bash
curl https://clickhouse.com/ | sh
```

## Create Table

```sql
CREATE TABLE IF NOT EXISTS `logs` (
  `host` String,
  `user-identifier` String,
  `datetime` DateTime,
  `method` String,
  `request` String,
  `protocol` String,
  `status` UInt64,
  `bytes` UInt64,
  `referer` String,
  `user-agent` String,
)
ENGINE = MergeTree
ORDER BY (toStartOfHour(datetime), status, request, host)
```

## Create Table with Codec

```sql
CREATE TABLE IF NOT EXISTS `logs` (
  `host` String CODEC(ZSTD),
  `user-identifier` String CODEC(ZSTD),
  `datetime` DateTime CODEC(Delta, ZSTD),
  `method` String CODEC(ZSTD),
  `request` String CODEC(ZSTD),
  `protocol` String CODEC(ZSTD),
  `status` UInt64 CODEC(ZSTD),
  `bytes` UInt64 CODEC(ZSTD),
  `referer` String CODEC(ZSTD),
  `user-agent` String CODEC(ZSTD),
)
ENGINE = MergeTree
ORDER BY (toStartOfHour(datetime), status, request, host)
```

## Import Data

```sql
INSERT INTO logs FROM INFILE '/home/weny/Projects/flog/test.log' FORMAT CSVWithNames;
```

