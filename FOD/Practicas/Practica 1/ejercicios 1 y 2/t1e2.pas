program t1e2;

var
  archivo: text;
  numero, cantMayor, cantTotal, sumTotal: Integer;
  avg : Real;
  nombreArchivo: String;

begin
  Write('Ingrese el nombre del archivo a leer: ');
  Readln(nombreArchivo);
  
  Assign(archivo, nombreArchivo);
  {$I-} // Desactiva el manejo de errores para la apertura del archivo
  Reset(archivo); // Abre el archivo en modo lectura
  {$I+} // Activa el manejo de errores nuevamente
  
  if IOResult <> 0 then // Verifica si hubo un error al abrir el archivo
  begin
    Writeln('Error al abrir el archivo.');
    case IOResult of
      2: Writeln('Archivo no encontrado.');
      5: Writeln('Permiso denegado.');
    else
      Writeln('Error desconocido.');
    end;
    Halt; // Finaliza el programa
  end;
  
  
  cantMayor := 0;
  cantTotal := 0;
  avg := 0;
  sumTotal := 0;
  
  Writeln('Contenido del archivo:');
  while not EOF(archivo) do // Mientras no se llegue al final del archivo
  begin
    Readln(archivo, numero); // Lee un número del archivo
    
    if (numero > 1500) then
		cantMayor := cantMayor +1;
	cantTotal := cantTotal + 1;
	sumTotal := sumTotal + numero;	
		   
    Writeln(numero); // Muestra el número en la consola
  end;  
  
  Close(archivo); // Cierra el archivo
  
  avg := (sumTotal / cantTotal);
  
  writeln('El promedio final fue de : ', avg:0:3); // Informa el promedio
  writeln('La cantidad de numeros mayores a 1500 fue de: ', cantMayor); // Informa la cantidad mayor a 1500
end.
