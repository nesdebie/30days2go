# 30days2go

My goal is to learn Go in 30 Days Through 30 Projects of Increasing Complexity

### Compilation
Every mini-project has its own Makefile compiling the project different way

```
# 1) make run --> directly execute the main.go file or the binary if possible (no args)
$> go run main.go
Hello World!

# 2) make --> create a binary called main
$> go build main.go
$> ./main
Hello World!

# 3) make rename --> create a named binary (without using mv)
$> go mod init hello-world
$> go build hello-world
$> ./hello-world
Hello-World!
```

<p align="center">
  <img src="https://go.dev/blog/go-brand/Go-Logo/SVG/Go-Logo_Blue.svg" width="160" alt="Go Logo">
</p>
