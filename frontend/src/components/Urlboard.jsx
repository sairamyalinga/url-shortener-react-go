import { useEffect, useState } from "react";
import axios from 'axios';
import PropTypes from 'prop-types';
import Header from "./Header";
import { handleCopyToClipboard } from "../lib/utils";

function Urlitem({url}){
    
    return(
    <div className="card border border-0 mb-3 mx-5">
      <div className="card-body rounded border border-info bg-info-subtle text-info-emphasis">
        <div className="d-flex justify-content-between align-items-center  ">
          <p className="mb-0">Short URL: <a href={`${url.ShortURL}`}>{`${url.ShortURL}`}</a></p>
          <div >
            <button className="btn btn-secondary btn-sm mr-2 mx-4" onClick={() => (
                handleCopyToClipboard(url.ShortURL)
            )}>
              Copy
            </button>
            <button className="btn btn-danger btn-sm">
              Delete
            </button>
          </div>
        </div>
        <p className="mb-0">Original URL: <a href={`${url.Url}`}>{url.Url}</a></p>
      </div>
    </div>
    )
}
function Urlboard({handleLogout}){
    const [data, setData] = useState([])
    useEffect(() => {
        let token = ""
        token = localStorage.getItem('token');
        axios.get("http://localhost:5050/api/geturls",{
            headers:{
                Authorization: `Bearer ${token}`
            }
        }).then(response => {
            setData(response.data)
        }).catch(Error => {
            console.log("Err", Error);
    })

    }, [])
    return (
        <div>
            <Header handleLogout={handleLogout} page = "urlboard"/>
            {data ? (data.map((url, index) => (
                
                <Urlitem
                key={index}
                url={url}/>
                
            ))):(<div className="d-flex align-items-center justify-content-center">
            <div className="badge text-bg-info text-wrap fs-5 fst-italic"> No shortURLs in your account.</div>
            </div>
            )}
        </div>
    );
}

Urlboard.propTypes = { 
    handleLogout: PropTypes.func.isRequired
  }

  Urlitem.propTypes ={
    url:PropTypes.object.isRequired
  }

export default Urlboard;