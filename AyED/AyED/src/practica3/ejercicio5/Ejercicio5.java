package practica3.ejercicio5;

import practica3.GeneralTree;
import resources.Util;

public class Ejercicio5 {

	public static void main(String[] args) {

		GeneralTree<Integer> root = Util.generarArbolGeneral(3, 4);
		root.visualize(3);
		System.out.println(root.esAncestro(3, 9));
		
	}

}
