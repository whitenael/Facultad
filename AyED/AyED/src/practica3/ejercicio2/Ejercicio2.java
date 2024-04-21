package practica3.ejercicio2;

import java.util.List;

import practica3.GeneralTree;
import resources.Util;

public class Ejercicio2 {
	
	public static void main(String args[]) {
		
		GeneralTree<Integer> tree = Util.generarArbolGeneral(3, 4);
		tree.visualize(3);
		
		Recorridos r = new Recorridos();
									
		for (Integer it : r.numerosImparesMayoresPreOrden(tree, 3)) {
			System.out.println(it);
		}
	}

}
