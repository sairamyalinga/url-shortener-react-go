import PropTypes from 'prop-types';
import { useState } from 'react'
import Header from './Header.jsx'
import API from '../lib/utils.js'
import { handleCopyToClipboard } from '../lib/utils.js';

function Dashboard({ handleLogout}){

    const [shortURL, setShortURL] = useState('');
    const [showURL, setShowURL] = useState(false);
    const [err, setErr] = useState('');

    const handleClick = () =>{
      const requestData = document.getElementById('urlinput').value;
      
      API.post('/url',{url:requestData})
        .then(response =>{
          setShortURL(response.data.data.shortURL)
          setShowURL(true)
        })
        .catch((error) =>{
          setShortURL('')
          setErr(error)
        });
    };
  
    
    const handleClose = () =>{
      setShowURL(false)
      setShortURL('')
    }

    return (
        <div>
        <Header handleLogout={handleLogout}/>
        <div className="container" style={{ marginTop: '200px' }}>
        <div className="row">
          <div className="col-md-8">
            <div className="input-group">
              <input type="url" id = "urlinput" className="form-control rounded-pill-8" placeholder="https://" />
              <button type="button" className="btn btn-info" onClick={handleClick} >Shorten!</button>
            </div>
            <div>
            {showURL ? 
            
              (<div className="short-url-container">
                  <p>Short URL: {shortURL}</p>
                  <button className="btn btn-secondary" onClick={() => ( handleCopyToClipboard(shortURL) )}>Copy to Clipboard</button>
                  <button className="btn btn-danger" onClick={handleClose}>Close</button>
              </div>):
              
             (err && (
             <div className="p-3  text-danger fw-bold">
                {err}
              </div>
              ))}
          </div>
          </div>
        </div>
      </div>
      </div>
    );  
}

Dashboard.propTypes = { 
  handleLogout: PropTypes.func.isRequired
}

export default Dashboard;