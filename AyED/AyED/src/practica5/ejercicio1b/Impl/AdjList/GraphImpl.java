package tp5.ejercicio1b.Impl.AdjList;

import tp5.ejercicio1b.Edge;
import tp5.ejercicio1b.Vertex;
import tp5.ejercicio1b.Graph;

import java.util.ArrayList;
import java.util.List;

public class GraphImpl<T> implements Graph<T> {

    private ArrayList<VertexImpl<T>> vertices;

    public GraphImpl(){
        this.vertices = new ArrayList<>();
    }

    @Override
    public Vertex<T> createVertex(T data) {
        int pos = this.vertices.size();
        VertexImpl<T> vertex = new VertexImpl<T>(data, pos);
        this.vertices.add(vertex);
        return vertex;
    }

    @Override
    public void removeVertex(Vertex<T> vertex) {
        int position = vertex.getPosition();
        if (this.vertices.get(position) != vertex)
            return;
        this.vertices.remove(vertex);
        for(; position < this.vertices.size(); position++){
            this.vertices.get(position).decrementPosition();
        }
        for(VertexImpl<T> other : this.vertices){
            this.disconnect(other, vertex);
        }
    }

    @Override
    public Vertex<T> search(T data) {
        for(Vertex<T> vertex : this.vertices){
            if (vertex.equals(data))
                return vertex;
        }

        return null;
    }

    public boolean belongs(Vertex<T> vertex){
        int pos = vertex.getPosition();
        boolean esValido = pos >= 0 && pos <= this.vertices.size();

        if (esValido)
            return this.vertices.get(pos).equals(vertex);
        else
            return false;
    }

    @Override
    public void connect(Vertex<T> origin, Vertex<T> destination) {
        if (this.belongs(origin) && this.belongs(destination)){
            ((VertexImpl<T>)origin).connect(destination);
        }
    }

    @Override
    public void connect(Vertex<T> origin, Vertex<T> destination, int weight) {
        if (this.belongs(origin) && this.belongs(destination)){
            ((VertexImpl<T>)origin).connect(destination, weight);
        }
    }

    @Override
    public void disconnect(Vertex<T> origin, Vertex<T> destination) {
        if (this.belongs(origin) && this.belongs(destination))
            ((VertexImpl<T>) origin).disconnect(destination);
    }

    @Override
    public boolean existsEdge(Vertex<T> origin, Vertex<T> destination) {
        if (this.belongs(origin)){
            Edge<T> edge = ((VertexImpl)origin).getEdge(destination);
            if (edge != null)
                return true;
        }

        return false;
    }

    @Override
    public boolean isEmpty() {
        return this.vertices.isEmpty();
    }

    @Override
    public List<Vertex<T>> getVertex() {
        return new ArrayList<>(this.vertices);
    }

    @Override
    public int weight(Vertex<T> origin, Vertex<T> destination) {
        Edge<T> edge = ((VertexImpl)origin).getEdge(destination);
        if (edge == null)
            return 0;
        return edge.getWeight();
    }

    @Override
    public List<Edge<T>> getEdges(Vertex<T> vertex) {
        if (belongs(vertex)){
            return new ArrayList<>(((VertexImpl<T>)vertex).getEdges());
        }
        return null;
    }

    @Override
    public Vertex<T> getVertex(int position) {
        if (position >= 0 && position <= this.vertices.size()){
            this.vertices.get(position);
        }
        return null;
    }

    @Override
    public int getSize() {
        return this.vertices.size();
    }

    @Override
    public String toString(){
        StringBuilder g = new StringBuilder();
        for(VertexImpl<T> vertex : this.vertices){
            g.append(vertex.getData()).append(" -> ");
            for(Edge<T> edge : vertex.getEdges()){
                g.append(edge.target().getData());
            }
            g.append("\n");
        }
        return g.toString();
    }
}
