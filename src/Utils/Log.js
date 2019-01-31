
class Log {

/**
 * Aca estaria la logica para mandar un request al backen con los datos que se esperan loggear, 
 * o hacerlo por firebase, por ahora solo lo tiene en consola
 */
  generateMessage(type, message) {
    let today = new Date().toLocaleString();
    //[INFO] | 30/1/2019 12:07:10 | El HTML fue modificado correctamente.
    console.log("["+type+"] | " + today.toString()+" | "+message);
  }

  trace(message) {
    return this.generateMessage('TRACE', message);
  }
  
  info(message) {
    return this.generateMessage('INFO', message);
  }
  
  warn(message) {
    return this.generateMessage('WARN', message);
  }
  
  error(message) {
    return this.generateMessage('ERROR', message);
  }
}

export default new Log();