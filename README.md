<p align="center">
  <img width="250" src="https://user-images.githubusercontent.com/40446720/185377062-4859d1f4-25a4-44b6-b0d2-e887ee1d1060.png" />
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

## Usage

### Create a new task

```bash
tsk new <title>
tsk new 'make coffee' # example
```

<img width="450" alt="Screenshot 2022-08-18 at 5 13 09 PM" src="https://user-images.githubusercontent.com/40446720/185358144-fd4f6cc2-ac2d-4f11-92dc-44292dec52f2.png">

### Create a new task with status and priority

```bash
tsk new 'feed my cat' -s doing -p high
```

<img width="445" alt="Screenshot 2022-08-18 at 5 14 11 PM" src="https://user-images.githubusercontent.com/40446720/185358344-a06cb6ea-28f2-4d11-a9a2-8d6aac999792.png">

### List tasks

```bash
tsk ls
```

<img width="463" alt="Screenshot 2022-08-18 at 5 15 00 PM" src="https://user-images.githubusercontent.com/40446720/185358509-10908435-daf2-4b4f-ad08-100e10086291.png">

### List tasks with filters

```bash
tsk ls -s doing -p high
```

<img width="451" alt="Screenshot 2022-08-18 at 6 36 28 PM" src="https://user-images.githubusercontent.com/40446720/185375242-2bb16bf3-b936-4d0c-b742-d2fa808e4a00.png">

### Find tasks with a keyword

```bash
tsk find <keyword>
tsk find cat # example
```

<img width="464" alt="Screenshot 2022-08-18 at 5 15 54 PM" src="https://user-images.githubusercontent.com/40446720/185358696-2e55d92a-21a3-49c0-a195-90e2490db808.png">

### View tasks in a Kanban board

```bash
tsk board
```

<img width="318" alt="Screenshot 2022-08-18 at 5 16 27 PM" src="https://user-images.githubusercontent.com/40446720/185358801-b5e0de6d-1244-4ac8-a65f-3e73bdd72d49.png">

### Mark task(s) as todo

```bash
tsk todo <id> <id2> ...
tsk todo 2 # example
```

<img width="208" alt="Screenshot 2022-08-18 at 5 17 09 PM" src="https://user-images.githubusercontent.com/40446720/185358924-89528adf-81f5-434e-8658-41117d8507e6.png">

### Mark task(s) as doing

```bash
tsk doing <id> <id2> ...
tsk doing 2 # example
```

<img width="204" alt="Screenshot 2022-08-18 at 5 17 42 PM" src="https://user-images.githubusercontent.com/40446720/185359025-55f5d4b1-09c1-48f8-9424-9fea5cabc638.png">

### Mark task(s) as done

```bash
tsk done <id> <id2> ...
tsk done 2 # example
```

<img width="206" alt="Screenshot 2022-08-18 at 5 17 58 PM" src="https://user-images.githubusercontent.com/40446720/185359098-3d385a9c-0043-493c-8c13-a83e7753df69.png">

### Modify an existing task

```bash
tsk mod <id> -s todo -p medium
tsk mod 2 -s todo -p low # example
```

<img width="458" alt="Screenshot 2022-08-18 at 5 18 49 PM" src="https://user-images.githubusercontent.com/40446720/185359281-8aa1fbcc-95e2-40b2-975e-c47e80c8809c.png">

### Add note(s) on a task

```bash
tsk note <id> <note1> <note2> ...
tsk note 2 'it still hungry' # example
```

### Remove task(s)

```bash
tsk rm <id> <id2> ...
tsk rm 1 2 # example
```

<img width="195" alt="Screenshot 2022-08-18 at 5 21 00 PM" src="https://user-images.githubusercontent.com/40446720/185359793-faa50ea3-9466-4b95-9dc7-b8ecdea0782d.png">
