#!/bin/sh
/usr/local/bin/swagger-cli bundle index.yml > combined_spec.json
/usr/local/bin/json2yml combined_spec.json >  combined_spec.yml
swagger generate server -A blackbox -t ../server -f ./combined_spec.yml