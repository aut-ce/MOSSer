#!/bin/bash

EXT=java
LANG=java
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

files=$(find out/ -name "*.$EXT")
mkdir final

IFS=$'\n' 
for filename in $files; do
  echo $filename
  better_name=$(echo "$filename" | sed  -r "s/([^a-zA-Z0-9\/\.])//g")
  folder_name=$(dirname final/$better_name)
  mkdir -p "$folder_name"
  mv "$filename" $folder_name
done

./moss.pl -l $LANG -d final/out/**/**/**/**/**/**/**.$EXT


