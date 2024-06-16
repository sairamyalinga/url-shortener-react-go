import { useEffect } from "react";
import axios from 'axios';

function Urlboard(){
    useEffect(() => {
        let token = ""
        token = localStorage.getItem('token');
        axios.get("http://localhost:5050/api/geturls",{
            headers:{
                Authorization: `Bearer ${token}`
            }
        }).then(response => {
            console.log(response)
        }).catch(Error => {
            console.log("Err", Error);
    })

    }, [])
    return(
        <p>hi</p>
    )

}

export default Urlboard;