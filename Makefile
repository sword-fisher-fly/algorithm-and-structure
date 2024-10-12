
.PHONY: all

all: test

# Check the operating system
ifeq ($(shell uname), Darwin)
    OS := macOS
else ifeq ($(shell uname), Linux)
    OS := Linux
else
    $(error Unknown operating system)
endif

# Define variables based on the operating system
ifeq ($(OS), macOS)
    YESTERDAY := $(shell date -v -1d +"%Y-%m-%d")
else ifeq ($(OS), Linux)
    YESTERDAY := $(shell date -d 'yesterday' +"%Y-%m-%d")
endif

test:
	go test -v -timeout 30s ./array ./backtracing ./binarysearch ./dynamic ./graph ./greedy ./hash ./list ./math ./sort ./stack ./str ./twopointer 

test-all:
	go test -v -timeout 30s ./array ./backtracing ./binarysearch ./dynamic ./graph ./greedy ./hash ./list ./math ./sort ./stack ./str ./twopointer ./interview/...

yesterday-stats:
	@echo "Fetching commit statistics for $(YESTERDAY)..."; \
	git log --numstat --since='$(YESTERDAY)'  | \
	awk 'BEGIN {total_additions = 0; total_deletions = 0;} \
	{if ($$1 != "" && $$1 != "<" && $$1 != "commit") {total_additions += $$1; total_deletions += $$2;}} \
	END {print "Total additions: " total_additions; print "Total deletions: " total_deletions;}'

stats:
	@echo "Fetching commit statistics from init ..."; \
	git log --numstat  | \
	awk 'BEGIN {total_additions = 0; total_deletions = 0;} \
	{if ($$1 != "" && $$1 != "<" && $$1 != "commit") {total_additions += $$1; total_deletions += $$2;}} \
	END {print "Total additions: " total_additions; print "Total deletions: " total_deletions;}'


