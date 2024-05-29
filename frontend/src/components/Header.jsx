import logo from '/logo.jpeg'
import {IoIosLink} from "react-icons/io";
import { useNavigate } from "react-router-dom";
function Header(){

  const navigate = useNavigate();
 
    const signInPage = () => {
      navigate("/signup")
    }
    
  return (
  <header className="header">
      <div>
        <img src={logo} alt="logo" className="img-fluid float-end" style ={{height:"550px"}} />
        <button className="sign-up-button" onClick = {signInPage}>Sign In</button>
      </div>
        <div className="card-body">
          <h1 className="card-title" >ShortURL <IoIosLink/></h1>
          <h4 className="card-subtitle mb-2 text-body-secondary">Make it easy to share.</h4>
        </div>
  </header>);
  
}

export default Header;