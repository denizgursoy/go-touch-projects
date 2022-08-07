#!/usr/local/bin/bash
echo "creating files"

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
