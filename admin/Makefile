.PHONY: init dev build clean

init:
	npm install

dev:
	npm run dev

build: clean
	npm run build
	mv -f dist/* ../blog/res/static/admin/

clean:
	rm -rf ../blog/res/static/admin/*
	touch ../blog/res/static/admin/.gitkeep

.DEFAULT_GOAL := build