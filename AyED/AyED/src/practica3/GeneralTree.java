package practica3;

import java.util.Iterator;
import java.util.LinkedList;
import java.util.List;

public class GeneralTree<T> {

	private T data;

	private List<GeneralTree<T>> children = new LinkedList<GeneralTree<T>>();

	public GeneralTree() {
	}

	public GeneralTree(T data) {
		this.data = data;
	}

	public GeneralTree(T data, List<GeneralTree<T>> children) {
		this.data = data;
		this.children = children;
	}

	public T getData() {
		return data;
	}

	public void setData(T data) {
		this.data = data;
	}

	public List<GeneralTree<T>> getChildren() {
		return children;
	}

	public void setChildren(List<GeneralTree<T>> children) {
		if (children != null)
			this.children = children;
	}

	public void addChild(GeneralTree<T> child) {
		getChildren().add(child);
	}

	public boolean isLeaf() {
		return !hasChildren();
	}

	public boolean hasChildren() {
		return children != null && !children.isEmpty();
	}

	public void removeChildren(GeneralTree<T> child) {
		if (hasChildren()) {
			children.remove(child);
		}
	}
	
	public boolean esAncestro(T a, T b){
		
		GeneralTree<T> nodeA = buscarNodo(this, a);
		
		if (nodeA == null)
			return false;	
		
		return esAncestroDe(nodeA, b);
	}
	

	public GeneralTree<T> buscarNodo(GeneralTree<T> root, T data) {
		// Vemos si estamos parados sobre el nodo y lo retornamos
		
		if (root == null)
			return null;
		
		if (root.getData().equals(data))
			return root;
		
		// Caso contrario, recorremos recursivamente, cuando lo encontremos, lo retornamos
		for(GeneralTree<T> child : root.getChildren()){
			GeneralTree<T> found = buscarNodo(child, data);
			if (found != null)
				return found;
		}
			
		
		return null;
	}
	
	private boolean esAncestroDe(GeneralTree<T> ancestro, T data){
		// Recorremos el nodo del ancestro y verificamos si en algun recorrido de ese ancestro encontramos el nodo
		
		if (ancestro == null)
			return false;
		
		if (ancestro.getData().equals(data))
			return true;
		
		// Recorremos recursivamente hasta verificar todos los hijos o encontrar el nodo
		for(GeneralTree<T> child : ancestro.getChildren()){
			if(esAncestroDe(child, data))
				return true;
		}
		
		return false;
	}

	public void preOrden() {
		System.out.println(getData());

		List<GeneralTree<T>> children = this.getChildren();
		for (GeneralTree<T> child : children) {
			if (hasChildren())
				child.preOrden();
		}
	}

	public void visualize(int depth) {
		// Imprimir espacios para representar la profundidad y la estructura del arbol
		for (int i = 0; i < depth - 1; i++) {
			System.out.print("   |");
		}

		// Imprimir guiones para indicar la profundidad y el nodo actual
		if (depth > 0) {
			System.out.print("---");
		}

		// Imprimir el dato del nodo actual
		System.out.println(getData());

		// Obtener los hijos del nodo actual
		List<GeneralTree<T>> children = this.getChildren();

		// Llamar recursivamente a preOrden() para cada hijo
		for (GeneralTree<T> child : children) {
			if (child != null)
				child.visualize(depth + 1); // Incrementar la profundidad para
											// los hijos
		}
	}
}
