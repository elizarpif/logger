## logger
wrapper over [logrus](https://github.com/sirupsen/logrus)

sends logger in the "context"

```go
log := logger.NewLogger()
ctx := logger.WithLogger(context.Background(), log) 
```

```go
logger.GetLogger(ctx).WithField("some_field", 256).Error("some error")
```

by default, logges using current datetime


### Installation

```shell
 go get github.com/elizarpif/logger 
 ```
