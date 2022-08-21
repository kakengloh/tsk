<p align="center">
  <img width="250" alt="tsk logo" src="https://user-images.githubusercontent.com/40446720/185574124-28e9e2b4-bdfc-4aa8-aaed-c82d91576a97.png" />
</p>
<h1 align="center"> tsk </h1>
<p align="center">
  <b>tsk is a terminal task management app with an emphasis on simplicity, efficiency and ease of use</b>
</p>
<p align="center">
<a href="https://github.com/kakengloh/tsk/actions/workflows/build.yml"><img src="https://github.com/kakengloh/tsk/actions/workflows/build.yml/badge.svg" /></a> <a href="https://goreportcard.com/report/github.com/kakengloh/tsk"><img src="https://goreportcard.com/badge/github.com/kakengloh/tsk" /></a> <a href="https://github.com/kakengloh/tsk"><img src="https://img.shields.io/github/go-mod/go-version/kakengloh/tsk.svg" /></a>
</p>

<br>

## Description

`tsk` allows you to create and manage your tasks efficiently your terminal, so that you can dump your mouse 🖱️

## Why
Shiny task management web apps often have tons of unnecessary features causing UI glitches which impact our UX 😵‍💫 <br> The huge JS chunks loading and network calls on every smallest action causing feedback delay and it is annoying ⌛

Personal task management should be as simple as possible, let's build a snappy cli app that doesn't get in our way ✌️

## Features

- Simple and concise commands ✅
- Works without internet connection ✅
- Data stored locally - powered by [BoltDB](https://github.com/etcd-io/bbolt) ✅

## Installation

### Go

```bash
# Go 1.16+
go install github.com/kakengloh/tsk@latest

# Go < 1.16
GO111MODULE=on go get github.com/kakengloh/tsk
```

> Ensure that `$GOPATH/bin` is in your `PATH`

### Executables

See [releases](https://github.com/kakengloh/tsk/releases) for executables

## Example

### Create a new task

```bash
tsk new 'make coffee'
```
<img width="479" alt="tsk new output" src="https://user-images.githubusercontent.com/40446720/185779302-427fce50-f7b2-42fe-8018-707937bfcfc6.png">

### Create a new task with status, priority and due

```bash
tsk new 'feed my cat' -s doing -p high -d 1h
```
<img width="523" alt="tsk new with options output" src="https://user-images.githubusercontent.com/40446720/185779314-05cd0a17-5a28-4cac-aa84-73754643e4af.png">

### List tasks

```bash
tsk ls
```
<img width="529" alt="tsk ls output" src="https://user-images.githubusercontent.com/40446720/185779321-cb894804-ab7f-4448-9671-2465d1c5652a.png">

### List tasks with filters

```bash
tsk ls -s doing -p high
```
<img width="533" alt="tsk ls with filters output" src="https://user-images.githubusercontent.com/40446720/185779328-8fc15724-5fd6-426b-81a0-536ec05bbf6a.png">

### List tasks with a keyword

```bash
tsk ls cat
```
<img width="527" alt="tsk ls with keyword output" src="https://user-images.githubusercontent.com/40446720/185779335-be1bf537-9b35-4021-9482-2f5cfbfbfabd.png">

### List tasks as JSON

```bash
tsk ls -f json
```
<img width="262" alt="tsk ls json format output" src="https://user-images.githubusercontent.com/40446720/185779832-0495fab5-1c91-4d75-ad1a-f1dc1ac1e6a0.png">

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
<img width="532" alt="tsk mod output" src="https://user-images.githubusercontent.com/40446720/185779357-c74fa73a-0f72-457b-8b96-3c2218a40f9a.png">

### Add note(s) on a task

```bash
tsk note 2 'it still hungry' 'meow...'
```
<img width="624" alt="tsk note output" src="https://user-images.githubusercontent.com/40446720/185779360-49eb2425-6491-48e4-b556-e92cd7a43aa1.png">

### Remove task(s)

```bash
tsk rm <id> <id2> ...
tsk rm 1 2 # example
```

<img width="195" alt="tsk rm output" src="https://user-images.githubusercontent.com/40446720/185359793-faa50ea3-9466-4b95-9dc7-b8ecdea0782d.png">

### Clean your data

```bash
tsk clean
```

## Todo
- [ ] Task deadline and reminder (via desktop notification)

## Contributing
We welcome all feature requests and pull requests! 🙋
