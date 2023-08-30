generate:
	vantagev1/generate.sh
	vantagev2/generate.sh
	git add .
	git commit -m "Generate clients."

diff:
	git diff HEAD^1..HEAD
