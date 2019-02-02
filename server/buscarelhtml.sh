#!/bin/bash
echo "Comienzo a ejeuctar el script de python" 
cd webpage2html
python webpage2html.py $1 > $2
echo "Terminado!!" 
