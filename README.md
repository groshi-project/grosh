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
* `grosh register <URL> [USERNAME]` create a new groshi user
* `grosh login <URL> [USERNAME]` login to groshi server and store credentials locally
* `grosh logout` remove locally stored credentials

### Transactions
* `grosh new [--timestamp=<TIME>] [--description=<TEXT>] <AMOUNT> <CURRENCY>` create a new transaction
* `grosh list [--uuid] [--currency=<CURRENCY>] [--end-time=<TIME>] <START-TIME>` list transactions in given period and optionally in given currency
* `grosh summary [--end-time=<END-TIME>] <START-TIME> <CURRENCY>` show summary of transactions for given period in given currency
* `grosh update [--amount=AMOUNT] [--currency=CURRENCY] [--description=DESCRIPTION] [--timestamp=TIME] <UUID>` update transaction
* `grosh remove <UUID>` remove transaction