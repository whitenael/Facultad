package practica2.ejercicio3;
import java.util.ArrayList;
import java.util.List;

import practica2.BinaryTree;

public class ContadorArbol {
	private BinaryTree<Integer> tree;
	
	private void numerosParesInOrden(BinaryTree<Integer> tree, ArrayList<Integer> lista) {					
		if (tree != null) {
			numerosParesInOrden(tree.getLeftChild(), lista);
			if (tree.getData() %2 == 0)
				lista.add(tree.getData());
			numerosParesInOrden(tree.getRightChild(), lista);				
		}
	}
	
	private void numerosParesPostOrden(BinaryTree<Integer> tree, ArrayList<Integer> lista){
		if (tree != null){
			numerosParesPostOrden(tree.getLeftChild(), lista);
			numerosParesPostOrden(tree.getRightChild(), lista);
			if (tree.getData() %2 == 0)
				lista.add(tree.getData());
		}
	}
	
	// el parametro orden sirve para saber que implementacion de recorrido tomar
	// 0 -> InOrden
	// 1 -> PostOrden
	public ArrayList<Integer> numerosPares(int orden) {
		ArrayList<Integer> lista = new ArrayList();
		
		switch(orden){
			case 0: numerosParesInOrden(tree, lista);
				break;
			case 1: numerosParesPostOrden(tree, lista);
				break;
			default: System.out.println("Parametro no válido");
		}
		
		return lista;
	}
	
	

	public BinaryTree<Integer> getTree() {
		return tree;
	}

	public void setTree(BinaryTree<Integer> tree) {
		this.tree = tree;
	}
	
}
