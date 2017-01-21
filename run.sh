#!/bin/bash

mkdir out

for filename in *.zip; do
  better_name=$(echo "$filename" | sed 's, ,\ ,g')
  echo unzip "$better_name" -d "out/$better_name"
  unzip -q "$better_name" -d "out/$better_name"
done

for filename in *.rar; do
  better_name=$(echo "$filename" | sed 's, ,\ ,g')
  mkdir "out/$better_name"
  echo unrar -inul x "$better_name" "out/$better_name"
  unrar x "$better_name" "out/$better_name"
done
