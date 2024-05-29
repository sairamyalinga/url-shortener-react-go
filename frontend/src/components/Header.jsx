import PropTypes from 'prop-types';
import logo from '/logo.jpeg'
import {IoIosLink} from "react-icons/io";
import { useNavigate } from "react-router-dom";

function Header({ logout, handleLogout }){

  const navigate = useNavigate();
 
    const signInPage = () => {
      navigate("/signup")
    }
    
  return (
  <header className="header">
      <div>
        <img src={logo} alt="logo" className="img-fluid float-end" style ={{height:"550px"}} />
        {logout?(<button className="sign-up-button" onClick={handleLogout}>Logout</button>):<button className="sign-up-button" onClick = {signInPage}>Sign In</button>}
        </div>
        <div className="card-body">
          <h1 className="card-title" >ShortURL <IoIosLink/></h1>
          <h4 className="card-subtitle mb-2 text-body-secondary">Make it easy to share.</h4>
        </div>
  </header>);
  
}

Header.propTypes = {
  logout: PropTypes.bool.isRequired,
  handleLogout: PropTypes.func.isRequired
};


export default Header;