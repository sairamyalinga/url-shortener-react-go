import {  useEffect} from 'react';
import './App.css'
import Dashboard from './components/Dashboard';
import { Routes, Route } from "react-router-dom"
import Signin from './components/Signin';
import { useNavigate, useLocation } from "react-router-dom";
import Urlboard from "./components/Urlboard"
import {  URLProvider } from './context/URLsContext';

function App() {

  const navigate = useNavigate();
  const location = useLocation();
  const curr_path = location.pathname;
  

  useEffect(() => {
    if (!localStorage.getItem('token') && curr_path !== "/") {
      navigate("/");
    }
  }, [curr_path, navigate]);


  return (
    <div>
      <URLProvider>
      <Routes>
      <Route path = "/" element = {<Signin /> }></Route>
      <Route path = "/dashboard" element = {<Dashboard  />}></Route>
      <Route path ="/urlboard" element={<Urlboard />}></Route>
      </Routes>  
      </URLProvider>
    </div>
  );
}

export default App
