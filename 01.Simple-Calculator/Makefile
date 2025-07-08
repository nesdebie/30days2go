TARGET = main
NAME = calculator
SRC = main.go

GREEN = \033[0;32m
RED = \033[0;31m
NC = \033[0m

all: $(TARGET)

$(TARGET):
	@if [ -f $(TARGET) ]; then \
		echo "$(TARGET) already exists, skipping build"; \
		exit 0; \
	elif [ -f $(NAME) ]; then \
		rm -f $(NAME) go.mod; \
	fi
	@echo "$(GREEN)[Running] $(RED)go build -o $(TARGET) $(SRC)$(NC)"
	go build -o $(TARGET) $(SRC)
	@echo "$(RED)Usage:$(NC) ./$(TARGET) <arguments>"

rename:
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
	@echo "$(RED)Usage:$(NC) ./$(NAME) <arguments>"

clean:
	@rm -f $(TARGET) $(NAME) go.mod

re: clean all

.PHONY: all clean rename re
