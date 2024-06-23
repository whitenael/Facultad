package tp5.ejercicio1b;

import java.util.List;

public interface Graph<T> {

    // creates vertex with the given data
    public Vertex<T> createVertex(T data);

    // removes vertex from the graph along its connections with others
    public void removeVertex(Vertex<T> vertex);

    // searches and returns the first vertex with given data
    public Vertex<T> search(T data);

    // verifies both vertex exists in the graph and connects them
    public void connect(Vertex<T> origin, Vertex<T> destination);

    public void connect(Vertex<T> origin, Vertex<T> destination, int weight);

    // verifies both exist and disconnect the source vertex with the destination
    public void disconnect(Vertex<T> origin, Vertex<T> destination);

    // returns true if an edge exists between origin and destination
    public boolean existsEdge(Vertex<T> origin, Vertex<T> destination);

    // returns true if the graph has no data (i.e., no vertex created)
    public boolean isEmpty();

    // returns vertex list
    public List<Vertex<T>> getVertex();

    // returns weight between 2 vertices
    public int weight(Vertex<T> origin, Vertex<T> destination);

    // returns a list with the vertex adjacent edges
    public List<Edge<T>> getEdges(Vertex<T> v);

    // returns the vertex with the given position
    public Vertex<T> getVertex(int position);

    // returns the amount of vertices in the graph
    public int getSize();

}
