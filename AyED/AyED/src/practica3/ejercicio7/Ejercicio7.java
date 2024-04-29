package practica3.ejercicio7;

import practica3.GeneralTree;
import resources.Util;

public class Ejercicio7 {

	public static void main(String[] args) {

		GeneralTree<Integer> tree = Util.generarArbolGeneral_Aleatorio(5, 5);		
		tree.visualize(3);
		
		Caminos caminos = new Caminos(tree);
		for(Integer path : caminos.longestPath()){			
			System.out.print(path + ", ");
		}
	}

}
