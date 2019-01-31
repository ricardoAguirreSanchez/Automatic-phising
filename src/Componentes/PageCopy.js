import React, { Component } from 'react';
import '../App.css';
import axios from 'axios';
import $ from 'jquery';
import HashMap from 'hashmap';
import logger from '../Utils/Log';


/**
 * Clase que representa una pagina clonada (con la captura de todo input en un log)
 */
class PageCopy extends Component {
    
    constructor(){
        super();
        this.state = {
            url : '',
            boddy: '',
            hasmapState: new HashMap()
        }
        this.modificado = false;
        this.cambioInput = this.cambioInput.bind(this);
        this.adaptarHTML = this.adaptarHTML.bind(this);
    }
    
    componentDidMount () {
        
        window.cambioInput = this.cambioInput;
        let url = this.props.location.pathname + this.props.location.search
        
        //arma la url para clonar
        let urlAClonar = 'https:/' + url
        logger.info('El url a consumir por axios es: ' + urlAClonar)
                
        const principal = this;
        axios.get(urlAClonar)
        .then(function (response) {
            //guardamos el response en el state
            logger.info('Recibimos el response correctamente.');
            debugger;
            principal.setState({
                ...principal.state,
                boddy: response.data,
                url : urlAClonar
            })
        })
        .catch(function (error) {
            logger.error(error);
        });
    }

    componentWillUnmount() {
        window.myfunc = null;
    }

    //capturamos las modificaciones en los input para luego logearlas
    cambioInput(e){
        let id = e.id;
        let value = e.value;
        this.setState({
            ...this.state,
            hasmapState: this.state.hasmapState.set(id,value),
        })
        logger.info(this.state.url + " se ingreso en el input id:" + id + " el valor:" + value);
    }

    //modifico el html del response, para todo input, si tiene id lo dejo, si no tiene le seteo un id y agrego el evento onChange
    adaptarHTML(){
        debugger;
        //let boddy = '<div><form><input type="search"/><input type="search" id="BBB"/></form></div>'; DUMMY
        let boddy = this.state.boddy;
        
        let processedHTML = '';
        
            let counter = 1;
            let $div = $('<div>').html(boddy);
                
                //busca los input y les agrega un id (si no lo tienen) y el evento onchange
                $div.find('input').attr({"onChange":"window.cambioInput(this)",'id':function(){
                    if (this.id === null || this.id ===""){
                        this.id = "ID-GENERADO-" + counter;
                        counter++;
                    }
                }});
                processedHTML = $div.html();
            if (boddy !== "" && this.modificado === false){
                
                if (counter> 1){
                    logger.info("El HTML fue modificado correctamente.");
                }else{
                    logger.info("El HTML no contiene inputs.");
                }
                this.modificado = true;
                
            }
            return processedHTML;
    }
    
  render() {
    return (
        <div className="container">
            {/* Esto devuelve el html de la pagina ya modificado */}
            <div dangerouslySetInnerHTML={{ __html: this.adaptarHTML() }}/>
        </div>
    );
  }
}

export default PageCopy;
