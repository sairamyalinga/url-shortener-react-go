import { useEffect} from 'react';
import './App.css'
import Dashboard from './components/Dashboard';
import { Routes, Route } from "react-router-dom"
import Signin from './components/Signin';
import { useNavigate, useLocation } from "react-router-dom";
import Urlboard from "./components/Urlboard"

function App() {

  const navigate = useNavigate();
  const location = useLocation();
  const curr_path = location.pathname;

  useEffect(() => {
    if (!localStorage.getItem('token') && curr_path !== "/") {
      navigate("/");
    }
  }, [curr_path, navigate]);

  const handleLogout = () => {
    localStorage.removeItem('token');
      navigate("/")    
  };

  return (
    <div>
      <Routes>
      <Route path = "/" element = {<Signin /> }></Route>
      <Route path = "/dashboard" element = {<Dashboard handleLogout={handleLogout}  />}></Route>
      <Route path ="/urlboard" element={<Urlboard handleLogout={handleLogout}/>}></Route>
      </Routes>  
    </div>
  );
}

export default App
