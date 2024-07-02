import { useEffect, useState } from "react";
import axios from "axios";
import PropTypes from "prop-types";
import Header from "./Header";
import { handleCopyToClipboard } from "../lib/utils";

function Urlitem({ url, handleDelete }) {
  const shorturl = `http://localhost:5050/${url.ShortID}`;
  
  return (
    <div className="card border border-0 mb-3 mx-5">
      <div className="card-body rounded border border-info bg-info-subtle text-info-emphasis">
        <div className="d-flex justify-content-between align-items-center  ">
          <p className="mb-0">
            Short URL: <a href = {shorturl} target="_blank">{shorturl}</a>
          </p>
          <div>
            <button
              className="btn btn-secondary btn-sm mr-2 mx-4"
              onClick={() => handleCopyToClipboard(shorturl)}
            >
              Copy
            </button>
            <button
              className="btn btn-danger btn-sm"
              onClick={() => handleDelete(url.ShortID)}
            >
              Delete
            </button>
          </div>
        </div>
        <p className="mb-0">
          Original URL: <a href={`${url.Url}`}>{url.URL}</a>
        </p>
      </div>
    </div>
  );
}

function Urlboard({ handleLogout }) {
  
  const [data, setData] = useState([]);
  const getURLs = () => {
    let token = "";
    token = localStorage.getItem("token");
    axios
      .get("http://localhost:5050/api/urls", {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      })
      .then((response) => {
        setData(response.data.data);
      })
      .catch((Error) => {
        console.log("Err", Error);
      });
  }
  useEffect(() => {
    getURLs();
  }, []);

  const handleDelete = (shortID) => {
    let token = "";
    if (localStorage.getItem("token") != null) {
      token = localStorage.getItem("token");
    }
    console.log(token);
    axios
      .delete("http://localhost:5050/api/url", {
        headers: {
          Authorization: `Bearer ${token}`,
        },
        data:{
          shortID,
        }
      })
      .then(() => {
        getURLs()
      })
      .catch((error) => {
        alert(error);
      });
  };
  
  return (
    <div>
      <Header handleLogout={handleLogout} page="urlboard" />
      {data ? (
        data.map((url, index) => <Urlitem key={index} url={url} handleDelete={() => handleDelete(url.ShortID)} />)
      ) : (
        <div className="d-flex align-items-center justify-content-center">
          <div className="badge text-bg-info text-wrap fs-5 fst-italic">
            {" "}
            No shortURLs in your account.
          </div>
        </div>
      )}
    </div>
  );
}

Urlboard.propTypes = {
  handleLogout: PropTypes.func.isRequired,
};

Urlitem.propTypes = {
  url: PropTypes.object.isRequired,
  handleDelete: PropTypes.func.isRequired
};

export default Urlboard;
