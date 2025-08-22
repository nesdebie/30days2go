NAME = ws-chat
SRC = main.go

GREEN = \033[0;32m
RED = \033[0;31m
NC = \033[0m


all: $(NAME)

$(NAME):
	@echo "$(GREEN)[Running] $(RED)go mod init $(NAME)$(NC)"
	go mod init $(NAME)
	go get -u github.com/gorilla/websocket
	@echo "$(GREEN)[Running] $(RED)go build -o $(NAME) $(SRC)$(NC)"
	go build -o $(NAME) $(SRC)
	@echo "$(RED)Usage: $(GREEN)./$(NAME)$(NC)"
	@echo "$(GREEN)--- How to join $(NAME) ? ---$(NC)"
	@echo "$(RED)Install:  $(GREEN)cargo install websocat$(NC)"
	@echo "$(RED)Run: $(GREEN)websocat ws://localhost:8080/ws$(NC)"

clean:
	@rm -rf $(NAME) go.mod go.sum

re: clean all

.PHONY: all clean re
