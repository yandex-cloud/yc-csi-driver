# Yandex.Cloud Go SDK

[![GoDoc](https://godoc.org/github.com/yandex-cloud/go-sdk?status.svg)](https://godoc.org/github.com/yandex-cloud/go-sdk)
[![CircleCI](https://circleci.com/gh/yandex-cloud/go-sdk.svg?style=shield)](https://circleci.com/gh/yandex-cloud/go-sdk)

Go SDK for Yandex.Cloud services.

**NOTE:** SDK is under development, and may make
backwards-incompatible changes.

## Installation

```bash
go get github.com/yandex-cloud/go-sdk
```

## Example usages

### Initializing SDK

```go
sdk, err := ycsdk.Build(ctx, ycsdk.Config{
	Credentials: ycsdk.OAuthToken(token),
})
if err != nil {
	log.Fatal(err)
}
```

### Retries

If you want to add retries to SDK, provide configuration using `retry/v1` package:

```go
import (
	...

	ycsdk "bb.yandex-team.ru/cloud/cloud-go/sdk"
	"bb.yandex-team.ru/cloud/cloud-go/sdk/pkg/retry/v1"
)

...

retriesDialOption, err := retry.RetryDialOption(
	retry.WithRetries(retry.DefaultNameConfig(), 2),
	retry.WithRetryableStatusCodes(retry.DefaultNameConfig(), codes.AlreadyExists, codes.Unavailable),
)
if err != nil {
	log.Fatal(err)
}

_, err = ycsdk.Build(
	ctx,
	ycsdk.Config{
		Credentials: ycsdk.OAuthToken(*token),
	},
	retriesDialOption,
)
```

It's **strongly recommended** to use provided default configuration to avoid retry amplification:

```go
import (
	...

	ycsdk "bb.yandex-team.ru/cloud/cloud-go/sdk"
	"bb.yandex-team.ru/cloud/cloud-go/sdk/pkg/retry/v1"
)

...

retriesDialOption, err := retry.DefaultRetryDialOption()
if err != nil {
	log.Fatal(err)
}

_, err = ycsdk.Build(
	ctx,
	ycsdk.Config{
		Credentials: ycsdk.OAuthToken(*token),
	},
	retriesDialOption,
)
```


### More examples

More examples can be found in [examples dir](examples).
