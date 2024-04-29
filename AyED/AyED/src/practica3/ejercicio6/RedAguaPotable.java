package practica3.ejercicio6;

import practica3.GeneralTree;

public class RedAguaPotable {

	public GeneralTree<Character> red;

	public RedAguaPotable(GeneralTree<Character> red) {
		super();
		this.red = red;
	}

	public double minimoCaudal(double caudal) {
									
		return minimoCaudalRecursivo(this.red, caudal);
	}
	
	private double minimoCaudalRecursivo(GeneralTree<Character> tree, double caudal){
		
		if (tree == null)
			return caudal;
		
		System.out.println("Caudal del arbol " + tree.getData() + ": " + caudal);
		
		for (GeneralTree<Character> child : tree.getChildren()){						
			return minimoCaudalRecursivo(child, caudal / tree.getChildren().size());			
		}
		
		return 0;
		
	}

}
