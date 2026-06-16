generate:
	for d in vantagev1/models vantagev1/vantage vantagev2/models vantagev2/vantage; do \
		if [ -d "$$d" ]; then rm -r "$$d"; fi; \
	done
	vantagev1/generate.sh
	vantagev2/generate.sh
	git add .
	# git commit -m "Generate clients."
	go mod tidy

diff:
	git diff HEAD^1..HEAD
