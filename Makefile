PREFIX ?= $(HOME)/.local/bin
export PATH := $(HOME)/.bun/bin:$(PATH)

.PHONY: build frontend install uninstall clean dev

frontend:
	cd web && bun install && bun run build

build: frontend
	go build -o syde ./cmd/syde/
	go build -o syded ./cmd/syded/

install: build
	install -d $(PREFIX)
	install -m 755 syde $(PREFIX)/syde
	install -m 755 syded $(PREFIX)/syded
	@if [ -d .syde ]; then \
		echo "Installing syde skill into this project..."; \
		$(PREFIX)/syde install-skill --all; \
	else \
		echo "No .syde/ in current directory — skipping skill install. Run 'syde init --install-skill' inside your project to set up the skill."; \
	fi

uninstall:
	rm -f $(PREFIX)/syde $(PREFIX)/syded

clean:
	rm -f syde syded
	rm -rf web/dist web/node_modules/.vite

dev:
	cd web && bun run dev
