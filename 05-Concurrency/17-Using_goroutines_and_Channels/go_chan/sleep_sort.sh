#!/bin/bash

for v in $@; do
  (sleep $v && echo $v)&
done
wait