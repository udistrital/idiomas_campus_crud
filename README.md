
# idiomas_crud
API para gestión de conocimientos en idiomas de una persona

Integración con

 - `CI`
 - `AWS Lambda - S3`
 - `Drone 1.x`
 - `idiomas_campus_crud master/develop`

## Requerimientos
Go version >= 1.8.

## Preparación
Para usar el API, usar el comando:

 - `go get github.com/udistrital/idiomas_campus_crud`

## Ejecución
Definir los valores de las siguientes variables de entorno:

 - `IDIOMAS_CAMPUS_CRUD_HTTP_PORT`: Puerto asignado para la ejecución del API
 - `IDIOMAS_CAMPUS_CRUD__PGUSER`: Usuario de la base de datos
 - `IDIOMAS_CAMPUS_CRUD__PGPASS`: Clave del usuario para la conexión a la base de datos  
 - `IDIOMAS_CAMPUS_CRUD__PGURLS`: Host de conexión
 - `IDIOMAS_CAMPUS_CRUD__PGDB`: Nombre de la base de datos
 - `IDIOMAS_CAMPUS_CRUD__SCHEMA`: Esquema a utilizar en la base de datos

## Ejemplo
IDIOMAS_CAMPUS_CRUD_HTTP_PORT=8103 IDIOMAS_CAMPUS_CRUD__PGUSER=user IDIOMAS_CAMPUS_CRUD__PGPASS=password IDIOMAS_CAMPUS_CRUD__PGURLS=localhost IDIOMAS_CAMPUS_CRUD__PGDB=bd IDIOMAS_CAMPUS_CRUD__SCHEMA=idiomas bee run

## Modelo BD
![image](https://github.com/udistrital/idiomas_crud/blob/develop/modelo_idiomas_crud.png).


