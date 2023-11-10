/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package agendagui.Data;

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
    
    private Map<String,Class> clasesCargadas;
    private Map<String,Class> clasesInstanciables;

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

    public Map<String, Class> getClasesCargadas() {
        return clasesCargadas;
    }

    public Map<String, Class> getClasesInstanciables() {
        return clasesInstanciables;
    }

    public void setClasesCargadas(Map<String, Class> clasesCargadas) {
        this.clasesCargadas = clasesCargadas;
    }

    public void setClasesInstanciables(Map<String, Class> clasesInstanciables) {
        this.clasesInstanciables = clasesInstanciables;
    }
    
    

    public LinkedList<Object> getListaObjetos() {
        return listaObjetos;
    }
    
    
    
}
