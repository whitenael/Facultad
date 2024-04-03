package practica2;



public class BinaryTree <T> {
	
	private T data;
	private BinaryTree<T> leftChild;   
	private BinaryTree<T> rightChild; 

	
	public BinaryTree() {
		super();
	}

	public BinaryTree(T data) {
		this.data = data;
	}

	public T getData() {
		return data;
	}

	public void setData(T data) {
		this.data = data;
	}
	/**
	 * Preguntar antes de invocar si hasLeftChild()
	 * @return
	 */
	public BinaryTree<T> getLeftChild() {
		return leftChild;
	}
	/**
	 * Preguntar antes de invocar si hasRightChild()
	 * @return
	 */
	public BinaryTree<T> getRightChild() {
		return this.rightChild;
	}

	public void addLeftChild(BinaryTree<T> child) {
		this.leftChild = child;
	}

	public void addRightChild(BinaryTree<T> child) {
		this.rightChild = child;
	}

	public void removeLeftChild() {
		this.leftChild = null;
	}

	public void removeRightChild() {
		this.rightChild = null;
	}

	public boolean isEmpty(){
		return (this.isLeaf() && this.getData() == null);
	}

	public boolean isLeaf() {
		return (!this.hasLeftChild() && !this.hasRightChild());

	}
		
	public boolean hasLeftChild() {
		return this.leftChild!=null;
	}

	public boolean hasRightChild() {
		return this.rightChild!=null;
	}
	
	
	@Override
	public String toString() {
		return this.getData().toString();
	}

	public int contarHojas() {
        return contarHojasRecursivo(this);
    }

    private int contarHojasRecursivo(BinaryTree<T> nodo) {
        if (nodo == null)
            return 0;
        if (nodo.isLeaf())
            return 1;
        return contarHojasRecursivo(nodo.getLeftChild()) + contarHojasRecursivo(nodo.getRightChild());
    }
				
    	 
    public BinaryTree<T> espejo(){
    	return espejoRecursivo(this);
    }
    
    private BinaryTree<T> espejoRecursivo(BinaryTree<T> nodo){
    	
    	if (nodo == null)
            return null;

        // Creamos un nuevo nodo con el mismo dato
        BinaryTree<T> espejo = new BinaryTree<>(nodo.getData());

        // Generamos el espejo de los hijos recursivamente
        espejo.addLeftChild(espejoRecursivo(nodo.getRightChild()));
        espejo.addRightChild(espejoRecursivo(nodo.getLeftChild()));

        return espejo;
    }
        
	// 0<=n<=m
	private void entreNivelesRecursivo(BinaryTree<T> nodo, int n, int m, int nivelActual)
	{
		if (nodo == null)
			return;
		
		if (nivelActual >= n && nivelActual <= m) 
		{
			System.out.println(nodo.getData());
		}
		
		if (nivelActual < m)
		{
			entreNivelesRecursivo(nodo.getLeftChild(), n, m, nivelActual+1);
			entreNivelesRecursivo(nodo.getRightChild(), n, m, nivelActual+1);
		}
	}
	
	public void entreNiveles(int n, int m) {
		entreNivelesRecursivo(this, n, m, 0);
	}
		
}

