package practica3.ejercicio3;

import practica3.GeneralTree;
import resources.Util;

public class Ejercicio3 {
	public static void main(String[] args){
		
		GeneralTree<Integer> tree = Util.generarArbolGeneral(3, 4);
		Recorridos rr = new Recorridos();
		tree.visualize(3);		
		
		System.out.println("------------------------");
		
		rr.inOrden(tree);
		
	}
}
