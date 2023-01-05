#!/bin/bash

mkdir output
mkdir csv
start go clean ./...
start go build -o ./output ./...
sleep 4
start mosquitto -v

cd ./output

while [ true ]
do
  files=$(ls)

  i=0
  for file in $files; do
    i=$((i+1))
    echo "$i) $file"
  done

  echo "Entrez le numéro du fichier à exécuter (0 pour tous les fichiers, -1 pour quitter) :"
  read file_number

  # Si l'utilisateur a saisi -1, on close
  if [ $file_number -eq -1 ]
  then
    exit
  fi

  # Si l'utilisateur a saisi 0, on exécute tous les fichiers
  if [ $file_number -eq 0 ]
  then
    echo $files
    for file in $files
    do
      start $file
    done
    exit
  fi

  selected_file=$(echo "$files" | sed -n "${file_number}p")
  start $selected_file
  read -p "Appuyez sur une touche pour continuer..."
done

