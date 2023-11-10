/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Main.java to edit this template
 */
package agendagui;

import agendagui.Data.Agenda;
import agendagui.Data.ObjAgenda;
import agendagui.GUI.AgendaMain;

import java.io.File;
import java.io.IOException;
import java.lang.reflect.InvocationTargetException;
import java.net.URL;
import java.net.URLClassLoader;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Iterator;
import java.util.LinkedList;
import java.util.Map;
import java.util.jar.JarFile;
import java.util.logging.Level;
import java.util.logging.Logger;

/**
 *
 * @author oviquez
 */
public class AgendaGUI {

    static final String PLUGIN_FOLDER="src/Plugins";
    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
        
        
        Agenda miAgenda = new Agenda();
        
        miAgenda.setClasesCargadas( cargarPlugins() );
        
        miAgenda.setClasesInstanciables( cargarClasesInstanciables(miAgenda.getClasesCargadas()) );
        
        Iterator<Map.Entry<String, Class>> iter = miAgenda.getClasesInstanciables().entrySet().iterator();
		if( iter != null ) {
			while( iter.hasNext() ) {
				Map.Entry<String, Class> entry = iter.next();
				miAgenda.getContactos().add(entry.getKey());
			}
		}
        
        AgendaMain guiMain = new AgendaMain(miAgenda);
        guiMain.setVisible(true);
    }
    
    private static Map cargarPlugins(){
        Map<String, Class> classList = new HashMap<>();
        //LinkedList<> classList = new LinkedList();
        File pluginFolder=new File(PLUGIN_FOLDER);
        if(!pluginFolder.exists())
        {
            if(pluginFolder.mkdirs())
            {
                System.out.println("Created plugin folder");
            }
        }
        File[] files=pluginFolder.listFiles((dir, name) -> name.endsWith(".jar"));
        ArrayList<URL> urls=new ArrayList<>();
        ArrayList<String> classes=new ArrayList<>();
        if(files!=null) {
            Arrays.stream(files).forEach(file -> {
                try {
                    JarFile jarFile=new JarFile(file);
                    urls.add(new URL("jar:file:"+PLUGIN_FOLDER+"/"+file.getName()+"!/"));
                    jarFile.stream().forEach(jarEntry -> {
                        if(jarEntry.getName().endsWith(".class"))
                        {
                            classes.add(jarEntry.getName());
                        }
                    });
                } catch (IOException e) {
                    e.printStackTrace();
                }
            });
            URLClassLoader pluginLoader=new URLClassLoader(urls.toArray(new URL[urls.size()]));
            
            classes.forEach(s -> {
                try {
                    StringBuilder s1 = new StringBuilder();
                    s1.append(s);
                    s1.reverse();
                    String name = s1.substring(0,s1.indexOf("/"));
                    s1 = new StringBuilder();
                    name = s1.append(name).substring(s1.indexOf(".")+1,s1.length());
                    s1 = new StringBuilder();
                    name = s1.append(name).reverse().toString();
                    Class classs=pluginLoader.loadClass(s.replaceAll("/",".").replace(".class",""));
                    classList.put(name,classs);
                }   catch (ClassNotFoundException ex) {
                    Logger.getLogger(AgendaGUI.class.getName()).log(Level.SEVERE, null, ex);
                }
            });
        }
        return classList;
    }
    
    public static Map cargarClasesInstanciables(Map listaClases){
        Map<String, Class> clasesInstanciables = new HashMap<String,Class>();
        //LinkedList clasesInstanciables = new LinkedList();
        listaClases.forEach((c,v) -> {
            Class[] interfaces=((Class)v).getInterfaces();
                        for (Class anInterface : interfaces) {
                            if(anInterface==ObjAgenda.class){
                                clasesInstanciables.put(((String)c),((Class)v));
                            }
                        }

                    });
        return clasesInstanciables;
    }
    
    
}
