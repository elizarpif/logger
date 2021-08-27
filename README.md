## logger
wrapper over [logrus](https://github.com/sirupsen/logrus)

sends logger in the "context"

```go
log := NewLogger()
ctx := WithLogger(context.Background, log) 
```

```go
GetLogger(ctx).WithField("some_field", 256).Error("some error")
```

by default, logges using current datetime

