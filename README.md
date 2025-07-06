# 30days2go

## Goal
Learn Go in 30 Days Through 30 Projects of Increasing Complexity

### Compilation
Every mini-project has its own Makefile compiling the project different way

```
# 1) make run --> directly execute the main.go file
$> go run main.go
Hello World!

# 2) make && make run --> create a binary called main then execute it
$> go build main.go
$> ./main
Hello World!

# 3) make rename && make run --> create a binary renamed  hello-world then execute it
$> go mod init hello-world
$> go build hello-world
$> ./hello-world
Hello-World!
```
