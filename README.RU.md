## logger
обертка над [logrus](https://github.com/sirupsen/logrus)

передает логгер в "контексте"

```go
log := NewLogger()
ctx := WithLogger(context.Background, log) 
```

```go
GetLogger(ctx).WithField("some_field", 256).Error("some error")
```

логирует по дефолту используя текущее время


