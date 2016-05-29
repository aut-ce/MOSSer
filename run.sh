#!/bin/bash
for filename in *.zip; do
  better_name=$(echo "$filename" | sed 's, ,\ ,g')
  echo unzip "$better_name" -d "out/$better_name"
  unzip "$better_name" -d "out/$better_name"
done

for filename in *.rar; do
  better_name=$(echo "$filename" | sed 's, ,\ ,g')
  mkdir "out/$better_name"
  echo unrar x "$better_name" "out/$better_name"
  unrar x "$better_name" "out/$better_name"
done
