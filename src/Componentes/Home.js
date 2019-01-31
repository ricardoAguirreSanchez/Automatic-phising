import React, { Component } from 'react';
import '../App.css';

class Home extends Component {
    
    constructor(){
        super();
        this.state = {
            titulo : ''

        }
        
        this.toClone = this.toClone.bind(this);
    }


    toClone(e){
        e.preventDefault();
        debugger;
        console.log("...clonado!")
    }

  render() {
    return (
        <div className="container">
            <form id="create-course-form" className="card-body" onSubmit={this.toClone}>
                <div className="form-group">
                    <input type="text" placeholder="Titulo" name="titulo" 
                    className="form-control" onChange={this.cambioInput}/>
                </div>
                
                <div className="form-group">
                    <button type="submit" className="btn btn-primary">Clone</button>
                </div>
                
            </form>
        </div>
    );
  }
}

export default Home;
