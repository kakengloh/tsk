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
<img width="600" alt="tsk new output" src="https://user-images.githubusercontent.com/40446720/186668426-a5908430-c1db-4529-9206-6033571cff85.png">

### Create a new task with status, priority and due

```bash
tsk new 'feed my cat' -s doing -p high -d 1h
```
<img width="600" alt="tsk new with options output" src="https://user-images.githubusercontent.com/40446720/186668696-6ba2e1b3-d2d2-4db9-953b-ac706876f365.png">

### List tasks

```bash
tsk ls
```
<img width="600" alt="tsk ls output" src="https://user-images.githubusercontent.com/40446720/186668844-d73e83dd-e334-403c-b59e-e8410984c994.png">

### List tasks with filters (status, priority, due)

```bash
tsk ls -s doing -p high -d 1h
```
<img width="600" alt="tsk ls with filters output" src="https://user-images.githubusercontent.com/40446720/186668966-12b472d3-b38e-449c-b1ad-eec70c86ac42.png">

### List tasks with a keyword

```bash
tsk ls cat
```
<img width="600" alt="tsk ls with keyword output" src="https://user-images.githubusercontent.com/40446720/186669061-d20d7a1a-7c75-4225-a0e3-f450dbb193af.png">

### List tasks as JSON

```bash
tsk ls -f json
```
<img width="250" alt="tsk ls json format output" src="https://user-images.githubusercontent.com/40446720/186669184-f25cb05d-6625-41db-ac87-3e12b2c03ae0.png">

### View tasks in a Kanban board

```bash
tsk board
```
<img width="400" alt="tsk board output" src="https://user-images.githubusercontent.com/40446720/186669288-670f387c-0da8-42cd-a348-502c50853d4c.png">

### Mark task(s) as todo

```bash
tsk todo 2
```
<img width="250" alt="tsk todo output" src="https://user-images.githubusercontent.com/40446720/186669381-e5bde5b1-84bd-4cf8-9721-564739930b1e.png">

### Mark task(s) as doing

```bash
tsk doing 2
```
<img width="250" alt="tsk doing output" src="https://user-images.githubusercontent.com/40446720/186669448-5eedb3d0-af4b-4074-a42a-d9daf387571c.png">

### Mark task(s) as done

```bash
tsk done 2
```
<img width="250" alt="tsk done output" src="https://user-images.githubusercontent.com/40446720/186669471-7ab542ad-ce34-495d-b5cc-8aeb36d086d9.png">

### Modify an existing task

```bash
tsk mod 2 -s todo -p low
```
<img width="600" alt="tsk mod output" src="https://user-images.githubusercontent.com/40446720/186669548-1be2b856-5f2a-4e34-8788-bbfdf15f58a9.png">

### Add note(s) on a task

```bash
tsk note 2 'it still hungry' 'meow...'
```
<img width="600" alt="tsk note output" src="https://user-images.githubusercontent.com/40446720/186669611-8a7c67aa-ac04-479d-b1c3-46fd2829d24c.png">

### Remove task(s)

```bash
tsk rm 1
```

### Clean your data

```bash
tsk clean
```

## Todo
- [x] Task due
- [ ] Due reminder (via desktop notification)

## Contributing
We welcome all feature requests and pull requests! 🙋

---

<a href="https://www.buymeacoffee.com/kakengloh" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png" alt="Buy Me A Coffee" style="height: 60px !important;width: 217px !important;" ></a>
