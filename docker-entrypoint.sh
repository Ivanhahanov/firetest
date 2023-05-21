#!/bin/sh
set -e

# untar rockyou
cd /wordlists
tar -xvzf rockyou.txt.tar.gz
rm rockyou.txt.tar.gz

# run app
firetest "$@"
