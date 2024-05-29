import { useState } from 'react';


function Signin() {
    const [changeForm, setChangeForm] = useState(false);
    
    return (
        <div className="bg-container">
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
                    <input type="text" className="login-input" autoComplete="off" placeholder="Enter User Name"/>
                    <input type="password" className="login-input" placeholder="Enter Password"/>
                    <button>Login</button>
                </form>

                <form className={`inpt-group ${changeForm ? "" : "hidden"}`}>
                    <input type="text" className="login-input" autoComplete="off" placeholder="Enter User Name"/>
                    <input type="password" className="login-input" placeholder="Enter Password"/>
                    <button>Register</button>
                </form>
            </div>
        </div>
    );
}

export default Signin;
