#/bin/bash

# Mac M1
GOOS=darwin
GOARCH=arm64
echo 'Building for Mac M1...'
go build -o ./bin/tsk-darwin-arm64
if [ $? -ne 0 ]; then
  echo 'Build failed'
  exit 1
fi

# Mac Intel
GOOS=darwin
GOARCH=amd64
echo 'Building for Mac Intel...'
go build -o ./bin/tsk-darwin-amd64
if [ $? -ne 0 ]; then
  echo 'Build failed'
  exit 1
fi

# Linux 64-bit
GOOS=linux
GOARCH=amd64
echo 'Building for Linux 64-bit...'
go build -o ./bin/tsk-linux-amd64
if [ $? -ne 0 ]; then
  echo 'Build failed'
  exit 1
fi

# Linux 32-bit
GOOS=linux
GOARCH=386
echo 'Building for Linux 32-bit...'
go build -o ./bin/tsk-linux-386
if [ $? -ne 0 ]; then
  echo 'Build failed'
  exit 1
fi

# Windows 64-bit
GOOS=windows
GOARCH=amd64
echo 'Building for Windows 64-bit...'
go build -o ./bin/tsk-windows-amd64.exe
if [ $? -ne 0 ]; then
  echo 'Build failed'
  exit 1
fi

# Windows 32-bit
GOOS=windows
GOARCH=386
echo 'Building for Windows 32-bit...'
go build -o ./bin/tsk-windows-386.exe
if [ $? -ne 0 ]; then
  echo 'Build failed'
  exit 1
fi

echo 'Build success ✨'
exit 0