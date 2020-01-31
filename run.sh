#!/bin/bash

EXT=c
LANG=c

if [ -d output ]; then
	echo "output folder exists so remove it first"
	rm -Rf output
fi
mkdir output

# loop over the student's submissions
for student in $(ls ./input); do
        # validate student identification
        if [ -d ./input/$student ] && [[ $student =~ ^[0-9]+$ ]]; then
                echo "student [$student]"
                i=0
                # loop over the problems
                # [ problem name can have spaces ]
                OIFS="$IFS"
                IFS=$'\n'
                for problem in $(ls ./input/$student); do
                        i=$((i+1))
                        printf "\t problem [$problem]\n"
                        mkdir -p "./output/${student}_$i"
                        # unzip each problem submission
                        for sub in $(ls ./input/$student/$problem/*.zip 2> /dev/null); do
                                echo $sub
                                printf "\t\t unzip $student - $problem [$i]\n"
                                unzip -q "$sub" -d "./output/${student}_$i"
                        done
                        # move raw files
                        for f in $(ls ./input/$student/$problem/*.$EXT 2> /dev/null); do
                                cp $f "./output/${student}_$i/"
                        done
                done
                IFS="$OIFS"
        fi
done

# remove cache files
find output -name __MACOSX -type d -print0 | xargs -0 rm -r --

# create the final directory put files into it with corrected filename
OIFS="$IFS"
IFS=$'\n'
for filename in $(find output -name "*.$EXT"); do
        better_name=$(echo "$filename" | sed -E  "s/([^a-zA-Z0-9\_\-\/\.])//g")
        better_name=$(echo $better_name | sed -E "s/output/final/")
        echo $better_name
        mkdir -p $(dirname $better_name)
        cp "$filename" "$better_name"
done
IFS="$OIFS"

./moss.pl -l $LANG `find final -name "*.$EXT" -type f -exec grep -Iq . {} \; -and  -print`

rm -Rf final # remove final directory
