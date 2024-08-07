package practica2.ejercicio5;

import java.util.Iterator;
import java.util.LinkedList;
import java.util.Queue;

import practica2.BinaryTree;

public class ProfundidadDeArbolBinario {
	BinaryTree<Integer> tree;
	
	public ProfundidadDeArbolBinario(){
		
	}
	
	public ProfundidadDeArbolBinario(BinaryTree<Integer> tree){
		this.tree = tree;
	}
	
	public int sumaElementosProfundidad(int p)
	{	
		// recorremos por niveles
	    if (this.tree == null)
	        return 0;
	
	    Queue<BinaryTree> queue = new LinkedList<>();
	    queue.offer(this.tree);
	    
	    // Para recordar:
	    
	    // offer : agrega un elemento al final de la lista (si la lista esta llena, devuelve false)
	    // poll: quita el primer elemento de la lista
	    // peek: devuelve el primer elemento de la lista sin retirarlo
	    
	    int currentLevel = 1;
	    int profundidad = 0;
	
	    while (!queue.isEmpty() && currentLevel <= p) {
            int levelSize = queue.size();
            for (int i = 0; i < levelSize; i++) {
            	BinaryTree node = queue.poll();
                System.out.print(node.getData() + " ");
                profundidad += (Integer)node.getData();
                // A�adir los hijos del nodo actual a la cola
                if (node.getLeftChild() != null)
                    queue.offer(node.getLeftChild());
                if (node.getRightChild() != null)
                    queue.offer(node.getRightChild());
            }
            System.out.println(); // Salto de l�nea despu�s de cada nivel
            currentLevel++;
        }
	
	
	    return profundidad;
	}
	
	public void queueExample(BinaryTree<Integer> tree) {
		
		if (tree == null)
			return;
		
		Queue<BinaryTree> queue = new LinkedList<>();
		queue.offer(tree);
		
		while(!queue.isEmpty()) {
			int levelSize = queue.size();
			
			for (int i = 0; i < levelSize; i++) {
				BinaryTree node = queue.poll();
				System.out.println(node.getData());
				
				if (node.hasLeftChild())
					queue.offer(node.getLeftChild());
				if(node.hasRightChild())
					queue.offer(node.getRightChild());					
			}
			System.out.println();
		}
		
		
		
	}

}
