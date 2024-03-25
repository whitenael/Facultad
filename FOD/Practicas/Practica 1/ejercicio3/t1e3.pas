program t1e3;

const
	tipo_archivo = '.dat';

type

	Empleado = record
		numeroEmpleado: Integer;
		apellido: String;
		nombre: String;
		edad: Integer;	
	end;	
		
	archivo = file of Empleado;
	
procedure handleException(e : integer; nombre : string);
begin
	case e of
		2: Writeln(nombre + ' no encontrado.');
		5: Writeln('Error desconocido');
		100: Writeln('Error de entrada/salida');
	end;		
end;

procedure buscarEmpleado();
var e : empleado;
	numeroEmpleado : integer;
	arc : archivo;
	nom : string;
begin	

	writeln('Ingrese el nombre del archivo al que desea agregar un registro: ');
	readln(nom);
	
	writeln('Ingrese el numero del empleado que desea buscar: ');
	readln(numeroEmpleado);
	
	// Abrimos el archivo, manejamos las excepciones y agregamos el empleado al archivo
	Assign(arc, nom+tipo_archivo);
	{$I-}
	Reset(arc);
	{$I+}
	
	if IOResult <> 0 then begin
		handleException(IOResult, nom);
	end;
		
	// Lo cerramos
		Close(arc);	

	while not EOF(arc) do begin
		Read(arc, e);
		if (e.numeroEmpleado = numeroEmpleado) then begin			
			writeln('Empleado encontrado: ');
			writeln('Nombre: ', e.nombre);
			writeln('Apellido: ', e.apellido);
			writeln('Edad: ', e.edad);
			break;
		end;	
		
		writeln('No se ha encontrado al empleado');
	end;
end;

function existeEmpleado(var arc: archivo; numeroEmpleado : integer): boolean;
var ok : boolean;
	e : empleado;
begin
	ok := false;	
	while not EOF(arc) do begin
		Read(arc, e);
		if (e.numeroEmpleado = numeroEmpleado) then		
			ok := true;
	end;
		
	existeEmpleado := ok;
end;
		
procedure leerEmpleado(var e : Empleado);
begin
	Write('Numero de empleado: ');
	Readln(e.numeroEmpleado);

	if e.numeroEmpleado <> 0 then
	begin	
		Write('Apellido: ');
		Readln(e.apellido);
		Write('Nombre: ');
		Readln(e.nombre);
		Write('Edad: ');
		Readln(e.edad);		
	end;
end;

procedure crearMenu(var menu : text);
var opcion : string[30];
begin
	opcion := '-';
	
	//Assign(menu, 'menu');
	Rewrite(menu); // Crea o sobrescribe el archivo
	
	writeln('Ingrese el formato de su menu, cuando desee terminarlo, simplemente ingrese 0:');
	while(opcion <> '0') do begin
		Readln(opcion);
		if opcion <> '0' then
			writeln(menu, opcion);
	end;			
end;

procedure buscarMenu(var menu : text);
var nom : string;
begin
	// Chequear si existe un archivo llamado menu	
	nom := 'menu';
	Assign(menu, nom);	
	{$I-} // Desactiva el manejo de errores para la apertura del archivo
	Reset(menu); // Abre el archivo en modo lectura
	{$I+} // Activa el manejo de errores nuevamente

	if IOResult <> 0 then // Verifica si hubo un error al abrir el archivo
	begin
		handleException(IOResult, nom); // Manejamos el error de lectura
		// Si no encontramos el menu, le damos la opcion de crearlo y acceder al el
		crearMenu(menu);			
		Reset(menu); // Abrimos el menu
	end;
end;

procedure agregarEmpleado();
var e : empleado;
	arc : archivo;
	nom : string;
begin	
	writeln('Ingrese el nombre del archivo al que desea agregar un registro: ');
	readln(nom);
	
	// Leemos el nuevo registro
	leerEmpleado(e);
	
	// Abrimos el archivo, manejamos las excepciones y agregamos el empleado al archivo
	Assign(arc, nom+tipo_archivo);
	{$I-}
	Reset(arc);
	{$I+}
	
	if IOResult <> 0 then begin
		handleException(IOResult, nom);
	end;
	
	// Hacemos control de unicidad
	if (existeEmpleado(arc, e.NumeroEmpleado) <> true) then begin			
		// Una vez nos aseguramos el archivo es legible, ubicamos el puntero al ultimo	
		Seek(arc, FileSize(arc));
		
		// Escribimos el nuevo registro
		Write(arc, e);
		
		writeln('Nuevo empleado agregado exitosamente.');	
	end
	
	else 
		writeln('El empleado ya existe!');
		
	// Lo cerramos
		Close(arc);	
end;
	
procedure registrarEmpleados();
var arc : archivo;
	nombre : string[20];
	e : Empleado;
begin

	Writeln('Ingrese un nombre para el archivo'); // Le asignamos un nombre
	Readln(nombre);
	
	Assign(arc, nombre+tipo_archivo);
	Rewrite(arc); // Crea o sobrescribe el archivo

	Writeln('Ingrese los datos de los empleados (Ingrese numero de empleado 0 para finalizar):');
	repeat
		leerEmpleado(e);
		if (e.numeroEmpleado <> 0) then
			Write(arc, e);
	until e.numeroEmpleado = 0;
		Close(arc);
	Writeln('Archivo de empleados creado exitosamente.');
end;

procedure leerEmpleados();
var arc : archivo;
	nombre : string[20];
	e : Empleado;
begin
	writeln('Ingrese el nombre del archivo a abrir: ');
	readln(nombre);
		
	Assign(arc, nombre+tipo_archivo);
	{$I-}	
	Reset(arc);
	{$I+}
	
	if IOResult <> 0 then begin
		handleException(IOResult, nombre);
	end;
	
	while not EOF(arc) do // Mientras no se llegue al final del archivo
	begin
		Read(arc, e); // Lee un registro de empleado del archivo
		Writeln('Numero de empleado: ', e.numeroEmpleado);
		Writeln('Apellido: ', e.apellido);
		Writeln('Nombre: ', e.nombre);
		Writeln('Edad: ', e.edad);		
		Writeln;
	end;

	Close(arc); // Cierra el archivo
end;

var
	menu : text;		
	opcion : string[30];	
	userInput : char;

begin

	// inicializacion de variables
	
	userInput := '-';
	opcion := '';	
	
	buscarMenu(menu);
	
	
	while not EOF(menu) do // Mientras no se llegue al final del archivo
	begin
		Readln(menu, opcion); // Lee un número del archivo		
		Writeln(opcion); // Escribimos la opcion del menu
	end;
	

	while(userInput <> '0') do begin		
		writeln('--------');
		writeln('Seleccione una opcion del menu, pulse 0 para salir: ');
		readln(userInput);
		
		case userInput of
			'1': registrarEmpleados();
			'2': leerEmpleados();	
			'3': agregarEmpleado();		
			'4': buscarEmpleado();
		end;
	end;
	
	
	Close(menu); // Cierra el archivo
end.
