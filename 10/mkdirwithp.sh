#!/bin/bash

# set -e

mkdir one

mkdir one 
mkdir one/two
mkdir one/two/three
mkdir one/two/three/four
mkdir one/two/three/four/five


mkdir -p one
echo $?

mkdir -p one/two/three/four/five/six/seven/eight/nine/ten
echo $?


mkdir two/three/four/five/six/seven/eight/nine/ten/eleven/twelve/thirteen/fourteen/fifteen/sixteen/seventeen/eighteen/nineteen/twenty
echo $?