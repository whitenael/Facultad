package practica3.ejercicio4;

import java.util.LinkedList;
import java.util.Queue;

import practica3.GeneralTree;

public class AnalizarArbol {
	
	public double devolverMaximoPromedio(GeneralTree<AreaEmpresa> tree){
		
		if (tree == null)
			return 0;
		
		Queue<GeneralTree<AreaEmpresa>> queue = new LinkedList<>();
		queue.offer(tree);
		
		double sum = 0;
		double avg = 0;
		double max = -1;
		int currentLevel = 0;
		
		while (!queue.isEmpty()){
			int level = queue.size();
			for (int i = 0; i < level; i++) {
				// Nos traemos el primer elemento de la cola			
				GeneralTree<AreaEmpresa> node = queue.poll();				
				sum+= node.getData().getDelay();
				
				// Agregamos sus hijos a la cola
				if (node.hasChildren()){					
					for(GeneralTree<AreaEmpresa> child : node.getChildren()){
						if (child != null)
							queue.offer(child);
					}					
				}				
			}	
			currentLevel++;
			avg = sum / level;		
			System.out.println("Promedio " + currentLevel + ": " + avg);
			
			if (avg > max){
				max = avg;
			}				
					
		}
		
		return max;
	}

}
