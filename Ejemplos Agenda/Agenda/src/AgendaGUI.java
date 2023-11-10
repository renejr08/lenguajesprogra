/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Main.java to edit this template
 */

import Data.Agenda;
import Data.ObjAgenda;
import GUI.AgendaMain;

/**
 *
 * @author oviquez
 */
public class AgendaGUI {
    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
        
        
        Agenda miAgenda = new Agenda();
        
        AgendaMain guiMain = new AgendaMain(miAgenda);
        guiMain.setVisible(true);
    }
    
}
