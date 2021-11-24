# CPH-Businnes-LS-Exam


# Table of Contents
  - [How to run](#how-to-run)
  - [The team](#the-team)
  - [Assignment](#assignment)
    - [Objective](#objective)
    - [Task](#task)
  - [Introduction](#introduction)
  - [Project Achitecture](#project-achitecture)

## How to run 

#### First start up af minikube cluster
```
$ minikube start
```
#### Now run Terraform to apply the infrasturcture
```
$ cd terraform
$ terraform apply --var-file=dev.tfvars -auto-approve
```

## The team
- [Andreas Vikke (cph-av105)](https://github.com/andreasvikke)
- [Sofus Nielsen (cph-sn311)](https://github.com/sofushn)
- [Martin Frederiksen (cph-mf237)](https://github.com/MartinFrederiksen)

## Assignment
TODO: Insert Assignment PDF
## Introduction
TODO: Write this

## Project Achitecture
![image](assets/achitecture.png)


av-105


python -m grpc_tools.protoc -I../../protos --python_out=. --grpc_python_out=. ../../protos/route_guide.proto