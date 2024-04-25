package practica3.ejercicio4;

import practica3.GeneralTree;
import resources.Util;

public class Ejercicio4 {

	public static void main(String[] args) {
		
		AnalizarArbol ab = new AnalizarArbol();
		GeneralTree<AreaEmpresa> tree = new GeneralTree<AreaEmpresa>();
		tree.visualize(3);
		
		System.out.println(ab.devolverMaximoPromedio(tree));

	}

}
