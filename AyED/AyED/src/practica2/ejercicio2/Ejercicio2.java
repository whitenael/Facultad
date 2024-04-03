package practica2.ejercicio2;
import practica2.BinaryTree;
import static resources.Util.generarArbolEnteros;
import resources.TreeTraversal;

import java.util.Random;

public class Ejercicio2 {
	
	public static void main(String[] args) {		
		BinaryTree<Integer> tree = generarArbolEnteros(4);
		BinaryTree<Integer> eert = tree.espejo();
		
		TreeTraversal<Integer> tt = new TreeTraversal<Integer>();
		tt.visualize(tree);
		
		
		System.out.println("Entre niveles: ");
		tree.entreNiveles(2, 3);
										
	}
	
	

}
