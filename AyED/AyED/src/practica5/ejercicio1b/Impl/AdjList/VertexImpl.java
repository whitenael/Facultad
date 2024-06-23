package tp5.ejercicio1b.Impl.AdjList;

import tp5.ejercicio1b.Edge;
import tp5.ejercicio1b.Vertex;

import java.util.ArrayList;
import java.util.List;

public class VertexImpl<T> implements Vertex<T> {

    private T data;
    private int position;
    private List<Edge<T>> edges;

    public VertexImpl(T data) {
        this.data = data;
    }

    public VertexImpl(T data, int position) {
        this.data = data;
        this.position = position;
        this.edges = new ArrayList<>();
    }

    public List<Edge<T>> getEdges(){
        return this.edges;
    }

    @Override
    public T getData() {
        return data;
    }

    @Override
    public void setData(T data) {
        this.data = data;
    }

    public void decrementPosition(){
        this.position--;
    }

    @Override
    public int getPosition() {
        return position;
    }

    protected void connect(Vertex<T> destination){
        this.connect(destination, 1);
    }

    protected void connect(Vertex<T> destination, int weight){
        Edge<T> edge = this.getEdge(destination);
        if (edge != null){
            // solo creamos el enlace si no existe
            this.edges.add(new EdgeImpl<T>(destination, weight));
        }
    }

    protected void disconnect(Vertex<T> destination){
        Edge<T> edge = this.getEdge(destination);
        if (edge != null){
            this.edges.remove(edge);
        }
    }

    private void disconnect(Vertex<T> destination, Vertex<T> origin){

    }

    protected Edge<T> getEdge(Vertex<T> target){
        if (target != null){
            for(Edge<T> edge : edges){
                if (edge == target)
                    return edge;
            }
        }
        return null;
    }
}
