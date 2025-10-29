package co.edu.unicauca.capaDeControladores;

import java.rmi.Remote;
import java.rmi.RemoteException;
import co.edu.unicauca.fachadaServices.DTO.PreferenciasDTORespuesta;

//hereda de la clase Remote,lo cual la convierte en itnerfaz remota
public interface ControladorPreferenciasUsuariosInt extends Remote {
    //definicion del metodo remoto
    public PreferenciasDTORespuesta  getReferencias(Integer id) throws RemoteException;
    //cada definicion del metodo debe especificar que puede lanzar la excepcion java.rmi.RemoteException
        
}
