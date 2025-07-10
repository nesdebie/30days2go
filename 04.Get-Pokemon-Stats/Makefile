TARGET = main
NAME = pokestats
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
	go get golang.org/x/text/unicode/norm
	@echo "$(GREEN)[Running] $(RED)go build -o $(NAME) $(SRC)$(NC)"
	go build -o $(NAME) $(SRC)
	@echo "$(RED)Usage: $(GREEN)./$(NAME) <pokemon>$(NC)"
	@echo "$(RED)Update csv: $(GREEN)make csv $(NC)"

csv:
	@if [ -f pokemon_names_multilang.csv ]; then \
		rm pokemon_names_multilang.csv; \
	fi
	go run utils/update_csv.go

clean:
	@rm -f $(NAME) go.mod go.sum

re: clean all

.PHONY: all clean csv re
