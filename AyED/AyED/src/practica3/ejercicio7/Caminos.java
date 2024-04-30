package practica3.ejercicio7;

import java.util.ArrayList;
import java.util.List;

import practica3.GeneralTree;

public class Caminos {
	
	public GeneralTree<Integer> tree;
	
	public Caminos(GeneralTree<Integer> tree) {
		super();
		this.tree = tree;
	}

	
	// Recorremos un camino y acumulamos el largo
	
	// Una vez recorrido el camino chequeamos si es mayor al anterior
	// 1. Si es mayor, durante la vuelva almacenamos los nodos en la lista
	// 2. Si es menor, ignoramos el camino y recorremos el siguiente
	
	// Retornamos la lista	
	

	public List<Integer> longestPath(){
		
		if (this.tree == null)
			return null;
		
		List<Integer> currentPath = new ArrayList<Integer>();
		List<Integer> longestPath = new ArrayList<Integer>();
		
		longestPath(this.tree, currentPath, longestPath);

		return longestPath;
	}
	
	private void longestPath(GeneralTree<Integer> tree, List<Integer> currentPath, List<Integer> longestPath){
		
		if (tree == null)
			return;				
		
		currentPath.add(tree.getData());
		
		if (currentPath.size() > longestPath.size()) {
			longestPath.clear();
			longestPath.addAll(currentPath);
		}
		
		for(GeneralTree<Integer> child : tree.getChildren()) {
			longestPath(child, currentPath, longestPath);
		}
		
		currentPath.remove(currentPath.size() - 1);
						
	}
}
