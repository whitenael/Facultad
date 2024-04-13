program crearEmpleados;

const

	tipo_archivo = '.dat';

type
	emp = record
		nombre : string[30];		
		cht: integer;
	end;
	
	empFile = file of emp;
	
procedure leerEmp(var e : emp);
begin
	write('Ingrese el numero del empleado: ');
	readln(e.cht);
	if (e.cht <> 0) then begin
		write('Ingrese el nombre del empleado: ');
		readln(e.nombre);		
	end;
end;

var
  archivo: empFile;
  e : emp;
  nombreArchivo: String;

begin
  Write('Ingrese el nombre del archivo: ');
  Readln(nombreArchivo);
  
  Assign(archivo, nombreArchivo+tipo_archivo);
  Rewrite(archivo); // Crea o sobrescribe el archivo
  
  leerEmp(e);    
  
  while(e.cht <> 0) do begin
	write(archivo, e);
    leerEmp(e);
  end;
  
  Close(archivo);
  Writeln('Archivo creado exitosamente.');
end.
