package practica2.ejercicio5;

import practica2.BinaryTree;
import resources.TreeTraversal;
import resources.Util;

public class Ejercicio5 {

	public static void main(String[] args) {
		BinaryTree<Integer> tree = Util.generarArbolEnteros(3);
		ProfundidadDeArbolBinario prof = new ProfundidadDeArbolBinario(tree);
		TreeTraversal<Integer> tt = new TreeTraversal<Integer>();
		
		tt.visualize(tree);
		System.out.println();			
		prof.queueExample(tree);
		
		
		
	}

}
