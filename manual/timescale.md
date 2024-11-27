# TimescaleDB

## Install TimescaleDB

```sh
docker run -d \
  --name timescaledb \
  -p 5433:5432 \
  -e POSTGRES_PASSWORD=password \
  timescale/timescaledb-ha:pg17
```

## Create Table

```sql
CREATE TABLE IF NOT EXISTS logs (
  host TEXT,
  "user-identifier" TEXT,
  datetime TIMESTAMPTZ NOT NULL,
  method TEXT,
  request TEXT,
  protocol TEXT,
  status BIGINT,
  bytes BIGINT,
  referer TEXT,
  "user-agent" TEXT
);
```

## Convert to Hypertable
```sql
SELECT create_hypertable('logs', 'datetime');
```

## Install timescaledb-parallel-copy

```sh
go install github.com/timescale/timescaledb-parallel-copy/cmd/timescaledb-parallel-copy@latest
```
