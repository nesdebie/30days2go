NAME = simple-cache
SRC = main.go

GREEN = \033[0;32m
RED = \033[0;31m
NC = \033[0m


all: $(NAME)

$(NAME):
	@echo "$(GREEN)[Running] $(RED)go mod init $(NAME)$(NC)"
	go mod init $(NAME)
	@echo "$(GREEN)[Running] $(RED)go build -o $(NAME) $(SRC)$(NC)"
	go build -o $(NAME) $(SRC)

clean:
	@rm -rf $(NAME) go.mod go.sum

re: clean all

.PHONY: all clean re
