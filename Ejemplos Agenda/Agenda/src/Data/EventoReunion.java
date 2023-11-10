/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package Data;

import GUI.AgregarEventoReunion;
import java.util.Date;
import java.util.LinkedList;

/**
 *
 * @author oviquez
 */
public class EventoReunion extends Evento{
    
    private int cantidadAsistentes;

    public EventoReunion(int cantidadAsistentes, Date fecha, String nombre) {
        super(fecha, nombre);
        this.cantidadAsistentes = cantidadAsistentes;
    }

    @Override
    public void imprimir() {
        System.out.println(this.toString());
    }

    @Override
    public void showGUI(LinkedList<Object> laAgenda) {
        (new AgregarEventoReunion(laAgenda)).setVisible(true);
    }
    
    

    @Override
    public String toString() {
        return "EventoReunion: \nNombre= " + getNombre() +
                                   "\nFecha= " + getFecha() +
                                   "\nCantidad de Asistentes= " + this.cantidadAsistentes;
    }
}
