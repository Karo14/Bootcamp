/* Ejercicio 1 */
use movies_db;
select * from movies;

/* Ejercicio 2 */
use movies_db;
select a.first_name nombre, a.last_name apellido, a.rating 
from actors a;

/* Ejercicio 3 */
use movies_db;
select s.title titulo 
from series s;

/* Ejercicio 4 */
use movies_db;
select a.first_name nombre, a.last_name apellido, a.rating 
from actors a
where rating > 7.5;

/* Ejercicio 5 */
use movies_db;
select m.title titulo, m.rating,m.awards premios
from movies m
where m.rating > 7.5 and m.awards > 2 ;

/* Ejercicio 6 */
use movies_db;
select m.title titulo, m.rating
from movies m
order by m.rating;

/* Ejercicio 7 */
use movies_db;
select m.title titulo
from movies m
limit 3;

/* Ejercicio 8 - Mostrar el top 5 de las películas con mayor rating. */
use movies_db;
select m.title titulo
from movies m
order by m.rating desc
limit 5;

/* Ejercicio 9 - Listar los primeros 10 actores. */
use movies_db;
select concat(a.first_name," ", a.last_name) Nombres
from actors a
limit 10;

/* Ejercicio 10 - Mostrar el título y rating de todas las películas cuyo título sea de Toy Story. */
use movies_db;
select m.title titulo,m.rating
from movies m
where m.title like "%Toy Story%";

/* Ejercicio 11 - Mostrar a todos los actores cuyos nombres empiezan con Sam. */
use movies_db;
select concat(a.first_name," ", a.last_name) Nombres
from actors a
where concat(a.first_name," ", a.last_name) like "Sam%";

/* Ejercicio 12 - Mostrar el título de las películas que salieron entre el 2004 y 2008. */
use movies_db;
select m.title titulo,extract(year from m.release_date) Año
from movies m
where extract(year from m.release_date) between "2004" and "2008";

/*Ejercicio 13 - Traer el título de las películas con el rating mayor a 3, con más de 1 premio y con fecha de lanzamiento entre el año 1988 al 2009. Ordenar los resultados por rating.*/
use movies_db;
select m.title titulo, m.rating
from movies m
where rating > 3 and awards > 1 and extract(year from m.release_date) between "1988" and "2009"
order by rating





