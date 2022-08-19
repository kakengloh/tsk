<p align="center">
  <img width="250" alt="tsk logo" src="https://user-images.githubusercontent.com/40446720/185574124-28e9e2b4-bdfc-4aa8-aaed-c82d91576a97.png" />
</p>
<h1 align="center"> tsk </h1>
<p align="center">
  <b>tsk is a simple cli personal task management tool</b>
</p>
<p align="center">
<a href="https://github.com/kakengloh/tsk/actions/workflows/build.yml"><img src="https://github.com/kakengloh/tsk/actions/workflows/build.yml/badge.svg" /></a> <a href="https://goreportcard.com/report/github.com/kakengloh/tsk"><img src="https://goreportcard.com/badge/github.com/kakengloh/tsk" /></a> <a href="https://github.com/kakengloh/tsk"><img src="https://img.shields.io/github/go-mod/go-version/kakengloh/tsk.svg" /></a>
</p>

<br>

## Description

`tsk` allows you to create and manage your tasks efficiently with few keystrokes.

Features:

- concise commands with various options
- shell commands auto completion
- works perfectly fine without internet connection
- you own your data (data is stored in your local machine)
- clean user interface

## Installation

Go:

```bash
# Go 1.16+
go install github.com/kakengloh/tsk@latest

# Go < 1.16
GO111MODULE=on go get github.com/kakengloh/tsk
```

> Ensure that `$GOPATH/bin` is in your `PATH`

## Example

### Create a new task

```bash
tsk new 'make coffee'
```
<img width="422" alt="tsk new output" src="https://user-images.githubusercontent.com/40446720/185561994-1be87426-0130-4e8f-953e-22035fc62c8c.png">

### Create a new task with status and priority

```bash
tsk new 'feed my cat' -s doing -p high
```
<img width="428" alt="tsk new with options output" src="https://user-images.githubusercontent.com/40446720/185562720-de877827-3547-4582-9f86-3d7843a3581b.png">

### List tasks

```bash
tsk ls
```
<img width="440" alt="tsk ls output" src="https://user-images.githubusercontent.com/40446720/185562098-bbfe2e4e-1718-43ed-b230-3619f8b0a89f.png">

### List tasks with filters

```bash
tsk ls -s doing -p high
```
<img width="446" alt="tsk ls with filters output" src="https://user-images.githubusercontent.com/40446720/185562147-b5a99efe-d2ba-467b-a4cf-c7c62478cb3c.png">

### List tasks with a keyword

```bash
tsk ls cat
```
<img width="438" alt="tsk ls with keyword output" src="https://user-images.githubusercontent.com/40446720/185562223-c40f92ab-43cb-480c-9b2e-686a226b8193.png">

### View tasks in a Kanban board

```bash
tsk board
```
<img width="311" alt="tsk board output" src="https://user-images.githubusercontent.com/40446720/185562238-6d245e95-303d-4f66-9c51-cc961ba55ddd.png">

### Mark task(s) as todo

```bash
tsk todo 2
```
<img width="208" alt="tsk todo output" src="https://user-images.githubusercontent.com/40446720/185358924-89528adf-81f5-434e-8658-41117d8507e6.png">

### Mark task(s) as doing

```bash
tsk doing 2
```
<img width="204" alt="tsk doing output" src="https://user-images.githubusercontent.com/40446720/185359025-55f5d4b1-09c1-48f8-9424-9fea5cabc638.png">

### Mark task(s) as done

```bash
tsk done 2
```
<img width="206" alt="tsk done output" src="https://user-images.githubusercontent.com/40446720/185359098-3d385a9c-0043-493c-8c13-a83e7753df69.png">

### Modify an existing task

```bash
tsk mod 2 -s todo -p low
```
<img width="443" alt="tsk mod output" src="https://user-images.githubusercontent.com/40446720/185562386-3c1d0b22-1ad3-4c9d-83e4-7f0344e5cdfe.png">

### Add note(s) on a task

```bash
tsk note 2 'it still hungry' 'meow...'
```
 <img width="536" alt="tsk note output" src="https://user-images.githubusercontent.com/40446720/185562436-656295d8-0285-4cd7-a329-55adfccfaeb8.png">

### Remove task(s)

```bash
tsk rm <id> <id2> ...
tsk rm 1 2 # example
```

<img width="195" alt="tsk rm output" src="https://user-images.githubusercontent.com/40446720/185359793-faa50ea3-9466-4b95-9dc7-b8ecdea0782d.png">
