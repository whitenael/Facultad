program testTextoBinario;

type 
	registroVotos = record
		codProv : integer;
		codLoc : integer;
		nroMesa : integer;
		cantVotos : integer;
		desc : string;
	end;
	
	archivoVotos = file of registroVotos;

var

	opc : Byte;
	nomArch, nomArch2 : string;
	archivo : tArchivoVotos;
	carga : text;
	votos : registroVotos;

begin
	
	writeln('Votos');}
	writeln;
	writeln('0. Terminar Programa');
	writeln('1. Crear un archivo binario desde un archivo de texto');
	writeln('2. Abrir un archivo binario y exportar a texto. ')
	repeat
		write('Ingrese el numero de opcion: ');
		if (opc=1) or (opc=2) then begin
			writeln;
			write('Nombre del archivo de votos: ');
			readln(nomArch);
			assign(arch, nomArch);
		end;
	
		case opc of 
			1: begin
				write('Nombre del archivo de carga: ');
				readln(nomArch2);
				assign(carga, nomArch2);
				reset(carga);
				rewrite(arch);
				while (not eof(carga)) do begin
					readln(carga, votos.codProv, votos.codLoc, votos.nroMesa, votos.cantVotos, votos.desc);
					write(arch, votos);
				end;
				write('Archivo cargado.');
				readln;
				close(arch); close(carga);
			end;
			2: begin
				write('Nombre del archivo de texto: ');
				readln(nomArch2);
				assign(carga, nomArch2);
				reset(arch);
				rewrite(carga);
				while (not eof(arch)) do begin
					read(arch, votos);
					writeln(carga, votos.codProv, ' ', votos.codLoc, ' ', votos.nroMesa, ' ', votos.cantVotos, ' ', votos.desc);
				end;
				close(arch); close(carga);
			end;
	Until (opc = 0);	
end.
