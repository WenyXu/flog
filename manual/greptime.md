# GreptimeDB

## Create Table

```sql
CREATE TABLE IF NOT EXISTS `logs` (
  `host` STRING NULL,
  `user-identifier` STRING NULL,
  `datetime` TIMESTAMP(0) NULL,
  `method` STRING NULL,
  `request` STRING NULL,
  `protocol` STRING NULL,
  `status` BIGINT NULL,
  `bytes` BIGINT NULL,
  `referer` STRING NULL,
  `user-agent` STRING NULL,
  TIME INDEX (`datetime`)
)
```

## Import Data

```sh
COPY logs FROM '/home/weny/Projects/flog/test.log' WITH (format = 'csv');
```
