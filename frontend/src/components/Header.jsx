import logo from '/logo.jpeg'
import {IoIosLink} from "react-icons/io";
function Header(){
    
  return (
  <header className="header">
      <div>
        <img src={logo} alt="logo" className="img-fluid float-end" style ={{height:"550px"}} />
      </div>
        <div className="card-body" style ={{padding:"60px", margin:"20px",}}>
          <h1 className="card-title" >ShortURL <IoIosLink/></h1>
          <h4 className="card-subtitle mb-2 text-body-secondary">Make it easy to share.</h4>
        </div>
  </header>);
  
}

export default Header;