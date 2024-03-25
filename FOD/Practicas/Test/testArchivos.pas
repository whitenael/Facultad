program testArchivo;

type
    persona = record
            dni : string[8];
            nombre_completo : string[30];
            direccion : string[40];
            sexo : char;
            salario : real;
    end;

    archivo = file of persona;
    
procedure readPersona(var p : persona; var output : archivo);
begin
	writeln('Ingrese el dni');
	readln(p.dni);
	writeln('Ingrese el nombre completo');
	readln(p.nombre_completo);
	writeln('Ingrese la direccion');
	readln(p.direccion);
	writeln('Ingrese el sexo');
	readln(p.sexo);
	writeln('Ingrese el salario');
	readln(p.salario);
	
	write(output, p);
	close(output);
end;

var
   personas : archivo;
   nombre_fisico : string[100];
   per : persona;

begin

     nombre_fisico := 'archivo_test.txt';
    
     {asignacion del nombre fisico}
     assign(personas, nombre_fisico);
     
     {apertura del archivo}
     rewrite(personas);
     
     {asignacion de los datos}
	 readPersona(per, personas);     
     
end.
