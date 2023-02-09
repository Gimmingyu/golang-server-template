OUTDIR = build
OUTFILE = appd
TARGET = ./cmd/main.go


all: build

build:
	@echo "\033[92mBuild server...\033[0m"
	@go build -o $(OUTDIR)/$(OUTFILE) $(TARGET)

clean:
	@echo "\033[92mClean daemon files...\033[0m"
	@rm $(OUTDIR)/*

fclean:
	@echo "\033[92mForce clean daemon files...\033[0m"
	@rm -f $(OUTDIR)/*

re:
	@make fclean
	@make all

.PHONY: appd clean all test build