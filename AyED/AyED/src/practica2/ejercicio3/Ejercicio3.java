package practica2.ejercicio3;
import resources.TreeTraversal;
import resources.Util;

public class Ejercicio3 {

	public static void main(String[] args) {
		
		ContadorArbol cont = new ContadorArbol();
		cont.setTree(Util.generarArbolEnteros(4));
		TreeTraversal<Integer> tt = new TreeTraversal<Integer>();
		
		tt.visualize(cont.getTree());		
		
		for (Integer i : cont.numerosPares(0)) {
			System.out.println(i);
		}
		
	}

}
