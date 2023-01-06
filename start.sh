#!/bin/bash

if [ ! -d "output" ]; then
  mkdir output
fi

if [ ! -d "csv" ]; then
  mkdir csv
fi

start go clean ./...

# Exécute la commande "go build" en arrière-plan et enregistre le PID dans une variable
go build -o ./output ./... &
pid=$!

# Attend que le processus avec le PID spécifié soit terminé
wait $pid

# Lance mosquitto
start mosquitto -v

cd ./output
files=$(ls)

while [ true ]
do

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

