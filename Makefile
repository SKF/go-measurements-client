.PHONY: clean
API_URL := "https://measurement-api.sandbox.iot.enlight.skf.com"

RM     ?= rm
WGET   ?= wget
MKDIR  ?= mkdir
DOCKER ?= docker

generate: rest/models

rest/models: rest/swagger.json
	$(RM) -rf "$@" && $(MKDIR) -p "$@"
	$(DOCKER) run --rm \
		--volume "$(shell pwd):/src" \
		--user "$(shell id -u):$(shell id -g)" \
		quay.io/goswagger/swagger:v0.25.0 \
			generate model --spec="/src/$<" --target="/src/$@/.." --skip-validation #TODO: remove this flag asap

rest/swagger.json:
	$(WGET) "$(API_URL)/docs/swagger/doc.json" -O "$@"
.PHONY: rest/swagger.json

clean:
	$(RM) -rf rest/models
