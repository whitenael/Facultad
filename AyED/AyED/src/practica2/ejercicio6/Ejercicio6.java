package practica2.ejercicio6;
import practica2.BinaryTree;
import resources.TreeTraversal;
import resources.Util;

public class Ejercicio6 {

	public static void main(String[] args) {

		BinaryTree<Integer> tree = Util.generarArbolEnteros(3);		
		TreeTraversal<Integer> tt = new TreeTraversal<Integer>();
		Transformacion tr = new Transformacion();
		
		tt.visualize(tree);
		System.out.println();
		tt.visualize(tr.suma(tree));
	}

}
