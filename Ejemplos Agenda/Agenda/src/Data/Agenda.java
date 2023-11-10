/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package Data;

import java.util.LinkedList;
import java.util.Map;

/**
 *
 * @author oviquez
 */
public class Agenda {
    private LinkedList<Object> listaObjetos;
    
    private LinkedList<String> contactos;
    private LinkedList<String> eventos;

    public Agenda() {
        this.listaObjetos=new LinkedList<Object>();
        this.contactos = new LinkedList<String>();
        this.eventos = new LinkedList<String>();
    }

    public LinkedList<String> getContactos() {
        return contactos;
    }

    public LinkedList<String> getEventos() {
        return eventos;
    }

    public LinkedList<Object> getListaObjetos() {
        return listaObjetos;
    }
    
    
    
}
