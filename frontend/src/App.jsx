import { useState } from 'react'
import Header from './components/Header'
import './App.css'
import axios from 'axios'

function App() {

  const [shortURL, setShortURL] = useState('');

  const handleClick = () =>{
    const requestData = document.getElementById('urlinput').value;
    console.log(requestData)
    axios.post('http://localhost:5050/api/shorturl',{url:requestData})
      .then(response =>{
        console.log(response.data)
        setShortURL(response.data.shortURL)
      })
      .catch(error =>{
        console.log('Error', error)
      });
  };


  return (
    <div>
      <Header />
      <div className="container" style={{ marginTop: '200px' }}>
        <div className="row">
          <div className="col-md-8">
            <div className="input-group">
              <input type="url" id = "urlinput" className="form-control rounded-pill-8" placeholder="https://" />
              <button type="button" className="btn btn-info" onClick={handleClick} >Shorten!</button>
            </div>
            <div>
            {shortURL && <p>Short URL: {shortURL}</p>}
          </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default App
