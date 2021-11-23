# time

## 获取现在时间

```
func Now() Time 
```

## 时间转化为指定格式

```
func (t Time) Format(layout string) string
```

```
func Parse(layout, value string) (Time, error)
```

## 比较时间大小

```
func (t Time) Before(u Time) bool
```

```
func (t Time) Equal(u Time) bool
```

```
func (t Time) After(u Time) bool
```

## 增加或减少某个时间

```
func ParseDuration(s string) (Duration, error)
```

```
func (t Time) Add(d Duration) Time
```

## 某个时间转化为0时0分0秒

```
func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
```

## 判断时间是否为空
```
func (t Time) IsZero() bool 
```