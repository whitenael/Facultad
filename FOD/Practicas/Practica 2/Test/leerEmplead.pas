program leerEmpleado;

type
	emp = record
		nombre : string[30];
		direccion: string[30];
		cht: integer;
	end;
	
	
	empFile = file of emp;
	
var
	arc : empFile;
	e : emp;

begin

	assign(arc, 'empleados.dat');	
	reset(arc);
	
	while(not eof(arc)) do begin
		read(arc, e);
		writeln('Nombre: ', e.nombre);	
		writeln('Direccion: ', e.direccion);	
		writeln('cht: ', e.cht);
		writeln;
	end;

end.
