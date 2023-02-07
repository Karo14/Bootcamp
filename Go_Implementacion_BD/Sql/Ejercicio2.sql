use Ejercicio6;
CREATE TABLE Departamento(
    depto_nro 	   	VARCHAR(10),	
    nombre_depto	VARCHAR(50),
    localidad		VARCHAR(50),
    PRIMARY KEY (depto_nro)
);

use Ejercicio6;
CREATE TABLE Empleado(
    cod_emp 	   	VARCHAR(10),	
    nombre			VARCHAR(50),
    apellido		VARCHAR(50),
    puesto			VARCHAR(50),
    fecha_alta 		DATE,
	salario			DECIMAL(15,2),
    comision		DECIMAL(15,2),
    depto_nro		VARCHAR(10),
    PRIMARY KEY (cod_emp),
    FOREIGN KEY (depto_nro) REFERENCES Departamento(depto_nro)
);

INSERT INTO `Ejercicio6`.`Departamento`
(`depto_nro`,`nombre_depto`,`localidad`)
VALUES
("D-000-1","Software","Los Tigres"),
("D-000-2","Sistemas","Guadalupe"),
("D-000-3","Contabilidad","La Roca"),
("D-000-4","Ventas","Plata");


use Ejercicio6;
INSERT INTO `Empleado`
(`cod_emp`,`nombre`,`apellido`,`puesto`,`fecha_alta`,`salario`,`comision`,`depto_nro`)
VALUES
("E-0001","César","Piñero","Vendedor","2018/05/12",80000,15000,"D-000-4"),
("E-0002","Yosep","Kowaleski","Analista","2015/07/14","140000","0","D-000-2"),
("E-0003","Mariela","Barrios","Director","2014/06/05","185000","0","D-000-3"),
("E-0004","Jonathan","Aguilera","Vendedor","2015/06/03","85000","10000","D-000-4"),
("E-0005","Daniel","Brezezicki","Vendedor","2018/03/03","83000","10000","D-000-4"),
("E-0006","Mito","Barchuk","Presidente","2014/06/05","190000","0","D-000-3"),
("E-0007","Emilio","Galarza","Desarrollador","2014/08/02","60000","0","D-000-1");


/*Ejercicio 1 - Seleccionar el nombre, el puesto y la localidad de los departamentos donde trabajan los vendedores.*/
use Ejercicio6;
SELECT e.nombre, e.puesto, d.localidad
FROM Empleado e
INNER JOIN Departamento d ON (e.depto_nro=d.depto_nro)
WHERE d.nombre_depto = "Ventas";

/*Ejercicio 2 - Visualizar los departamentos con más de cinco empleados.*/
use Ejercicio6;
SELECT count(d.depto_nro)empleados,d.nombre_depto departamento
FROM Empleado e
INNER JOIN Departamento d ON (e.depto_nro=d.depto_nro)
group by 2
having empleados>5;

/*Ejercicio 3 - Mostrar el nombre, salario y nombre del departamento de los empleados que tengan el mismo puesto que ‘Mito Barchuk’.*/
use Ejercicio6;
SELECT CONCAT(e.nombre," ",e.apellido)Nombres, e.salario, d.nombre_depto
FROM Empleado e
INNER JOIN Departamento d ON (e.depto_nro=d.depto_nro)
WHERE e.puesto = (SELECT puesto FROM Empleado WHERE nombre="Mito" AND apellido="Barchuk");

/*Ejercicio 4 - Mostrar los datos de los empleados que trabajan en el departamento de contabilidad, ordenados por nombre.*/
use Ejercicio6;
SELECT e.*
FROM Empleado e
INNER JOIN Departamento d ON (e.depto_nro=d.depto_nro)
WHERE d.nombre_depto = "Contabilidad"
ORDER BY e.nombre;

/*Ejercicio 5 - Mostrar el nombre del empleado que tiene el salario más bajo.*/
use Ejercicio6;
SELECT concat(e.nombre," ",e.apellido) Nombre, MIN(e.salario) sueldo
FROM Empleado e
GROUP BY 1 
ORDER BY 2
LIMIT 1;

/*Ejercicio 6 - Mostrar los datos del empleado que tiene el salario más alto en el departamento de ‘Ventas’.*/
use Ejercicio6;
SELECT e.*
FROM Empleado e
INNER JOIN Departamento d ON (e.depto_nro=d.depto_nro)
WHERE d.nombre_depto = "Ventas"
ORDER BY e.salario desc
limit 1;




