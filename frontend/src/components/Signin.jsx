import { useContext, useState } from 'react';
import { useNavigate } from "react-router-dom";
import API from '../lib/utils';
import { UrlDispatchContext } from '../context/URLsContext';

function Signin() {

    const [changeForm, setChangeForm] = useState(false);
    const dispatch = useContext(UrlDispatchContext)
    const navigate = useNavigate();
   
    const handlelogin = (event) => {
        event.preventDefault();
        const username = document.getElementById("loginuser").value;
        const pwd = document.getElementById("loginpwd").value;
        API
        .post('/login',{user_name:username, password:pwd})
        .then(response =>{
          const jwtToken = response.data.data.token;  
        //   console.log(response.data.data.token)
          localStorage.setItem('token', jwtToken);
          dispatch({type: 'SIGN_IN', payload:true})
          navigate("/dashboard");  
        })
        .catch(error =>{
          console.log('Error', error)
        });

    }

    const handlesignup = (event) => {
        event.preventDefault();
        const username = document.getElementById("signupuser").value;
        const pwd = document.getElementById("signuppwd").value; 
        API.post('/signup',{user_name:username, password:pwd})
        .then(response =>{    
          alert(response.data?.message || 'Sign Up Success!');
          setTimeout(() => {
            setChangeForm(false);
          }, 1000);
        })
        .catch(error =>{
          console.log('Error', error)
          alert('Sign Up Failed! Try again.');
        });
    }
    
    return (
        <div className="bg-container">
            <h1 className='position-absolute top-0 start-5 m-3 text-black'>ShortURL</h1>
            <div className="form-box">
                <div className="toggle-area">
                    <button 
                        className={`toggle-button ${changeForm ? "" : "active"}`} 
                        onClick={() => setChangeForm(false)}
                    >
                        Login
                    </button>
                    <button 
                        className={`toggle-button ${changeForm ? "active" : ""}`} 
                        onClick={() => setChangeForm(true)}
                    >
                        SignUp
                    </button>
                </div>

                <form className={`inpt-group ${changeForm ? "hidden" : ""}`}>
                    <input type="text" id="loginuser" className="login-input" autoComplete="off" placeholder="Enter Username"/>
                    <input type="password" id="loginpwd" className="login-input" placeholder="Enter Password"/>
                    <button onClick = {handlelogin}>Login</button>
                </form>

                <form className={`inpt-group ${changeForm ? "" : "hidden"}`}>
                    <input type="text"  id="signupuser"className="login-input" autoComplete="off" placeholder="Enter Username"/>
                    <input type="password" id="signuppwd"className="login-input" placeholder="Enter Password"/>
                    <button onClick = {handlesignup}>Register</button>
                </form>
            </div>
        </div>
    );
}

export default Signin;
