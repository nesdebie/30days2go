TARGET = main
HELLO_WORLD = hello-world
SRC = main.go
GREEN = \033[0;32m
RED = \033[0;31m
NC = \033[0m # No Color
all: $(TARGET)

$(TARGET):
	@if [ -f $(TARGET) ]; then \
		@echo "hello-world already exists, skipping build"; \
		exit 0; \
	fi
	@echo "$(GREEN)[Running] $(RED)go build -o $(TARGET) $(SRC)$(NC)";
	go build -o $(TARGET) $(SRC)

run:
	@if [ -f $(TARGET) ]; then \
		echo "$(GREEN)[Running] $(RED)./$(TARGET)$(NC)"; \
		./$(TARGET); \
	elif [ -f $(HELLO_WORLD) ]; then \
		echo "$(GREEN)[Running] $(RED)./$(HELLO_WORLD)$(NC)"; \
		./$(HELLO_WORLD); \
	else \
		echo "$(GREEN)-> Running: $(RED)go run $(SRC)$(NC)"; \
		go run $(SRC); \
	fi

rename:
	@if [ -f $(HELLO_WORLD) ]; then \
		@echo "hello-world already exists, skipping rename"; \
		exit 0; \
	fi
	@echo "$(GREEN)[Running] $(RED)go mod init $(HELLO_WORLD)$(NC)";
	@go mod init $(HELLO_WORLD)
	@echo "$(GREEN)[Running] $(RED)go build -o $(HELLO_WORLD) $(SRC)$(NC)";
	go build -o $(HELLO_WORLD) $(SRC)

clean:
	rm -f $(TARGET) hello-world go.mod go.sum

re: clean all

.PHONY: all clean rename run re
