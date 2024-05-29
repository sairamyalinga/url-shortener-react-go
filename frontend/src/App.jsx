
import './App.css'
import Dashboard from './components/Dashboard';
import { BrowserRouter, Routes, Route } from "react-router-dom"
import Signin from './components/Signin';
function App() {

 

  return (
    <BrowserRouter>
    <div>
      <Routes>
      <Route path = "/SignUp" element = {<Dashboard/>}></Route>
      <Route path = "/login" element = {<Dashboard/>}></Route>
      <Route path = "/" element = {<Signin/> }></Route>
      </Routes>
      
    </div>
    </BrowserRouter>
  );
}

export default App
