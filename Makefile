GIT				:= git
GO				:= go
FIRST_GOPATH	:= $(firstword $(subst :, ,$(shell $(GO) env GOPATH)))
GOAUTO          ?= $(FIRST_GOPATH)/bin/auto
MESSAGES		= $(shell $(GO) run ./cmd/main.go message)
REMARK 			?= ""
CM 				?= ""



$(GOAUTO): go.sum
	@echo "> installing auto"
	@$(GO) install "github.com/R0L/core/cmd/auto"

.PHONY: commit
commit: $(GOAUTO)
	$(GOAUTO)
	$(GIT) add .
	GIT_COMMITTER_DATE="$(REMARK)" $(GIT) commit -m "$(CM)" --date="$(REMARK)"


.PHONY: commits
commits: $(commit)
	@$(foreach message, $(MESSAGES), REMARK="$(shell echo "$(message)"|awk -F '|' '{print $$1}')"; CM="$(shell echo "$(message)"|awk -F '|' '{print $$2}')"; export REMARK; export CM; $(MAKE) commit;)