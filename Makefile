PROJECT_NAME = trendyol
OUT_DIR = bin
MAIN_FILE_DIR = cmd/$(PROJECT_NAME)/trendyol.go

build:
	@go build -o $(OUT_DIR)/$(PROJECT_NAME) $(MAIN_FILE_DIR)

run:
	@go run $(MAIN_FILE_DIR)

clear:
	@rm -rf $(OUT_DIR)/$(PROJECT_NAME)