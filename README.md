Proyecto para clonar paginas a partir de una url

Etapas del proyecto:
1.- Clonado de paginas a partir del response usando la url (sin el https://)
2.- Logear los datos de los inputs que se ingresan al momento de usar la pagina clonada (por ahora en la consola del          navegador)
3.- Guardar el historial de log en firebase (para evitar tener un backend) o tener un backend con golang
4.- Ver la posibilidad de obtener los estilos para tener una mejor visualizacion de la pagina
5.- Mejorar las exepciones
6.- Agregar gift de cargando...
7.- Agregar semaforos para el manejo del script y los log en el servidor

Pasos:
1 .- A partir de la url 

### http://localhost:3000/<pagina sin el https//>

### Ejemplo http://localhost:3000/stackoverflow.com/questions/50420248/how-to-parse-a-string-containing-html-to-react-components#


Se debe mostrar la version clonada (posiblemente algunas carezcan de estilos).

2.- Ingresar algunos caracteres en los inputs que se encuentren.

3.- Ver en la consola que se este loggeando el registro de dichos valores.
