.PHONY: dev install-tools

dev:
	@echo "Starting development server with auto-reload (air)"
	@if command -v air >/dev/null 2>&1; then \
		air; \
	else \
		echo "air is not installed. Run 'make install-tools' or install via 'brew install air' or 'go install github.com/cosmtrek/air@latest'"; \
	fi

install-tools:
	@echo "Installing air..."
	@go install github.com/cosmtrek/air@latest
