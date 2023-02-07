/*
¿A qué se denomina JOIN en una base de datos y para qué se utiliza?
R/ Se utiliza para obtener datos de varias tablas relacionadas entre sí. 
Consiste en combinar datos de una tabla con datos de la otra tabla, a partir de una o varias condiciones en común.

Explicar dos tipos de JOIN.
R/  - Inner Join se utiliza para traer los datos relacionados de dos o más tablas.
	- Left Join se utiliza para traer los datos de la tabla izquierda más los relacionados de la tabla derecha.

¿Para qué se utiliza el GROUP BY?
R/ - Agrupa los resultados según las columnas indicadas. 
   - Genera un solo registro por cada grupo de filas que compartan las columnas indicadas.
   - Reduce la cantidad de filas de la consulta.
   - Se suele utilizar en conjunto con funciones de agregación, para obtener datos resumidos y agrupados por las columnas que se necesiten.

¿Para qué se utiliza el HAVING? 
R/  - La cláusula HAVING se utiliza para incluir condiciones con algunas funciones SQL.
	- Afecta a los resultados traidos por Group By.

¿En qué casos utilizarías una tabla temporal?
R/ - Se utiliza Para evitar multiples join en una misma consulta

¿En qué situaciones utilizamos índices? ¿Por qué no utilizarlos siempre?
R/  - Se utilizan para tener una PK y FK esto ayuda con el rendimiento de las consultas de la BD
	- si no se utiliza ocupa espacio en memoria cuando no está en constante uso 

Escribir una consulta genérica para cada uno de los siguientes diagramas:

- SELECT * FROM tablaA a 
  INNER JOIN tablaB b ON (a.id=b.id)

- SELECT * FROM tablaA a
  LEFT JOIN tablaB b ON (a.id=b.id)

*/

/* Ejercicio 1 - Mostrar el título y el nombre del género de todas las series.*/
use movies_db;
SELECT s.title titulo, g.name nombre
FROM series s 
INNER JOIN genres g on (s.genre_id=g.id);


/* Ejercicio 2 - Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos. */
use movies_db;
SELECT e.title titulo,ac.first_name nombre, ac.last_name apellido
FROM episodes e
INNER JOIN actor_episode a on (e.id=a.episode_id)
INNER JOIN actors ac on (a.actor_id=ac.id);

/* Ejercicio 3 - Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas. */
USE movies_db;
SELECT s.title titulo, count(se.id)temporadas
FROM series s
INNER JOIN seasons se on (s.id=se.serie_id)
GROUP BY 1;

/* Ejercicio 4 - Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3. */
USE movies_db;
SELECT gen.name nombre, count(m.id) Cantidad
FROM genres gen
INNER JOIN movies m ON (gen.id=m.genre_id)
GROUP BY 1
HAVING Cantidad>=3;

/* Ejercicio 5 - Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan. */
USE movies_db;
SELECT distinct a.first_name nombre, last_name apellido
FROM actors a 
INNER JOIN actor_movie am ON (a.id = am.actor_id)
INNER JOIN movies m ON (am.movie_id = m.id)
WHERE m.title like "La Guerra de las galaxias%"
ORDER BY 1,2;
