package tp5.ejercicio1b;

public interface Edge<T> {

    // returns the target vertex of the edge
    public Vertex<T> target();

    // returns edge weight
    public int getWeight();

}
