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
	@printf "$(GREEN)[Running] $(RED)go build -o $(TARGET) $(SRC)$(NC)\n"
	go build -o $(TARGET) $(SRC)
	@printf "$(RED)Usage:$(NC) ./$(TARGET) <arguments>\n"

rename:
	@if [ -f $(NAME) ]; then \
		echo "$(NAME) already exists, skipping rename"; \
		exit 0; \
	elif [ -f $(TARGET) ]; then \
		rm -f $(TARGET); \
	fi
	@printf "$(GREEN)[Running] $(RED)go mod init $(NAME)$(NC)\n"
	go mod init $(NAME)
	@printf "$(GREEN)[Running] $(RED)go build -o $(NAME) $(SRC)$(NC)\n"
	go build -o $(NAME) $(SRC)
	@printf "$(RED)Usage:$(NC) ./$(NAME) <arguments>\n"

clean:
	@rm -f $(TARGET) $(NAME) go.mod

re: clean all

.PHONY: all clean rename re
