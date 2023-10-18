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
Create a new groshi user:
* `grosh register <URL> [USERNAME]`

Login to groshi server and store credentials locally:
* `grosh login <URL> [USERNAME]` 

Remove locally stored credentials:
* `grosh logout` 

### Transactions
Create a new transaction:
* `grosh new [--timestamp=<TIME>] [--description=<TEXT>] <AMOUNT> <CURRENCY>`

List transactions in given period and optionally in given currency:
* `grosh list [--uuid] [--currency=<CURRENCY>] [--end-time=<TIME>] <START-TIME>` 

Show summary of transactions for given period in given currency:
* `grosh summary [--end-time=<END-TIME>] <START-TIME> <CURRENCY>`

Update transaction:
* `grosh update [--amount=AMOUNT] [--currency=CURRENCY] [--description=DESCRIPTION] [--timestamp=TIME] <UUID>` 

Remove transaction:
* `grosh remove <UUID>` 
