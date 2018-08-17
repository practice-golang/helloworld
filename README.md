# helloworld

## 내려받기, 빌드, 실행
```sh
$ cd your workspace
workspace$ git clone https://github.com/practice-golang/helloworld.git
workspace$ cd helloworld
hello$ go build hello
hello$ ./hello.exe
```

## Package
* =라이브러리
* [Golang 기본 제공하는 것](https://golang.org/pkg/)도 있고 제3자가 개발하여 배포하기도 한다.
* Package는 사용되려면 아래와 같이 되어있어야 한다.
  * %Gopath%/src 하위에 존재해야 사용할 수 있다.
  * main과 같은 depth의 경로에 존재해야한다.
  * 아니면, main 또는 package 하위에 넣으려면 vendor 폴더 밑에 있어야 한다.
    * 따라서 이 레포지터리 소스에 있는 spkg라는 package는 쓸 수 없는 위치에 있으므로 사용 불가
    * [dep](https://github.com/golang/dep)으로 vendor 폴더에 들어갈 패키지를 관리할 수 있다.

## import
```go
// Package 불러오기
import "fmt"

import (
    "fmt"
    "mypkg"
)

// Package 사용
fmt.Println("안녕 세상!")
```
