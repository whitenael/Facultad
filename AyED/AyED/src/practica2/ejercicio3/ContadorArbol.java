package practica2.ejercicio3;
import java.util.ArrayList;
import java.util.List;

import practica2.BinaryTree;

public class ContadorArbol {
	private BinaryTree<Integer> tree;
	
	private void numerosParesRecursivo(BinaryTree<Integer> tree, ArrayList<Integer> lista) {					
		if (tree != null) {
			numerosParesRecursivo(tree.getLeftChild(), lista);
			if (tree.getData() %2 == 0)
				lista.add(tree.getData());
			numerosParesRecursivo(tree.getRightChild(), lista);				
		}
	}
	
	public ArrayList<Integer> numerosPares() {
		ArrayList<Integer> lista = new ArrayList();
		numerosParesRecursivo(this.tree, lista);
		return lista;
	}

	public BinaryTree<Integer> getTree() {
		return tree;
	}

	public void setTree(BinaryTree<Integer> tree) {
		this.tree = tree;
	}
	
}
