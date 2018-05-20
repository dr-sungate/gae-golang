#!/bin/sh

GOPATH=`pwd`/gopath /google-cloud-sdk/platform/google_appengine/dev_appserver.py app/app.yaml --host=0.0.0.0 --admin_host=0.0.0.0

