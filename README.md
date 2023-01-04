# adb
Airdb Development Builder

## Quick Install

```
go install -ldflags -X=github.com/airdb/adb/internal/adblib.BuildTime=$(date +%s) github.com/airdb/adb@latest
```

## Download

```
https://github.com/airdb/adb/releases/latest/download/adb
https://github.com/airdb/adb/releases/latest/download/adb-darwin
https://github.com/airdb/adb/releases/latest/download/adb-linux-amd64.zip


https://github.com/airdb/adb/releases/download/v1.0.0/adb
```


# For Template Maintaining

```bash
go get github.com/rakyll/statik
statik -include='*' -src gin-template -f
```
