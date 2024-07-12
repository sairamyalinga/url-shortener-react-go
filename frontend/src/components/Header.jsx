import PropTypes from 'prop-types';
import logo from '/logo.jpeg'
import {IoIosLink} from "react-icons/io";
import { useNavigate } from 'react-router-dom';
import { useContext } from 'react';
import { UrlDispatchContext } from '../context/URLsContext';


function Header({ page }){

  const navigate = useNavigate()
  const dispatch = useContext(UrlDispatchContext)
  const handleBoards = () => {
    if (page !== "urlboard"){
    navigate("/urlboard")
    }
    else{
      navigate("/dashboard")
    }


  }

  const handleLogout = () => {
    dispatch({ type: 'LOG_OUT' });
    localStorage.removeItem('token');
    navigate("/")    
  }
   
  return (
    <header className="header">
      <div>
        <img src={logo} alt="logo" className="img-fluid float-end" style ={{height:"550px"}} />
        <button className="sign-up-button" onClick={handleLogout}>Logout</button>
        <button className='urls-button' onClick={handleBoards}>{page === "urlboard"?"Create ShortURL": "Get URLS"}</button>
        </div>
        <div className="card-body">
          <h1 className="card-title" >ShortURL <IoIosLink/></h1>
          <h4 className="card-subtitle mb-2 text-body-secondary">Make it easy to share.</h4>
      </div>
    </header>); 
}

Header.propTypes = {
  page: PropTypes.string
};

export default Header;