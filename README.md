# confunc

[![Build Status](https://travis-ci.org/alperkose/confunc.svg?branch=master)](https://travis-ci.org/alperkose/confunc) [![Coverage Status](https://coveralls.io/repos/github/alperkose/confunc/badge.svg)](https://coveralls.io/github/alperkose/confunc)

`confunc` provides a functional approach to configure an application. With `confunc`, each configuraion value can be defined as a function, providing a dynamic value rather than a static one that is assigned once the application starts up.

## Basic Usage
Example:
```go
// declare your configuration value
maxNumberOfOpenConnections := confunc.From(confunc.Env()).Int("MAX_OPEN_CONNECTIONS")
...
// when you need to access the value
if connectionCount < maxNumberOfOpenConnections() {
    // do your thing
}
```

## Source

Source is where the configuration is retrieved. It is defined by the `Source` interface where a string value is returned for a specified key. Once it is passed to `confunc.From()` function; a struct that can return configuration functions is created. Available sources are:
- `Env()` : retrieves Environment Variables
- `Map(map[string]string)` : retrieves configuration from a map
- `cfconsul.Source(*api.Config)` : retrieves configuration from consul. You need to import `confunc/cfconsul` package

## Interceptor
`confunc` provides a middleware mechanism via `Interceptor` type. It enables handling of various cases like default value

```go
maxNumberOfOpenConnections := confunc.From(confunc.Env()).Int("MAX_OPEN_CONNECTIONS", confunc.Default(10))
...
// when you need to access the value
if connectionCount < maxNumberOfOpenConnections() {
    // do your thing
}
```

`Interceptor` is a simple function that accepts a function that returns a string and also returns string.

```go
type String func() string

type Interceptor func(String) string
```

Interceptors process the configuration value before it is accessed, every time. Custom interceptors can be easily introduces by providing a function matching the above signature. Available interceptors are:
- `Default(defaultValue string)` : returns a default value if there is no configuration made
- `CacheOnce()` : caches the value retrieved once and prevents accessing the configuration source each time
- `CacheFor(cacheDuration time.Duration)` : caches the value for a duration provided.