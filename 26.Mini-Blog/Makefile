NAME = mini-blog
SRC = main.go

GREEN = \033[0;32m
RED = \033[0;31m
BLUE = \033[0;34m
YELLOW = \033[0;33m
NC = \033[0m


all: $(NAME)

$(NAME):
	@echo "$(GREEN)[Running] $(RED)go mod init $(NAME)$(NC)"
	go mod init $(NAME)
	@echo "$(GREEN)[Running] $(RED)go build -o $(NAME) $(SRC)$(NC)"
	go get github.com/mattn/go-sqlite3
	go build -o $(NAME) $(SRC)
	@echo "$(RED)---- $(GREEN)USAGES$(RED) ----$(NC)"
	@echo "curl -X $(RED)GET$(NC) http://localhost:8080/articles"
	@echo "curl -X $(GREEN)DELETE$(NC) http://localhost:8080/articles/$(YELLOW){article_id}$(NC)"
	@echo "curl -X $(BLUE)POST$(NC) http://localhost:8080/articles \\"
	@echo "    -H \"Content-Type: application/json\" \\"
	@echo "    -d '{\"author\": \"nesdebie\", \"title\": \"Go and SQLite\", \"content\": \"First!\"}'"
	@echo "$(RED)----- $(GREEN)END$(RED) ------$(NC)"

clean:
	@rm -rf $(NAME) go.mod go.sum database.db

re: clean all

.PHONY: all clean re
