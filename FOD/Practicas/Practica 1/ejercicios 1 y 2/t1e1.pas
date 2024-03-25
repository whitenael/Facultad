program t1e1;

var
  archivo: text;
  numero: Integer;
  nombreArchivo: String;

begin
  Write('Ingrese el nombre del archivo: ');
  Readln(nombreArchivo);
  
  Assign(archivo, nombreArchivo);
  Rewrite(archivo); // Crea o sobrescribe el archivo
  
  Writeln('Ingrese los numeros enteros (Ingrese 30000 para finalizar):');
  repeat
    Write('Numero: ');
    Readln(numero);
    
    if numero <> 30000 then
      writeln(archivo, numero);
  until numero = 30000;
  
  Close(archivo);
  Writeln('Archivo creado exitosamente.');
end.
