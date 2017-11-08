#!/bin/sh
if [ $SW_ENV == "utests" ]; then
    go test -v -short;
else
    if [ $SW_ENV == "itests" ]; then
        go test -v
    else
        auth;
    fi
fi