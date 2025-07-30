NAME = files-organizer
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
	@echo "$(RED)Usage: $(GREEN)./$(NAME) path/to/directory$(NC)"
	@echo "$(RED)Test: $(GREEN)make test$(NC)"

test:
	@if [ -d testfiles ]; then \
		rm -rf testfiles; \
	fi
	@mkdir testfiles
	@echo "Hello World" > testfiles/test.txt
	@echo "Hello World" > testfiles/test.jpg
	@echo "Hello World" > testfiles/test.pdf
	@echo "Hello World" > testfiles/test.doc
	@echo "Hello World" > testfiles/test.docx
	@echo "Hello World" > testfiles/test.mkv
	@echo "Hello World" > testfiles/test.mp4
	@echo "Hello World" > testfiles/test.lol
	@echo "Hello World" > testfiles/test.mov
	@echo "$(GREEN)Directory $(RED)testfiles$(GREEN) created !$(NC)"

clean:
	@rm -rf $(NAME) go.mod go.sum testfiles

re: clean all

.PHONY: all clean re test
