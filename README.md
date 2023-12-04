# internettime
Package to convert time to Internet Time (.beats) and the other way

## Convert A Duration String to Internet Time
```go
func DurationToInternetTime(durationStr string) (float64, error)
```
Example:
```go
internetTimeDuration := DurationToInternetTime("1h10m12s")
```

## Convert Internet Time to Standard Time
```go
func InternetToStandardTime(beats float64) time.Time
```

Example:
```go
standardTime := InternetToStandardTime(413)
```

## Convert Standard Time to Internet Time
```go
func StandardToInternetTime(t time.Time) float64
```

Example:
```go
now := time.Now()
internetTime := StandardToInternetTime(now)
```
