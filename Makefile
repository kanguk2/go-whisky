PROJECT_PATH := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
GENERATED_SRC_DIR ?= ${PROJECT_PATH}/generated

flywayClean:
	flyway -configFiles=$(PROJECT_PATH)/db/flyway.conf -locations=filesystem:$(PROJECT_PATH)/db/migration clean

flywayMigrate:
	 flyway -configFiles=$(PROJECT_PATH)/db/flyway.conf -locations=filesystem:$(PROJECT_PATH)/db/migration migrate

flywayInfo:
	flyway -configFiles=$(PROJECT_PATH)/db/flyway.conf -locations=filesystem:$(PROJECT_PATH)/db/migration info

flywayValidate:
	flyway -configFiles=$(PROJECT_PATH)/db/flyway.conf -locations=filesystem:$(PROJECT_PATH)/db/migration validate

clean: cleanGen
	rm -rf $(PROJECT_PATH)/bin

cleanGen:
	rm -rf $(PROJECT_PATH)/generated

genBoiler: cleanGen
	@echo "Generating sqlboiler models..."
	-@PSQL_DBNAME=go_whisky \
	PSQL_HOST=localhost \
	PSQL_PORT=5432 \
	PSQL_USER=go_whisky \
	PSQL_SSLMODE=disable \
	PSQL_PASSWORD= \
	PSQL_BLACKLIST=flyway_schema_history \
	sqlboiler --wipe --no-tests \
	-o ${GENERATED_SRC_DIR}/models psql $+ > /dev/null || \
		printf "\033[0;33m[WARN] Fail to generating sqlboiler models!\033[0m\n"
