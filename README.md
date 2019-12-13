# vartf2tfvar

[![GoDoc](https://godoc.org/github.com/zaap59/vartf2tfvar?status.svg)](https://godoc.org/github.com/zaap59/vartf2tfvar)

Convert `variable.tf` file to tfvars format.

`./vartf2tfvar -file ./variables-sample.tf` will produce

```hcl
alb_name = ""

alb_sg = ""

alb_target_health_check = ""

alb_target_port = ""

app_name = ""

app_private_dns = ""

owner = ""

user_sg_ids = ""
```

## Usage

```sh
./vartf2tfvar
Usage: ./vartf2tfvar [OPTIONS] argument ...
  -file string
        a string var (default "variables.tf")
```

## Install

```sh
go get -u github.com/zaap59/vartf2tfvar
```

## Build

```sh
go build -o vartf2tfvar .
```
