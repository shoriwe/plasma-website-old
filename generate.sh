#!/bin/bash
go build -mod vendor -o /tmp/website

/tmp/website www
