NAME = qr-code-generator
SRC = main.go

GREEN = \033[0;32m
RED = \033[0;31m
NC = \033[0m

all: $(NAME)

$(NAME):
	@echo "$(GREEN)[Running] $(RED)go mod init $(NAME)$(NC)"
	go mod init $(NAME)
	mkdir -p img
	go get github.com/skip2/go-qrcode
	@echo "$(GREEN)[Running] $(RED)go build -o $(NAME) $(SRC)$(NC)"
	go build -o $(NAME) $(SRC)
	@echo "$(RED)Usage: $(GREEN)./$(NAME) <text_or_url>$(NC)"

clean:
	@rm -rf $(NAME) go.mod go.sum img

re: clean all

.PHONY: all clean csv re
