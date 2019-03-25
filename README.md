#  Kubernetes Days India 2019

## Using Kubernetes API effectively with Golang

Slides can be found [at this link](https://docs.google.com/presentation/d/1QiU_yeVTw7H5sYollZD2FWXmNQZ0KO1t3Ca25lhwYyk/edit?usp=sharing)

To build:

```
glide install
./build.sh
```

To get help, just run binary and to to run a specific example pass the argument:

```
./kdi 
NAME:
   k8sdaysindia - Using client-go effectively with Kubernetes api

USAGE:
   kdi [global options] command [command options] [arguments...]

VERSION:
   0.2

COMMANDS:
     crud       Run CRUD example
     lister     Run lister example
     informer   Run informer example
     workqueue  Run workqueue example
     help, h    Shows a list of commands or help for one command

./kdi crud
./kdi lister
```
