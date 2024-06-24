package tp5.ejercicio2;

import tp5.ejercicio1.Edge;
import tp5.ejercicio1.Graph;
import tp5.ejercicio1.Vertex;
import tp5.ejercicio1.adjList.AdjListGraph;
import tp5.ejercicio1.adjList.AdjListVertex;

import java.util.*;

public class Recorridos<T> {

    public List<T> dfs(Graph<T> graph){
        Deque<Vertex<T>> stack = new LinkedList<>();
        stack.push(graph.getVertex(0));

        List<T> vertices = new ArrayList<>();

        while(!stack.isEmpty()){
            Vertex<T> current = stack.pop();
            if (!current.isVisited()){
                ((AdjListVertex<T>)current).setVisited(true);

                vertices.add(current.getData());

                List<Edge<T>> edges = ((AdjListVertex<T>) current).getEdges();
                for(Edge<T> edge : edges){
                    stack.push(edge.getTarget());
                }
            }
        }

        return vertices;
    }

    public List<T> bfs(Graph<T> graph)
    {
        Queue<Vertex<T>> queue = new LinkedList<>();
        queue.add(graph.getVertex(0));
        List<T> vertices = new ArrayList<>();

        while(!queue.isEmpty()){
            Vertex<T> current = queue.poll();
            if (!current.isVisited()){

                ((AdjListVertex<T>)current).setVisited(true);
                vertices.add(current.getData());

                List<Edge<T>> edges = ((AdjListVertex<T>) current).getEdges();
                for (Edge<T> edge : edges){
                    queue.add(edge.getTarget());
                }

            }

        }

        return vertices;
    }
}
