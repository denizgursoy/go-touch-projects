#!/usr/local/bin/bash

for d in */; do
  base=$(basename "$d")

  if [ $base == "package" ]  || [ $base == "compressed" ] || [ $base == "test" ]; then
    continue
  fi

  cd $base
  zip -r -qq $base *
  mv "${base}.zip" ../compressed
  cd ..
done
