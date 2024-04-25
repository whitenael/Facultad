package practica3.ejercicio3;

import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

import practica3.GeneralTree;

public class Recorridos {
	
	public void inOrden(GeneralTree<Integer> tree){
		
		if (tree == null)
			return;
		
		List<GeneralTree<Integer>> children = tree.getChildren();
		
		if (!children.isEmpty()){
			inOrden(children.get(0));
		}
		
		System.out.println(tree.getData());
		
		for (int i = 1; i < children.size(); i++) {
			inOrden(children.get(i));
		}		
	}
	
	public void preOrden(GeneralTree<Integer> tree){
		
		if (tree == null)
			return;
		
		System.out.println(tree.getData());
		
		List<GeneralTree<Integer>> children = tree.getChildren();
		
		for(GeneralTree<Integer> child : children){			
			preOrden(child);
		}
	}
	
	public void niveles(GeneralTree<Integer> tree){
		if (tree == null)
			return;
		
		Queue<GeneralTree<Integer>> queue = new LinkedList<>();
		queue.offer(tree);
		
		while (!queue.isEmpty()){
			int level = queue.size();
			for (int i = 0; i < level; i++) {
				// Nos traemos el primer elemento de la cola			
				GeneralTree<Integer> node = queue.poll();				
				System.out.print(node.getData());
				
				// Agregamos sus hijos a la cola
				if (node.hasChildren()){					
					for(GeneralTree<Integer> child : node.getChildren()){
						if (child != null)
							queue.offer(child);
					}					
				}				
			}	
			System.out.println();
		}
	}

}
