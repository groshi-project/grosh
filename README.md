# grosh
grosh is a simple command-line client for [groshi](https://github.com/groshi-project).
> **grosh** is not released yet but has some stable useful functional already.
> Use `make build` command to build it, run `./build/grosh --help` to get some help.

## Installation
Build grosh binary:
```shell
make build
```

And then install it:
```shell
sudo make install
```


## Implemented commands
### User
* `grosh register <URL> [USERNAME]` create new groshi user
* `grosh login <URL> [USERNAME]` login to groshi server and store credentials
* `grosh logout` remove locally stored credentials

### Transactions
* `groshi new [--description=<TEXT>] [--timestamp=<TIME>] <AMOUNT> <CURRENCY>` create new transaction
* `groshi list --currency=<CURRENCY> --end-time=<TIME> <START-TIME>` list transactions for given period and optionally in given currency
* `grosh summary --currency=<CURRENCY>` show summary of transactions for given period and optionally in given currency