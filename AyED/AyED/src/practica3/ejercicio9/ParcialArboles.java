package practica3.ejercicio9;

import java.util.Comparator;
import java.util.List;
import java.util.NoSuchElementException;

import practica3.GeneralTree;
import resources.Util;

public class ParcialArboles {

	public static void main(String[] args) {
		
		GeneralTree<Integer> tree = Util.generarArbolGeneral(3, 2);
		tree.visualize(3);
		
		System.out.println();
		
		//System.out.println(esDeSeleccion(tree));
		
		esDeSeleccion(tree);
	}
	
	public static boolean esDeSeleccion(GeneralTree<Integer> tree) {		
		System.out.print(tree.getData() == getMin(tree));
		System.out.println(getMin(tree));
		
		// De esta manera podemos arrastrar una condicion boolean de forma recursiva
	    if (tree.getData() == getMin(tree)) {
	        for (GeneralTree<Integer> child : tree.getChildren()) {
	        	if (!child.isLeaf()) {
	        		if (!esDeSeleccion(child))
		                return false;
	        	}	            
	        }
	        return true;
	    }
	    return false;
	}
	
	public static void printMin(GeneralTree<Integer> tree) {	
		System.out.print(tree.getData() == getMin(tree));
		System.out.println(getMin(tree));
		for (GeneralTree<Integer> child : tree.getChildren()) {		
			if (!child.isLeaf())
				printMin(child);
		}
	}
	
	private static Integer getMin(GeneralTree<Integer> child) {						
		GeneralTree<Integer> min = child.getChildren()
			      .stream()
			      .min(Comparator.comparing(GeneralTree<Integer>::getData))
			      .orElseThrow(NoSuchElementException::new);
		
		return min.getData();
			    
	}

}
