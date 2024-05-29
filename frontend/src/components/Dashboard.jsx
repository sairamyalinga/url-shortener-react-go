import PropTypes from 'prop-types';
import { useState } from 'react'
import Header from './Header.jsx'
import axios from 'axios'

function Dashboard({ logout, handleLogout }){
    const [shortURL, setShortURL] = useState('');
    const [showURL, setShowURL] = useState(false);
  
    const handleClick = () =>{
      const requestData = document.getElementById('urlinput').value;
      console.log(requestData)
      axios.post('http://localhost:5050/api/shorturl',{url:requestData})
        .then(response =>{
          console.log(response.data)
          setShortURL(response.data.shortURL)
          setShowURL(true)
        })
        .catch(error =>{
          console.log('Error', error)
        });
    };
  
    const handleCopyToClipboard = () =>{
      navigator.clipboard.writeText(shortURL)
      .then(() => {
        console.log('Text copied to clipboard');
        alert('ShortURL copied to clipboard!');
      })
      .catch(err => {
        console.error('Failed to copy text: ', err);
      });
    };
  
    const handleClose = () =>{
      setShowURL(false)
      setShortURL('')
    }

    return (
        <div>
        <Header logout={logout} handleLogout={handleLogout}/>
        
        <div className="container" style={{ marginTop: '200px' }}>
        <div className="row">
          <div className="col-md-8">
            <div className="input-group">
              <input type="url" id = "urlinput" className="form-control rounded-pill-8" placeholder="https://" />
              <button type="button" className="btn btn-info" onClick={handleClick} >Shorten!</button>
            </div>
            <div>
            {showURL && (
        <div className="short-url-container">
          <p>Short URL: {shortURL}</p>
          <button className="btn btn-secondary" onClick={handleCopyToClipboard}>Copy to Clipboard</button>
          <button className="btn btn-danger" onClick={handleClose}>Close</button>
        </div>
      )}
          </div>
          </div>
        </div>
      </div>
      </div>
    );
  
}

Dashboard.propTypes = {
  logout: PropTypes.bool.isRequired,
  handleLogout: PropTypes.func.isRequired
};

export default Dashboard;