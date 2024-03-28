//import { useState } from 'react'
import Header from './components/Header'
import './App.css'
//import inputUrl from './components/input'

function App() {
  return (
    <div>
      <Header />
      <div className="container" style={{ marginTop: '200px' }}>
        <div className="row">
          <div className="col-md-8">
            <div className="input-group">
              <input type="url" className="form-control rounded-pill-8" placeholder="https://" />
              <button type="button" className="btn btn-info">Shorten!!</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default App
