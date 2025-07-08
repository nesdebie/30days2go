TARGET = main
NAME = password-generator
SRC = main.go
GREEN = \033[0;32m
RED = \033[0;31m
NC = \033[0m # No Color
all: $(TARGET)

$(TARGET):
	@if [ -f $(TARGET) ]; then \
		@echo "$(NAME) already exists, skipping build"; \
		exit 0; \
	elif [ -f $(NAME) ]; then \
		rm -f $(NAME) go.mod; \
	fi
	@echo "$(GREEN)[Running] $(RED)go build -o $(TARGET) $(SRC)$(NC)";
	go build -o $(TARGET) $(SRC)

run:
	@if [ -f $(TARGET) ]; then \
		echo "$(GREEN)[Running] $(RED)./$(TARGET)$(NC)"; \
		./$(TARGET); \
	elif [ -f $(NAME) ]; then \
		echo "$(GREEN)[Running] $(RED)./$(NAME)$(NC)"; \
		./$(NAME); \
	else \
		echo "$(GREEN)-> Running: $(RED)go run $(SRC)$(NC)"; \
		go run $(SRC); \
	fi

rename:
	@if [ -f $(NAME) ]; then \
		@echo "$(NAME) already exists, skipping rename"; \
		exit 0; \
	elif [ -f $(TARGET) ]; then \
		rm -f $(TARGET); \
	fi
	@echo "$(GREEN)[Running] $(RED)go mod init $(NAME)$(NC)";
	@go mod init $(NAME)
	@echo "$(GREEN)[Running] $(RED)go build -o $(NAME) $(SRC)$(NC)";
	go build -o $(NAME) $(SRC)

clean:
	rm -f $(TARGET) $(NAME) go.mod

re: clean all

.PHONY: all clean rename run re
.PHONY: all clean rename run re