package practica3.ejercicio7;

import java.util.ArrayList;
import java.util.List;

import com.sun.xml.internal.bind.v2.runtime.reflect.ListTransducedAccessorImpl;

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
		
		List<Integer> list = new ArrayList<Integer>();
		List<Integer> auxList = new ArrayList<Integer>();
		auxList.add(this.tree.getData());

		return longestPath(this.tree, list, auxList);
	}
	
	private List<Integer> longestPath(GeneralTree<Integer> tree, List<Integer> list, List<Integer> auxList){
		
		if (tree == null)
			return null;				
		
		for (GeneralTree<Integer> child : tree.getChildren()){
			
			if (child == null)
				break;
			
			auxList.add(child.getData());
			
			if (auxList.size() > list.size())
				list = new ArrayList<Integer>(auxList);									
			
			return longestPath(child, list, auxList);			
		}
					
		return list;
	}
}
