program ConvertirTextoABinario;

type
    archivoTexto = text; // Definición del tipo de archivo de texto
    archivoBinario = file of string; // Definición del tipo de archivo binario

procedure ConvertirABinario(nomArchTexto: string; nomArchBinario: string);
var
    archTexto: archivoTexto;
    archBinario: archivoBinario;
    linea: string;
begin
    // Abre el archivo de texto para lectura
    assign(archTexto, nomArchTexto);
    reset(archTexto);

    // Crea el archivo binario para escritura
    assign(archBinario, nomArchBinario);
    rewrite(archBinario);

    // Lee línea por línea del archivo de texto y escribe en el archivo binario
    while not eof(archTexto) do
    begin
        readln(archTexto, linea);
        write(archBinario, linea);
    end;

    // Cierra los archivos
    close(archTexto);
    close(archBinario);

    writeln('Archivo de texto convertido a binario exitosamente.');
end;

procedure ConvertirATexto(nomArchBinario: string; nomArchTexto: string);
var
    archTexto: archivoTexto;
    archBinario: archivoBinario;
    linea: string;
begin
    // Abre el archivo binario para lectura
    assign(archBinario, nomArchBinario);
    reset(archBinario);

    // Crea el archivo de texto para escritura
    assign(archTexto, nomArchTexto);
    rewrite(archTexto);

    // Lee cada registro del archivo binario y escribe en el archivo de texto
    while not eof(archBinario) do
    begin
        read(archBinario, linea);
        writeln(archTexto, linea);
    end;

    // Cierra los archivos
    close(archTexto);
    close(archBinario);

    writeln('Archivo binario convertido a texto exitosamente.');
end;

procedure LeerArchivoBinario(nomArchBinario: string);
var    
    archBinario: archivoBinario;
    linea: string;
begin
    // Abre el archivo binario para lectura
    assign(archBinario, nomArchBinario);
    reset(archBinario);
    
    // Lee cada registro del archivo binario y escribe en consola
    while not eof(archBinario) do
    begin
        read(archBinario, linea);
        writeln(linea);
    end;

    // Cierra los archivos    
    close(archBinario);   
end;

begin
    LeerArchivoBinario('archivo.bin');
end.
