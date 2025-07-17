NAME = zip-file-creator
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
	@echo "$(RED)Usage: $(GREEN)./$(NAME) <file1> <file2> ... <fileN>$(NC)"

unzip:
	@echo "$(GREEN)[Running] $(RED)unzip compressed.zip$(NC)"
	unzip compressed.zip

clean:
	@rm -f $(NAME) go.mod go.sum compressed.zip

re: clean all

.PHONY: all clean re
