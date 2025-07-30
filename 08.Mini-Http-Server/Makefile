NAME = web-server-go
SRC = main.go

GREEN = \033[0;32m
RED = \033[0;31m
NC = \033[0m

all: $(NAME)

$(NAME):
	@echo "$(GREEN)[Running] $(RED)go mod init $(NAME)$(NC)"
	go mod init $(NAME)
	@echo "$(GREEN)[Running] $(RED)go build -o $(NAME) $(SRC)$(NC)"
	go get github.com/go-chi/chi/v5
	go get github.com/go-chi/chi/v5/middleware

	go build -o $(NAME) $(SRC)
	@echo "$(RED)Usage: $(GREEN)./$(NAME)$(NC)"
	@echo "$(RED)Visit: $(GREEN)curl localhost:8080/hello/<your_name>$(NC)"
	@echo "$(RED)Visit: $(GREEN)curl localhost:8080/time$(NC)"

clean:
	@rm -f $(NAME) go.mod go.sum

re: clean all

.PHONY: all clean re
