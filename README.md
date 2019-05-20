# go-endpoint
[![CircleCI](https://circleci.com/gh/rpsraghu/go-endpoint/tree/master.svg?style=svg)](https://circleci.com/gh/rpsraghu/go-endpoint/tree/master)

Welcome to go-endpoint, a tool to monitor response time for a given endpoint

## Prerequisites
This requires Docker as a CI tool.

## Compile
```
$ git clone https://github.com/rpsraghu/go-endpoint.git
$ make build
```
Above will generate a docker image named rpsraghu/go-endpoint:X.X.X

## Usage

Start using the generated docker image in dev/prod environment as required.

## Development

```
$ docker run --rm --name=monitor-endpoint -e TARGET_HOST=yourdomain.com -e TARGET_PROTO=https -e GO_ENVIRONMENT=development rpsraghu/go-endpoint:X.X.X
```

## Production
```
$ docker run --rm --name=monitor-endpoint -e TARGET_HOST=yourdomain.com -e TARGET_PROTO=https rpsraghu/go-endpoint:X.X.X
```

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/rpsraghu/go-endpoint. This project is intended to be a safe, welcoming space for collaboration, and contributors are expected to adhere to the [Contributor Covenant](http://contributor-covenant.org) code of conduct.

## License

The Utility is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).

## Code of Conduct

Everyone interacting in the go-endpoint project’s codebases, issue trackers, chat rooms and mailing lists is expected to follow the [code of conduct](https://github.com/rpsraghu/endpoint_benchmark/blob/master/CODE_OF_CONDUCT.md).
