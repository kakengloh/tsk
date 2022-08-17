<!-- <p align="center">
  <img width="250" src="">
</p> -->
<h1 align="center"> tsk </h1>
<p align="center">
  <b>tsk is a simple cli personal task management tool</b>
</p>

<br>

## Description

`tsk` allows you to create and manage your tasks efficiently with few keystrokes.

Features:

- straightforward commands with various options
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

## Usage

### Create a new task

```bash
tsk new 'make coffee'
```

### Create a new task with status and priority

```bash
tsk new 'feed my cat' -s doing -p high
```

### List tasks

```bash
tsk ls
```

### Find tasks with a keyword

```bash
tsk find <keyword>
```

### View tasks in a Kanban board

```bash
tsk board
```

### Mark task(s) as todo

```bash
tsk todo <id> <id2> <id3>
```

### Mark task(s) as doing

```bash
tsk doing <id> <id2> <id3>
```

### Mark task(s) as done

```bash
tsk done <id> <id2> <id3>
```

### Modify an existing task

```bash
tsk mod <id> -n workout -s todo -p medium
```

### Comment on a task

```bash
tsk cmt <id> '3 sets of push ups'
```

### Remove task(s)

```bash
tsk rm <id> <id2> <id3>
```
