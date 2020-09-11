Proyecto para clonar paginas a partir de una url

<ul>
  <li>Frontend con reactjs</li>
  <li>Backend con Golang (pendiente)</li>
</ul>


Etapas del proyecto:
<ol>
  <li>Clonado de paginas a partir del response usando la url (sin el https://)</li>
  <li>Logear los datos de los inputs que se ingresan al momento de usar la pagina clonada (por ahora en la consola del          navegador)</li>
  <li>Guardar el historial de log en firebase (para evitar tener un backend) o tener un backend con golang</li>
  
  <li>Mejorar las excepciones</li>
  <li>Agregar gift de cargando...</li>
  <li>Guardar el historial de log en firebase (para evitar tener un backend) o tener un backend con golang</li>
  
  <li>Agregar semaforos para el manejo del script y los log en el servidor</li>
</ol>

<h2>Pasos:</h2>
1 .- Copiar cualquier url a continuación del localhost

### http://localhost:3000/<pagina sin el https//>

![Alt text](capturas/captura_pagina_copiada.png?raw=true "Title")

2.- Ingresar algunos caracteres en los inputs que se encuentren.

![Alt text](capturas/captura_dato_input.png?raw=true "Title")

3.- Ver en el servidor un archivo log que llevara el registro de cada cambio

![Alt text](capturas/captura_log.png?raw=true "Title")
