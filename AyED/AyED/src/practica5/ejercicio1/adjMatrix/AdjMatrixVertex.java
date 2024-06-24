package tp5.ejercicio1.adjMatrix;

import tp5.ejercicio1.Vertex;

public class AdjMatrixVertex<T> implements Vertex<T> {
	private T data;
	private int position;
	private boolean visited;
	
	/**
	 * Constructor del vértices.  Solo se crean desde el grafo.
	 */
	AdjMatrixVertex(T data, int position) {
		this.data = data;
		this.position = position;
	}
	
	public T getData() {
		return this.data;
	}
	
	public void setData(T data) {
		this.data = data;
	}

	public int getPosition() {
		return position;
	}

	@Override
	public void setVisited(boolean visited) {
		this.visited = visited;
	}

	@Override
	public boolean isVisited() {
		return visited;
	}

	void setPosition(int position) {
		this.position = position;
	}

	void decrementPosition() {
		this.position--;
	}
}
