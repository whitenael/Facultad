program p2ej8;

// procedimiento que busca los precios mas baratos

	// num y codigo son parametros de entrada || codigo2 y codigo1 devuelven los precios mas baratos || min2, min1 y aux controlan los minimos (a corregir)
	// min2, min1 y aux estan declarados como parametros para que se referencien con las variables declaradas en el programa principal, y no se pierda su valor una vez finalizado el proceso
	procedure productosBaratos(num:integer; codigo:integer; var codigo2, codigo1, min2, min1, aux: integer);
	
	begin
		
		if ((num < min2)and(aux=0)) then
		begin
			min2:=num;
			aux:=min2;
			codigo2:=codigo;
			write('Configurando los precios mas baratos1...'); //solo para debugging
		end;
		
		if ((num < min1)and(num<aux)) then
		begin
			min1:=num;
			aux:=0;
			codigo1:=codigo;
			write('Configurando los precios mas baratos2...'); // solo para debugging
		end;
		
					
	end;

// procedimiento que busca el codigo de un producto x mas caro (pantalon)

	// precio, codigo y  nombre son parametros de entrada, codigoFinal devuelve el codigo del pantalon mas caro, y max controla el maximo
	procedure pantalonMasCaro (precio, codigo: integer; nombre: String; var codigoFinal, max: integer);
	begin	
		if (nombre= 'Pantalon') then
		begin
			if (max<precio) then
			begin
				max:=precio;
				codigoFinal:=max;
			end;
		end;
	end;

// procedimiento que calcula el promedio de precios de nuestra ropa

	function precioPromedio (suma, cant: integer): integer;
	begin
	
		precioPromedio := suma div cant;
		
	end;

var 
	precio, codigo: integer;
	barato1, barato2, codigoPant, promedio, suma, productosTotales, min2, min1, aux, max: integer;
	nombre : String;

begin
	// variables locaes / referenciadas a los procesos
	productosTotales:=1;
	suma:=0;
	min1:=32000;
	min2:=32000;
	max:=-1;
	aux:=0;

	while (productosTotales <= 3) do
	begin
	
		// entrada de datos
		writeln('Producto #', productosTotales);
	
		writeln('Ingrese el precio del producto: ');
		readln(precio);
		writeln('Ingrese el codigo del producto: ');
		readln(codigo);
		writeln('Ingrese la descripcion del producto: ');
		readln(nombre);
		
		productosBaratos(precio, codigo, barato2, barato1, min2, min1, aux);	//llamada + parametros
		pantalonMasCaro(precio, codigo, nombre, codigoPant, max);				// llamada + parametros
		
		suma:=suma+precio; //suma de los precios totales
		productosTotales:=productosTotales+1;	// suma de productos totales
		
		promedio:=precioPromedio(suma, productosTotales); //calcula el promedio
		
		writeln();
	end;
	
	// informe
	writeln('Los 2 codigos de los productos mas barato son: ', barato2, ' y ', barato1);
	writeln('El codigo del pantalon mas caro es: ', codigoPant);
	writeln('El promedio de precios es: ', promedio);

end.
