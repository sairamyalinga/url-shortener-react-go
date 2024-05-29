import { useEffect, useState } from 'react';
import './App.css'
import Dashboard from './components/Dashboard';
import { BrowserRouter, Routes, Route } from "react-router-dom"
import Signin from './components/Signin';


function App() {

  const [logout, setLogout] = useState(false);

  useEffect(()=>{
    if (localStorage.getItem('token') == null){
      setLogout(false)
    }else{
      setLogout(true)
    }
  }, []);

  const handleLogout = () => {
    localStorage.removeItem('token');
      setLogout(false)
  };

  return (
    <BrowserRouter>
    <div>
      <Routes>
      <Route path = "/SignUp" element = {<Signin setLogout={setLogout} /> }></Route>
      <Route path = "/" element = {<Dashboard logout={logout} handleLogout={handleLogout}  setLogout={setLogout}/>}></Route>
      </Routes>
      
    </div>
    </BrowserRouter>
  );
}

export default App
