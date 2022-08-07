#!/usr/local/bin/bash

for d in */; do
  base=$(basename "$d")

  if [ $base == "compressed" ]; then
    continue
  fi

  cd $base
  zip -r -qq $base *
  mv "${base}.zip" ../compressed
  cd ..
done
