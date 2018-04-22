#!/bin/bash

EXT=java
LANG=java

if [ -d out ]; then
	echo "out folder is exist"
	rm -Rf out
fi
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
find out -name __MACOSX -type d -print0 | xargs -0 rm -r --

files=$(find out/ -name "*.$EXT")
mkdir final

IFS=$'\n'
for filename in $files; do
  echo $filename
  better_name=$(echo "$filename" | sed  "s/([^a-zA-Z0-9\/\.])//g")
  folder_name=$(dirname final/$better_name)
  mkdir -p "$folder_name"
  mv "$filename" $folder_name
done

./moss.pl -l $LANG -d `find final/out -name "*.$EXT" -type f -exec grep -Iq . {} \; -and  -print`
