NAME = log-analyzer
SRC = main.go

GREEN = \033[0;32m
RED = \033[0;31m
NC = \033[0m

all: $(NAME)

$(NAME):
	@echo "$(GREEN)[Running] $(RED)go mod init $(NAME)$(NC)"
	go mod init $(NAME)
	go build -o $(NAME) $(SRC)
	@echo "$(RED)Usage: $(GREEN)./$(NAME) <logfile>$(NC)"

clean:
	@rm -rf $(NAME) go.mod go.sum

re: clean all

.PHONY: all clean csv re
