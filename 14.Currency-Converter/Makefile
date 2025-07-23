TARGET = main
NAME = currency-converter
SRC = main.go

GREEN = \033[0;32m
RED = \033[0;31m
NC = \033[0m

all: $(NAME)

$(NAME):
	@if [ -f $(NAME) ]; then \
		echo "$(NAME) already exists, skipping rename"; \
		exit 0; \
	elif [ -f $(TARGET) ]; then \
		rm -f $(TARGET); \
	fi
	@echo "$(GREEN)[Running] $(RED)go mod init $(NAME)$(NC)"
	go mod init $(NAME)
	@echo "$(GREEN)[Running] $(RED)go build -o $(NAME) $(SRC)$(NC)"
	go build -o $(NAME) $(SRC)
	@echo "$(RED)Usage: $(GREEN)./$(NAME) <amount> <from_currency> <to_currency>$(NC)"

clean:
	@rm -f $(NAME) go.mod go.sum

re: clean all

.PHONY: all clean csv re
