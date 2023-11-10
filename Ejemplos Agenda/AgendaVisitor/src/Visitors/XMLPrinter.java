package Visitors;

import Data.Contacto;
import Data.ContactoFamiliar;
import Data.EventoReunion;
import Data.Visitor;


public class XMLPrinter implements Visitor<String> {
    @Override
    public String visitContactoFamiliar(ContactoFamiliar ctx) {
        StringBuffer xmlString = new StringBuffer();

        xmlString.append("contactoFamiliar en versión XML");

        return xmlString.toString();
    }

    @Override
    public String visitEventoReunion(EventoReunion ctx) {
        StringBuffer xmlString = new StringBuffer();

        xmlString.append("EventoReunión en versión XML");

        return xmlString.toString();
    }
}
